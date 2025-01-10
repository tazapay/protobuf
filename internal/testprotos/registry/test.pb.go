// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Different proto type definitions for testing the Types registry.

// Code generated by protoc-gen-go. DO NOT EDIT.
// source: internal/testprotos/registry/test.proto

package registry

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

type Enum1 int32

const (
	Enum1_ONE Enum1 = 1
)

// Enum value maps for Enum1.
var (
	Enum1_name = map[int32]string{
		1: "ONE",
	}
	Enum1_value = map[string]int32{
		"ONE": 1,
	}
)

func (x Enum1) Enum() *Enum1 {
	p := new(Enum1)
	*p = x
	return p
}

func (x Enum1) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Enum1) Descriptor() protoreflect.EnumDescriptor {
	return file_internal_testprotos_registry_test_proto_enumTypes[0].Descriptor()
}

func (Enum1) Type() protoreflect.EnumType {
	return &file_internal_testprotos_registry_test_proto_enumTypes[0]
}

func (x Enum1) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *Enum1) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = Enum1(num)
	return nil
}

// Deprecated: Use Enum1.Descriptor instead.
func (Enum1) EnumDescriptor() ([]byte, []int) {
	return file_internal_testprotos_registry_test_proto_rawDescGZIP(), []int{0}
}

type Enum2 int32

const (
	Enum2_UNO Enum2 = 1
)

// Enum value maps for Enum2.
var (
	Enum2_name = map[int32]string{
		1: "UNO",
	}
	Enum2_value = map[string]int32{
		"UNO": 1,
	}
)

func (x Enum2) Enum() *Enum2 {
	p := new(Enum2)
	*p = x
	return p
}

func (x Enum2) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Enum2) Descriptor() protoreflect.EnumDescriptor {
	return file_internal_testprotos_registry_test_proto_enumTypes[1].Descriptor()
}

func (Enum2) Type() protoreflect.EnumType {
	return &file_internal_testprotos_registry_test_proto_enumTypes[1]
}

func (x Enum2) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *Enum2) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = Enum2(num)
	return nil
}

// Deprecated: Use Enum2.Descriptor instead.
func (Enum2) EnumDescriptor() ([]byte, []int) {
	return file_internal_testprotos_registry_test_proto_rawDescGZIP(), []int{1}
}

type Enum3 int32

const (
	Enum3_YI Enum3 = 1
)

// Enum value maps for Enum3.
var (
	Enum3_name = map[int32]string{
		1: "YI",
	}
	Enum3_value = map[string]int32{
		"YI": 1,
	}
)

func (x Enum3) Enum() *Enum3 {
	p := new(Enum3)
	*p = x
	return p
}

func (x Enum3) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Enum3) Descriptor() protoreflect.EnumDescriptor {
	return file_internal_testprotos_registry_test_proto_enumTypes[2].Descriptor()
}

func (Enum3) Type() protoreflect.EnumType {
	return &file_internal_testprotos_registry_test_proto_enumTypes[2]
}

func (x Enum3) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *Enum3) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = Enum3(num)
	return nil
}

// Deprecated: Use Enum3.Descriptor instead.
func (Enum3) EnumDescriptor() ([]byte, []int) {
	return file_internal_testprotos_registry_test_proto_rawDescGZIP(), []int{2}
}

type Message1 struct {
	state           protoimpl.MessageState `protogen:"open.v1"`
	extensionFields protoimpl.ExtensionFields
	unknownFields   protoimpl.UnknownFields
	sizeCache       protoimpl.SizeCache
}

func (x *Message1) Reset() {
	*x = Message1{}
	mi := &file_internal_testprotos_registry_test_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Message1) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message1) ProtoMessage() {}

func (x *Message1) ProtoReflect() protoreflect.Message {
	mi := &file_internal_testprotos_registry_test_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Message1.ProtoReflect.Descriptor instead.
func (*Message1) Descriptor() ([]byte, []int) {
	return file_internal_testprotos_registry_test_proto_rawDescGZIP(), []int{0}
}

type Message2 struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Message2) Reset() {
	*x = Message2{}
	mi := &file_internal_testprotos_registry_test_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Message2) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message2) ProtoMessage() {}

