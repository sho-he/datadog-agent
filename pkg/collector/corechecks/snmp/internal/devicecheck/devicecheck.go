// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2016-present Datadog, Inc.

//nolint:revive // TODO(NDM) Fix revive linter
package devicecheck

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"runtime"
	"strings"
	"time"

	"go.uber.org/atomic"

	"github.com/DataDog/datadog-agent/comp/core/config"
	haagenthelpers "github.com/DataDog/datadog-agent/comp/haagent/helpers"
	"github.com/DataDog/datadog-agent/pkg/collector/externalhost"
	configUtils "github.com/DataDog/datadog-agent/pkg/config/utils"
	"github.com/DataDog/datadog-agent/pkg/metrics/servicecheck"
	"github.com/DataDog/datadog-agent/pkg/util/hostname/validate"
	"github.com/DataDog/datadog-agent/pkg/util/log"

	"github.com/DataDog/datadog-agent/pkg/networkdevice/metadata"
	"github.com/DataDog/datadog-agent/pkg/networkdevice/pinger"
	"github.com/DataDog/datadog-agent/pkg/networkdevice/utils"
	coresnmp "github.com/DataDog/datadog-agent/pkg/snmp"

	"github.com/DataDog/datadog-agent/pkg/collector/corechecks/snmp/internal/checkconfig"
	"github.com/DataDog/datadog-agent/pkg/collector/corechecks/snmp/internal/common"
	"github.com/DataDog/datadog-agent/pkg/collector/corechecks/snmp/internal/fetch"
	"github.com/DataDog/datadog-agent/pkg/collector/corechecks/snmp/internal/report"
	"github.com/DataDog/datadog-agent/pkg/collector/corechecks/snmp/internal/session"
	"github.com/DataDog/datadog-agent/pkg/collector/corechecks/snmp/internal/valuestore"
	"github.com/DataDog/datadog-agent/pkg/diagnose/diagnosis"
	"github.com/DataDog/datadog-agent/pkg/networkdevice/diagnoses"
	"github.com/DataDog/datadog-agent/pkg/persistentcache"
	"github.com/DataDog/datadog-agent/pkg/snmp/gosnmplib"
)

const (
	snmpLoaderTag           = "loader:core"
	serviceCheckName        = "snmp.can_check"
	deviceReachableMetric   = "snmp.device.reachable"
	deviceUnreachableMetric = "snmp.device.unreachable"
	pingReachableMetric     = "networkdevice.ping.reachable"
	pingUnreachableMetric   = "networkdevice.ping.unreachable"
	pingPacketLoss          = "networkdevice.ping.packet_loss"
	pingAvgRttMetric        = "networkdevice.ping.avg_rtt"
	deviceHostnamePrefix    = "device:"
	checkDurationThreshold  = 30 // Thirty seconds
)

// DeviceCheck hold info necessary to collect info for a single device
type DeviceCheck struct {
	config                  *checkconfig.CheckConfig
	sender                  *report.MetricSender
	session                 session.Session
	sessionFactory          session.Factory
	devicePinger            pinger.Pinger
	sessionCloseErrorCount  *atomic.Uint64
	savedDynamicTags        []string
	diagnoses               *diagnoses.Diagnoses
	interfaceBandwidthState report.InterfaceBandwidthState
	cacheKey                string
	agentConfig             config.Component
}

const cacheKeyPrefix = "snmp-tags"

// NewDeviceCheck returns a new DeviceCheck
func NewDeviceCheck(config *checkconfig.CheckConfig, ipAddress string, sessionFactory session.Factory, agentConfig config.Component) (*DeviceCheck, error) {
	newConfig := config.CopyWithNewIP(ipAddress)

	var devicePinger pinger.Pinger
	var err error
	if newConfig.PingEnabled {
		devicePinger, err = createPinger(newConfig.PingConfig)
		if err != nil {
			return nil, fmt.Errorf("failed to create pinger: %s", err)
		}
	}

	configHash := newConfig.DeviceDigest(newConfig.IPAddress)
	cacheKey := fmt.Sprintf("%s:%s", cacheKeyPrefix, configHash)

	d := DeviceCheck{
		config:                  newConfig,
		sessionFactory:          sessionFactory,
		devicePinger:            devicePinger,
		sessionCloseErrorCount:  atomic.NewUint64(0),
		diagnoses:               diagnoses.NewDeviceDiagnoses(newConfig.DeviceID),
		interfaceBandwidthState: report.MakeInterfaceBandwidthState(),
		cacheKey:                cacheKey,
		agentConfig:             agentConfig,
	}

	d.readTagsFromCache()

	return &d, nil
}

