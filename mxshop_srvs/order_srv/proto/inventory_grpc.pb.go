// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v5.27.1
// source: inventory.proto

package proto

import (
	context "context"

	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	Inventory_SetStocks_FullMethodName = "/Inventory/SetStocks"
	Inventory_InvDetail_FullMethodName = "/Inventory/InvDetail"
	Inventory_Sell_FullMethodName      = "/Inventory/Sell"
	Inventory_ReBack_FullMethodName    = "/Inventory/ReBack"
)

// InventoryClient is the client API for Inventory service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type InventoryClient interface {
	SetStocks(ctx context.Context, in *GoodsInvInfo, opts ...grpc.CallOption) (*emptypb.Empty, error)
	InvDetail(ctx context.Context, in *GoodsInvInfo, opts ...grpc.CallOption) (*GoodsInvInfo, error)
	Sell(ctx context.Context, in *SellInfo, opts ...grpc.CallOption) (*emptypb.Empty, error)
	ReBack(ctx context.Context, in *SellInfo, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type inventoryClient struct {
	cc grpc.ClientConnInterface
}

func NewInventoryClient(cc grpc.ClientConnInterface) InventoryClient {
	return &inventoryClient{cc}
}

func (c *inventoryClient) SetStocks(ctx context.Context, in *GoodsInvInfo, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, Inventory_SetStocks_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *inventoryClient) InvDetail(ctx context.Context, in *GoodsInvInfo, opts ...grpc.CallOption) (*GoodsInvInfo, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GoodsInvInfo)
	err := c.cc.Invoke(ctx, Inventory_InvDetail_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *inventoryClient) Sell(ctx context.Context, in *SellInfo, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, Inventory_Sell_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *inventoryClient) ReBack(ctx context.Context, in *SellInfo, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, Inventory_ReBack_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// InventoryServer is the server API for Inventory service.
// All implementations must embed UnimplementedInventoryServer
// for forward compatibility
type InventoryServer interface {
	SetStocks(context.Context, *GoodsInvInfo) (*emptypb.Empty, error)
	InvDetail(context.Context, *GoodsInvInfo) (*GoodsInvInfo, error)
	Sell(context.Context, *SellInfo) (*emptypb.Empty, error)
	ReBack(context.Context, *SellInfo) (*emptypb.Empty, error)
	mustEmbedUnimplementedInventoryServer()
}

// UnimplementedInventoryServer must be embedded to have forward compatible implementations.
type UnimplementedInventoryServer struct {
}

func (UnimplementedInventoryServer) SetStocks(context.Context, *GoodsInvInfo) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetStocks not implemented")
}
func (UnimplementedInventoryServer) InvDetail(context.Context, *GoodsInvInfo) (*GoodsInvInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InvDetail not implemented")
}
func (UnimplementedInventoryServer) Sell(context.Context, *SellInfo) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Sell not implemented")
}
func (UnimplementedInventoryServer) ReBack(context.Context, *SellInfo) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReBack not implemented")
}
func (UnimplementedInventoryServer) mustEmbedUnimplementedInventoryServer() {}

// UnsafeInventoryServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to InventoryServer will
// result in compilation errors.
type UnsafeInventoryServer interface {
	mustEmbedUnimplementedInventoryServer()
}

func RegisterInventoryServer(s grpc.ServiceRegistrar, srv InventoryServer) {
	s.RegisterService(&Inventory_ServiceDesc, srv)
}

func _Inventory_SetStocks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GoodsInvInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InventoryServer).SetStocks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Inventory_SetStocks_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InventoryServer).SetStocks(ctx, req.(*GoodsInvInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _Inventory_InvDetail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GoodsInvInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InventoryServer).InvDetail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Inventory_InvDetail_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InventoryServer).InvDetail(ctx, req.(*GoodsInvInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _Inventory_Sell_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SellInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InventoryServer).Sell(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Inventory_Sell_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InventoryServer).Sell(ctx, req.(*SellInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _Inventory_ReBack_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SellInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InventoryServer).ReBack(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Inventory_ReBack_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InventoryServer).ReBack(ctx, req.(*SellInfo))
	}
	return interceptor(ctx, in, info, handler)
}

// Inventory_ServiceDesc is the grpc.ServiceDesc for Inventory service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Inventory_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Inventory",
	HandlerType: (*InventoryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SetStocks",
			Handler:    _Inventory_SetStocks_Handler,
		},
		{
			MethodName: "InvDetail",
			Handler:    _Inventory_InvDetail_Handler,
		},
		{
			MethodName: "Sell",
			Handler:    _Inventory_Sell_Handler,
		},
		{
			MethodName: "ReBack",
			Handler:    _Inventory_ReBack_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "inventory.proto",
}