// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.12.4
// source: administration_service.proto

package users_service

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

const (
	AdministrationService_Create_FullMethodName  = "/users_service.AdministrationService/Create"
	AdministrationService_GetById_FullMethodName = "/users_service.AdministrationService/GetById"
	AdministrationService_GetList_FullMethodName = "/users_service.AdministrationService/GetList"
	AdministrationService_Update_FullMethodName  = "/users_service.AdministrationService/Update"
	AdministrationService_Delete_FullMethodName  = "/users_service.AdministrationService/Delete"
)

// AdministrationServiceClient is the client API for AdministrationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AdministrationServiceClient interface {
	Create(ctx context.Context, in *CreateAdministration, opts ...grpc.CallOption) (*Administration, error)
	GetById(ctx context.Context, in *AdministrationPrimaryKey, opts ...grpc.CallOption) (*Administration, error)
	GetList(ctx context.Context, in *GetListAdministrationRequest, opts ...grpc.CallOption) (*GetListAdministrationResponse, error)
	Update(ctx context.Context, in *UpdateAdministration, opts ...grpc.CallOption) (*Administration, error)
	Delete(ctx context.Context, in *AdministrationPrimaryKey, opts ...grpc.CallOption) (*Empty, error)
}

type administrationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAdministrationServiceClient(cc grpc.ClientConnInterface) AdministrationServiceClient {
	return &administrationServiceClient{cc}
}

func (c *administrationServiceClient) Create(ctx context.Context, in *CreateAdministration, opts ...grpc.CallOption) (*Administration, error) {
	out := new(Administration)
	err := c.cc.Invoke(ctx, AdministrationService_Create_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *administrationServiceClient) GetById(ctx context.Context, in *AdministrationPrimaryKey, opts ...grpc.CallOption) (*Administration, error) {
	out := new(Administration)
	err := c.cc.Invoke(ctx, AdministrationService_GetById_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *administrationServiceClient) GetList(ctx context.Context, in *GetListAdministrationRequest, opts ...grpc.CallOption) (*GetListAdministrationResponse, error) {
	out := new(GetListAdministrationResponse)
	err := c.cc.Invoke(ctx, AdministrationService_GetList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *administrationServiceClient) Update(ctx context.Context, in *UpdateAdministration, opts ...grpc.CallOption) (*Administration, error) {
	out := new(Administration)
	err := c.cc.Invoke(ctx, AdministrationService_Update_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *administrationServiceClient) Delete(ctx context.Context, in *AdministrationPrimaryKey, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, AdministrationService_Delete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AdministrationServiceServer is the server API for AdministrationService service.
// All implementations must embed UnimplementedAdministrationServiceServer
// for forward compatibility
type AdministrationServiceServer interface {
	Create(context.Context, *CreateAdministration) (*Administration, error)
	GetById(context.Context, *AdministrationPrimaryKey) (*Administration, error)
	GetList(context.Context, *GetListAdministrationRequest) (*GetListAdministrationResponse, error)
	Update(context.Context, *UpdateAdministration) (*Administration, error)
	Delete(context.Context, *AdministrationPrimaryKey) (*Empty, error)
	mustEmbedUnimplementedAdministrationServiceServer()
}

// UnimplementedAdministrationServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAdministrationServiceServer struct {
}

func (UnimplementedAdministrationServiceServer) Create(context.Context, *CreateAdministration) (*Administration, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedAdministrationServiceServer) GetById(context.Context, *AdministrationPrimaryKey) (*Administration, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetById not implemented")
}
func (UnimplementedAdministrationServiceServer) GetList(context.Context, *GetListAdministrationRequest) (*GetListAdministrationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetList not implemented")
}
func (UnimplementedAdministrationServiceServer) Update(context.Context, *UpdateAdministration) (*Administration, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedAdministrationServiceServer) Delete(context.Context, *AdministrationPrimaryKey) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedAdministrationServiceServer) mustEmbedUnimplementedAdministrationServiceServer() {}

// UnsafeAdministrationServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AdministrationServiceServer will
// result in compilation errors.
type UnsafeAdministrationServiceServer interface {
	mustEmbedUnimplementedAdministrationServiceServer()
}

func RegisterAdministrationServiceServer(s grpc.ServiceRegistrar, srv AdministrationServiceServer) {
	s.RegisterService(&AdministrationService_ServiceDesc, srv)
}

func _AdministrationService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAdministration)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdministrationServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AdministrationService_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdministrationServiceServer).Create(ctx, req.(*CreateAdministration))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdministrationService_GetById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AdministrationPrimaryKey)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdministrationServiceServer).GetById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AdministrationService_GetById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdministrationServiceServer).GetById(ctx, req.(*AdministrationPrimaryKey))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdministrationService_GetList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetListAdministrationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdministrationServiceServer).GetList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AdministrationService_GetList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdministrationServiceServer).GetList(ctx, req.(*GetListAdministrationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdministrationService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateAdministration)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdministrationServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AdministrationService_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdministrationServiceServer).Update(ctx, req.(*UpdateAdministration))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdministrationService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AdministrationPrimaryKey)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdministrationServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AdministrationService_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdministrationServiceServer).Delete(ctx, req.(*AdministrationPrimaryKey))
	}
	return interceptor(ctx, in, info, handler)
}

// AdministrationService_ServiceDesc is the grpc.ServiceDesc for AdministrationService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AdministrationService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "users_service.AdministrationService",
	HandlerType: (*AdministrationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _AdministrationService_Create_Handler,
		},
		{
			MethodName: "GetById",
			Handler:    _AdministrationService_GetById_Handler,
		},
		{
			MethodName: "GetList",
			Handler:    _AdministrationService_GetList_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _AdministrationService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _AdministrationService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "administration_service.proto",
}
