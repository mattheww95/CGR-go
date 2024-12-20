// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.28.1
// source: cgr.proto

package main

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CGR struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Cgr  []byte `protobuf:"bytes,2,opt,name=Cgr,proto3" json:"Cgr,omitempty"`
	Size uint64 `protobuf:"varint,3,opt,name=Size,proto3" json:"Size,omitempty"`
}

func (x *CGR) Reset() {
	*x = CGR{}
	mi := &file_cgr_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CGR) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CGR) ProtoMessage() {}

func (x *CGR) ProtoReflect() protoreflect.Message {
	mi := &file_cgr_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CGR.ProtoReflect.Descriptor instead.
func (*CGR) Descriptor() ([]byte, []int) {
	return file_cgr_proto_rawDescGZIP(), []int{0}
}

func (x *CGR) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CGR) GetCgr() []byte {
	if x != nil {
		return x.Cgr
	}
	return nil
}

func (x *CGR) GetSize() uint64 {
	if x != nil {
		return x.Size
	}
	return 0
}

var File_cgr_proto protoreflect.FileDescriptor

var file_cgr_proto_rawDesc = []byte{
	0x0a, 0x09, 0x63, 0x67, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x6d, 0x61, 0x69,
	0x6e, 0x22, 0x3f, 0x0a, 0x03, 0x43, 0x47, 0x52, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03,
	0x43, 0x67, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x03, 0x43, 0x67, 0x72, 0x12, 0x12,
	0x0a, 0x04, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x04, 0x53, 0x69,
	0x7a, 0x65, 0x42, 0x04, 0x5a, 0x02, 0x2e, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cgr_proto_rawDescOnce sync.Once
	file_cgr_proto_rawDescData = file_cgr_proto_rawDesc
)

func file_cgr_proto_rawDescGZIP() []byte {
	file_cgr_proto_rawDescOnce.Do(func() {
		file_cgr_proto_rawDescData = protoimpl.X.CompressGZIP(file_cgr_proto_rawDescData)
	})
	return file_cgr_proto_rawDescData
}

var file_cgr_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_cgr_proto_goTypes = []any{
	(*CGR)(nil), // 0: main.CGR
}
var file_cgr_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_cgr_proto_init() }
func file_cgr_proto_init() {
	if File_cgr_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_cgr_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cgr_proto_goTypes,
		DependencyIndexes: file_cgr_proto_depIdxs,
		MessageInfos:      file_cgr_proto_msgTypes,
	}.Build()
	File_cgr_proto = out.File
	file_cgr_proto_rawDesc = nil
	file_cgr_proto_goTypes = nil
	file_cgr_proto_depIdxs = nil
}
