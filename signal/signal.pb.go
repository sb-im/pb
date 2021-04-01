// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.12.3
// source: signal.proto

package signal

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type TrackSource int32

const (
	TrackSource_DRONE   TrackSource = 0
	TrackSource_MONITOR TrackSource = 1
)

// Enum value maps for TrackSource.
var (
	TrackSource_name = map[int32]string{
		0: "DRONE",
		1: "MONITOR",
	}
	TrackSource_value = map[string]int32{
		"DRONE":   0,
		"MONITOR": 1,
	}
)

func (x TrackSource) Enum() *TrackSource {
	p := new(TrackSource)
	*p = x
	return p
}

func (x TrackSource) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TrackSource) Descriptor() protoreflect.EnumDescriptor {
	return file_signal_proto_enumTypes[0].Descriptor()
}

func (TrackSource) Type() protoreflect.EnumType {
	return &file_signal_proto_enumTypes[0]
}

func (x TrackSource) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TrackSource.Descriptor instead.
func (TrackSource) EnumDescriptor() ([]byte, []int) {
	return file_signal_proto_rawDescGZIP(), []int{0}
}

type SessionDescription struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sdp         *SDP        `protobuf:"bytes,1,opt,name=sdp,proto3" json:"sdp,omitempty"`
	Id          string      `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	TrackSource TrackSource `protobuf:"varint,3,opt,name=track_source,json=trackSource,proto3,enum=pb.TrackSource" json:"track_source,omitempty"`
}

func (x *SessionDescription) Reset() {
	*x = SessionDescription{}
	if protoimpl.UnsafeEnabled {
		mi := &file_signal_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SessionDescription) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SessionDescription) ProtoMessage() {}

func (x *SessionDescription) ProtoReflect() protoreflect.Message {
	mi := &file_signal_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SessionDescription.ProtoReflect.Descriptor instead.
func (*SessionDescription) Descriptor() ([]byte, []int) {
	return file_signal_proto_rawDescGZIP(), []int{0}
}

func (x *SessionDescription) GetSdp() *SDP {
	if x != nil {
		return x.Sdp
	}
	return nil
}

func (x *SessionDescription) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *SessionDescription) GetTrackSource() TrackSource {
	if x != nil {
		return x.TrackSource
	}
	return TrackSource_DRONE
}

type SDP struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type        int32  `protobuf:"varint,1,opt,name=type,proto3" json:"type,omitempty"`
	Description []byte `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *SDP) Reset() {
	*x = SDP{}
	if protoimpl.UnsafeEnabled {
		mi := &file_signal_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SDP) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SDP) ProtoMessage() {}

func (x *SDP) ProtoReflect() protoreflect.Message {
	mi := &file_signal_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SDP.ProtoReflect.Descriptor instead.
func (*SDP) Descriptor() ([]byte, []int) {
	return file_signal_proto_rawDescGZIP(), []int{1}
}

func (x *SDP) GetType() int32 {
	if x != nil {
		return x.Type
	}
	return 0
}

func (x *SDP) GetDescription() []byte {
	if x != nil {
		return x.Description
	}
	return nil
}

var File_signal_proto protoreflect.FileDescriptor

var file_signal_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02,
	0x70, 0x62, 0x22, 0x73, 0x0a, 0x12, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x44, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x19, 0x0a, 0x03, 0x73, 0x64, 0x70, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x44, 0x50, 0x52, 0x03,
	0x73, 0x64, 0x70, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x32, 0x0a, 0x0c, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x5f, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0f, 0x2e, 0x70, 0x62, 0x2e, 0x54,
	0x72, 0x61, 0x63, 0x6b, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x0b, 0x74, 0x72, 0x61, 0x63,
	0x6b, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x22, 0x3b, 0x0a, 0x03, 0x53, 0x44, 0x50, 0x12, 0x12,
	0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x2a, 0x25, 0x0a, 0x0b, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x53, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x12, 0x09, 0x0a, 0x05, 0x44, 0x52, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x0b,
	0x0a, 0x07, 0x4d, 0x4f, 0x4e, 0x49, 0x54, 0x4f, 0x52, 0x10, 0x01, 0x42, 0x1c, 0x5a, 0x1a, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x53, 0x42, 0x2d, 0x49, 0x4d, 0x2f,
	0x70, 0x62, 0x2f, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_signal_proto_rawDescOnce sync.Once
	file_signal_proto_rawDescData = file_signal_proto_rawDesc
)

func file_signal_proto_rawDescGZIP() []byte {
	file_signal_proto_rawDescOnce.Do(func() {
		file_signal_proto_rawDescData = protoimpl.X.CompressGZIP(file_signal_proto_rawDescData)
	})
	return file_signal_proto_rawDescData
}

var file_signal_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_signal_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_signal_proto_goTypes = []interface{}{
	(TrackSource)(0),           // 0: pb.TrackSource
	(*SessionDescription)(nil), // 1: pb.SessionDescription
	(*SDP)(nil),                // 2: pb.SDP
}
var file_signal_proto_depIdxs = []int32{
	2, // 0: pb.SessionDescription.sdp:type_name -> pb.SDP
	0, // 1: pb.SessionDescription.track_source:type_name -> pb.TrackSource
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_signal_proto_init() }
func file_signal_proto_init() {
	if File_signal_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_signal_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SessionDescription); i {
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
		file_signal_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SDP); i {
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
			RawDescriptor: file_signal_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_signal_proto_goTypes,
		DependencyIndexes: file_signal_proto_depIdxs,
		EnumInfos:         file_signal_proto_enumTypes,
		MessageInfos:      file_signal_proto_msgTypes,
	}.Build()
	File_signal_proto = out.File
	file_signal_proto_rawDesc = nil
	file_signal_proto_goTypes = nil
	file_signal_proto_depIdxs = nil
}
