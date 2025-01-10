// Copyright 2024 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This testproto explicitly configures the API level of each message.
//
// This allows creating mixed trees of proto messages on different API levels.

// Code generated by protoc-gen-go. DO NOT EDIT.
// source: internal/testprotos/mixed/mixed.proto

package mixed

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/gofeaturespb"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

type Open struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// These fields allow for arbitrary mixing.
	Open          *Open   `protobuf:"bytes,1,opt,name=open" json:"open,omitempty"`
	Hybrid        *Hybrid `protobuf:"bytes,2,opt,name=hybrid" json:"hybrid,omitempty"`
	Opaque        *Opaque `protobuf:"bytes,3,opt,name=opaque" json:"opaque,omitempty"`
	OptionalInt32 *int32  `protobuf:"varint,4,opt,name=optional_int32,json=optionalInt32" json:"optional_int32,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Open) Reset() {
	*x = Open{}
	mi := &file_internal_testprotos_mixed_mixed_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Open) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Open) ProtoMessage() {}

func (x *Open) ProtoReflect() protoreflect.Message {
	mi := &file_internal_testprotos_mixed_mixed_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Open.ProtoReflect.Descriptor instead.
func (*Open) Descriptor() ([]byte, []int) {
	return file_internal_testprotos_mixed_mixed_proto_rawDescGZIP(), []int{0}
}

func (x *Open) GetOpen() *Open {
	if x != nil {
		return x.Open
	}
	return nil
}

func (x *Open) GetHybrid() *Hybrid {
	if x != nil {
		return x.Hybrid
	}
	return nil
}

func (x *Open) GetOpaque() *Opaque {
	if x != nil {
		return x.Opaque
	}
	return nil
}

func (x *Open) GetOptionalInt32() int32 {
	if x != nil && x.OptionalInt32 != nil {
		return *x.OptionalInt32
	}
	return 0
}

type Hybrid struct {
	state protoimpl.MessageState `protogen:"hybrid.v1"`
	// These fields allow for arbitrary mixing.
	Open          *Open   `protobuf:"bytes,1,opt,name=open" json:"open,omitempty"`
	Hybrid        *Hybrid `protobuf:"bytes,2,opt,name=hybrid" json:"hybrid,omitempty"`
	Opaque        *Opaque `protobuf:"bytes,3,opt,name=opaque" json:"opaque,omitempty"`
	OptionalInt32 *int32  `protobuf:"varint,4,opt,name=optional_int32,json=optionalInt32" json:"optional_int32,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Hybrid) Reset() {
	*x = Hybrid{}
	mi := &file_internal_testprotos_mixed_mixed_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Hybrid) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Hybrid) ProtoMessage() {}

func (x *Hybrid) ProtoReflect() protoreflect.Message {
	mi := &file_internal_testprotos_mixed_mixed_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *Hybrid) GetOpen() *Open {
	if x != nil {
		return x.Open
	}
	return nil
}

func (x *Hybrid) GetHybrid() *Hybrid {
	if x != nil {
		return x.Hybrid
	}
	return nil
}

func (x *Hybrid) GetOpaque() *Opaque {
	if x != nil {
		return x.Opaque
	}
	return nil
}

func (x *Hybrid) GetOptionalInt32() int32 {
	if x != nil && x.OptionalInt32 != nil {
		return *x.OptionalInt32
	}
	return 0
}

func (x *Hybrid) SetOpen(v *Open) {
	x.Open = v
}

func (x *Hybrid) SetHybrid(v *Hybrid) {
	x.Hybrid = v
}

func (x *Hybrid) SetOpaque(v *Opaque) {
	x.Opaque = v
}

func (x *Hybrid) SetOptionalInt32(v int32) {
	x.OptionalInt32 = &v
}

func (x *Hybrid) HasOpen() bool {
	if x == nil {
		return false
	}
	return x.Open != nil
}

func (x *Hybrid) HasHybrid() bool {
	if x == nil {
		return false
	}
	return x.Hybrid != nil
}

func (x *Hybrid) HasOpaque() bool {
	if x == nil {
		return false
	}
	return x.Opaque != nil
}