// SetSender sets the current sender
func (d *DeviceCheck) SetSender(sender *report.MetricSender) {
	d.sender = sender
}

// SetInterfaceBandwidthState sets the interface bandwidth state
func (d *DeviceCheck) SetInterfaceBandwidthState(state report.InterfaceBandwidthState) {
	d.interfaceBandwidthState = state
}

// GetInterfaceBandwidthState returns interface bandwidth state
func (d *DeviceCheck) GetInterfaceBandwidthState() report.InterfaceBandwidthState {
	return d.interfaceBandwidthState
}

// GetIPAddress returns device IP
func (d *DeviceCheck) GetIPAddress() string {
	return d.config.IPAddress
}

// GetDeviceID returns device ID
func (d *DeviceCheck) GetDeviceID() string {
	return d.config.DeviceID
}

// GetIDTags returns device IDTags
func (d *DeviceCheck) GetIDTags() []string {
	return d.config.DeviceIDTags
}

// GetDeviceHostname returns DeviceID as hostname if UseDeviceIDAsHostname is true
func (d *DeviceCheck) GetDeviceHostname() (string, error) {
	if d.config.UseDeviceIDAsHostname {
		hostname := deviceHostnamePrefix + d.config.DeviceID
		normalizedHostname, err := validate.NormalizeHost(hostname)
		if err != nil {
			return "", err
		}
		return normalizedHostname, nil
	}
	return "", nil
}

// Run executes the check
func (d *DeviceCheck) Run(collectionTime time.Time) error {
	startTime := time.Now()
	staticTags := append(d.config.GetStaticTags(), d.config.GetNetworkTags()...)

	var err error
	d.session, err = d.sessionFactory(d.config)
	if err != nil {
		return err
	}

	// Fetch and report metrics
	var checkErr error
	var deviceStatus metadata.DeviceStatus
	var pingStatus metadata.DeviceStatus

	deviceReachable, dynamicTags, values, checkErr := d.getValuesAndTags()

	tags := utils.CopyStrings(staticTags)
	if checkErr != nil {
		tags = append(tags, d.savedDynamicTags...)
		d.sender.ServiceCheck(serviceCheckName, servicecheck.ServiceCheckCritical, tags, checkErr.Error())
	} else {
		if !reflect.DeepEqual(d.savedDynamicTags, dynamicTags) {
			d.savedDynamicTags = dynamicTags
			d.writeTagsInCache()
		}

		tags = append(tags, dynamicTags...)
		d.sender.ServiceCheck(serviceCheckName, servicecheck.ServiceCheckOK, tags, "")
	}

	metricTags := append(tags, "dd.internal.resource:ndm_device_user_tags:"+d.GetDeviceID())
	d.sender.Gauge(deviceReachableMetric, utils.BoolToFloat64(deviceReachable), metricTags)
	d.sender.Gauge(deviceUnreachableMetric, utils.BoolToFloat64(!deviceReachable), metricTags)
	if values != nil {
		d.sender.ReportMetrics(d.config.Metrics, values, metricTags, d.config.DeviceID)
	}

	// Get a system appropriate ping check
	if d.devicePinger != nil {
		log.Tracef("%s: pinging host", d.config.IPAddress)
		pingResult, err := d.devicePinger.Ping(d.config.IPAddress)
		if err != nil {
			// if the ping fails, send no metrics/metadata, log and add diagnosis
			log.Errorf("%s: failed to ping device: %s", d.config.IPAddress, err.Error())
			pingStatus = metadata.DeviceStatusUnreachable
			d.diagnoses.Add("error", "SNMP_FAILED_TO_PING_DEVICE", "Agent encountered an error when pinging this network device. Check agent logs for more details.")
			d.sender.Gauge(pingReachableMetric, utils.BoolToFloat64(false), metricTags)
			d.sender.Gauge(pingUnreachableMetric, utils.BoolToFloat64(true), metricTags)
		} else {
			// if ping succeeds, set pingCanConnect for use in metadata and send metrics
			log.Debugf("%s: ping returned: %+v", d.config.IPAddress, pingResult)
			if pingResult.CanConnect {
				pingStatus = metadata.DeviceStatusReachable
			} else {
				pingStatus = metadata.DeviceStatusUnreachable
			}
			d.submitPingMetrics(pingResult, metricTags)
		}
	} else {
		log.Tracef("%s: SNMP ping disabled for host", d.config.IPAddress)
	}

	if d.config.CollectDeviceMetadata {
		if deviceReachable {
			deviceStatus = metadata.DeviceStatusReachable
		} else {
			deviceStatus = metadata.DeviceStatusUnreachable
		}

		checkDuration := time.Since(startTime).Seconds()

		if checkDuration > checkDurationThreshold {
			d.diagnoses.Add("warn", "SNMP_HIGH_CHECK_DURATION", fmt.Sprintf("Check duration is high for this device, last check took %.2f seconds.", checkDuration))
		}

		// We include instance tags to `deviceMetadataTags` since device metadata tags are not enriched with `checkSender.checkTags`.
		// `checkSender.checkTags` are added for metrics, service checks, events only.
		// Note that we don't add some extra tags like `service` tag that might be present in `checkSender.checkTags`.
		deviceMetadataTags := append(utils.CopyStrings(tags), d.config.InstanceTags...)
		deviceMetadataTags = append(deviceMetadataTags, utils.GetAgentVersionTag())

		deviceDiagnosis := d.diagnoses.Report()

		d.sender.ReportNetworkDeviceMetadata(d.config, values, deviceMetadataTags, collectionTime, deviceStatus, pingStatus, deviceDiagnosis)
	}

	d.submitTelemetryMetrics(startTime, metricTags)
	d.setDeviceHostExternalTags()
	d.interfaceBandwidthState.RemoveExpiredBandwidthUsageRates(startTime.UnixNano())

	return checkErr
}