func (x *Message2) ProtoReflect() protoreflect.Message {
	mi := &file_internal_testprotos_registry_test_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Message2.ProtoReflect.Descriptor instead.
func (*Message2) Descriptor() ([]byte, []int) {
	return file_internal_testprotos_registry_test_proto_rawDescGZIP(), []int{1}
}

type Message3 struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Message3) Reset() {
	*x = Message3{}
	mi := &file_internal_testprotos_registry_test_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Message3) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message3) ProtoMessage() {}

func (x *Message3) ProtoReflect() protoreflect.Message {
	mi := &file_internal_testprotos_registry_test_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Message3.ProtoReflect.Descriptor instead.
func (*Message3) Descriptor() ([]byte, []int) {
	return file_internal_testprotos_registry_test_proto_rawDescGZIP(), []int{2}
}

type Message4 struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	BoolField     *bool                  `protobuf:"varint,30,opt,name=bool_field,json=boolField" json:"bool_field,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Message4) Reset() {
	*x = Message4{}
	mi := &file_internal_testprotos_registry_test_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Message4) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message4) ProtoMessage() {}

func (x *Message4) ProtoReflect() protoreflect.Message {
	mi := &file_internal_testprotos_registry_test_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Message4.ProtoReflect.Descriptor instead.
func (*Message4) Descriptor() ([]byte, []int) {
	return file_internal_testprotos_registry_test_proto_rawDescGZIP(), []int{3}
}

func (x *Message4) GetBoolField() bool {
	if x != nil && x.BoolField != nil {
		return *x.BoolField
	}
	return false
}

var file_internal_testprotos_registry_test_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*Message1)(nil),
		ExtensionType: (*string)(nil),
		Field:         11,
		Name:          "testprotos.string_field",
		Tag:           "bytes,11,opt,name=string_field",
		Filename:      "internal/testprotos/registry/test.proto",
	},
	{
		ExtendedType:  (*Message1)(nil),
		ExtensionType: (*Enum1)(nil),
		Field:         12,
		Name:          "testprotos.enum_field",
		Tag:           "varint,12,opt,name=enum_field,enum=testprotos.Enum1",
		Filename:      "internal/testprotos/registry/test.proto",
	},
	{
		ExtendedType:  (*Message1)(nil),
		ExtensionType: (*Message2)(nil),
		Field:         13,
		Name:          "testprotos.message_field",
		Tag:           "bytes,13,opt,name=message_field",
		Filename:      "internal/testprotos/registry/test.proto",
	},
	{
		ExtendedType:  (*Message1)(nil),
		ExtensionType: (*Message2)(nil),
		Field:         21,
		Name:          "testprotos.Message4.message_field",
		Tag:           "bytes,21,opt,name=message_field",
		Filename:      "internal/testprotos/registry/test.proto",
	},
	{
		ExtendedType:  (*Message1)(nil),
		ExtensionType: (*Enum1)(nil),
		Field:         22,
		Name:          "testprotos.Message4.enum_field",
		Tag:           "varint,22,opt,name=enum_field,enum=testprotos.Enum1",
		Filename:      "internal/testprotos/registry/test.proto",
	},
	{
		ExtendedType:  (*Message1)(nil),
		ExtensionType: (*string)(nil),
		Field:         23,
		Name:          "testprotos.Message4.string_field",
		Tag:           "bytes,23,opt,name=string_field",
		Filename:      "internal/testprotos/registry/test.proto",
	},
}

// Extension fields to Message1.
var (
	// optional string string_field = 11;
	E_StringField = &file_internal_testprotos_registry_test_proto_extTypes[0]
	// optional testprotos.Enum1 enum_field = 12;
	E_EnumField = &file_internal_testprotos_registry_test_proto_extTypes[1]
	// optional testprotos.Message2 message_field = 13;
	E_MessageField = &file_internal_testprotos_registry_test_proto_extTypes[2]
	// optional testprotos.Message2 message_field = 21;
	E_Message4_MessageField = &file_internal_testprotos_registry_test_proto_extTypes[3]
	// optional testprotos.Enum1 enum_field = 22;
	E_Message4_EnumField = &file_internal_testprotos_registry_test_proto_extTypes[4]
	// optional string string_field = 23;
	E_Message4_StringField = &file_internal_testprotos_registry_test_proto_extTypes[5]
)

var File_internal_testprotos_registry_test_proto protoreflect.FileDescriptor

const file_internal_testprotos_registry_test_proto_rawDesc = "\n'internal/testprotos/registry/test.proto\ntestprotos\"\nMessage1*\n\x80\x80\x80\x80\"\n\nMessage2\"\n\nMessage3\"\xfb\nMessage4\n\nbool_field (R	boolField2O\n\rmessage_field.testprotos.Message1 (2.testprotos.Message2RmessageField2F\n\nenum_field.testprotos.Message1 (2.testprotos.Enum1R	enumField27\nstring_field.testprotos.Message1 (	RstringField*\nEnum1\nONE*\nEnum2\nUNO*\nEnum3\nYI:7\nstring_field.testprotos.Message1 (	RstringField:F\n\nenum_field.testprotos.Message1 (2.testprotos.Enum1R	enumField:O\n\rmessage_field.testprotos.Message1\r (2.testprotos.Message2RmessageFieldB9Z7google.golang.org/protobuf/internal/testprotos/registry"

var (
	file_internal_testprotos_registry_test_proto_rawDescOnce sync.Once
	file_internal_testprotos_registry_test_proto_rawDescData = file_internal_testprotos_registry_test_proto_rawDesc
)

func file_internal_testprotos_registry_test_proto_rawDescGZIP() []byte {
	file_internal_testprotos_registry_test_proto_rawDescOnce.Do(func() {
		file_internal_testprotos_registry_test_proto_rawDescData = string(protoimpl.X.CompressGZIP([]byte(file_internal_testprotos_registry_test_proto_rawDescData)))
	})
	return []byte(file_internal_testprotos_registry_test_proto_rawDescData)
}

var file_internal_testprotos_registry_test_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_internal_testprotos_registry_test_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_internal_testprotos_registry_test_proto_goTypes = []any{
	(Enum1)(0),       // 0: testprotos.Enum1
	(Enum2)(0),       // 1: testprotos.Enum2
	(Enum3)(0),       // 2: testprotos.Enum3
	(*Message1)(nil), // 3: testprotos.Message1
	(*Message2)(nil), // 4: testprotos.Message2
	(*Message3)(nil), // 5: testprotos.Message3
	(*Message4)(nil), // 6: testprotos.Message4
}
var file_internal_testprotos_registry_test_proto_depIdxs = []int32{
	3,  // 0: testprotos.string_field:extendee -> testprotos.Message1
	3,  // 1: testprotos.enum_field:extendee -> testprotos.Message1
	3,  // 2: testprotos.message_field:extendee -> testprotos.Message1
	3,  // 3: testprotos.Message4.message_field:extendee -> testprotos.Message1
	3,  // 4: testprotos.Message4.enum_field:extendee -> testprotos.Message1
	3,  // 5: testprotos.Message4.string_field:extendee -> testprotos.Message1
	0,  // 6: testprotos.enum_field:type_name -> testprotos.Enum1
	4,  // 7: testprotos.message_field:type_name -> testprotos.Message2
	4,  // 8: testprotos.Message4.message_field:type_name -> testprotos.Message2
	0,  // 9: testprotos.Message4.enum_field:type_name -> testprotos.Enum1
	10, // [10:10] is the sub-list for method output_type
	10, // [10:10] is the sub-list for method input_type
	6,  // [6:10] is the sub-list for extension type_name
	0,  // [0:6] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_internal_testprotos_registry_test_proto_init() }
func file_internal_testprotos_registry_test_proto_init() {
	if File_internal_testprotos_registry_test_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_internal_testprotos_registry_test_proto_rawDesc), len(file_internal_testprotos_registry_test_proto_rawDesc)),
			NumEnums:      3,
			NumMessages:   4,
			NumExtensions: 6,
			NumServices:   0,
		},
		GoTypes:           file_internal_testprotos_registry_test_proto_goTypes,
		DependencyIndexes: file_internal_testprotos_registry_test_proto_depIdxs,
		EnumInfos:         file_internal_testprotos_registry_test_proto_enumTypes,
		MessageInfos:      file_internal_testprotos_registry_test_proto_msgTypes,
		ExtensionInfos:    file_internal_testprotos_registry_test_proto_extTypes,
	}.Build()
	File_internal_testprotos_registry_test_proto = out.File
	file_internal_testprotos_registry_test_proto_goTypes = nil
	file_internal_testprotos_registry_test_proto_depIdxs = nil
}