func (x *Hybrid) HasOptionalInt32() bool {
	if x == nil {
		return false
	}
	return x.OptionalInt32 != nil
}

func (x *Hybrid) ClearOpen() {
	x.Open = nil
}

func (x *Hybrid) ClearHybrid() {
	x.Hybrid = nil
}

func (x *Hybrid) ClearOpaque() {
	x.Opaque = nil
}

func (x *Hybrid) ClearOptionalInt32() {
	x.OptionalInt32 = nil
}

type Hybrid_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	// These fields allow for arbitrary mixing.
	Open          *Open
	Hybrid        *Hybrid
	Opaque        *Opaque
	OptionalInt32 *int32
}

func (b0 Hybrid_builder) Build() *Hybrid {
	m0 := &Hybrid{}
	b, x := &b0, m0
	_, _ = b, x
	x.Open = b.Open
	x.Hybrid = b.Hybrid
	x.Opaque = b.Opaque
	x.OptionalInt32 = b.OptionalInt32
	return m0
}

type Opaque struct {
	state                    protoimpl.MessageState `protogen:"opaque.v1"`
	xxx_hidden_Open          *Open                  `protobuf:"bytes,1,opt,name=open"`
	xxx_hidden_Hybrid        *Hybrid                `protobuf:"bytes,2,opt,name=hybrid"`
	xxx_hidden_Opaque        *Opaque                `protobuf:"bytes,3,opt,name=opaque"`
	xxx_hidden_OptionalInt32 int32                  `protobuf:"varint,4,opt,name=optional_int32,json=optionalInt32"`
	XXX_raceDetectHookData   protoimpl.RaceDetectHookData
	XXX_presence             [1]uint32
	unknownFields            protoimpl.UnknownFields
	sizeCache                protoimpl.SizeCache
}

func (x *Opaque) Reset() {
	*x = Opaque{}
	mi := &file_internal_testprotos_mixed_mixed_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Opaque) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Opaque) ProtoMessage() {}

func (x *Opaque) ProtoReflect() protoreflect.Message {
	mi := &file_internal_testprotos_mixed_mixed_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *Opaque) GetOpen() *Open {
	if x != nil {
		return x.xxx_hidden_Open
	}
	return nil
}

func (x *Opaque) GetHybrid() *Hybrid {
	if x != nil {
		return x.xxx_hidden_Hybrid
	}
	return nil
}

func (x *Opaque) GetOpaque() *Opaque {
	if x != nil {
		return x.xxx_hidden_Opaque
	}
	return nil
}

func (x *Opaque) GetOptionalInt32() int32 {
	if x != nil {
		return x.xxx_hidden_OptionalInt32
	}
	return 0
}

func (x *Opaque) SetOpen(v *Open) {
	x.xxx_hidden_Open = v
}

func (x *Opaque) SetHybrid(v *Hybrid) {
	x.xxx_hidden_Hybrid = v
}

func (x *Opaque) SetOpaque(v *Opaque) {
	x.xxx_hidden_Opaque = v
}

func (x *Opaque) SetOptionalInt32(v int32) {
	x.xxx_hidden_OptionalInt32 = v
	protoimpl.X.SetPresent(&(x.XXX_presence[0]), 3, 4)
}

func (x *Opaque) HasOpen() bool {
	if x == nil {
		return false
	}
	return x.xxx_hidden_Open != nil
}

func (x *Opaque) HasHybrid() bool {
	if x == nil {
		return false
	}
	return x.xxx_hidden_Hybrid != nil
}

func (x *Opaque) HasOpaque() bool {
	if x == nil {
		return false
	}
	return x.xxx_hidden_Opaque != nil
}

func (x *Opaque) HasOptionalInt32() bool {
	if x == nil {
		return false
	}
	return protoimpl.X.Present(&(x.XXX_presence[0]), 3)
}

func (x *Opaque) ClearOpen() {
	x.xxx_hidden_Open = nil
}

func (x *Opaque) ClearHybrid() {
	x.xxx_hidden_Hybrid = nil
}

func (x *Opaque) ClearOpaque() {
	x.xxx_hidden_Opaque = nil
}

