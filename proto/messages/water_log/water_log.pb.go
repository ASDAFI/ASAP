// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.15.8
// source: messages/water_log.proto

package water_log

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CreateWaterLogRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DeviceSerial string                 `protobuf:"bytes,1,opt,name=device_serial,json=deviceSerial,proto3" json:"device_serial,omitempty"`
	UserId       uint32                 `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Volume       float32                `protobuf:"fixed32,3,opt,name=volume,proto3" json:"volume,omitempty"`
	EntryTime    *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=entry_time,json=entryTime,proto3" json:"entry_time,omitempty"`
}

func (x *CreateWaterLogRequest) Reset() {
	*x = CreateWaterLogRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_messages_water_log_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateWaterLogRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateWaterLogRequest) ProtoMessage() {}

func (x *CreateWaterLogRequest) ProtoReflect() protoreflect.Message {
	mi := &file_messages_water_log_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateWaterLogRequest.ProtoReflect.Descriptor instead.
func (*CreateWaterLogRequest) Descriptor() ([]byte, []int) {
	return file_messages_water_log_proto_rawDescGZIP(), []int{0}
}

func (x *CreateWaterLogRequest) GetDeviceSerial() string {
	if x != nil {
		return x.DeviceSerial
	}
	return ""
}

func (x *CreateWaterLogRequest) GetUserId() uint32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *CreateWaterLogRequest) GetVolume() float32 {
	if x != nil {
		return x.Volume
	}
	return 0
}

func (x *CreateWaterLogRequest) GetEntryTime() *timestamppb.Timestamp {
	if x != nil {
		return x.EntryTime
	}
	return nil
}

type CreateWaterLogResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message uint32 `protobuf:"varint,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *CreateWaterLogResponse) Reset() {
	*x = CreateWaterLogResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_messages_water_log_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateWaterLogResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateWaterLogResponse) ProtoMessage() {}

func (x *CreateWaterLogResponse) ProtoReflect() protoreflect.Message {
	mi := &file_messages_water_log_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateWaterLogResponse.ProtoReflect.Descriptor instead.
func (*CreateWaterLogResponse) Descriptor() ([]byte, []int) {
	return file_messages_water_log_proto_rawDescGZIP(), []int{1}
}

func (x *CreateWaterLogResponse) GetMessage() uint32 {
	if x != nil {
		return x.Message
	}
	return 0
}

var File_messages_water_log_proto protoreflect.FileDescriptor

var file_messages_water_log_proto_rawDesc = []byte{
	0x0a, 0x18, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2f, 0x77, 0x61, 0x74, 0x65, 0x72,
	0x5f, 0x6c, 0x6f, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x73, 0x2e, 0x77, 0x61, 0x74, 0x65, 0x72, 0x4c, 0x6f, 0x67, 0x1a, 0x1f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa8,
	0x01, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x57, 0x61, 0x74, 0x65, 0x72, 0x4c, 0x6f,
	0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x64, 0x65, 0x76, 0x69,
	0x63, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0c, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x53, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x12, 0x17, 0x0a,
	0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06,
	0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x06, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x12, 0x39,
	0x0a, 0x0a, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09,
	0x65, 0x6e, 0x74, 0x72, 0x79, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x32, 0x0a, 0x16, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x57, 0x61, 0x74, 0x65, 0x72, 0x4c, 0x6f, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x42, 0x1a, 0x5a,
	0x18, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2f,
	0x77, 0x61, 0x74, 0x65, 0x72, 0x5f, 0x6c, 0x6f, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_messages_water_log_proto_rawDescOnce sync.Once
	file_messages_water_log_proto_rawDescData = file_messages_water_log_proto_rawDesc
)

func file_messages_water_log_proto_rawDescGZIP() []byte {
	file_messages_water_log_proto_rawDescOnce.Do(func() {
		file_messages_water_log_proto_rawDescData = protoimpl.X.CompressGZIP(file_messages_water_log_proto_rawDescData)
	})
	return file_messages_water_log_proto_rawDescData
}

var file_messages_water_log_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_messages_water_log_proto_goTypes = []interface{}{
	(*CreateWaterLogRequest)(nil),  // 0: messages.waterLog.CreateWaterLogRequest
	(*CreateWaterLogResponse)(nil), // 1: messages.waterLog.CreateWaterLogResponse
	(*timestamppb.Timestamp)(nil),  // 2: google.protobuf.Timestamp
}
var file_messages_water_log_proto_depIdxs = []int32{
	2, // 0: messages.waterLog.CreateWaterLogRequest.entry_time:type_name -> google.protobuf.Timestamp
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_messages_water_log_proto_init() }
func file_messages_water_log_proto_init() {
	if File_messages_water_log_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_messages_water_log_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateWaterLogRequest); i {
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
		file_messages_water_log_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateWaterLogResponse); i {
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
			RawDescriptor: file_messages_water_log_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_messages_water_log_proto_goTypes,
		DependencyIndexes: file_messages_water_log_proto_depIdxs,
		MessageInfos:      file_messages_water_log_proto_msgTypes,
	}.Build()
	File_messages_water_log_proto = out.File
	file_messages_water_log_proto_rawDesc = nil
	file_messages_water_log_proto_goTypes = nil
	file_messages_water_log_proto_depIdxs = nil
}