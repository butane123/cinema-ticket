// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: order.proto

package order

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

// OrderClient is the client API for Order service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OrderClient interface {
	SetPaid(ctx context.Context, in *SetPaidReq, opts ...grpc.CallOption) (*SetPaidReply, error)
	JudgeExist(ctx context.Context, in *JudgeExistReq, opts ...grpc.CallOption) (*JudgeExistReply, error)
	FindById(ctx context.Context, in *FindReq, opts ...grpc.CallOption) (*FindReply, error)
	Create(ctx context.Context, in *CreateReq, opts ...grpc.CallOption) (*CreateReply, error)
	CreateRevert(ctx context.Context, in *CreateReq, opts ...grpc.CallOption) (*CreateReply, error)
}

type orderClient struct {
	cc grpc.ClientConnInterface
}

func NewOrderClient(cc grpc.ClientConnInterface) OrderClient {
	return &orderClient{cc}
}

func (c *orderClient) SetPaid(ctx context.Context, in *SetPaidReq, opts ...grpc.CallOption) (*SetPaidReply, error) {
	out := new(SetPaidReply)
	err := c.cc.Invoke(ctx, "/order.order/setPaid", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderClient) JudgeExist(ctx context.Context, in *JudgeExistReq, opts ...grpc.CallOption) (*JudgeExistReply, error) {
	out := new(JudgeExistReply)
	err := c.cc.Invoke(ctx, "/order.order/judgeExist", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderClient) FindById(ctx context.Context, in *FindReq, opts ...grpc.CallOption) (*FindReply, error) {
	out := new(FindReply)
	err := c.cc.Invoke(ctx, "/order.order/findById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderClient) Create(ctx context.Context, in *CreateReq, opts ...grpc.CallOption) (*CreateReply, error) {
	out := new(CreateReply)
	err := c.cc.Invoke(ctx, "/order.order/create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderClient) CreateRevert(ctx context.Context, in *CreateReq, opts ...grpc.CallOption) (*CreateReply, error) {
	out := new(CreateReply)
	err := c.cc.Invoke(ctx, "/order.order/createRevert", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OrderServer is the server API for Order service.
// All implementations must embed UnimplementedOrderServer
// for forward compatibility
type OrderServer interface {
	SetPaid(context.Context, *SetPaidReq) (*SetPaidReply, error)
	JudgeExist(context.Context, *JudgeExistReq) (*JudgeExistReply, error)
	FindById(context.Context, *FindReq) (*FindReply, error)
	Create(context.Context, *CreateReq) (*CreateReply, error)
	CreateRevert(context.Context, *CreateReq) (*CreateReply, error)
	mustEmbedUnimplementedOrderServer()
}

// UnimplementedOrderServer must be embedded to have forward compatible implementations.
type UnimplementedOrderServer struct {
}

func (UnimplementedOrderServer) SetPaid(context.Context, *SetPaidReq) (*SetPaidReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetPaid not implemented")
}
func (UnimplementedOrderServer) JudgeExist(context.Context, *JudgeExistReq) (*JudgeExistReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method JudgeExist not implemented")
}
func (UnimplementedOrderServer) FindById(context.Context, *FindReq) (*FindReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindById not implemented")
}
func (UnimplementedOrderServer) Create(context.Context, *CreateReq) (*CreateReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedOrderServer) CreateRevert(context.Context, *CreateReq) (*CreateReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateRevert not implemented")
}
func (UnimplementedOrderServer) mustEmbedUnimplementedOrderServer() {}

// UnsafeOrderServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OrderServer will
// result in compilation errors.
type UnsafeOrderServer interface {
	mustEmbedUnimplementedOrderServer()
}

func RegisterOrderServer(s grpc.ServiceRegistrar, srv OrderServer) {
	s.RegisterService(&Order_ServiceDesc, srv)
}

func _Order_SetPaid_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetPaidReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServer).SetPaid(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/order.order/setPaid",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServer).SetPaid(ctx, req.(*SetPaidReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Order_JudgeExist_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JudgeExistReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServer).JudgeExist(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/order.order/judgeExist",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServer).JudgeExist(ctx, req.(*JudgeExistReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Order_FindById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServer).FindById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/order.order/findById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServer).FindById(ctx, req.(*FindReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Order_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/order.order/create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServer).Create(ctx, req.(*CreateReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Order_CreateRevert_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServer).CreateRevert(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/order.order/createRevert",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServer).CreateRevert(ctx, req.(*CreateReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Order_ServiceDesc is the grpc.ServiceDesc for Order service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Order_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "order.order",
	HandlerType: (*OrderServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "setPaid",
			Handler:    _Order_SetPaid_Handler,
		},
		{
			MethodName: "judgeExist",
			Handler:    _Order_JudgeExist_Handler,
		},
		{
			MethodName: "findById",
			Handler:    _Order_FindById_Handler,
		},
		{
			MethodName: "create",
			Handler:    _Order_Create_Handler,
		},
		{
			MethodName: "createRevert",
			Handler:    _Order_CreateRevert_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "order.proto",
}