func (x *Opaque) ClearOptionalInt32() {
	protoimpl.X.ClearPresent(&(x.XXX_presence[0]), 3)
	x.xxx_hidden_OptionalInt32 = 0
}

type Opaque_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	// These fields allow for arbitrary mixing.
	Open          *Open
	Hybrid        *Hybrid
	Opaque        *Opaque
	OptionalInt32 *int32
}

func (b0 Opaque_builder) Build() *Opaque {
	m0 := &Opaque{}
	b, x := &b0, m0
	_, _ = b, x
	x.xxx_hidden_Open = b.Open
	x.xxx_hidden_Hybrid = b.Hybrid
	x.xxx_hidden_Opaque = b.Opaque
	if b.OptionalInt32 != nil {
		protoimpl.X.SetPresentNonAtomic(&(x.XXX_presence[0]), 3, 4)
		x.xxx_hidden_OptionalInt32 = *b.OptionalInt32
	}
	return m0
}

type OpenLazy struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// These fields allow for arbitrary mixing.
	Open          *OpenLazy   `protobuf:"bytes,1,opt,name=open" json:"open,omitempty"`
	Hybrid        *HybridLazy `protobuf:"bytes,2,opt,name=hybrid" json:"hybrid,omitempty"`
	Opaque        *OpaqueLazy `protobuf:"bytes,3,opt,name=opaque" json:"opaque,omitempty"`
	OptionalInt32 *int32      `protobuf:"varint,4,opt,name=optional_int32,json=optionalInt32" json:"optional_int32,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *OpenLazy) Reset() {
	*x = OpenLazy{}
	mi := &file_internal_testprotos_mixed_mixed_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *OpenLazy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OpenLazy) ProtoMessage() {}

func (x *OpenLazy) ProtoReflect() protoreflect.Message {
	mi := &file_internal_testprotos_mixed_mixed_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OpenLazy.ProtoReflect.Descriptor instead.
func (*OpenLazy) Descriptor() ([]byte, []int) {
	return file_internal_testprotos_mixed_mixed_proto_rawDescGZIP(), []int{3}
}

func (x *OpenLazy) GetOpen() *OpenLazy {
	if x != nil {
		return x.Open
	}
	return nil
}

func (x *OpenLazy) GetHybrid() *HybridLazy {
	if x != nil {
		return x.Hybrid
	}
	return nil
}

func (x *OpenLazy) GetOpaque() *OpaqueLazy {
	if x != nil {
		return x.Opaque
	}
	return nil
}

func (x *OpenLazy) GetOptionalInt32() int32 {
	if x != nil && x.OptionalInt32 != nil {
		return *x.OptionalInt32
	}
	return 0
}

type HybridLazy struct {
	state protoimpl.MessageState `protogen:"hybrid.v1"`
	// These fields allow for arbitrary mixing.
	Open          *OpenLazy   `protobuf:"bytes,1,opt,name=open" json:"open,omitempty"`
	Hybrid        *HybridLazy `protobuf:"bytes,2,opt,name=hybrid" json:"hybrid,omitempty"`
	Opaque        *OpaqueLazy `protobuf:"bytes,3,opt,name=opaque" json:"opaque,omitempty"`
	OptionalInt32 *int32      `protobuf:"varint,4,opt,name=optional_int32,json=optionalInt32" json:"optional_int32,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *HybridLazy) Reset() {
	*x = HybridLazy{}
	mi := &file_internal_testprotos_mixed_mixed_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *HybridLazy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HybridLazy) ProtoMessage() {}

func (x *HybridLazy) ProtoReflect() protoreflect.Message {
	mi := &file_internal_testprotos_mixed_mixed_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *HybridLazy) GetOpen() *OpenLazy {
	if x != nil {
		return x.Open
	}
	return nil
}

func (x *HybridLazy) GetHybrid() *HybridLazy {
	if x != nil {
		return x.Hybrid
	}
	return nil
}

func (x *HybridLazy) GetOpaque() *OpaqueLazy {
	if x != nil {
		return x.Opaque
	}
	return nil
}