func (d *DeviceCheck) setDeviceHostExternalTags() {
	deviceHostname, err := d.GetDeviceHostname()
	if deviceHostname == "" || err != nil {
		return
	}
	agentTags := d.buildExternalTags()
	log.Debugf("Set external tags for device host, host=`%s`, agentTags=`%v`", deviceHostname, agentTags)
	externalhost.SetExternalTags(deviceHostname, common.SnmpExternalTagsSourceType, agentTags)
}

func (d *DeviceCheck) buildExternalTags() []string {
	agentTags := configUtils.GetConfiguredTags(d.agentConfig, false)
	if haagenthelpers.IsEnabled(d.agentConfig) {
		agentTags = append(agentTags, haagenthelpers.GetHaAgentTags(d.agentConfig)...)
	}
	return agentTags
}

func (d *DeviceCheck) getValuesAndTags() (bool, []string, *valuestore.ResultValueStore, error) {
	var deviceReachable bool
	var checkErrors []string
	var tags []string

	// Create connection
	connErr := d.session.Connect()
	if connErr != nil {
		d.diagnoses.Add("error", "SNMP_FAILED_TO_OPEN_CONNECTION", "Agent failed to open connection.")
		return false, tags, nil, fmt.Errorf("snmp connection error: %s", connErr)
	}
	defer func() {
		err := d.session.Close()
		if err != nil {
			d.sessionCloseErrorCount.Inc()
			log.Warnf("failed to close session (count: %d): %v", d.sessionCloseErrorCount.Load(), err)
		}
	}()

	// Check if the device is reachable
	getNextValue, err := d.session.GetNext([]string{coresnmp.DeviceReachableGetNextOid})
	if err != nil {
		deviceReachable = false
		d.diagnoses.Add("error", "SNMP_FAILED_TO_POLL_DEVICE", "Agent failed to poll this network device. Check the authentication method and ensure the agent can ping it.")
		checkErrors = append(checkErrors, fmt.Sprintf("check device reachable: failed: %s", err))
	} else {
		deviceReachable = true
		if log.ShouldLog(log.DebugLvl) {
			log.Debugf("check device reachable: success: %v", gosnmplib.PacketAsString(getNextValue))
		}
	}

	err = d.detectMetricsToMonitor(d.session)
	if err != nil {
		d.diagnoses.Add("error", "SNMP_FAILED_TO_DETECT_PROFILE", "Agent failed to detect a profile for this network device.")
		checkErrors = append(checkErrors, fmt.Sprintf("failed to autodetect profile: %s", err))
	}

	tags = append(tags, d.config.ProfileTags...)

	valuesStore, err := fetch.Fetch(d.session, d.config)
	if log.ShouldLog(log.DebugLvl) {
		log.Debugf("fetched values: %v", valuestore.ResultValueStoreAsString(valuesStore))
	}

	if err != nil {
		checkErrors = append(checkErrors, fmt.Sprintf("failed to fetch values: %s", err))
	} else {
		tags = append(tags, d.sender.GetCheckInstanceMetricTags(d.config.MetricTags, valuesStore)...)
	}

	var joinedError error
	if len(checkErrors) > 0 {
		joinedError = errors.New(strings.Join(checkErrors, "; "))
	}
	return deviceReachable, tags, valuesStore, joinedError
}

