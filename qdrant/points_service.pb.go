// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.12.4
// source: points_service.proto

package go_client

import (
	_ "github.com/golang/protobuf/ptypes/struct"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
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
	0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xba, 0x08, 0x0a, 0x06, 0x50,
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
	0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x49, 0x0a, 0x0a,
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
	0x00, 0x12, 0x38, 0x0a, 0x06, 0x53, 0x63, 0x72, 0x6f, 0x6c, 0x6c, 0x12, 0x14, 0x2e, 0x71, 0x64,
	0x72, 0x61, 0x6e, 0x74, 0x2e, 0x53, 0x63, 0x72, 0x6f, 0x6c, 0x6c, 0x50, 0x6f, 0x69, 0x6e, 0x74,
	0x73, 0x1a, 0x16, 0x2e, 0x71, 0x64, 0x72, 0x61, 0x6e, 0x74, 0x2e, 0x53, 0x63, 0x72, 0x6f, 0x6c,
	0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x41, 0x0a, 0x09, 0x52,
	0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x12, 0x17, 0x2e, 0x71, 0x64, 0x72, 0x61, 0x6e,
	0x74, 0x2e, 0x52, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x50, 0x6f, 0x69, 0x6e, 0x74,
	0x73, 0x1a, 0x19, 0x2e, 0x71, 0x64, 0x72, 0x61, 0x6e, 0x74, 0x2e, 0x52, 0x65, 0x63, 0x6f, 0x6d,
	0x6d, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x50,
	0x0a, 0x0e, 0x52, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x42, 0x61, 0x74, 0x63, 0x68,
	0x12, 0x1c, 0x2e, 0x71, 0x64, 0x72, 0x61, 0x6e, 0x74, 0x2e, 0x52, 0x65, 0x63, 0x6f, 0x6d, 0x6d,
	0x65, 0x6e, 0x64, 0x42, 0x61, 0x74, 0x63, 0x68, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x1a, 0x1e,
	0x2e, 0x71, 0x64, 0x72, 0x61, 0x6e, 0x74, 0x2e, 0x52, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e,
	0x64, 0x42, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x35, 0x0a, 0x05, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x13, 0x2e, 0x71, 0x64, 0x72, 0x61,
	0x6e, 0x74, 0x2e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x1a, 0x15,
	0x2e, 0x71, 0x64, 0x72, 0x61, 0x6e, 0x74, 0x2e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_points_service_proto_goTypes = []interface{}{
	(*UpsertPoints)(nil),               // 0: qdrant.UpsertPoints
	(*DeletePoints)(nil),               // 1: qdrant.DeletePoints
	(*GetPoints)(nil),                  // 2: qdrant.GetPoints
	(*SetPayloadPoints)(nil),           // 3: qdrant.SetPayloadPoints
	(*DeletePayloadPoints)(nil),        // 4: qdrant.DeletePayloadPoints
	(*ClearPayloadPoints)(nil),         // 5: qdrant.ClearPayloadPoints
	(*CreateFieldIndexCollection)(nil), // 6: qdrant.CreateFieldIndexCollection
	(*DeleteFieldIndexCollection)(nil), // 7: qdrant.DeleteFieldIndexCollection
	(*SearchPoints)(nil),               // 8: qdrant.SearchPoints
	(*SearchBatchPoints)(nil),          // 9: qdrant.SearchBatchPoints
	(*ScrollPoints)(nil),               // 10: qdrant.ScrollPoints
	(*RecommendPoints)(nil),            // 11: qdrant.RecommendPoints
	(*RecommendBatchPoints)(nil),       // 12: qdrant.RecommendBatchPoints
	(*CountPoints)(nil),                // 13: qdrant.CountPoints
	(*PointsOperationResponse)(nil),    // 14: qdrant.PointsOperationResponse
	(*GetResponse)(nil),                // 15: qdrant.GetResponse
	(*SearchResponse)(nil),             // 16: qdrant.SearchResponse
	(*SearchBatchResponse)(nil),        // 17: qdrant.SearchBatchResponse
	(*ScrollResponse)(nil),             // 18: qdrant.ScrollResponse
	(*RecommendResponse)(nil),          // 19: qdrant.RecommendResponse
	(*RecommendBatchResponse)(nil),     // 20: qdrant.RecommendBatchResponse
	(*CountResponse)(nil),              // 21: qdrant.CountResponse
}
var file_points_service_proto_depIdxs = []int32{
	0,  // 0: qdrant.Points.Upsert:input_type -> qdrant.UpsertPoints
	1,  // 1: qdrant.Points.Delete:input_type -> qdrant.DeletePoints
	2,  // 2: qdrant.Points.Get:input_type -> qdrant.GetPoints
	3,  // 3: qdrant.Points.SetPayload:input_type -> qdrant.SetPayloadPoints
	3,  // 4: qdrant.Points.OverwritePayload:input_type -> qdrant.SetPayloadPoints
	4,  // 5: qdrant.Points.DeletePayload:input_type -> qdrant.DeletePayloadPoints
	5,  // 6: qdrant.Points.ClearPayload:input_type -> qdrant.ClearPayloadPoints
	6,  // 7: qdrant.Points.CreateFieldIndex:input_type -> qdrant.CreateFieldIndexCollection
	7,  // 8: qdrant.Points.DeleteFieldIndex:input_type -> qdrant.DeleteFieldIndexCollection
	8,  // 9: qdrant.Points.Search:input_type -> qdrant.SearchPoints
	9,  // 10: qdrant.Points.SearchBatch:input_type -> qdrant.SearchBatchPoints
	10, // 11: qdrant.Points.Scroll:input_type -> qdrant.ScrollPoints
	11, // 12: qdrant.Points.Recommend:input_type -> qdrant.RecommendPoints
	12, // 13: qdrant.Points.RecommendBatch:input_type -> qdrant.RecommendBatchPoints
	13, // 14: qdrant.Points.Count:input_type -> qdrant.CountPoints
	14, // 15: qdrant.Points.Upsert:output_type -> qdrant.PointsOperationResponse
	14, // 16: qdrant.Points.Delete:output_type -> qdrant.PointsOperationResponse
	15, // 17: qdrant.Points.Get:output_type -> qdrant.GetResponse
	14, // 18: qdrant.Points.SetPayload:output_type -> qdrant.PointsOperationResponse
	14, // 19: qdrant.Points.OverwritePayload:output_type -> qdrant.PointsOperationResponse
	14, // 20: qdrant.Points.DeletePayload:output_type -> qdrant.PointsOperationResponse
	14, // 21: qdrant.Points.ClearPayload:output_type -> qdrant.PointsOperationResponse
	14, // 22: qdrant.Points.CreateFieldIndex:output_type -> qdrant.PointsOperationResponse
	14, // 23: qdrant.Points.DeleteFieldIndex:output_type -> qdrant.PointsOperationResponse
	16, // 24: qdrant.Points.Search:output_type -> qdrant.SearchResponse
	17, // 25: qdrant.Points.SearchBatch:output_type -> qdrant.SearchBatchResponse
	18, // 26: qdrant.Points.Scroll:output_type -> qdrant.ScrollResponse
	19, // 27: qdrant.Points.Recommend:output_type -> qdrant.RecommendResponse
	20, // 28: qdrant.Points.RecommendBatch:output_type -> qdrant.RecommendBatchResponse
	21, // 29: qdrant.Points.Count:output_type -> qdrant.CountResponse
	15, // [15:30] is the sub-list for method output_type
	0,  // [0:15] is the sub-list for method input_type
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