func (x *HybridLazy) GetOptionalInt32() int32 {
	if x != nil && x.OptionalInt32 != nil {
		return *x.OptionalInt32
	}
	return 0
}

func (x *HybridLazy) SetOpen(v *OpenLazy) {
	x.Open = v
}

func (x *HybridLazy) SetHybrid(v *HybridLazy) {
	x.Hybrid = v
}

func (x *HybridLazy) SetOpaque(v *OpaqueLazy) {
	x.Opaque = v
}

func (x *HybridLazy) SetOptionalInt32(v int32) {
	x.OptionalInt32 = &v
}

func (x *HybridLazy) HasOpen() bool {
	if x == nil {
		return false
	}
	return x.Open != nil
}

func (x *HybridLazy) HasHybrid() bool {
	if x == nil {
		return false
	}
	return x.Hybrid != nil
}

func (x *HybridLazy) HasOpaque() bool {
	if x == nil {
		return false
	}
	return x.Opaque != nil
}

func (x *HybridLazy) HasOptionalInt32() bool {
	if x == nil {
		return false
	}
	return x.OptionalInt32 != nil
}

func (x *HybridLazy) ClearOpen() {
	x.Open = nil
}

func (x *HybridLazy) ClearHybrid() {
	x.Hybrid = nil
}

func (x *HybridLazy) ClearOpaque() {
	x.Opaque = nil
}

func (x *HybridLazy) ClearOptionalInt32() {
	x.OptionalInt32 = nil
}

type HybridLazy_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	// These fields allow for arbitrary mixing.
	Open          *OpenLazy
	Hybrid        *HybridLazy
	Opaque        *OpaqueLazy
	OptionalInt32 *int32
}

func (b0 HybridLazy_builder) Build() *HybridLazy {
	m0 := &HybridLazy{}
	b, x := &b0, m0
	_, _ = b, x
	x.Open = b.Open
	x.Hybrid = b.Hybrid
	x.Opaque = b.Opaque
	x.OptionalInt32 = b.OptionalInt32
	return m0
}

type OpaqueLazy struct {
	state                    protoimpl.MessageState `protogen:"opaque.v1"`
	xxx_hidden_Open          *OpenLazy              `protobuf:"bytes,1,opt,name=open"`
	xxx_hidden_Hybrid        *HybridLazy            `protobuf:"bytes,2,opt,name=hybrid"`
	xxx_hidden_Opaque        *OpaqueLazy            `protobuf:"bytes,3,opt,name=opaque"`
	xxx_hidden_OptionalInt32 int32                  `protobuf:"varint,4,opt,name=optional_int32,json=optionalInt32"`
	// Deprecated: Do not use. This will be deleted in the near future.
	XXX_lazyUnmarshalInfo  protoimpl.LazyUnmarshalInfo
	XXX_raceDetectHookData protoimpl.RaceDetectHookData
	XXX_presence           [1]uint32
	unknownFields          protoimpl.UnknownFields
	sizeCache              protoimpl.SizeCache
}

func (x *OpaqueLazy) Reset() {
	*x = OpaqueLazy{}
	mi := &file_internal_testprotos_mixed_mixed_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *OpaqueLazy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OpaqueLazy) ProtoMessage() {}

func (x *OpaqueLazy) ProtoReflect() protoreflect.Message {
	mi := &file_internal_testprotos_mixed_mixed_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *OpaqueLazy) GetOpen() *OpenLazy {
	if x != nil {
		if protoimpl.X.Present(&(x.XXX_presence[0]), 0) {
			if protoimpl.X.AtomicCheckPointerIsNil(&x.xxx_hidden_Open) {
				protoimpl.X.UnmarshalField(x, 1)
			}
			var rv *OpenLazy
			protoimpl.X.AtomicLoadPointer(protoimpl.Pointer(&x.xxx_hidden_Open), protoimpl.Pointer(&rv))
			return rv
		}
	}
	return nil
}

