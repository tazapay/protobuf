// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Code generated by protoc-gen-go. DO NOT EDIT.
// source: cmd/protoc-gen-go/testdata/imports/test_a_2/m3.proto

package test_a_2

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

type M3 struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *M3) Reset() {
	*x = M3{}
	mi := &file_cmd_protoc_gen_go_testdata_imports_test_a_2_m3_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *M3) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*M3) ProtoMessage() {}

func (x *M3) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_protoc_gen_go_testdata_imports_test_a_2_m3_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use M3.ProtoReflect.Descriptor instead.
func (*M3) Descriptor() ([]byte, []int) {
	return file_cmd_protoc_gen_go_testdata_imports_test_a_2_m3_proto_rawDescGZIP(), []int{0}
}

var File_cmd_protoc_gen_go_testdata_imports_test_a_2_m3_proto protoreflect.FileDescriptor

const file_cmd_protoc_gen_go_testdata_imports_test_a_2_m3_proto_rawDesc = "\n4cmd/protoc-gen-go/testdata/imports/test_a_2/m3.prototest.a\"\nM3BHZFgoogle.golang.org/protobuf/cmd/protoc-gen-go/testdata/imports/test_a_2bproto3"

var (
	file_cmd_protoc_gen_go_testdata_imports_test_a_2_m3_proto_rawDescOnce sync.Once
	file_cmd_protoc_gen_go_testdata_imports_test_a_2_m3_proto_rawDescData = file_cmd_protoc_gen_go_testdata_imports_test_a_2_m3_proto_rawDesc
)

func file_cmd_protoc_gen_go_testdata_imports_test_a_2_m3_proto_rawDescGZIP() []byte {
	file_cmd_protoc_gen_go_testdata_imports_test_a_2_m3_proto_rawDescOnce.Do(func() {
		file_cmd_protoc_gen_go_testdata_imports_test_a_2_m3_proto_rawDescData = string(protoimpl.X.CompressGZIP([]byte(file_cmd_protoc_gen_go_testdata_imports_test_a_2_m3_proto_rawDescData)))
	})
	return []byte(file_cmd_protoc_gen_go_testdata_imports_test_a_2_m3_proto_rawDescData)
}

var file_cmd_protoc_gen_go_testdata_imports_test_a_2_m3_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_cmd_protoc_gen_go_testdata_imports_test_a_2_m3_proto_goTypes = []any{
	(*M3)(nil), // 0: test.a.M3
}
var file_cmd_protoc_gen_go_testdata_imports_test_a_2_m3_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_cmd_protoc_gen_go_testdata_imports_test_a_2_m3_proto_init() }
func file_cmd_protoc_gen_go_testdata_imports_test_a_2_m3_proto_init() {
	if File_cmd_protoc_gen_go_testdata_imports_test_a_2_m3_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_cmd_protoc_gen_go_testdata_imports_test_a_2_m3_proto_rawDesc), len(file_cmd_protoc_gen_go_testdata_imports_test_a_2_m3_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cmd_protoc_gen_go_testdata_imports_test_a_2_m3_proto_goTypes,
		DependencyIndexes: file_cmd_protoc_gen_go_testdata_imports_test_a_2_m3_proto_depIdxs,
		MessageInfos:      file_cmd_protoc_gen_go_testdata_imports_test_a_2_m3_proto_msgTypes,
	}.Build()
	File_cmd_protoc_gen_go_testdata_imports_test_a_2_m3_proto = out.File
	file_cmd_protoc_gen_go_testdata_imports_test_a_2_m3_proto_goTypes = nil
	file_cmd_protoc_gen_go_testdata_imports_test_a_2_m3_proto_depIdxs = nil
}
