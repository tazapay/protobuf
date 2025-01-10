// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Code generated by protoc-gen-go. DO NOT EDIT.
// source: internal/testprotos/enums/enums_hybrid/enums.hybrid.proto

//go:build !protoopaque

package enums_hybrid

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/gofeaturespb"
	reflect "reflect"
	unsafe "unsafe"
)

type Enum int32

const (
	Enum_DEFAULT     Enum = 1337
	Enum_ZERO        Enum = 0
	Enum_ONE         Enum = 1
	Enum_ELEVENT     Enum = 11
	Enum_SEVENTEEN   Enum = 17
	Enum_THIRTYSEVEN Enum = 37
	Enum_SIXTYSEVEN  Enum = 67
	Enum_NEGATIVE    Enum = -1
)

// Enum value maps for Enum.
var (
	Enum_name = map[int32]string{
		1337: "DEFAULT",
		0:    "ZERO",
		1:    "ONE",
		11:   "ELEVENT",
		17:   "SEVENTEEN",
		37:   "THIRTYSEVEN",
		67:   "SIXTYSEVEN",
		-1:   "NEGATIVE",
	}
	Enum_value = map[string]int32{
		"DEFAULT":     1337,
		"ZERO":        0,
		"ONE":         1,
		"ELEVENT":     11,
		"SEVENTEEN":   17,
		"THIRTYSEVEN": 37,
		"SIXTYSEVEN":  67,
		"NEGATIVE":    -1,
	}
)

func (x Enum) Enum() *Enum {
	p := new(Enum)
	*p = x
	return p
}

func (x Enum) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Enum) Descriptor() protoreflect.EnumDescriptor {
	return file_internal_testprotos_enums_enums_hybrid_enums_hybrid_proto_enumTypes[0].Descriptor()
}

func (Enum) Type() protoreflect.EnumType {
	return &file_internal_testprotos_enums_enums_hybrid_enums_hybrid_proto_enumTypes[0]
}

func (x Enum) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

var File_internal_testprotos_enums_enums_hybrid_enums_hybrid_proto protoreflect.FileDescriptor

const file_internal_testprotos_enums_enums_hybrid_enums_hybrid_proto_rawDesc = "\n9internal/testprotos/enums/enums_hybrid/enums.hybrid.protohybrid.goproto.proto.enums!google/protobuf/go_features.proto*{\nEnum\nDEFAULT\xb9\n\nZERO\x00\nONE\nELEVENT\r\n	SEVENTEEN\nTHIRTYSEVEN%\n\nSIXTYSEVENC\nNEGATIVE\xff\xff\xff\xff\xff\xff\xff\xff\xffBMZAgoogle.golang.org/protobuf/internal/testprotos/enums/enums_hybrid\x92\xd2>beditionsp\xe8"

var file_internal_testprotos_enums_enums_hybrid_enums_hybrid_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_internal_testprotos_enums_enums_hybrid_enums_hybrid_proto_goTypes = []any{
	(Enum)(0), // 0: hybrid.goproto.proto.enums.Enum
}
var file_internal_testprotos_enums_enums_hybrid_enums_hybrid_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_internal_testprotos_enums_enums_hybrid_enums_hybrid_proto_init() }
func file_internal_testprotos_enums_enums_hybrid_enums_hybrid_proto_init() {
	if File_internal_testprotos_enums_enums_hybrid_enums_hybrid_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_internal_testprotos_enums_enums_hybrid_enums_hybrid_proto_rawDesc), len(file_internal_testprotos_enums_enums_hybrid_enums_hybrid_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_internal_testprotos_enums_enums_hybrid_enums_hybrid_proto_goTypes,
		DependencyIndexes: file_internal_testprotos_enums_enums_hybrid_enums_hybrid_proto_depIdxs,
		EnumInfos:         file_internal_testprotos_enums_enums_hybrid_enums_hybrid_proto_enumTypes,
	}.Build()
	File_internal_testprotos_enums_enums_hybrid_enums_hybrid_proto = out.File
	file_internal_testprotos_enums_enums_hybrid_enums_hybrid_proto_goTypes = nil
	file_internal_testprotos_enums_enums_hybrid_enums_hybrid_proto_depIdxs = nil
}
