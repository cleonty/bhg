//implant.proto

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.13.0
// source: implant.proto

package grpcapi

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// Command defines a with both input and output fields
type Command struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	In  string `protobuf:"bytes,1,opt,name=In,proto3" json:"In,omitempty"`
	Out string `protobuf:"bytes,2,opt,name=Out,proto3" json:"Out,omitempty"`
}

func (x *Command) Reset() {
	*x = Command{}
	if protoimpl.UnsafeEnabled {
		mi := &file_implant_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Command) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Command) ProtoMessage() {}

func (x *Command) ProtoReflect() protoreflect.Message {
	mi := &file_implant_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Command.ProtoReflect.Descriptor instead.
func (*Command) Descriptor() ([]byte, []int) {
	return file_implant_proto_rawDescGZIP(), []int{0}
}

func (x *Command) GetIn() string {
	if x != nil {
		return x.In
	}
	return ""
}

func (x *Command) GetOut() string {
	if x != nil {
		return x.Out
	}
	return ""
}

// Empty defines an empty message used in place of null
type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_implant_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_implant_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_implant_proto_rawDescGZIP(), []int{1}
}

var File_implant_proto protoreflect.FileDescriptor

var file_implant_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x69, 0x6d, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x67, 0x72, 0x70, 0x63, 0x61, 0x70, 0x69, 0x22, 0x2b, 0x0a, 0x07, 0x43, 0x6f, 0x6d, 0x6d,
	0x61, 0x6e, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x49, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x4f, 0x75, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x4f, 0x75, 0x74, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x32, 0x6b,
	0x0a, 0x07, 0x49, 0x6d, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x12, 0x30, 0x0a, 0x0c, 0x46, 0x65, 0x74,
	0x63, 0x68, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x12, 0x0e, 0x2e, 0x67, 0x72, 0x70, 0x63,
	0x61, 0x70, 0x69, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x10, 0x2e, 0x67, 0x72, 0x70, 0x63,
	0x61, 0x70, 0x69, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x12, 0x2e, 0x0a, 0x0a, 0x53,
	0x65, 0x6e, 0x64, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x12, 0x10, 0x2e, 0x67, 0x72, 0x70, 0x63,
	0x61, 0x70, 0x69, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x1a, 0x0e, 0x2e, 0x67, 0x72,
	0x70, 0x63, 0x61, 0x70, 0x69, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x32, 0x39, 0x0a, 0x05, 0x41,
	0x64, 0x6d, 0x69, 0x6e, 0x12, 0x30, 0x0a, 0x0a, 0x52, 0x75, 0x6e, 0x43, 0x6f, 0x6d, 0x6d, 0x61,
	0x6e, 0x64, 0x12, 0x10, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x6f, 0x6d,
	0x6d, 0x61, 0x6e, 0x64, 0x1a, 0x10, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x61, 0x70, 0x69, 0x2e, 0x43,
	0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x42, 0x15, 0x5a, 0x13, 0x63, 0x6c, 0x65, 0x6f, 0x6e, 0x74,
	0x79, 0x2f, 0x72, 0x61, 0x74, 0x3b, 0x67, 0x72, 0x70, 0x63, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_implant_proto_rawDescOnce sync.Once
	file_implant_proto_rawDescData = file_implant_proto_rawDesc
)

func file_implant_proto_rawDescGZIP() []byte {
	file_implant_proto_rawDescOnce.Do(func() {
		file_implant_proto_rawDescData = protoimpl.X.CompressGZIP(file_implant_proto_rawDescData)
	})
	return file_implant_proto_rawDescData
}