func (x *OpaqueLazy) GetHybrid() *HybridLazy {
	if x != nil {
		if protoimpl.X.Present(&(x.XXX_presence[0]), 1) {
			if protoimpl.X.AtomicCheckPointerIsNil(&x.xxx_hidden_Hybrid) {
				protoimpl.X.UnmarshalField(x, 2)
			}
			var rv *HybridLazy
			protoimpl.X.AtomicLoadPointer(protoimpl.Pointer(&x.xxx_hidden_Hybrid), protoimpl.Pointer(&rv))
			return rv
		}
	}
	return nil
}

func (x *OpaqueLazy) GetOpaque() *OpaqueLazy {
	if x != nil {
		if protoimpl.X.Present(&(x.XXX_presence[0]), 2) {
			if protoimpl.X.AtomicCheckPointerIsNil(&x.xxx_hidden_Opaque) {
				protoimpl.X.UnmarshalField(x, 3)
			}
			var rv *OpaqueLazy
			protoimpl.X.AtomicLoadPointer(protoimpl.Pointer(&x.xxx_hidden_Opaque), protoimpl.Pointer(&rv))
			return rv
		}
	}
	return nil
}

func (x *OpaqueLazy) GetOptionalInt32() int32 {
	if x != nil {
		return x.xxx_hidden_OptionalInt32
	}
	return 0
}

func (x *OpaqueLazy) SetOpen(v *OpenLazy) {
	protoimpl.X.AtomicSetPointer(&x.xxx_hidden_Open, v)
	if v == nil {
		protoimpl.X.ClearPresent(&(x.XXX_presence[0]), 0)
	} else {
		protoimpl.X.SetPresent(&(x.XXX_presence[0]), 0, 4)
	}
}

func (x *OpaqueLazy) SetHybrid(v *HybridLazy) {
	protoimpl.X.AtomicSetPointer(&x.xxx_hidden_Hybrid, v)
	if v == nil {
		protoimpl.X.ClearPresent(&(x.XXX_presence[0]), 1)
	} else {
		protoimpl.X.SetPresent(&(x.XXX_presence[0]), 1, 4)
	}
}

func (x *OpaqueLazy) SetOpaque(v *OpaqueLazy) {
	protoimpl.X.AtomicSetPointer(&x.xxx_hidden_Opaque, v)
	if v == nil {
		protoimpl.X.ClearPresent(&(x.XXX_presence[0]), 2)
	} else {
		protoimpl.X.SetPresent(&(x.XXX_presence[0]), 2, 4)
	}
}

func (x *OpaqueLazy) SetOptionalInt32(v int32) {
	x.xxx_hidden_OptionalInt32 = v
	protoimpl.X.SetPresent(&(x.XXX_presence[0]), 3, 4)
}

func (x *OpaqueLazy) HasOpen() bool {
	if x == nil {
		return false
	}
	return protoimpl.X.Present(&(x.XXX_presence[0]), 0)
}

func (x *OpaqueLazy) HasHybrid() bool {
	if x == nil {
		return false
	}
	return protoimpl.X.Present(&(x.XXX_presence[0]), 1)
}

func (x *OpaqueLazy) HasOpaque() bool {
	if x == nil {
		return false
	}
	return protoimpl.X.Present(&(x.XXX_presence[0]), 2)
}

func (x *OpaqueLazy) HasOptionalInt32() bool {
	if x == nil {
		return false
	}
	return protoimpl.X.Present(&(x.XXX_presence[0]), 3)
}

func (x *OpaqueLazy) ClearOpen() {
	protoimpl.X.ClearPresent(&(x.XXX_presence[0]), 0)
	protoimpl.X.AtomicSetPointer(&x.xxx_hidden_Open, (*OpenLazy)(nil))
}

func (x *OpaqueLazy) ClearHybrid() {
	protoimpl.X.ClearPresent(&(x.XXX_presence[0]), 1)
	protoimpl.X.AtomicSetPointer(&x.xxx_hidden_Hybrid, (*HybridLazy)(nil))
}

func (x *OpaqueLazy) ClearOpaque() {
	protoimpl.X.ClearPresent(&(x.XXX_presence[0]), 2)
	protoimpl.X.AtomicSetPointer(&x.xxx_hidden_Opaque, (*OpaqueLazy)(nil))
}

func (x *OpaqueLazy) ClearOptionalInt32() {
	protoimpl.X.ClearPresent(&(x.XXX_presence[0]), 3)
	x.xxx_hidden_OptionalInt32 = 0
}

