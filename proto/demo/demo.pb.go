// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.0
// 	protoc        v5.26.1
// source: demo/demo.proto

package demo

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

type Ping struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Seq       int64 `protobuf:"zigzag64,1,opt,name=seq,proto3" json:"seq,omitempty"`
	Timestamp int64 `protobuf:"zigzag64,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *Ping) Reset() {
	*x = Ping{}
	if protoimpl.UnsafeEnabled {
		mi := &file_demo_demo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Ping) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ping) ProtoMessage() {}

func (x *Ping) ProtoReflect() protoreflect.Message {
	mi := &file_demo_demo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Ping.ProtoReflect.Descriptor instead.
func (*Ping) Descriptor() ([]byte, []int) {
	return file_demo_demo_proto_rawDescGZIP(), []int{0}
}

func (x *Ping) GetSeq() int64 {
	if x != nil {
		return x.Seq
	}
	return 0
}

func (x *Ping) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

type Pong struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ack       int64 `protobuf:"zigzag64,1,opt,name=ack,proto3" json:"ack,omitempty"`
	Timestamp int64 `protobuf:"zigzag64,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *Pong) Reset() {
	*x = Pong{}
	if protoimpl.UnsafeEnabled {
		mi := &file_demo_demo_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Pong) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Pong) ProtoMessage() {}

func (x *Pong) ProtoReflect() protoreflect.Message {
	mi := &file_demo_demo_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Pong.ProtoReflect.Descriptor instead.
func (*Pong) Descriptor() ([]byte, []int) {
	return file_demo_demo_proto_rawDescGZIP(), []int{1}
}

func (x *Pong) GetAck() int64 {
	if x != nil {
		return x.Ack
	}
	return 0
}

func (x *Pong) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

var File_demo_demo_proto protoreflect.FileDescriptor

var file_demo_demo_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x64, 0x65, 0x6d, 0x6f, 0x2f, 0x64, 0x65, 0x6d, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x36, 0x0a, 0x04, 0x50, 0x69, 0x6e, 0x67, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x65, 0x71,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x12, 0x52, 0x03, 0x73, 0x65, 0x71, 0x12, 0x1c, 0x0a, 0x09, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x12, 0x52, 0x09,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x22, 0x36, 0x0a, 0x04, 0x50, 0x6f, 0x6e,
	0x67, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x63, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x12, 0x52, 0x03,
	0x61, 0x63, 0x6b, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x12, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x32, 0x26, 0x0a, 0x04, 0x44, 0x65, 0x6d, 0x6f, 0x12, 0x1e, 0x0a, 0x0c, 0x43, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x05, 0x2e, 0x50, 0x69, 0x6e, 0x67,
	0x1a, 0x05, 0x2e, 0x50, 0x6f, 0x6e, 0x67, 0x28, 0x01, 0x42, 0x2c, 0x5a, 0x2a, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x75, 0x62, 0x62, 0x65, 0x72, 0x64, 0x75,
	0x63, 0x6b, 0x6b, 0x6b, 0x2f, 0x67, 0x6f, 0x2d, 0x64, 0x65, 0x6d, 0x6f, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x64, 0x65, 0x6d, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_demo_demo_proto_rawDescOnce sync.Once
	file_demo_demo_proto_rawDescData = file_demo_demo_proto_rawDesc
)

func file_demo_demo_proto_rawDescGZIP() []byte {
	file_demo_demo_proto_rawDescOnce.Do(func() {
		file_demo_demo_proto_rawDescData = protoimpl.X.CompressGZIP(file_demo_demo_proto_rawDescData)
	})
	return file_demo_demo_proto_rawDescData
}

var file_demo_demo_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_demo_demo_proto_goTypes = []interface{}{
	(*Ping)(nil), // 0: Ping
	(*Pong)(nil), // 1: Pong
}
var file_demo_demo_proto_depIdxs = []int32{
	0, // 0: Demo.ClientStream:input_type -> Ping
	1, // 1: Demo.ClientStream:output_type -> Pong
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_demo_demo_proto_init() }
func file_demo_demo_proto_init() {
	if File_demo_demo_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_demo_demo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Ping); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_demo_demo_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Pong); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_demo_demo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_demo_demo_proto_goTypes,
		DependencyIndexes: file_demo_demo_proto_depIdxs,
		MessageInfos:      file_demo_demo_proto_msgTypes,
	}.Build()
	File_demo_demo_proto = out.File
	file_demo_demo_proto_rawDesc = nil
	file_demo_demo_proto_goTypes = nil
	file_demo_demo_proto_depIdxs = nil
}