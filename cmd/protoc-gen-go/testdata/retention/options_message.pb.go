// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Code generated by protoc-gen-go. DO NOT EDIT.
// source: cmd/protoc-gen-go/testdata/retention/options_message.proto

package retention

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

// Retention attributes set on fields nested within a message
type OptionsMessage struct {
	state                 protoimpl.MessageState `protogen:"open.v1"`
	PlainField            *int32                 `protobuf:"varint,1,opt,name=plain_field,json=plainField" json:"plain_field,omitempty"`
	RuntimeRetentionField *int32                 `protobuf:"varint,2,opt,name=runtime_retention_field,json=runtimeRetentionField" json:"runtime_retention_field,omitempty"`
	SourceRetentionField  *int32                 `protobuf:"varint,3,opt,name=source_retention_field,json=sourceRetentionField" json:"source_retention_field,omitempty"`
	unknownFields         protoimpl.UnknownFields
	sizeCache             protoimpl.SizeCache
}

func (x *OptionsMessage) Reset() {
	*x = OptionsMessage{}
	mi := &file_cmd_protoc_gen_go_testdata_retention_options_message_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *OptionsMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OptionsMessage) ProtoMessage() {}

func (x *OptionsMessage) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_protoc_gen_go_testdata_retention_options_message_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OptionsMessage.ProtoReflect.Descriptor instead.
func (*OptionsMessage) Descriptor() ([]byte, []int) {
	return file_cmd_protoc_gen_go_testdata_retention_options_message_proto_rawDescGZIP(), []int{0}
}

func (x *OptionsMessage) GetPlainField() int32 {
	if x != nil && x.PlainField != nil {
		return *x.PlainField
	}
	return 0
}

func (x *OptionsMessage) GetRuntimeRetentionField() int32 {
	if x != nil && x.RuntimeRetentionField != nil {
		return *x.RuntimeRetentionField
	}
	return 0
}

func (x *OptionsMessage) GetSourceRetentionField() int32 {
	if x != nil && x.SourceRetentionField != nil {
		return *x.SourceRetentionField
	}
	return 0
}

var file_cmd_protoc_gen_go_testdata_retention_options_message_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptorpb.FileOptions)(nil),
		ExtensionType: (*int32)(nil),
		Field:         511807920,
		Name:          "testretention.imported_plain_option",
		Tag:           "varint,511807920,opt,name=imported_plain_option",
		Filename:      "cmd/protoc-gen-go/testdata/retention/options_message.proto",
	},
	{
		ExtendedType:  (*descriptorpb.FileOptions)(nil),
		ExtensionType: (*int32)(nil),
		Field:         512484074,
		Name:          "testretention.imported_runtime_retention_option",
		Tag:           "varint,512484074,opt,name=imported_runtime_retention_option",
		Filename:      "cmd/protoc-gen-go/testdata/retention/options_message.proto",
	},
	{
		ExtendedType:  (*descriptorpb.FileOptions)(nil),
		ExtensionType: (*int32)(nil),
		Field:         512645287,
		Name:          "testretention.imported_source_retention_option",
		Tag:           "varint,512645287,opt,name=imported_source_retention_option",
		Filename:      "cmd/protoc-gen-go/testdata/retention/options_message.proto",
	},
	{
		ExtendedType:  (*descriptorpb.FileOptions)(nil),
		ExtensionType: (*OptionsMessage)(nil),
		Field:         504871168,
		Name:          "testretention.file_option",
		Tag:           "bytes,504871168,opt,name=file_option",
		Filename:      "cmd/protoc-gen-go/testdata/retention/options_message.proto",
	},
}

// Extension fields to descriptorpb.FileOptions.
var (
	// optional int32 imported_plain_option = 511807920;
	E_ImportedPlainOption = &file_cmd_protoc_gen_go_testdata_retention_options_message_proto_extTypes[0]
	// optional int32 imported_runtime_retention_option = 512484074;
	E_ImportedRuntimeRetentionOption = &file_cmd_protoc_gen_go_testdata_retention_options_message_proto_extTypes[1]
	// optional int32 imported_source_retention_option = 512645287;
	E_ImportedSourceRetentionOption = &file_cmd_protoc_gen_go_testdata_retention_options_message_proto_extTypes[2]
	// optional testretention.OptionsMessage file_option = 504871168;
	E_FileOption = &file_cmd_protoc_gen_go_testdata_retention_options_message_proto_extTypes[3]
)

var File_cmd_protoc_gen_go_testdata_retention_options_message_proto protoreflect.FileDescriptor

