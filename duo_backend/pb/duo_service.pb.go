// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: duo_service.proto

package pb

import (
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

var File_duo_service_proto protoreflect.FileDescriptor

var file_duo_service_proto_rawDesc = []byte{
	0x0a, 0x11, 0x64, 0x75, 0x6f, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a, 0x13, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0a, 0x76, 0x6f,
	0x69, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xcc, 0x01, 0x0a, 0x0a, 0x44, 0x55, 0x4f,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x37, 0x0a, 0x08, 0x52, 0x65, 0x67, 0x69, 0x73,
	0x74, 0x65, 0x72, 0x12, 0x13, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x65,
	0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x41, 0x0a, 0x10, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x43, 0x68, 0x61, 0x6c, 0x6c,
	0x65, 0x6e, 0x67, 0x65, 0x12, 0x10, 0x2e, 0x70, 0x62, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x70, 0x62, 0x2e, 0x4c, 0x6f, 0x67, 0x69,
	0x6e, 0x43, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x22, 0x00, 0x12, 0x42, 0x0a, 0x0f, 0x73, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x43, 0x68, 0x61,
	0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x12, 0x1a, 0x2e, 0x70, 0x62, 0x2e, 0x4c, 0x6f, 0x67, 0x69,
	0x6e, 0x43, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x1a, 0x11, 0x2e, 0x70, 0x62, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x13, 0x5a, 0x11, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x75, 0x6f, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var file_duo_service_proto_goTypes = []interface{}{
	(*RegisterRequest)(nil),        // 0: pb.RegisterRequest
	(*LoginRequest)(nil),           // 1: pb.LoginRequest
	(*LoginChallengeResponse)(nil), // 2: pb.LoginChallengeResponse
	(*RegisterResponse)(nil),       // 3: pb.RegisterResponse
	(*LoginChallengeRequest)(nil),  // 4: pb.LoginChallengeRequest
	(*LoginResponse)(nil),          // 5: pb.LoginResponse
}
var file_duo_service_proto_depIdxs = []int32{
	0, // 0: pb.DUOService.Register:input_type -> pb.RegisterRequest
	1, // 1: pb.DUOService.requestChallenge:input_type -> pb.LoginRequest
	2, // 2: pb.DUOService.submitChallenge:input_type -> pb.LoginChallengeResponse
	3, // 3: pb.DUOService.Register:output_type -> pb.RegisterResponse
	4, // 4: pb.DUOService.requestChallenge:output_type -> pb.LoginChallengeRequest
	5, // 5: pb.DUOService.submitChallenge:output_type -> pb.LoginResponse
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_duo_service_proto_init() }
func file_duo_service_proto_init() {
	if File_duo_service_proto != nil {
		return
	}
	file_auth_messages_proto_init()
	file_void_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_duo_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_duo_service_proto_goTypes,
		DependencyIndexes: file_duo_service_proto_depIdxs,
	}.Build()
	File_duo_service_proto = out.File
	file_duo_service_proto_rawDesc = nil
	file_duo_service_proto_goTypes = nil
	file_duo_service_proto_depIdxs = nil
}