type OpaqueLazy_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	// These fields allow for arbitrary mixing.
	Open          *OpenLazy
	Hybrid        *HybridLazy
	Opaque        *OpaqueLazy
	OptionalInt32 *int32
}

func (b0 OpaqueLazy_builder) Build() *OpaqueLazy {
	m0 := &OpaqueLazy{}
	b, x := &b0, m0
	_, _ = b, x
	if b.Open != nil {
		protoimpl.X.SetPresentNonAtomic(&(x.XXX_presence[0]), 0, 4)
		x.xxx_hidden_Open = b.Open
	}
	if b.Hybrid != nil {
		protoimpl.X.SetPresentNonAtomic(&(x.XXX_presence[0]), 1, 4)
		x.xxx_hidden_Hybrid = b.Hybrid
	}
	if b.Opaque != nil {
		protoimpl.X.SetPresentNonAtomic(&(x.XXX_presence[0]), 2, 4)
		x.xxx_hidden_Opaque = b.Opaque
	}
	if b.OptionalInt32 != nil {
		protoimpl.X.SetPresentNonAtomic(&(x.XXX_presence[0]), 3, 4)
		x.xxx_hidden_OptionalInt32 = *b.OptionalInt32
	}
	return m0
}

var File_internal_testprotos_mixed_mixed_proto protoreflect.FileDescriptor

const file_internal_testprotos_mixed_mixed_proto_rawDesc = "\n%internal/testprotos/mixed/mixed.protogoproto.proto.test!google/protobuf/go_features.proto\"\xcc\nOpen,\nopen (2.goproto.proto.test.OpenRopen2\nhybrid (2.goproto.proto.test.HybridRhybrid2\nopaque (2.goproto.proto.test.OpaqueRopaque%\noptional_int32 (R\roptionalInt32:b\xd2>\"\xce\nHybrid,\nopen (2.goproto.proto.test.OpenRopen2\nhybrid (2.goproto.proto.test.HybridRhybrid2\nopaque (2.goproto.proto.test.OpaqueRopaque%\noptional_int32 (R\roptionalInt32:b\xd2>\"\xce\nOpaque,\nopen (2.goproto.proto.test.OpenRopen2\nhybrid (2.goproto.proto.test.HybridRhybrid2\nopaque (2.goproto.proto.test.OpaqueRopaque%\noptional_int32 (R\roptionalInt32:b\xd2>\"\xe8\nOpenLazy4\nopen (2.goproto.proto.test.OpenLazyB(Ropen:\nhybrid (2.goproto.proto.test.HybridLazyB(Rhybrid:\nopaque (2.goproto.proto.test.OpaqueLazyB(Ropaque%\noptional_int32 (R\roptionalInt32:b\xd2>\"\xea\n\nHybridLazy4\nopen (2.goproto.proto.test.OpenLazyB(Ropen:\nhybrid (2.goproto.proto.test.HybridLazyB(Rhybrid:\nopaque (2.goproto.proto.test.OpaqueLazyB(Ropaque%\noptional_int32 (R\roptionalInt32:b\xd2>\"\xea\n\nOpaqueLazy4\nopen (2.goproto.proto.test.OpenLazyB(Ropen:\nhybrid (2.goproto.proto.test.HybridLazyB(Rhybrid:\nopaque (2.goproto.proto.test.OpaqueLazyB(Ropaque%\noptional_int32 (R\roptionalInt32:b\xd2>B6Z4google.golang.org/protobuf/internal/testprotos/mixedbeditionsp\xe8"

var (
	file_internal_testprotos_mixed_mixed_proto_rawDescOnce sync.Once
	file_internal_testprotos_mixed_mixed_proto_rawDescData = file_internal_testprotos_mixed_mixed_proto_rawDesc
)

func file_internal_testprotos_mixed_mixed_proto_rawDescGZIP() []byte {
	file_internal_testprotos_mixed_mixed_proto_rawDescOnce.Do(func() {
		file_internal_testprotos_mixed_mixed_proto_rawDescData = string(protoimpl.X.CompressGZIP([]byte(file_internal_testprotos_mixed_mixed_proto_rawDescData)))
	})
	return []byte(file_internal_testprotos_mixed_mixed_proto_rawDescData)
}