var file_cmd_protoc_gen_go_testdata_retention_options_message_proto_rawDesc = string([]byte{
	0x0a, 0x3a, 0x63, 0x6d, 0x64, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e,
	0x2d, 0x67, 0x6f, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x72, 0x65, 0x74,
	0x65, 0x6e, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x5f, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x74, 0x65,
	0x73, 0x74, 0x72, 0x65, 0x74, 0x65, 0x6e, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x20, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa9, 0x01,
	0x0a, 0x0e, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x6c, 0x61, 0x69, 0x6e, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x70, 0x6c, 0x61, 0x69, 0x6e, 0x46, 0x69, 0x65, 0x6c,
	0x64, 0x12, 0x3b, 0x0a, 0x17, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x72, 0x65, 0x74,
	0x65, 0x6e, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x42, 0x03, 0x88, 0x01, 0x01, 0x52, 0x15, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65,
	0x52, 0x65, 0x74, 0x65, 0x6e, 0x74, 0x69, 0x6f, 0x6e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x39,
	0x0a, 0x16, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x72, 0x65, 0x74, 0x65, 0x6e, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x42, 0x03,
	0x88, 0x01, 0x02, 0x52, 0x14, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x65, 0x74, 0x65, 0x6e,
	0x74, 0x69, 0x6f, 0x6e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x3a, 0x54, 0x0a, 0x15, 0x69, 0x6d, 0x70,
	0x6f, 0x72, 0x74, 0x65, 0x64, 0x5f, 0x70, 0x6c, 0x61, 0x69, 0x6e, 0x5f, 0x6f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x18, 0xb0, 0xa3, 0x86, 0xf4, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x13, 0x69, 0x6d, 0x70, 0x6f,
	0x72, 0x74, 0x65, 0x64, 0x50, 0x6c, 0x61, 0x69, 0x6e, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x3a,
	0x70, 0x0a, 0x21, 0x69, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x64, 0x5f, 0x72, 0x75, 0x6e, 0x74,
	0x69, 0x6d, 0x65, 0x5f, 0x72, 0x65, 0x74, 0x65, 0x6e, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x18, 0xea, 0xc5, 0xaf, 0xf4, 0x01, 0x20, 0x01, 0x28, 0x05, 0x42, 0x03, 0x88, 0x01,
	0x01, 0x52, 0x1e, 0x69, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x64, 0x52, 0x75, 0x6e, 0x74, 0x69,
	0x6d, 0x65, 0x52, 0x65, 0x74, 0x65, 0x6e, 0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x3a, 0x6e, 0x0a, 0x20, 0x69, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x64, 0x5f, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x5f, 0x72, 0x65, 0x74, 0x65, 0x6e, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x4f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x18, 0xa7, 0xb1, 0xb9, 0xf4, 0x01, 0x20, 0x01, 0x28, 0x05, 0x42, 0x03, 0x88,
	0x01, 0x02, 0x52, 0x1d, 0x69, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x64, 0x53, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x52, 0x65, 0x74, 0x65, 0x6e, 0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x3a, 0x60, 0x0a, 0x0b, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x80,
	0xf2, 0xde, 0xf0, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x72,
	0x65, 0x74, 0x65, 0x6e, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x0a, 0x66, 0x69, 0x6c, 0x65, 0x4f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x42, 0x41, 0x5a, 0x3f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f,
	0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x63, 0x6d, 0x64, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e,
	0x2d, 0x67, 0x6f, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x72, 0x65, 0x74,
	0x65, 0x6e, 0x74, 0x69, 0x6f, 0x6e,
})

var (
	file_cmd_protoc_gen_go_testdata_retention_options_message_proto_rawDescOnce sync.Once
	file_cmd_protoc_gen_go_testdata_retention_options_message_proto_rawDescData []byte
)

func file_cmd_protoc_gen_go_testdata_retention_options_message_proto_rawDescGZIP() []byte {
	file_cmd_protoc_gen_go_testdata_retention_options_message_proto_rawDescOnce.Do(func() {
		file_cmd_protoc_gen_go_testdata_retention_options_message_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_cmd_protoc_gen_go_testdata_retention_options_message_proto_rawDesc), len(file_cmd_protoc_gen_go_testdata_retention_options_message_proto_rawDesc)))
	})
	return file_cmd_protoc_gen_go_testdata_retention_options_message_proto_rawDescData
}

var file_cmd_protoc_gen_go_testdata_retention_options_message_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_cmd_protoc_gen_go_testdata_retention_options_message_proto_goTypes = []any{
	(*OptionsMessage)(nil),           // 0: testretention.OptionsMessage
	(*descriptorpb.FileOptions)(nil), // 1: google.protobuf.FileOptions
}
var file_cmd_protoc_gen_go_testdata_retention_options_message_proto_depIdxs = []int32{
	1, // 0: testretention.imported_plain_option:extendee -> google.protobuf.FileOptions
	1, // 1: testretention.imported_runtime_retention_option:extendee -> google.protobuf.FileOptions
	1, // 2: testretention.imported_source_retention_option:extendee -> google.protobuf.FileOptions
	1, // 3: testretention.file_option:extendee -> google.protobuf.FileOptions
	0, // 4: testretention.file_option:type_name -> testretention.OptionsMessage
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	4, // [4:5] is the sub-list for extension type_name
	0, // [0:4] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_cmd_protoc_gen_go_testdata_retention_options_message_proto_init() }
func file_cmd_protoc_gen_go_testdata_retention_options_message_proto_init() {
	if File_cmd_protoc_gen_go_testdata_retention_options_message_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_cmd_protoc_gen_go_testdata_retention_options_message_proto_rawDesc), len(file_cmd_protoc_gen_go_testdata_retention_options_message_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 4,
			NumServices:   0,
		},
		GoTypes:           file_cmd_protoc_gen_go_testdata_retention_options_message_proto_goTypes,
		DependencyIndexes: file_cmd_protoc_gen_go_testdata_retention_options_message_proto_depIdxs,
		MessageInfos:      file_cmd_protoc_gen_go_testdata_retention_options_message_proto_msgTypes,
		ExtensionInfos:    file_cmd_protoc_gen_go_testdata_retention_options_message_proto_extTypes,
	}.Build()
	File_cmd_protoc_gen_go_testdata_retention_options_message_proto = out.File
	file_cmd_protoc_gen_go_testdata_retention_options_message_proto_goTypes = nil
	file_cmd_protoc_gen_go_testdata_retention_options_message_proto_depIdxs = nil
}
