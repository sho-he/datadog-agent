// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.2
// 	protoc        v5.29.1
// source: datadog/api/v1/api.proto

package core

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_datadog_api_v1_api_proto protoreflect.FileDescriptor

var file_datadog_api_v1_api_proto_rawDesc = []byte{
	0x0a, 0x18, 0x64, 0x61, 0x74, 0x61, 0x64, 0x6f, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31,
	0x2f, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x64, 0x61, 0x74, 0x61,
	0x64, 0x6f, 0x67, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x64, 0x61, 0x74, 0x61,
	0x64, 0x6f, 0x67, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x25, 0x64, 0x61, 0x74, 0x61, 0x64, 0x6f,
	0x67, 0x2f, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2f, 0x72, 0x65,
	0x6d, 0x6f, 0x74, 0x65, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x27, 0x64, 0x61, 0x74, 0x61, 0x64, 0x6f, 0x67, 0x2f, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x27, 0x64, 0x61, 0x74, 0x61, 0x64, 0x6f,
	0x67, 0x2f, 0x77, 0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x6d, 0x65, 0x74, 0x61, 0x2f, 0x77,
	0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x29, 0x64, 0x61, 0x74, 0x61, 0x64, 0x6f, 0x67, 0x2f, 0x61, 0x75, 0x74, 0x6f, 0x64,
	0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x2f, 0x61, 0x75, 0x74, 0x6f, 0x64, 0x69, 0x73,
	0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74,
	0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x71, 0x0a, 0x05, 0x41, 0x67, 0x65, 0x6e, 0x74,
	0x12, 0x68, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x48, 0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x21, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x64, 0x6f, 0x67, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e,
	0x76, 0x31, 0x2e, 0x48, 0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x64, 0x6f, 0x67, 0x2e, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x48, 0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x22, 0x15, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0f, 0x12, 0x0d, 0x2f, 0x76, 0x31,
	0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x68, 0x6f, 0x73, 0x74, 0x32, 0xd5, 0x0f, 0x0a, 0x0b, 0x41,
	0x67, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x63, 0x75, 0x72, 0x65, 0x12, 0x8f, 0x01, 0x0a, 0x14, 0x54,
	0x61, 0x67, 0x67, 0x65, 0x72, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x45, 0x6e, 0x74, 0x69, 0x74,
	0x69, 0x65, 0x73, 0x12, 0x23, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x64, 0x6f, 0x67, 0x2e, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x54, 0x61, 0x67,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x64,
	0x6f, 0x67, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x54, 0x61, 0x67, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x2a,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x24, 0x3a, 0x01, 0x2a, 0x22, 0x1f, 0x2f, 0x76, 0x31, 0x2f, 0x67,
	0x72, 0x70, 0x63, 0x2f, 0x74, 0x61, 0x67, 0x67, 0x65, 0x72, 0x2f, 0x73, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x5f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x30, 0x01, 0x12, 0xe5, 0x01, 0x0a,
	0x27, 0x54, 0x61, 0x67, 0x67, 0x65, 0x72, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x43,
	0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x49, 0x44, 0x46, 0x72, 0x6f, 0x6d, 0x4f, 0x72,
	0x69, 0x67, 0x69, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x3a, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x64,
	0x6f, 0x67, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x6e, 0x65,
	0x72, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x49, 0x44, 0x46,
	0x72, 0x6f, 0x6d, 0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x3b, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x64, 0x6f, 0x67, 0x2e, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65,
	0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x49, 0x44, 0x46, 0x72, 0x6f, 0x6d, 0x4f,
	0x72, 0x69, 0x67, 0x69, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x41, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x3b, 0x3a, 0x01, 0x2a, 0x22, 0x36, 0x2f, 0x76,
	0x31, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x74, 0x61, 0x67, 0x67, 0x65, 0x72, 0x2f, 0x67, 0x65,
	0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x5f, 0x66, 0x72, 0x6f, 0x6d, 0x5f, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x5f,
	0x69, 0x6e, 0x66, 0x6f, 0x12, 0x89, 0x01, 0x0a, 0x11, 0x54, 0x61, 0x67, 0x67, 0x65, 0x72, 0x46,
	0x65, 0x74, 0x63, 0x68, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x24, 0x2e, 0x64, 0x61, 0x74,
	0x61, 0x64, 0x6f, 0x67, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x65,
	0x74, 0x63, 0x68, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x25, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x64, 0x6f, 0x67, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x2e, 0x76, 0x31, 0x2e, 0x46, 0x65, 0x74, 0x63, 0x68, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x27, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x21, 0x3a,
	0x01, 0x2a, 0x22, 0x1c, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x74, 0x61, 0x67,
	0x67, 0x65, 0x72, 0x2f, 0x66, 0x65, 0x74, 0x63, 0x68, 0x5f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x12, 0x9b, 0x01, 0x0a, 0x17, 0x44, 0x6f, 0x67, 0x73, 0x74, 0x61, 0x74, 0x73, 0x64, 0x43, 0x61,
	0x70, 0x74, 0x75, 0x72, 0x65, 0x54, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x12, 0x27, 0x2e, 0x64,
	0x61, 0x74, 0x61, 0x64, 0x6f, 0x67, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x2e,
	0x43, 0x61, 0x70, 0x74, 0x75, 0x72, 0x65, 0x54, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x64, 0x6f, 0x67, 0x2e,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x61, 0x70, 0x74, 0x75, 0x72, 0x65,
	0x54, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x2d, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x27, 0x3a, 0x01, 0x2a, 0x22, 0x22, 0x2f, 0x76, 0x31, 0x2f,
	0x67, 0x72, 0x70, 0x63, 0x2f, 0x64, 0x6f, 0x67, 0x73, 0x74, 0x61, 0x74, 0x73, 0x64, 0x2f, 0x63,
	0x61, 0x70, 0x74, 0x75, 0x72, 0x65, 0x2f, 0x74, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x12, 0x8c,
	0x01, 0x0a, 0x17, 0x44, 0x6f, 0x67, 0x73, 0x74, 0x61, 0x74, 0x73, 0x64, 0x53, 0x65, 0x74, 0x54,
	0x61, 0x67, 0x67, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x1d, 0x2e, 0x64, 0x61, 0x74,
	0x61, 0x64, 0x6f, 0x67, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x61,
	0x67, 0x67, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x65, 0x1a, 0x25, 0x2e, 0x64, 0x61, 0x74, 0x61,
	0x64, 0x6f, 0x67, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x61, 0x67,
	0x67, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x2b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x25, 0x3a, 0x01, 0x2a, 0x22, 0x20, 0x2f, 0x76, 0x31,
	0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x64, 0x6f, 0x67, 0x73, 0x74, 0x61, 0x74, 0x73, 0x64, 0x2f,
	0x63, 0x61, 0x70, 0x74, 0x75, 0x72, 0x65, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x8f, 0x01,
	0x0a, 0x10, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x73, 0x12, 0x27, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x64, 0x6f, 0x67, 0x2e, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x64, 0x61,
	0x74, 0x61, 0x64, 0x6f, 0x67, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x43, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x28, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x22, 0x3a, 0x01, 0x2a,
	0x22, 0x1d, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x72, 0x65, 0x6d, 0x6f, 0x74,
	0x65, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x12,
	0x78, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x53, 0x74, 0x61, 0x74,
	0x65, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x26, 0x2e, 0x64, 0x61, 0x74, 0x61,
	0x64, 0x6f, 0x67, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x74,
	0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x26, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x20, 0x3a, 0x01, 0x2a, 0x22, 0x1b, 0x2f, 0x76,
	0x31, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x94, 0x01, 0x0a, 0x12, 0x43, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x48, 0x41,
	0x12, 0x27, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x64, 0x6f, 0x67, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x64, 0x61, 0x74, 0x61,
	0x64, 0x6f, 0x67, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x2b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x25, 0x3a, 0x01, 0x2a, 0x22, 0x20,
	0x2f, 0x76, 0x31, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x5f, 0x68, 0x61,
	0x12, 0x7d, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x53, 0x74, 0x61,
	0x74, 0x65, 0x48, 0x41, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x26, 0x2e, 0x64,
	0x61, 0x74, 0x61, 0x64, 0x6f, 0x67, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x47, 0x65,
	0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x29, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x23, 0x3a, 0x01, 0x2a, 0x22,
	0x1e, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x5f, 0x68, 0x61, 0x12,
	0xb3, 0x01, 0x0a, 0x1a, 0x57, 0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x6d, 0x65, 0x74, 0x61,
	0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x12, 0x2f,
	0x2e, 0x64, 0x61, 0x74, 0x61, 0x64, 0x6f, 0x67, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61,
	0x64, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x6d, 0x65,
	0x74, 0x61, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x30, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x64, 0x6f, 0x67, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x6c, 0x6f,
	0x61, 0x64, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x6d,
	0x65, 0x74, 0x61, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x30, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2a, 0x3a, 0x01, 0x2a, 0x22, 0x25, 0x2f, 0x76,
	0x31, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x77, 0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x6d,
	0x65, 0x74, 0x61, 0x2f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x5f, 0x65, 0x6e, 0x74, 0x69, 0x74,
	0x69, 0x65, 0x73, 0x30, 0x01, 0x12, 0xaf, 0x01, 0x0a, 0x13, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74,
	0x65, 0x72, 0x52, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x12, 0x2f, 0x2e,
	0x64, 0x61, 0x74, 0x61, 0x64, 0x6f, 0x67, 0x2e, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x61, 0x67,
	0x65, 0x6e, 0x74, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x6d, 0x6f,
	0x74, 0x65, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x30,
	0x2e, 0x64, 0x61, 0x74, 0x61, 0x64, 0x6f, 0x67, 0x2e, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x61,
	0x67, 0x65, 0x6e, 0x74, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x6d,
	0x6f, 0x74, 0x65, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x35, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2f, 0x3a, 0x01, 0x2a, 0x22, 0x2a, 0x2f, 0x76, 0x31,
	0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x61, 0x67, 0x65, 0x6e,
	0x74, 0x2f, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x72, 0x65, 0x6d, 0x6f, 0x74,
	0x65, 0x5f, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x12, 0x9b, 0x01, 0x0a, 0x19, 0x41, 0x75, 0x74, 0x6f,
	0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x32, 0x2e,
	0x64, 0x61, 0x74, 0x61, 0x64, 0x6f, 0x67, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x64, 0x69, 0x73, 0x63,
	0x6f, 0x76, 0x65, 0x72, 0x79, 0x2e, 0x41, 0x75, 0x74, 0x6f, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76,
	0x65, 0x72, 0x79, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x30, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2a, 0x3a, 0x01, 0x2a, 0x22, 0x25, 0x2f, 0x76,
	0x31, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x61, 0x75, 0x74, 0x6f, 0x64, 0x69, 0x73, 0x63, 0x6f,
	0x76, 0x65, 0x72, 0x79, 0x2f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x5f, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x73, 0x30, 0x01, 0x12, 0x6b, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x48, 0x6f, 0x73, 0x74,
	0x54, 0x61, 0x67, 0x73, 0x12, 0x20, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x64, 0x6f, 0x67, 0x2e, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x48, 0x6f, 0x73, 0x74, 0x54, 0x61, 0x67, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x64, 0x6f, 0x67,
	0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x48, 0x6f, 0x73, 0x74, 0x54, 0x61,
	0x67, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x1a, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x14, 0x12, 0x12,
	0x2f, 0x76, 0x31, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x68, 0x6f, 0x73, 0x74, 0x5f, 0x74, 0x61,
	0x67, 0x73, 0x32, 0xe6, 0x01, 0x0a, 0x0b, 0x52, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x41, 0x67, 0x65,
	0x6e, 0x74, 0x12, 0x6f, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x44,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x2c, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x64, 0x6f, 0x67,
	0x2e, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x47, 0x65, 0x74,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x2d, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x64, 0x6f, 0x67, 0x2e, 0x72,
	0x65, 0x6d, 0x6f, 0x74, 0x65, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x66, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x46, 0x6c, 0x61, 0x72, 0x65, 0x46,
	0x69, 0x6c, 0x65, 0x73, 0x12, 0x29, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x64, 0x6f, 0x67, 0x2e, 0x72,
	0x65, 0x6d, 0x6f, 0x74, 0x65, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x46, 0x6c,
	0x61, 0x72, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x2a, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x64, 0x6f, 0x67, 0x2e, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65,
	0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x46, 0x6c, 0x61, 0x72, 0x65, 0x46, 0x69,
	0x6c, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x15, 0x5a, 0x13, 0x70,
	0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x62, 0x67, 0x6f, 0x2f, 0x63, 0x6f,
	0x72, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_datadog_api_v1_api_proto_goTypes = []any{
	(*HostnameRequest)(nil),                           // 0: datadog.model.v1.HostnameRequest
	(*StreamTagsRequest)(nil),                         // 1: datadog.model.v1.StreamTagsRequest
	(*GenerateContainerIDFromOriginInfoRequest)(nil),  // 2: datadog.model.v1.GenerateContainerIDFromOriginInfoRequest
	(*FetchEntityRequest)(nil),                        // 3: datadog.model.v1.FetchEntityRequest
	(*CaptureTriggerRequest)(nil),                     // 4: datadog.model.v1.CaptureTriggerRequest
	(*TaggerState)(nil),                               // 5: datadog.model.v1.TaggerState
	(*ClientGetConfigsRequest)(nil),                   // 6: datadog.config.ClientGetConfigsRequest
	(*emptypb.Empty)(nil),                             // 7: google.protobuf.Empty
	(*WorkloadmetaStreamRequest)(nil),                 // 8: datadog.workloadmeta.WorkloadmetaStreamRequest
	(*RegisterRemoteAgentRequest)(nil),                // 9: datadog.remoteagent.RegisterRemoteAgentRequest
	(*HostTagRequest)(nil),                            // 10: datadog.model.v1.HostTagRequest
	(*GetStatusDetailsRequest)(nil),                   // 11: datadog.remoteagent.GetStatusDetailsRequest
	(*GetFlareFilesRequest)(nil),                      // 12: datadog.remoteagent.GetFlareFilesRequest
	(*HostnameReply)(nil),                             // 13: datadog.model.v1.HostnameReply
	(*StreamTagsResponse)(nil),                        // 14: datadog.model.v1.StreamTagsResponse
	(*GenerateContainerIDFromOriginInfoResponse)(nil), // 15: datadog.model.v1.GenerateContainerIDFromOriginInfoResponse
	(*FetchEntityResponse)(nil),                       // 16: datadog.model.v1.FetchEntityResponse
	(*CaptureTriggerResponse)(nil),                    // 17: datadog.model.v1.CaptureTriggerResponse
	(*TaggerStateResponse)(nil),                       // 18: datadog.model.v1.TaggerStateResponse
	(*ClientGetConfigsResponse)(nil),                  // 19: datadog.config.ClientGetConfigsResponse
	(*GetStateConfigResponse)(nil),                    // 20: datadog.config.GetStateConfigResponse
	(*WorkloadmetaStreamResponse)(nil),                // 21: datadog.workloadmeta.WorkloadmetaStreamResponse
	(*RegisterRemoteAgentResponse)(nil),               // 22: datadog.remoteagent.RegisterRemoteAgentResponse
	(*AutodiscoveryStreamResponse)(nil),               // 23: datadog.autodiscovery.AutodiscoveryStreamResponse
	(*HostTagReply)(nil),                              // 24: datadog.model.v1.HostTagReply
	(*GetStatusDetailsResponse)(nil),                  // 25: datadog.remoteagent.GetStatusDetailsResponse
	(*GetFlareFilesResponse)(nil),                     // 26: datadog.remoteagent.GetFlareFilesResponse
}
var file_datadog_api_v1_api_proto_depIdxs = []int32{
	0,  // 0: datadog.api.v1.Agent.GetHostname:input_type -> datadog.model.v1.HostnameRequest
	1,  // 1: datadog.api.v1.AgentSecure.TaggerStreamEntities:input_type -> datadog.model.v1.StreamTagsRequest
	2,  // 2: datadog.api.v1.AgentSecure.TaggerGenerateContainerIDFromOriginInfo:input_type -> datadog.model.v1.GenerateContainerIDFromOriginInfoRequest
	3,  // 3: datadog.api.v1.AgentSecure.TaggerFetchEntity:input_type -> datadog.model.v1.FetchEntityRequest
	4,  // 4: datadog.api.v1.AgentSecure.DogstatsdCaptureTrigger:input_type -> datadog.model.v1.CaptureTriggerRequest
	5,  // 5: datadog.api.v1.AgentSecure.DogstatsdSetTaggerState:input_type -> datadog.model.v1.TaggerState
	6,  // 6: datadog.api.v1.AgentSecure.ClientGetConfigs:input_type -> datadog.config.ClientGetConfigsRequest
	7,  // 7: datadog.api.v1.AgentSecure.GetConfigState:input_type -> google.protobuf.Empty
	6,  // 8: datadog.api.v1.AgentSecure.ClientGetConfigsHA:input_type -> datadog.config.ClientGetConfigsRequest
	7,  // 9: datadog.api.v1.AgentSecure.GetConfigStateHA:input_type -> google.protobuf.Empty
	8,  // 10: datadog.api.v1.AgentSecure.WorkloadmetaStreamEntities:input_type -> datadog.workloadmeta.WorkloadmetaStreamRequest
	9,  // 11: datadog.api.v1.AgentSecure.RegisterRemoteAgent:input_type -> datadog.remoteagent.RegisterRemoteAgentRequest
	7,  // 12: datadog.api.v1.AgentSecure.AutodiscoveryStreamConfig:input_type -> google.protobuf.Empty
	10, // 13: datadog.api.v1.AgentSecure.GetHostTags:input_type -> datadog.model.v1.HostTagRequest
	11, // 14: datadog.api.v1.RemoteAgent.GetStatusDetails:input_type -> datadog.remoteagent.GetStatusDetailsRequest
	12, // 15: datadog.api.v1.RemoteAgent.GetFlareFiles:input_type -> datadog.remoteagent.GetFlareFilesRequest
	13, // 16: datadog.api.v1.Agent.GetHostname:output_type -> datadog.model.v1.HostnameReply
	14, // 17: datadog.api.v1.AgentSecure.TaggerStreamEntities:output_type -> datadog.model.v1.StreamTagsResponse
	15, // 18: datadog.api.v1.AgentSecure.TaggerGenerateContainerIDFromOriginInfo:output_type -> datadog.model.v1.GenerateContainerIDFromOriginInfoResponse
	16, // 19: datadog.api.v1.AgentSecure.TaggerFetchEntity:output_type -> datadog.model.v1.FetchEntityResponse
	17, // 20: datadog.api.v1.AgentSecure.DogstatsdCaptureTrigger:output_type -> datadog.model.v1.CaptureTriggerResponse
	18, // 21: datadog.api.v1.AgentSecure.DogstatsdSetTaggerState:output_type -> datadog.model.v1.TaggerStateResponse
	19, // 22: datadog.api.v1.AgentSecure.ClientGetConfigs:output_type -> datadog.config.ClientGetConfigsResponse
	20, // 23: datadog.api.v1.AgentSecure.GetConfigState:output_type -> datadog.config.GetStateConfigResponse
	19, // 24: datadog.api.v1.AgentSecure.ClientGetConfigsHA:output_type -> datadog.config.ClientGetConfigsResponse
	20, // 25: datadog.api.v1.AgentSecure.GetConfigStateHA:output_type -> datadog.config.GetStateConfigResponse
	21, // 26: datadog.api.v1.AgentSecure.WorkloadmetaStreamEntities:output_type -> datadog.workloadmeta.WorkloadmetaStreamResponse
	22, // 27: datadog.api.v1.AgentSecure.RegisterRemoteAgent:output_type -> datadog.remoteagent.RegisterRemoteAgentResponse
	23, // 28: datadog.api.v1.AgentSecure.AutodiscoveryStreamConfig:output_type -> datadog.autodiscovery.AutodiscoveryStreamResponse
	24, // 29: datadog.api.v1.AgentSecure.GetHostTags:output_type -> datadog.model.v1.HostTagReply
	25, // 30: datadog.api.v1.RemoteAgent.GetStatusDetails:output_type -> datadog.remoteagent.GetStatusDetailsResponse
	26, // 31: datadog.api.v1.RemoteAgent.GetFlareFiles:output_type -> datadog.remoteagent.GetFlareFilesResponse
	16, // [16:32] is the sub-list for method output_type
	0,  // [0:16] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_datadog_api_v1_api_proto_init() }
func file_datadog_api_v1_api_proto_init() {
	if File_datadog_api_v1_api_proto != nil {
		return
	}
	file_datadog_model_v1_model_proto_init()
	file_datadog_remoteagent_remoteagent_proto_init()
	file_datadog_remoteconfig_remoteconfig_proto_init()
	file_datadog_workloadmeta_workloadmeta_proto_init()
	file_datadog_autodiscovery_autodiscovery_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_datadog_api_v1_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   3,
		},
		GoTypes:           file_datadog_api_v1_api_proto_goTypes,
		DependencyIndexes: file_datadog_api_v1_api_proto_depIdxs,
	}.Build()
	File_datadog_api_v1_api_proto = out.File
	file_datadog_api_v1_api_proto_rawDesc = nil
	file_datadog_api_v1_api_proto_goTypes = nil
	file_datadog_api_v1_api_proto_depIdxs = nil
}