var file_internal_testprotos_mixed_mixed_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_internal_testprotos_mixed_mixed_proto_goTypes = []any{
	(*Open)(nil),       // 0: goproto.proto.test.Open
	(*Hybrid)(nil),     // 1: goproto.proto.test.Hybrid
	(*Opaque)(nil),     // 2: goproto.proto.test.Opaque
	(*OpenLazy)(nil),   // 3: goproto.proto.test.OpenLazy
	(*HybridLazy)(nil), // 4: goproto.proto.test.HybridLazy
	(*OpaqueLazy)(nil), // 5: goproto.proto.test.OpaqueLazy
}
var file_internal_testprotos_mixed_mixed_proto_depIdxs = []int32{
	0,  // 0: goproto.proto.test.Open.open:type_name -> goproto.proto.test.Open
	1,  // 1: goproto.proto.test.Open.hybrid:type_name -> goproto.proto.test.Hybrid
	2,  // 2: goproto.proto.test.Open.opaque:type_name -> goproto.proto.test.Opaque
	0,  // 3: goproto.proto.test.Hybrid.open:type_name -> goproto.proto.test.Open
	1,  // 4: goproto.proto.test.Hybrid.hybrid:type_name -> goproto.proto.test.Hybrid
	2,  // 5: goproto.proto.test.Hybrid.opaque:type_name -> goproto.proto.test.Opaque
	0,  // 6: goproto.proto.test.Opaque.open:type_name -> goproto.proto.test.Open
	1,  // 7: goproto.proto.test.Opaque.hybrid:type_name -> goproto.proto.test.Hybrid
	2,  // 8: goproto.proto.test.Opaque.opaque:type_name -> goproto.proto.test.Opaque
	3,  // 9: goproto.proto.test.OpenLazy.open:type_name -> goproto.proto.test.OpenLazy
	4,  // 10: goproto.proto.test.OpenLazy.hybrid:type_name -> goproto.proto.test.HybridLazy
	5,  // 11: goproto.proto.test.OpenLazy.opaque:type_name -> goproto.proto.test.OpaqueLazy
	3,  // 12: goproto.proto.test.HybridLazy.open:type_name -> goproto.proto.test.OpenLazy
	4,  // 13: goproto.proto.test.HybridLazy.hybrid:type_name -> goproto.proto.test.HybridLazy
	5,  // 14: goproto.proto.test.HybridLazy.opaque:type_name -> goproto.proto.test.OpaqueLazy
	3,  // 15: goproto.proto.test.OpaqueLazy.open:type_name -> goproto.proto.test.OpenLazy
	4,  // 16: goproto.proto.test.OpaqueLazy.hybrid:type_name -> goproto.proto.test.HybridLazy
	5,  // 17: goproto.proto.test.OpaqueLazy.opaque:type_name -> goproto.proto.test.OpaqueLazy
	18, // [18:18] is the sub-list for method output_type
	18, // [18:18] is the sub-list for method input_type
	18, // [18:18] is the sub-list for extension type_name
	18, // [18:18] is the sub-list for extension extendee
	0,  // [0:18] is the sub-list for field type_name
}

func init() { file_internal_testprotos_mixed_mixed_proto_init() }
func file_internal_testprotos_mixed_mixed_proto_init() {
	if File_internal_testprotos_mixed_mixed_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_internal_testprotos_mixed_mixed_proto_rawDesc), len(file_internal_testprotos_mixed_mixed_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_internal_testprotos_mixed_mixed_proto_goTypes,
		DependencyIndexes: file_internal_testprotos_mixed_mixed_proto_depIdxs,
		MessageInfos:      file_internal_testprotos_mixed_mixed_proto_msgTypes,
	}.Build()
	File_internal_testprotos_mixed_mixed_proto = out.File
	file_internal_testprotos_mixed_mixed_proto_goTypes = nil
	file_internal_testprotos_mixed_mixed_proto_depIdxs = nil
}
