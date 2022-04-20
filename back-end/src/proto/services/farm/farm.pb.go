// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.15.8
// source: services/farm.proto

package farm

import (
	context "context"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	devices "messages/devices"
	user "messages/user"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_services_farm_proto protoreflect.FileDescriptor

var file_services_farm_proto_rawDesc = []byte{
	0x0a, 0x13, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x66, 0x61, 0x72, 0x6d, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x66,
	0x61, 0x72, 0x6d, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x13, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2f, 0x75, 0x73, 0x65, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73,
	0x2f, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xd8,
	0x01, 0x0a, 0x04, 0x46, 0x61, 0x72, 0x6d, 0x12, 0x55, 0x0a, 0x05, 0x4c, 0x6f, 0x67, 0x69, 0x6e,
	0x12, 0x1b, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x75, 0x73, 0x65, 0x72,
	0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x4c, 0x6f,
	0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x11, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x0b, 0x22, 0x06, 0x2f, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x3a, 0x01, 0x2a, 0x12, 0x79,
	0x0a, 0x0c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x12, 0x25,
	0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73,
	0x2e, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44,
	0x65, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1a, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x14, 0x22, 0x0f, 0x2f, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x3a, 0x01, 0x2a, 0x42, 0x0f, 0x5a, 0x0d, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x66, 0x61, 0x72, 0x6d, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var file_services_farm_proto_goTypes = []interface{}{
	(*user.LoginRequest)(nil),            // 0: messages.user.LoginRequest
	(*devices.CreateDeviceRequest)(nil),  // 1: messages.devices.CreateDeviceRequest
	(*user.LoginResponse)(nil),           // 2: messages.user.LoginResponse
	(*devices.CreateDeviceResponse)(nil), // 3: messages.devices.CreateDeviceResponse
}
var file_services_farm_proto_depIdxs = []int32{
	0, // 0: service.farm.Farm.Login:input_type -> messages.user.LoginRequest
	1, // 1: service.farm.Farm.CreateDevice:input_type -> messages.devices.CreateDeviceRequest
	2, // 2: service.farm.Farm.Login:output_type -> messages.user.LoginResponse
	3, // 3: service.farm.Farm.CreateDevice:output_type -> messages.devices.CreateDeviceResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_services_farm_proto_init() }
func file_services_farm_proto_init() {
	if File_services_farm_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_services_farm_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_services_farm_proto_goTypes,
		DependencyIndexes: file_services_farm_proto_depIdxs,
	}.Build()
	File_services_farm_proto = out.File
	file_services_farm_proto_rawDesc = nil
	file_services_farm_proto_goTypes = nil
	file_services_farm_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// FarmClient is the client API for Farm service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type FarmClient interface {
	Login(ctx context.Context, in *user.LoginRequest, opts ...grpc.CallOption) (*user.LoginResponse, error)
	CreateDevice(ctx context.Context, in *devices.CreateDeviceRequest, opts ...grpc.CallOption) (*devices.CreateDeviceResponse, error)
}

type farmClient struct {
	cc grpc.ClientConnInterface
}

func NewFarmClient(cc grpc.ClientConnInterface) FarmClient {
	return &farmClient{cc}
}

func (c *farmClient) Login(ctx context.Context, in *user.LoginRequest, opts ...grpc.CallOption) (*user.LoginResponse, error) {
	out := new(user.LoginResponse)
	err := c.cc.Invoke(ctx, "/service.farm.Farm/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *farmClient) CreateDevice(ctx context.Context, in *devices.CreateDeviceRequest, opts ...grpc.CallOption) (*devices.CreateDeviceResponse, error) {
	out := new(devices.CreateDeviceResponse)
	err := c.cc.Invoke(ctx, "/service.farm.Farm/CreateDevice", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FarmServer is the server API for Farm service.
type FarmServer interface {
	Login(context.Context, *user.LoginRequest) (*user.LoginResponse, error)
	CreateDevice(context.Context, *devices.CreateDeviceRequest) (*devices.CreateDeviceResponse, error)
}

// UnimplementedFarmServer can be embedded to have forward compatible implementations.
type UnimplementedFarmServer struct {
}

func (*UnimplementedFarmServer) Login(context.Context, *user.LoginRequest) (*user.LoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (*UnimplementedFarmServer) CreateDevice(context.Context, *devices.CreateDeviceRequest) (*devices.CreateDeviceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateDevice not implemented")
}

func RegisterFarmServer(s *grpc.Server, srv FarmServer) {
	s.RegisterService(&_Farm_serviceDesc, srv)
}

func _Farm_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(user.LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FarmServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.farm.Farm/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FarmServer).Login(ctx, req.(*user.LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Farm_CreateDevice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(devices.CreateDeviceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FarmServer).CreateDevice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.farm.Farm/CreateDevice",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FarmServer).CreateDevice(ctx, req.(*devices.CreateDeviceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Farm_serviceDesc = grpc.ServiceDesc{
	ServiceName: "service.farm.Farm",
	HandlerType: (*FarmServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _Farm_Login_Handler,
		},
		{
			MethodName: "CreateDevice",
			Handler:    _Farm_CreateDevice_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "services/farm.proto",
}
