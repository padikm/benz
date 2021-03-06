// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package data

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

// CreateEmpServiceClient is the client API for CreateEmpService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CreateEmpServiceClient interface {
	Create(ctx context.Context, in *EmpReq, opts ...grpc.CallOption) (*EmpResp, error)
	Get(ctx context.Context, in *NoArg, opts ...grpc.CallOption) (*GetResp, error)
	Edit(ctx context.Context, in *EmpReq, opts ...grpc.CallOption) (*EmpResp, error)
}

type createEmpServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCreateEmpServiceClient(cc grpc.ClientConnInterface) CreateEmpServiceClient {
	return &createEmpServiceClient{cc}
}

func (c *createEmpServiceClient) Create(ctx context.Context, in *EmpReq, opts ...grpc.CallOption) (*EmpResp, error) {
	out := new(EmpResp)
	err := c.cc.Invoke(ctx, "/data.CreateEmpService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *createEmpServiceClient) Get(ctx context.Context, in *NoArg, opts ...grpc.CallOption) (*GetResp, error) {
	out := new(GetResp)
	err := c.cc.Invoke(ctx, "/data.CreateEmpService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *createEmpServiceClient) Edit(ctx context.Context, in *EmpReq, opts ...grpc.CallOption) (*EmpResp, error) {
	out := new(EmpResp)
	err := c.cc.Invoke(ctx, "/data.CreateEmpService/Edit", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CreateEmpServiceServer is the server API for CreateEmpService service.
// All implementations must embed UnimplementedCreateEmpServiceServer
// for forward compatibility
type CreateEmpServiceServer interface {
	Create(context.Context, *EmpReq) (*EmpResp, error)
	Get(context.Context, *NoArg) (*GetResp, error)
	Edit(context.Context, *EmpReq) (*EmpResp, error)
	mustEmbedUnimplementedCreateEmpServiceServer()
}

// UnimplementedCreateEmpServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCreateEmpServiceServer struct {
}

func (UnimplementedCreateEmpServiceServer) Create(context.Context, *EmpReq) (*EmpResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedCreateEmpServiceServer) Get(context.Context, *NoArg) (*GetResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedCreateEmpServiceServer) Edit(context.Context, *EmpReq) (*EmpResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Edit not implemented")
}
func (UnimplementedCreateEmpServiceServer) mustEmbedUnimplementedCreateEmpServiceServer() {}

// UnsafeCreateEmpServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CreateEmpServiceServer will
// result in compilation errors.
type UnsafeCreateEmpServiceServer interface {
	mustEmbedUnimplementedCreateEmpServiceServer()
}

func RegisterCreateEmpServiceServer(s grpc.ServiceRegistrar, srv CreateEmpServiceServer) {
	s.RegisterService(&CreateEmpService_ServiceDesc, srv)
}

func _CreateEmpService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmpReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CreateEmpServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/data.CreateEmpService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CreateEmpServiceServer).Create(ctx, req.(*EmpReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _CreateEmpService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NoArg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CreateEmpServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/data.CreateEmpService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CreateEmpServiceServer).Get(ctx, req.(*NoArg))
	}
	return interceptor(ctx, in, info, handler)
}

func _CreateEmpService_Edit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmpReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CreateEmpServiceServer).Edit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/data.CreateEmpService/Edit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CreateEmpServiceServer).Edit(ctx, req.(*EmpReq))
	}
	return interceptor(ctx, in, info, handler)
}

// CreateEmpService_ServiceDesc is the grpc.ServiceDesc for CreateEmpService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CreateEmpService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "data.CreateEmpService",
	HandlerType: (*CreateEmpServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _CreateEmpService_Create_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _CreateEmpService_Get_Handler,
		},
		{
			MethodName: "Edit",
			Handler:    _CreateEmpService_Edit_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "data/emp_request.proto",
}