var file_implant_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_implant_proto_goTypes = []interface{}{
	(*Command)(nil), // 0: grpcapi.Command
	(*Empty)(nil),   // 1: grpcapi.Empty
}
var file_implant_proto_depIdxs = []int32{
	1, // 0: grpcapi.Implant.FetchCommand:input_type -> grpcapi.Empty
	0, // 1: grpcapi.Implant.SendOutput:input_type -> grpcapi.Command
	0, // 2: grpcapi.Admin.RunCommand:input_type -> grpcapi.Command
	0, // 3: grpcapi.Implant.FetchCommand:output_type -> grpcapi.Command
	1, // 4: grpcapi.Implant.SendOutput:output_type -> grpcapi.Empty
	0, // 5: grpcapi.Admin.RunCommand:output_type -> grpcapi.Command
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_implant_proto_init() }
func file_implant_proto_init() {
	if File_implant_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_implant_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Command); i {
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
		file_implant_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
			RawDescriptor: file_implant_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_implant_proto_goTypes,
		DependencyIndexes: file_implant_proto_depIdxs,
		MessageInfos:      file_implant_proto_msgTypes,
	}.Build()
	File_implant_proto = out.File
	file_implant_proto_rawDesc = nil
	file_implant_proto_goTypes = nil
	file_implant_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ImplantClient is the client API for Implant service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ImplantClient interface {
	FetchCommand(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Command, error)
	SendOutput(ctx context.Context, in *Command, opts ...grpc.CallOption) (*Empty, error)
}

type implantClient struct {
	cc grpc.ClientConnInterface
}

func NewImplantClient(cc grpc.ClientConnInterface) ImplantClient {
	return &implantClient{cc}
}

func (c *implantClient) FetchCommand(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Command, error) {
	out := new(Command)
	err := c.cc.Invoke(ctx, "/grpcapi.Implant/FetchCommand", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *implantClient) SendOutput(ctx context.Context, in *Command, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/grpcapi.Implant/SendOutput", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ImplantServer is the server API for Implant service.
type ImplantServer interface {
	FetchCommand(context.Context, *Empty) (*Command, error)
	SendOutput(context.Context, *Command) (*Empty, error)
}

// UnimplementedImplantServer can be embedded to have forward compatible implementations.
type UnimplementedImplantServer struct {
}

func (*UnimplementedImplantServer) FetchCommand(context.Context, *Empty) (*Command, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FetchCommand not implemented")
}
func (*UnimplementedImplantServer) SendOutput(context.Context, *Command) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendOutput not implemented")
}

func RegisterImplantServer(s *grpc.Server, srv ImplantServer) {
	s.RegisterService(&_Implant_serviceDesc, srv)
}

func _Implant_FetchCommand_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImplantServer).FetchCommand(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpcapi.Implant/FetchCommand",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImplantServer).FetchCommand(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Implant_SendOutput_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Command)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImplantServer).SendOutput(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpcapi.Implant/SendOutput",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImplantServer).SendOutput(ctx, req.(*Command))
	}
	return interceptor(ctx, in, info, handler)
}

var _Implant_serviceDesc = grpc.ServiceDesc{
	ServiceName: "grpcapi.Implant",
	HandlerType: (*ImplantServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FetchCommand",
			Handler:    _Implant_FetchCommand_Handler,
		},
		{
			MethodName: "SendOutput",
			Handler:    _Implant_SendOutput_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "implant.proto",
}

// AdminClient is the client API for Admin service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AdminClient interface {
	RunCommand(ctx context.Context, in *Command, opts ...grpc.CallOption) (*Command, error)
}

type adminClient struct {
	cc grpc.ClientConnInterface
}

func NewAdminClient(cc grpc.ClientConnInterface) AdminClient {
	return &adminClient{cc}
}

func (c *adminClient) RunCommand(ctx context.Context, in *Command, opts ...grpc.CallOption) (*Command, error) {
	out := new(Command)
	err := c.cc.Invoke(ctx, "/grpcapi.Admin/RunCommand", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AdminServer is the server API for Admin service.
type AdminServer interface {
	RunCommand(context.Context, *Command) (*Command, error)
}

// UnimplementedAdminServer can be embedded to have forward compatible implementations.
type UnimplementedAdminServer struct {
}

func (*UnimplementedAdminServer) RunCommand(context.Context, *Command) (*Command, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RunCommand not implemented")
}

func RegisterAdminServer(s *grpc.Server, srv AdminServer) {
	s.RegisterService(&_Admin_serviceDesc, srv)
}

func _Admin_RunCommand_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Command)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServer).RunCommand(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpcapi.Admin/RunCommand",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServer).RunCommand(ctx, req.(*Command))
	}
	return interceptor(ctx, in, info, handler)
}

var _Admin_serviceDesc = grpc.ServiceDesc{
	ServiceName: "grpcapi.Admin",
	HandlerType: (*AdminServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RunCommand",
			Handler:    _Admin_RunCommand_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "implant.proto",
}
