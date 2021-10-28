// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.13.0
// source: trip.proto

package trippb

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

type Location struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Latitude  float64 `protobuf:"fixed64,1,opt,name=latitude,proto3" json:"latitude,omitempty"`
	Longitude float64 `protobuf:"fixed64,2,opt,name=longitude,proto3" json:"longitude,omitempty"`
}

func (x *Location) Reset() {
	*x = Location{}
	if protoimpl.UnsafeEnabled {
		mi := &file_trip_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Location) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Location) ProtoMessage() {}

func (x *Location) ProtoReflect() protoreflect.Message {
	mi := &file_trip_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Location.ProtoReflect.Descriptor instead.
func (*Location) Descriptor() ([]byte, []int) {
	return file_trip_proto_rawDescGZIP(), []int{0}
}

func (x *Location) GetLatitude() float64 {
	if x != nil {
		return x.Latitude
	}
	return 0
}

func (x *Location) GetLongitude() float64 {
	if x != nil {
		return x.Longitude
	}
	return 0
}

type Trip struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Start         string      `protobuf:"bytes,1,opt,name=start,proto3" json:"start,omitempty"`                                      // 第一个字段是 start
	StartPos      *Location   `protobuf:"bytes,5,opt,name=start_pos,json=startPos,proto3" json:"start_pos,omitempty"`                // 第5个字段  而不是按照 写的顺序来的
	EndPos        *Location   `protobuf:"bytes,6,opt,name=end_pos,json=endPos,proto3" json:"end_pos,omitempty"`                      // 第5个字段  而不是按照 写的顺序来的
	PathLocaitons []*Location `protobuf:"bytes,7,rep,name=path_locaitons,json=pathLocaitons,proto3" json:"path_locaitons,omitempty"` // repeated 一般用复数形式
	End           string      `protobuf:"bytes,2,opt,name=end,proto3" json:"end,omitempty"`                                          // 第二个字段是 end
	DurationSec   int64       `protobuf:"varint,3,opt,name=duration_sec,json=durationSec,proto3" json:"duration_sec,omitempty"`      // duration in second
	FeeCent       int64       `protobuf:"varint,4,opt,name=fee_cent,json=feeCent,proto3" json:"fee_cent,omitempty"`
}

func (x *Trip) Reset() {
	*x = Trip{}
	if protoimpl.UnsafeEnabled {
		mi := &file_trip_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Trip) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Trip) ProtoMessage() {}

func (x *Trip) ProtoReflect() protoreflect.Message {
	mi := &file_trip_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Trip.ProtoReflect.Descriptor instead.
func (*Trip) Descriptor() ([]byte, []int) {
	return file_trip_proto_rawDescGZIP(), []int{1}
}

func (x *Trip) GetStart() string {
	if x != nil {
		return x.Start
	}
	return ""
}

func (x *Trip) GetStartPos() *Location {
	if x != nil {
		return x.StartPos
	}
	return nil
}

func (x *Trip) GetEndPos() *Location {
	if x != nil {
		return x.EndPos
	}
	return nil
}

func (x *Trip) GetPathLocaitons() []*Location {
	if x != nil {
		return x.PathLocaitons
	}
	return nil
}

func (x *Trip) GetEnd() string {
	if x != nil {
		return x.End
	}
	return ""
}

func (x *Trip) GetDurationSec() int64 {
	if x != nil {
		return x.DurationSec
	}
	return 0
}

func (x *Trip) GetFeeCent() int64 {
	if x != nil {
		return x.FeeCent
	}
	return 0
}

var File_trip_proto protoreflect.FileDescriptor

var file_trip_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x74, 0x72, 0x69, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x63, 0x6f,
	0x6f, 0x6c, 0x63, 0x61, 0x72, 0x22, 0x44, 0x0a, 0x08, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x01, 0x52, 0x08, 0x6c, 0x61, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x12, 0x1c, 0x0a,
	0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01,
	0x52, 0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75, 0x64, 0x65, 0x22, 0x82, 0x02, 0x0a, 0x04,
	0x54, 0x72, 0x69, 0x70, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x12, 0x2e, 0x0a, 0x09, 0x73, 0x74,
	0x61, 0x72, 0x74, 0x5f, 0x70, 0x6f, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e,
	0x63, 0x6f, 0x6f, 0x6c, 0x63, 0x61, 0x72, 0x2e, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x08, 0x73, 0x74, 0x61, 0x72, 0x74, 0x50, 0x6f, 0x73, 0x12, 0x2a, 0x0a, 0x07, 0x65, 0x6e,
	0x64, 0x5f, 0x70, 0x6f, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x63, 0x6f,
	0x6f, 0x6c, 0x63, 0x61, 0x72, 0x2e, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x06,
	0x65, 0x6e, 0x64, 0x50, 0x6f, 0x73, 0x12, 0x38, 0x0a, 0x0e, 0x70, 0x61, 0x74, 0x68, 0x5f, 0x6c,
	0x6f, 0x63, 0x61, 0x69, 0x74, 0x6f, 0x6e, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11,
	0x2e, 0x63, 0x6f, 0x6f, 0x6c, 0x63, 0x61, 0x72, 0x2e, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x0d, 0x70, 0x61, 0x74, 0x68, 0x4c, 0x6f, 0x63, 0x61, 0x69, 0x74, 0x6f, 0x6e, 0x73,
	0x12, 0x10, 0x0a, 0x03, 0x65, 0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x65,
	0x6e, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73,
	0x65, 0x63, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x53, 0x65, 0x63, 0x12, 0x19, 0x0a, 0x08, 0x66, 0x65, 0x65, 0x5f, 0x63, 0x65, 0x6e,
	0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x66, 0x65, 0x65, 0x43, 0x65, 0x6e, 0x74,
	0x42, 0x1c, 0x5a, 0x1a, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x3b, 0x74, 0x72, 0x69, 0x70, 0x70, 0x62, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_trip_proto_rawDescOnce sync.Once
	file_trip_proto_rawDescData = file_trip_proto_rawDesc
)

func file_trip_proto_rawDescGZIP() []byte {
	file_trip_proto_rawDescOnce.Do(func() {
		file_trip_proto_rawDescData = protoimpl.X.CompressGZIP(file_trip_proto_rawDescData)
	})
	return file_trip_proto_rawDescData
}

var file_trip_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_trip_proto_goTypes = []interface{}{
	(*Location)(nil), // 0: coolcar.Location
	(*Trip)(nil),     // 1: coolcar.Trip
}
var file_trip_proto_depIdxs = []int32{
	0, // 0: coolcar.Trip.start_pos:type_name -> coolcar.Location
	0, // 1: coolcar.Trip.end_pos:type_name -> coolcar.Location
	0, // 2: coolcar.Trip.path_locaitons:type_name -> coolcar.Location
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_trip_proto_init() }
func file_trip_proto_init() {
	if File_trip_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_trip_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Location); i {
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
		file_trip_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Trip); i {
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
			RawDescriptor: file_trip_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_trip_proto_goTypes,
		DependencyIndexes: file_trip_proto_depIdxs,
		MessageInfos:      file_trip_proto_msgTypes,
	}.Build()
	File_trip_proto = out.File
	file_trip_proto_rawDesc = nil
	file_trip_proto_goTypes = nil
	file_trip_proto_depIdxs = nil
}