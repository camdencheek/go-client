// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.21.12
// source: points_service.proto

package go_client

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/structpb"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_points_service_proto protoreflect.FileDescriptor

var file_points_service_proto_rawDesc = []byte{
	0x0a, 0x14, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x71, 0x64, 0x72, 0x61, 0x6e, 0x74, 0x1a, 0x0c,
	0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74,
	0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xf9, 0x0a, 0x0a, 0x06, 0x50,
	0x6f, 0x69, 0x6e, 0x74, 0x73, 0x12, 0x41, 0x0a, 0x06, 0x55, 0x70, 0x73, 0x65, 0x72, 0x74, 0x12,
	0x14, 0x2e, 0x71, 0x64, 0x72, 0x61, 0x6e, 0x74, 0x2e, 0x55, 0x70, 0x73, 0x65, 0x72, 0x74, 0x50,
	0x6f, 0x69, 0x6e, 0x74, 0x73, 0x1a, 0x1f, 0x2e, 0x71, 0x64, 0x72, 0x61, 0x6e, 0x74, 0x2e, 0x50,
	0x6f, 0x69, 0x6e, 0x74, 0x73, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x41, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x12, 0x14, 0x2e, 0x71, 0x64, 0x72, 0x61, 0x6e, 0x74, 0x2e, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x1a, 0x1f, 0x2e, 0x71, 0x64, 0x72, 0x61, 0x6e,
	0x74, 0x2e, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x2f, 0x0a, 0x03, 0x47,
	0x65, 0x74, 0x12, 0x11, 0x2e, 0x71, 0x64, 0x72, 0x61, 0x6e, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x50,
	0x6f, 0x69, 0x6e, 0x74, 0x73, 0x1a, 0x13, 0x2e, 0x71, 0x64, 0x72, 0x61, 0x6e, 0x74, 0x2e, 0x47,
	0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4e, 0x0a, 0x0d,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x12, 0x1a, 0x2e,
	0x71, 0x64, 0x72, 0x61, 0x6e, 0x74, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x69,
	0x6e, 0x74, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x1a, 0x1f, 0x2e, 0x71, 0x64, 0x72, 0x61,
	0x6e, 0x74, 0x2e, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4e, 0x0a, 0x0d,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x12, 0x1a, 0x2e,
	0x71, 0x64, 0x72, 0x61, 0x6e, 0x74, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x6f, 0x69,
	0x6e, 0x74, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x1a, 0x1f, 0x2e, 0x71, 0x64, 0x72, 0x61,
	0x6e, 0x74, 0x2e, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x49, 0x0a, 0x0a,
	0x53, 0x65, 0x74, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x18, 0x2e, 0x71, 0x64, 0x72,
	0x61, 0x6e, 0x74, 0x2e, 0x53, 0x65, 0x74, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x50, 0x6f,
	0x69, 0x6e, 0x74, 0x73, 0x1a, 0x1f, 0x2e, 0x71, 0x64, 0x72, 0x61, 0x6e, 0x74, 0x2e, 0x50, 0x6f,
	0x69, 0x6e, 0x74, 0x73, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4f, 0x0a, 0x10, 0x4f, 0x76, 0x65, 0x72, 0x77,
	0x72, 0x69, 0x74, 0x65, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x18, 0x2e, 0x71, 0x64,
	0x72, 0x61, 0x6e, 0x74, 0x2e, 0x53, 0x65, 0x74, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x50,
	0x6f, 0x69, 0x6e, 0x74, 0x73, 0x1a, 0x1f, 0x2e, 0x71, 0x64, 0x72, 0x61, 0x6e, 0x74, 0x2e, 0x50,
	0x6f, 0x69, 0x6e, 0x74, 0x73, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4f, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x1b, 0x2e, 0x71, 0x64, 0x72, 0x61,
	0x6e, 0x74, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64,
	0x50, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x1a, 0x1f, 0x2e, 0x71, 0x64, 0x72, 0x61, 0x6e, 0x74, 0x2e,
	0x50, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4d, 0x0a, 0x0c, 0x43, 0x6c, 0x65,
	0x61, 0x72, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x1a, 0x2e, 0x71, 0x64, 0x72, 0x61,
	0x6e, 0x74, 0x2e, 0x43, 0x6c, 0x65, 0x61, 0x72, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x50,
	0x6f, 0x69, 0x6e, 0x74, 0x73, 0x1a, 0x1f, 0x2e, 0x71, 0x64, 0x72, 0x61, 0x6e, 0x74, 0x2e, 0x50,
	0x6f, 0x69, 0x6e, 0x74, 0x73, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x59, 0x0a, 0x10, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x22, 0x2e, 0x71,
	0x64, 0x72, 0x61, 0x6e, 0x74, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x46, 0x69, 0x65, 0x6c,
	0x64, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x1a, 0x1f, 0x2e, 0x71, 0x64, 0x72, 0x61, 0x6e, 0x74, 0x2e, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x73,
	0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x59, 0x0a, 0x10, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x46, 0x69, 0x65,
	0x6c, 0x64, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x22, 0x2e, 0x71, 0x64, 0x72, 0x61, 0x6e, 0x74,
	0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x49, 0x6e, 0x64, 0x65,
	0x78, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x1f, 0x2e, 0x71, 0x64,
	0x72, 0x61, 0x6e, 0x74, 0x2e, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x4f, 0x70, 0x65, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x38,
	0x0a, 0x06, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x12, 0x14, 0x2e, 0x71, 0x64, 0x72, 0x61, 0x6e,
	0x74, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x1a, 0x16,
	0x2e, 0x71, 0x64, 0x72, 0x61, 0x6e, 0x74, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x47, 0x0a, 0x0b, 0x53, 0x65, 0x61, 0x72,
	0x63, 0x68, 0x42, 0x61, 0x74, 0x63, 0x68, 0x12, 0x19, 0x2e, 0x71, 0x64, 0x72, 0x61, 0x6e, 0x74,
	0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x42, 0x61, 0x74, 0x63, 0x68, 0x50, 0x6f, 0x69, 0x6e,
	0x74, 0x73, 0x1a, 0x1b, 0x2e, 0x71, 0x64, 0x72, 0x61, 0x6e, 0x74, 0x2e, 0x53, 0x65, 0x61, 0x72,
	0x63, 0x68, 0x42, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x49, 0x0a, 0x0c, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x73, 0x12, 0x19, 0x2e, 0x71, 0x64, 0x72, 0x61, 0x6e, 0x74, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63,
	0x68, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x1a, 0x1c, 0x2e, 0x71,
	0x64, 0x72, 0x61, 0x6e, 0x74, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x38, 0x0a, 0x06,
	0x53, 0x63, 0x72, 0x6f, 0x6c, 0x6c, 0x12, 0x14, 0x2e, 0x71, 0x64, 0x72, 0x61, 0x6e, 0x74, 0x2e,
	0x53, 0x63, 0x72, 0x6f, 0x6c, 0x6c, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x1a, 0x16, 0x2e, 0x71,
	0x64, 0x72, 0x61, 0x6e, 0x74, 0x2e, 0x53, 0x63, 0x72, 0x6f, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x41, 0x0a, 0x09, 0x52, 0x65, 0x63, 0x6f, 0x6d, 0x6d,
	0x65, 0x6e, 0x64, 0x12, 0x17, 0x2e, 0x71, 0x64, 0x72, 0x61, 0x6e, 0x74, 0x2e, 0x52, 0x65, 0x63,
	0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x1a, 0x19, 0x2e, 0x71,
	0x64, 0x72, 0x61, 0x6e, 0x74, 0x2e, 0x52, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x50, 0x0a, 0x0e, 0x52, 0x65, 0x63,
	0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x42, 0x61, 0x74, 0x63, 0x68, 0x12, 0x1c, 0x2e, 0x71, 0x64,
	0x72, 0x61, 0x6e, 0x74, 0x2e, 0x52, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x42, 0x61,
	0x74, 0x63, 0x68, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x1a, 0x1e, 0x2e, 0x71, 0x64, 0x72, 0x61,
	0x6e, 0x74, 0x2e, 0x52, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x42, 0x61, 0x74, 0x63,
	0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x52, 0x0a, 0x0f, 0x52,
	0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x12, 0x1c,
	0x2e, 0x71, 0x64, 0x72, 0x61, 0x6e, 0x74, 0x2e, 0x52, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e,
	0x64, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x1a, 0x1f, 0x2e, 0x71,
	0x64, 0x72, 0x61, 0x6e, 0x74, 0x2e, 0x52, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x47,
	0x72, 0x6f, 0x75, 0x70, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x35, 0x0a, 0x05, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x13, 0x2e, 0x71, 0x64, 0x72, 0x61, 0x6e,
	0x74, 0x2e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x1a, 0x15, 0x2e,
	0x71, 0x64, 0x72, 0x61, 0x6e, 0x74, 0x2e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_points_service_proto_goTypes = []interface{}{
	(*UpsertPoints)(nil),               // 0: qdrant.UpsertPoints
	(*DeletePoints)(nil),               // 1: qdrant.DeletePoints
	(*GetPoints)(nil),                  // 2: qdrant.GetPoints
	(*UpdatePointVectors)(nil),         // 3: qdrant.UpdatePointVectors
	(*DeletePointVectors)(nil),         // 4: qdrant.DeletePointVectors
	(*SetPayloadPoints)(nil),           // 5: qdrant.SetPayloadPoints
	(*DeletePayloadPoints)(nil),        // 6: qdrant.DeletePayloadPoints
	(*ClearPayloadPoints)(nil),         // 7: qdrant.ClearPayloadPoints
	(*CreateFieldIndexCollection)(nil), // 8: qdrant.CreateFieldIndexCollection
	(*DeleteFieldIndexCollection)(nil), // 9: qdrant.DeleteFieldIndexCollection
	(*SearchPoints)(nil),               // 10: qdrant.SearchPoints
	(*SearchBatchPoints)(nil),          // 11: qdrant.SearchBatchPoints
	(*SearchPointGroups)(nil),          // 12: qdrant.SearchPointGroups
	(*ScrollPoints)(nil),               // 13: qdrant.ScrollPoints
	(*RecommendPoints)(nil),            // 14: qdrant.RecommendPoints
	(*RecommendBatchPoints)(nil),       // 15: qdrant.RecommendBatchPoints
	(*RecommendPointGroups)(nil),       // 16: qdrant.RecommendPointGroups
	(*CountPoints)(nil),                // 17: qdrant.CountPoints
	(*PointsOperationResponse)(nil),    // 18: qdrant.PointsOperationResponse
	(*GetResponse)(nil),                // 19: qdrant.GetResponse
	(*SearchResponse)(nil),             // 20: qdrant.SearchResponse
	(*SearchBatchResponse)(nil),        // 21: qdrant.SearchBatchResponse
	(*SearchGroupsResponse)(nil),       // 22: qdrant.SearchGroupsResponse
	(*ScrollResponse)(nil),             // 23: qdrant.ScrollResponse
	(*RecommendResponse)(nil),          // 24: qdrant.RecommendResponse
	(*RecommendBatchResponse)(nil),     // 25: qdrant.RecommendBatchResponse
	(*RecommendGroupsResponse)(nil),    // 26: qdrant.RecommendGroupsResponse
	(*CountResponse)(nil),              // 27: qdrant.CountResponse
}
var file_points_service_proto_depIdxs = []int32{
	0,  // 0: qdrant.Points.Upsert:input_type -> qdrant.UpsertPoints
	1,  // 1: qdrant.Points.Delete:input_type -> qdrant.DeletePoints
	2,  // 2: qdrant.Points.Get:input_type -> qdrant.GetPoints
	3,  // 3: qdrant.Points.UpdateVectors:input_type -> qdrant.UpdatePointVectors
	4,  // 4: qdrant.Points.DeleteVectors:input_type -> qdrant.DeletePointVectors
	5,  // 5: qdrant.Points.SetPayload:input_type -> qdrant.SetPayloadPoints
	5,  // 6: qdrant.Points.OverwritePayload:input_type -> qdrant.SetPayloadPoints
	6,  // 7: qdrant.Points.DeletePayload:input_type -> qdrant.DeletePayloadPoints
	7,  // 8: qdrant.Points.ClearPayload:input_type -> qdrant.ClearPayloadPoints
	8,  // 9: qdrant.Points.CreateFieldIndex:input_type -> qdrant.CreateFieldIndexCollection
	9,  // 10: qdrant.Points.DeleteFieldIndex:input_type -> qdrant.DeleteFieldIndexCollection
	10, // 11: qdrant.Points.Search:input_type -> qdrant.SearchPoints
	11, // 12: qdrant.Points.SearchBatch:input_type -> qdrant.SearchBatchPoints
	12, // 13: qdrant.Points.SearchGroups:input_type -> qdrant.SearchPointGroups
	13, // 14: qdrant.Points.Scroll:input_type -> qdrant.ScrollPoints
	14, // 15: qdrant.Points.Recommend:input_type -> qdrant.RecommendPoints
	15, // 16: qdrant.Points.RecommendBatch:input_type -> qdrant.RecommendBatchPoints
	16, // 17: qdrant.Points.RecommendGroups:input_type -> qdrant.RecommendPointGroups
	17, // 18: qdrant.Points.Count:input_type -> qdrant.CountPoints
	18, // 19: qdrant.Points.Upsert:output_type -> qdrant.PointsOperationResponse
	18, // 20: qdrant.Points.Delete:output_type -> qdrant.PointsOperationResponse
	19, // 21: qdrant.Points.Get:output_type -> qdrant.GetResponse
	18, // 22: qdrant.Points.UpdateVectors:output_type -> qdrant.PointsOperationResponse
	18, // 23: qdrant.Points.DeleteVectors:output_type -> qdrant.PointsOperationResponse
	18, // 24: qdrant.Points.SetPayload:output_type -> qdrant.PointsOperationResponse
	18, // 25: qdrant.Points.OverwritePayload:output_type -> qdrant.PointsOperationResponse
	18, // 26: qdrant.Points.DeletePayload:output_type -> qdrant.PointsOperationResponse
	18, // 27: qdrant.Points.ClearPayload:output_type -> qdrant.PointsOperationResponse
	18, // 28: qdrant.Points.CreateFieldIndex:output_type -> qdrant.PointsOperationResponse
	18, // 29: qdrant.Points.DeleteFieldIndex:output_type -> qdrant.PointsOperationResponse
	20, // 30: qdrant.Points.Search:output_type -> qdrant.SearchResponse
	21, // 31: qdrant.Points.SearchBatch:output_type -> qdrant.SearchBatchResponse
	22, // 32: qdrant.Points.SearchGroups:output_type -> qdrant.SearchGroupsResponse
	23, // 33: qdrant.Points.Scroll:output_type -> qdrant.ScrollResponse
	24, // 34: qdrant.Points.Recommend:output_type -> qdrant.RecommendResponse
	25, // 35: qdrant.Points.RecommendBatch:output_type -> qdrant.RecommendBatchResponse
	26, // 36: qdrant.Points.RecommendGroups:output_type -> qdrant.RecommendGroupsResponse
	27, // 37: qdrant.Points.Count:output_type -> qdrant.CountResponse
	19, // [19:38] is the sub-list for method output_type
	0,  // [0:19] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_points_service_proto_init() }
func file_points_service_proto_init() {
	if File_points_service_proto != nil {
		return
	}
	file_points_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_points_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_points_service_proto_goTypes,
		DependencyIndexes: file_points_service_proto_depIdxs,
	}.Build()
	File_points_service_proto = out.File
	file_points_service_proto_rawDesc = nil
	file_points_service_proto_goTypes = nil
	file_points_service_proto_depIdxs = nil
}
