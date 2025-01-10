// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Code generated by protoc-gen-go. DO NOT EDIT.
// cmd/protoc-gen-go/testdata/comments/deprecated.proto is a deprecated file.

package comments

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

// Deprecated: The entire proto file cmd/protoc-gen-go/testdata/comments/deprecated.proto is marked as deprecated.
type DeprecatedEnum int32

const (
	// Deprecated: The entire proto file cmd/protoc-gen-go/testdata/comments/deprecated.proto is marked as deprecated.
	DeprecatedEnum_DEPRECATED DeprecatedEnum = 0
)

// Enum value maps for DeprecatedEnum.
var (
	DeprecatedEnum_name = map[int32]string{
		0: "DEPRECATED",
	}
	DeprecatedEnum_value = map[string]int32{
		"DEPRECATED": 0,
	}
)

func (x DeprecatedEnum) Enum() *DeprecatedEnum {
	p := new(DeprecatedEnum)
	*p = x
	return p
}

func (x DeprecatedEnum) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DeprecatedEnum) Descriptor() protoreflect.EnumDescriptor {
	return file_cmd_protoc_gen_go_testdata_comments_deprecated_proto_enumTypes[0].Descriptor()
}

func (DeprecatedEnum) Type() protoreflect.EnumType {
	return &file_cmd_protoc_gen_go_testdata_comments_deprecated_proto_enumTypes[0]
}

func (x DeprecatedEnum) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DeprecatedEnum.Descriptor instead.
func (DeprecatedEnum) EnumDescriptor() ([]byte, []int) {
	return file_cmd_protoc_gen_go_testdata_comments_deprecated_proto_rawDescGZIP(), []int{0}
}

// Deprecated: The entire proto file cmd/protoc-gen-go/testdata/comments/deprecated.proto is marked as deprecated.
type DeprecatedMessage struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Deprecated: The entire proto file cmd/protoc-gen-go/testdata/comments/deprecated.proto is marked as deprecated.
	DeprecatedField string `protobuf:"bytes,1,opt,name=deprecated_field,json=deprecatedField,proto3" json:"deprecated_field,omitempty"`
	unknownFields   protoimpl.UnknownFields
	sizeCache       protoimpl.SizeCache
}

func (x *DeprecatedMessage) Reset() {
	*x = DeprecatedMessage{}
	mi := &file_cmd_protoc_gen_go_testdata_comments_deprecated_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeprecatedMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeprecatedMessage) ProtoMessage() {}

func (x *DeprecatedMessage) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_protoc_gen_go_testdata_comments_deprecated_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeprecatedMessage.ProtoReflect.Descriptor instead.
func (*DeprecatedMessage) Descriptor() ([]byte, []int) {
	return file_cmd_protoc_gen_go_testdata_comments_deprecated_proto_rawDescGZIP(), []int{0}
}

// Deprecated: The entire proto file cmd/protoc-gen-go/testdata/comments/deprecated.proto is marked as deprecated.
func (x *DeprecatedMessage) GetDeprecatedField() string {
	if x != nil {
		return x.DeprecatedField
	}
	return ""
}

var File_cmd_protoc_gen_go_testdata_comments_deprecated_proto protoreflect.FileDescriptor

const file_cmd_protoc_gen_go_testdata_comments_deprecated_proto_rawDesc = "\n4cmd/protoc-gen-go/testdata/comments/deprecated.protogoproto.protoc.comments\"F\nDeprecatedMessage-\ndeprecated_field (	BRdeprecatedField:*(\nDeprecatedEnum\n\nDEPRECATED\x00BCZ>google.golang.org/protobuf/cmd/protoc-gen-go/testdata/comments\xb8bproto3"

var (
	file_cmd_protoc_gen_go_testdata_comments_deprecated_proto_rawDescOnce sync.Once
	file_cmd_protoc_gen_go_testdata_comments_deprecated_proto_rawDescData = file_cmd_protoc_gen_go_testdata_comments_deprecated_proto_rawDesc
)

func file_cmd_protoc_gen_go_testdata_comments_deprecated_proto_rawDescGZIP() []byte {
	file_cmd_protoc_gen_go_testdata_comments_deprecated_proto_rawDescOnce.Do(func() {
		file_cmd_protoc_gen_go_testdata_comments_deprecated_proto_rawDescData = string(protoimpl.X.CompressGZIP([]byte(file_cmd_protoc_gen_go_testdata_comments_deprecated_proto_rawDescData)))
	})
	return []byte(file_cmd_protoc_gen_go_testdata_comments_deprecated_proto_rawDescData)
}

var file_cmd_protoc_gen_go_testdata_comments_deprecated_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_cmd_protoc_gen_go_testdata_comments_deprecated_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_cmd_protoc_gen_go_testdata_comments_deprecated_proto_goTypes = []any{
	(DeprecatedEnum)(0),       // 0: goproto.protoc.comments.DeprecatedEnum
	(*DeprecatedMessage)(nil), // 1: goproto.protoc.comments.DeprecatedMessage
}
var file_cmd_protoc_gen_go_testdata_comments_deprecated_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_cmd_protoc_gen_go_testdata_comments_deprecated_proto_init() }
func file_cmd_protoc_gen_go_testdata_comments_deprecated_proto_init() {
	if File_cmd_protoc_gen_go_testdata_comments_deprecated_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_cmd_protoc_gen_go_testdata_comments_deprecated_proto_rawDesc), len(file_cmd_protoc_gen_go_testdata_comments_deprecated_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cmd_protoc_gen_go_testdata_comments_deprecated_proto_goTypes,
		DependencyIndexes: file_cmd_protoc_gen_go_testdata_comments_deprecated_proto_depIdxs,
		EnumInfos:         file_cmd_protoc_gen_go_testdata_comments_deprecated_proto_enumTypes,
		MessageInfos:      file_cmd_protoc_gen_go_testdata_comments_deprecated_proto_msgTypes,
	}.Build()
	File_cmd_protoc_gen_go_testdata_comments_deprecated_proto = out.File
	file_cmd_protoc_gen_go_testdata_comments_deprecated_proto_goTypes = nil
	file_cmd_protoc_gen_go_testdata_comments_deprecated_proto_depIdxs = nil
}
