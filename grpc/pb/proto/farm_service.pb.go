// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.15.8
// source: proto/farm_service.proto

package ASAP

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

type CreateConnectionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MessageRequest *Connection `protobuf:"bytes,1,opt,name=messageRequest,proto3" json:"messageRequest,omitempty"`
}

func (x *CreateConnectionRequest) Reset() {
	*x = CreateConnectionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_farm_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateConnectionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateConnectionRequest) ProtoMessage() {}

func (x *CreateConnectionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_farm_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateConnectionRequest.ProtoReflect.Descriptor instead.
func (*CreateConnectionRequest) Descriptor() ([]byte, []int) {
	return file_proto_farm_service_proto_rawDescGZIP(), []int{0}
}

func (x *CreateConnectionRequest) GetMessageRequest() *Connection {
	if x != nil {
		return x.MessageRequest
	}
	return nil
}

type CreateConnectionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MessageResponse *Connection `protobuf:"bytes,1,opt,name=messageResponse,proto3" json:"messageResponse,omitempty"`
}

func (x *CreateConnectionResponse) Reset() {
	*x = CreateConnectionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_farm_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateConnectionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateConnectionResponse) ProtoMessage() {}

func (x *CreateConnectionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_farm_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateConnectionResponse.ProtoReflect.Descriptor instead.
func (*CreateConnectionResponse) Descriptor() ([]byte, []int) {
	return file_proto_farm_service_proto_rawDescGZIP(), []int{1}
}

func (x *CreateConnectionResponse) GetMessageResponse() *Connection {
	if x != nil {
		return x.MessageResponse
	}
	return nil
}

var File_proto_farm_service_proto protoreflect.FileDescriptor

var file_proto_farm_service_proto_rawDesc = []byte{
	0x0a, 0x18, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x66, 0x61, 0x72, 0x6d, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x41, 0x53, 0x41, 0x50,
	0x1a, 0x18, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x66, 0x61, 0x72, 0x6d, 0x5f, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x53, 0x0a, 0x17, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x38, 0x0a, 0x0e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e,
	0x41, 0x53, 0x41, 0x50, 0x2e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x0e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22,
	0x56, 0x0a, 0x18, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3a, 0x0a, 0x0f, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x41, 0x53, 0x41, 0x50, 0x2e, 0x43, 0x6f, 0x6e, 0x6e,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0x5f, 0x0a, 0x11, 0x43, 0x6f, 0x6e, 0x6e, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4a, 0x0a, 0x07,
	0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x12, 0x1d, 0x2e, 0x41, 0x53, 0x41, 0x50, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x41, 0x53, 0x41, 0x50, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x27, 0x5a, 0x25, 0x2f, 0x68, 0x6f, 0x6d,
	0x65, 0x2f, 0x73, 0x61, 0x79, 0x6e, 0x61, 0x2f, 0x44, 0x65, 0x73, 0x6b, 0x74, 0x6f, 0x70, 0x2f,
	0x63, 0x61, 0x72, 0x72, 0x69, 0x6f, 0x74, 0x2f, 0x41, 0x53, 0x41, 0x50, 0x2f, 0x41, 0x53, 0x41,
	0x50, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_farm_service_proto_rawDescOnce sync.Once
	file_proto_farm_service_proto_rawDescData = file_proto_farm_service_proto_rawDesc
)

func file_proto_farm_service_proto_rawDescGZIP() []byte {
	file_proto_farm_service_proto_rawDescOnce.Do(func() {
		file_proto_farm_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_farm_service_proto_rawDescData)
	})
	return file_proto_farm_service_proto_rawDescData
}

var file_proto_farm_service_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_farm_service_proto_goTypes = []interface{}{
	(*CreateConnectionRequest)(nil),  // 0: GRPC&Protobuf.CreateConnectionRequest
	(*CreateConnectionResponse)(nil), // 1: GRPC&Protobuf.CreateConnectionResponse
	(*Connection)(nil),               // 2: GRPC&Protobuf.Connection
}
var file_proto_farm_service_proto_depIdxs = []int32{
	2, // 0: GRPC&Protobuf.CreateConnectionRequest.messageRequest:type_name -> GRPC&Protobuf.Connection
	2, // 1: GRPC&Protobuf.CreateConnectionResponse.messageResponse:type_name -> GRPC&Protobuf.Connection
	0, // 2: GRPC&Protobuf.ConnectionService.Connect:input_type -> GRPC&Protobuf.CreateConnectionRequest
	1, // 3: GRPC&Protobuf.ConnectionService.Connect:output_type -> GRPC&Protobuf.CreateConnectionResponse
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_proto_farm_service_proto_init() }
func file_proto_farm_service_proto_init() {
	if File_proto_farm_service_proto != nil {
		return
	}
	file_proto_farm_message_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_proto_farm_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateConnectionRequest); i {
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
		file_proto_farm_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateConnectionResponse); i {
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
			RawDescriptor: file_proto_farm_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_farm_service_proto_goTypes,
		DependencyIndexes: file_proto_farm_service_proto_depIdxs,
		MessageInfos:      file_proto_farm_service_proto_msgTypes,
	}.Build()
	File_proto_farm_service_proto = out.File
	file_proto_farm_service_proto_rawDesc = nil
	file_proto_farm_service_proto_goTypes = nil
	file_proto_farm_service_proto_depIdxs = nil
}
