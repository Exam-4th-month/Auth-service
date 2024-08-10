// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v3.21.12
// source: goal-service/goal-service.proto

package goal

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	GoalService_CreateGoal_FullMethodName  = "/goal.GoalService/CreateGoal"
	GoalService_GetGoals_FullMethodName    = "/goal.GoalService/GetGoals"
	GoalService_GetGoalById_FullMethodName = "/goal.GoalService/GetGoalById"
	GoalService_UpdateGoal_FullMethodName  = "/goal.GoalService/UpdateGoal"
	GoalService_DeleteGoal_FullMethodName  = "/goal.GoalService/DeleteGoal"
)

// GoalServiceClient is the client API for GoalService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GoalServiceClient interface {
	CreateGoal(ctx context.Context, in *CreateGoalRequest, opts ...grpc.CallOption) (*GoalResponse, error)
	GetGoals(ctx context.Context, in *GetGoalsRequest, opts ...grpc.CallOption) (*GoalsResponse, error)
	GetGoalById(ctx context.Context, in *GetGoalByIdRequest, opts ...grpc.CallOption) (*GoalResponse, error)
	UpdateGoal(ctx context.Context, in *UpdateGoalRequest, opts ...grpc.CallOption) (*GoalResponse, error)
	DeleteGoal(ctx context.Context, in *DeleteGoalRequest, opts ...grpc.CallOption) (*Empty, error)
}

type goalServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewGoalServiceClient(cc grpc.ClientConnInterface) GoalServiceClient {
	return &goalServiceClient{cc}
}

func (c *goalServiceClient) CreateGoal(ctx context.Context, in *CreateGoalRequest, opts ...grpc.CallOption) (*GoalResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GoalResponse)
	err := c.cc.Invoke(ctx, GoalService_CreateGoal_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *goalServiceClient) GetGoals(ctx context.Context, in *GetGoalsRequest, opts ...grpc.CallOption) (*GoalsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GoalsResponse)
	err := c.cc.Invoke(ctx, GoalService_GetGoals_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *goalServiceClient) GetGoalById(ctx context.Context, in *GetGoalByIdRequest, opts ...grpc.CallOption) (*GoalResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GoalResponse)
	err := c.cc.Invoke(ctx, GoalService_GetGoalById_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *goalServiceClient) UpdateGoal(ctx context.Context, in *UpdateGoalRequest, opts ...grpc.CallOption) (*GoalResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GoalResponse)
	err := c.cc.Invoke(ctx, GoalService_UpdateGoal_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *goalServiceClient) DeleteGoal(ctx context.Context, in *DeleteGoalRequest, opts ...grpc.CallOption) (*Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Empty)
	err := c.cc.Invoke(ctx, GoalService_DeleteGoal_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GoalServiceServer is the server API for GoalService service.
// All implementations must embed UnimplementedGoalServiceServer
// for forward compatibility
type GoalServiceServer interface {
	CreateGoal(context.Context, *CreateGoalRequest) (*GoalResponse, error)
	GetGoals(context.Context, *GetGoalsRequest) (*GoalsResponse, error)
	GetGoalById(context.Context, *GetGoalByIdRequest) (*GoalResponse, error)
	UpdateGoal(context.Context, *UpdateGoalRequest) (*GoalResponse, error)
	DeleteGoal(context.Context, *DeleteGoalRequest) (*Empty, error)
	mustEmbedUnimplementedGoalServiceServer()
}

// UnimplementedGoalServiceServer must be embedded to have forward compatible implementations.
type UnimplementedGoalServiceServer struct {
}

func (UnimplementedGoalServiceServer) CreateGoal(context.Context, *CreateGoalRequest) (*GoalResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateGoal not implemented")
}
func (UnimplementedGoalServiceServer) GetGoals(context.Context, *GetGoalsRequest) (*GoalsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGoals not implemented")
}
func (UnimplementedGoalServiceServer) GetGoalById(context.Context, *GetGoalByIdRequest) (*GoalResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGoalById not implemented")
}
func (UnimplementedGoalServiceServer) UpdateGoal(context.Context, *UpdateGoalRequest) (*GoalResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateGoal not implemented")
}
func (UnimplementedGoalServiceServer) DeleteGoal(context.Context, *DeleteGoalRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteGoal not implemented")
}
func (UnimplementedGoalServiceServer) mustEmbedUnimplementedGoalServiceServer() {}

// UnsafeGoalServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GoalServiceServer will
// result in compilation errors.
type UnsafeGoalServiceServer interface {
	mustEmbedUnimplementedGoalServiceServer()
}

func RegisterGoalServiceServer(s grpc.ServiceRegistrar, srv GoalServiceServer) {
	s.RegisterService(&GoalService_ServiceDesc, srv)
}

func _GoalService_CreateGoal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateGoalRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoalServiceServer).CreateGoal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GoalService_CreateGoal_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoalServiceServer).CreateGoal(ctx, req.(*CreateGoalRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GoalService_GetGoals_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetGoalsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoalServiceServer).GetGoals(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GoalService_GetGoals_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoalServiceServer).GetGoals(ctx, req.(*GetGoalsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GoalService_GetGoalById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetGoalByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoalServiceServer).GetGoalById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GoalService_GetGoalById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoalServiceServer).GetGoalById(ctx, req.(*GetGoalByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GoalService_UpdateGoal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateGoalRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoalServiceServer).UpdateGoal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GoalService_UpdateGoal_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoalServiceServer).UpdateGoal(ctx, req.(*UpdateGoalRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GoalService_DeleteGoal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteGoalRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoalServiceServer).DeleteGoal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GoalService_DeleteGoal_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoalServiceServer).DeleteGoal(ctx, req.(*DeleteGoalRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// GoalService_ServiceDesc is the grpc.ServiceDesc for GoalService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GoalService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "goal.GoalService",
	HandlerType: (*GoalServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateGoal",
			Handler:    _GoalService_CreateGoal_Handler,
		},
		{
			MethodName: "GetGoals",
			Handler:    _GoalService_GetGoals_Handler,
		},
		{
			MethodName: "GetGoalById",
			Handler:    _GoalService_GetGoalById_Handler,
		},
		{
			MethodName: "UpdateGoal",
			Handler:    _GoalService_UpdateGoal_Handler,
		},
		{
			MethodName: "DeleteGoal",
			Handler:    _GoalService_DeleteGoal_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "goal-service/goal-service.proto",
}