func (d *DeviceCheck) detectMetricsToMonitor(sess session.Session) error {
	if d.config.AutodetectProfile {
		// detect using sysObjectID
		sysObjectID, err := session.FetchSysObjectID(sess)
		if err != nil {
			return fmt.Errorf("failed to fetch sysobjectid: %s", err)
		}
		profile, err := d.config.ProfileProvider.GetProfileNameForSysObjectID(sysObjectID)
		if err != nil {
			return fmt.Errorf("failed to get profile sys object id for `%s`: %s", sysObjectID, err)
		}
		if profile != d.config.ProfileName {
			log.Debugf("detected profile change: %s -> %s", d.config.ProfileName, profile)
			err = d.config.SetProfile(profile)
			if err != nil {
				// Should not happen since the profile is one of those we matched in GetProfileNameForSysObjectID
				return fmt.Errorf("failed to refresh with profile `%s` detected using sysObjectID `%s`: %s", profile, sysObjectID, err)
			}
		}
	}
	return nil
}

func (d *DeviceCheck) submitTelemetryMetrics(startTime time.Time, tags []string) {
	newTags := append(utils.CopyStrings(tags), snmpLoaderTag, utils.GetAgentVersionTag())

	d.sender.Gauge("snmp.devices_monitored", float64(1), newTags)

	// SNMP Performance metrics
	d.sender.MonotonicCount("datadog.snmp.check_interval", time.Duration(startTime.UnixNano()).Seconds(), newTags)
	d.sender.Gauge("datadog.snmp.check_duration", time.Since(startTime).Seconds(), newTags)
	d.sender.Gauge("datadog.snmp.submitted_metrics", float64(d.sender.GetSubmittedMetrics()), newTags)
}

// GetDiagnoses collects diagnoses for diagnose CLI
func (d *DeviceCheck) GetDiagnoses() []diagnosis.Diagnosis {
	return d.diagnoses.ReportAsAgentDiagnoses()
}

// createPinger creates a pinger using the passed configuration
func createPinger(cfg pinger.Config) (pinger.Pinger, error) {
	// if OS is Windows or Mac, we should override UseRawSocket
	if runtime.GOOS == "windows" {
		cfg.UseRawSocket = true
	} else if runtime.GOOS == "darwin" {
		cfg.UseRawSocket = false
	}
	return pinger.New(cfg)
}

func (d *DeviceCheck) submitPingMetrics(pingResult *pinger.Result, tags []string) {
	d.sender.Gauge(pingAvgRttMetric, float64(pingResult.AvgRtt/time.Millisecond), tags)
	d.sender.Gauge(pingReachableMetric, utils.BoolToFloat64(pingResult.CanConnect), tags)
	d.sender.Gauge(pingUnreachableMetric, utils.BoolToFloat64(!pingResult.CanConnect), tags)
	d.sender.Gauge(pingPacketLoss, pingResult.PacketLoss, tags)
}

func (d *DeviceCheck) readTagsFromCache() {
	cacheValue, err := persistentcache.Read(d.cacheKey)
	if err != nil {
		log.Errorf("couldn't read cache for %s: %s", d.cacheKey, err)
	}
	if cacheValue == "" {
		d.savedDynamicTags = []string{}
		return
	}
	var tags []string
	if err = json.Unmarshal([]byte(cacheValue), &tags); err != nil {
		log.Errorf("couldn't unmarshal cache for %s: %s", d.cacheKey, err)
		return
	}
	d.savedDynamicTags = tags
}

func (d *DeviceCheck) writeTagsInCache() {
	cacheValue, err := json.Marshal(d.savedDynamicTags)
	if err != nil {
		log.Errorf("SNMP tags %s: Couldn't marshal cache: %s", d.config.Network, err)
		return
	}

	if err = persistentcache.Write(d.cacheKey, string(cacheValue)); err != nil {
		log.Errorf("SNMP tags %s: Couldn't write cache: %s", d.config.Network, err)
	}
}
