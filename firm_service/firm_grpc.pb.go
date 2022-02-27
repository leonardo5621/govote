// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package firm_service

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// FirmServiceClient is the client API for FirmService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FirmServiceClient interface {
	GetFirm(ctx context.Context, in *GetFirmRequest, opts ...grpc.CallOption) (*Firm, error)
	CreateFirm(ctx context.Context, in *CreateFirmRequest, opts ...grpc.CallOption) (*Firm, error)
}

type firmServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFirmServiceClient(cc grpc.ClientConnInterface) FirmServiceClient {
	return &firmServiceClient{cc}
}

func (c *firmServiceClient) GetFirm(ctx context.Context, in *GetFirmRequest, opts ...grpc.CallOption) (*Firm, error) {
	out := new(Firm)
	err := c.cc.Invoke(ctx, "/firm.FirmService/GetFirm", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *firmServiceClient) CreateFirm(ctx context.Context, in *CreateFirmRequest, opts ...grpc.CallOption) (*Firm, error) {
	out := new(Firm)
	err := c.cc.Invoke(ctx, "/firm.FirmService/CreateFirm", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FirmServiceServer is the server API for FirmService service.
// All implementations must embed UnimplementedFirmServiceServer
// for forward compatibility
type FirmServiceServer interface {
	GetFirm(context.Context, *GetFirmRequest) (*Firm, error)
	CreateFirm(context.Context, *CreateFirmRequest) (*Firm, error)
	mustEmbedUnimplementedFirmServiceServer()
}

// UnimplementedFirmServiceServer must be embedded to have forward compatible implementations.
type UnimplementedFirmServiceServer struct {
}

func (UnimplementedFirmServiceServer) GetFirm(context.Context, *GetFirmRequest) (*Firm, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFirm not implemented")
}
func (UnimplementedFirmServiceServer) CreateFirm(context.Context, *CreateFirmRequest) (*Firm, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateFirm not implemented")
}
func (UnimplementedFirmServiceServer) mustEmbedUnimplementedFirmServiceServer() {}

// UnsafeFirmServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FirmServiceServer will
// result in compilation errors.
type UnsafeFirmServiceServer interface {
	mustEmbedUnimplementedFirmServiceServer()
}

func RegisterFirmServiceServer(s grpc.ServiceRegistrar, srv FirmServiceServer) {
	s.RegisterService(&FirmService_ServiceDesc, srv)
}

func _FirmService_GetFirm_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFirmRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FirmServiceServer).GetFirm(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/firm.FirmService/GetFirm",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FirmServiceServer).GetFirm(ctx, req.(*GetFirmRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FirmService_CreateFirm_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateFirmRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FirmServiceServer).CreateFirm(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/firm.FirmService/CreateFirm",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FirmServiceServer).CreateFirm(ctx, req.(*CreateFirmRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// FirmService_ServiceDesc is the grpc.ServiceDesc for FirmService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FirmService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "firm.FirmService",
	HandlerType: (*FirmServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetFirm",
			Handler:    _FirmService_GetFirm_Handler,
		},
		{
			MethodName: "CreateFirm",
			Handler:    _FirmService_CreateFirm_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "firm.proto",
}