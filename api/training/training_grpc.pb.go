// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: api/training/training.proto

package training

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

// TrainingClient is the client API for Training service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TrainingClient interface {
	CreateTraining(ctx context.Context, in *CreateTrainingRequest, opts ...grpc.CallOption) (*CreateTrainingReply, error)
	UpdateTraining(ctx context.Context, in *UpdateTrainingRequest, opts ...grpc.CallOption) (*UpdateTrainingReply, error)
	DeleteTraining(ctx context.Context, in *DeleteTrainingRequest, opts ...grpc.CallOption) (*DeleteTrainingReply, error)
	GetTraining(ctx context.Context, in *GetTrainingRequest, opts ...grpc.CallOption) (*GetTrainingReply, error)
	ListTraining(ctx context.Context, in *ListTrainingRequest, opts ...grpc.CallOption) (*ListTrainingReply, error)
}

type trainingClient struct {
	cc grpc.ClientConnInterface
}

func NewTrainingClient(cc grpc.ClientConnInterface) TrainingClient {
	return &trainingClient{cc}
}

func (c *trainingClient) CreateTraining(ctx context.Context, in *CreateTrainingRequest, opts ...grpc.CallOption) (*CreateTrainingReply, error) {
	out := new(CreateTrainingReply)
	err := c.cc.Invoke(ctx, "/api.training.Training/CreateTraining", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *trainingClient) UpdateTraining(ctx context.Context, in *UpdateTrainingRequest, opts ...grpc.CallOption) (*UpdateTrainingReply, error) {
	out := new(UpdateTrainingReply)
	err := c.cc.Invoke(ctx, "/api.training.Training/UpdateTraining", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *trainingClient) DeleteTraining(ctx context.Context, in *DeleteTrainingRequest, opts ...grpc.CallOption) (*DeleteTrainingReply, error) {
	out := new(DeleteTrainingReply)
	err := c.cc.Invoke(ctx, "/api.training.Training/DeleteTraining", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *trainingClient) GetTraining(ctx context.Context, in *GetTrainingRequest, opts ...grpc.CallOption) (*GetTrainingReply, error) {
	out := new(GetTrainingReply)
	err := c.cc.Invoke(ctx, "/api.training.Training/GetTraining", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *trainingClient) ListTraining(ctx context.Context, in *ListTrainingRequest, opts ...grpc.CallOption) (*ListTrainingReply, error) {
	out := new(ListTrainingReply)
	err := c.cc.Invoke(ctx, "/api.training.Training/ListTraining", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TrainingServer is the server API for Training service.
// All implementations must embed UnimplementedTrainingServer
// for forward compatibility
type TrainingServer interface {
	CreateTraining(context.Context, *CreateTrainingRequest) (*CreateTrainingReply, error)
	UpdateTraining(context.Context, *UpdateTrainingRequest) (*UpdateTrainingReply, error)
	DeleteTraining(context.Context, *DeleteTrainingRequest) (*DeleteTrainingReply, error)
	GetTraining(context.Context, *GetTrainingRequest) (*GetTrainingReply, error)
	ListTraining(context.Context, *ListTrainingRequest) (*ListTrainingReply, error)
	mustEmbedUnimplementedTrainingServer()
}

// UnimplementedTrainingServer must be embedded to have forward compatible implementations.
type UnimplementedTrainingServer struct {
}

func (UnimplementedTrainingServer) CreateTraining(context.Context, *CreateTrainingRequest) (*CreateTrainingReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTraining not implemented")
}
func (UnimplementedTrainingServer) UpdateTraining(context.Context, *UpdateTrainingRequest) (*UpdateTrainingReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateTraining not implemented")
}
func (UnimplementedTrainingServer) DeleteTraining(context.Context, *DeleteTrainingRequest) (*DeleteTrainingReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTraining not implemented")
}
func (UnimplementedTrainingServer) GetTraining(context.Context, *GetTrainingRequest) (*GetTrainingReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTraining not implemented")
}
func (UnimplementedTrainingServer) ListTraining(context.Context, *ListTrainingRequest) (*ListTrainingReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListTraining not implemented")
}
func (UnimplementedTrainingServer) mustEmbedUnimplementedTrainingServer() {}

// UnsafeTrainingServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TrainingServer will
// result in compilation errors.
type UnsafeTrainingServer interface {
	mustEmbedUnimplementedTrainingServer()
}

func RegisterTrainingServer(s grpc.ServiceRegistrar, srv TrainingServer) {
	s.RegisterService(&Training_ServiceDesc, srv)
}

func _Training_CreateTraining_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTrainingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TrainingServer).CreateTraining(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.training.Training/CreateTraining",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TrainingServer).CreateTraining(ctx, req.(*CreateTrainingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Training_UpdateTraining_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateTrainingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TrainingServer).UpdateTraining(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.training.Training/UpdateTraining",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TrainingServer).UpdateTraining(ctx, req.(*UpdateTrainingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Training_DeleteTraining_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteTrainingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TrainingServer).DeleteTraining(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.training.Training/DeleteTraining",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TrainingServer).DeleteTraining(ctx, req.(*DeleteTrainingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Training_GetTraining_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTrainingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TrainingServer).GetTraining(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.training.Training/GetTraining",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TrainingServer).GetTraining(ctx, req.(*GetTrainingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Training_ListTraining_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListTrainingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TrainingServer).ListTraining(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.training.Training/ListTraining",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TrainingServer).ListTraining(ctx, req.(*ListTrainingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Training_ServiceDesc is the grpc.ServiceDesc for Training service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Training_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.training.Training",
	HandlerType: (*TrainingServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateTraining",
			Handler:    _Training_CreateTraining_Handler,
		},
		{
			MethodName: "UpdateTraining",
			Handler:    _Training_UpdateTraining_Handler,
		},
		{
			MethodName: "DeleteTraining",
			Handler:    _Training_DeleteTraining_Handler,
		},
		{
			MethodName: "GetTraining",
			Handler:    _Training_GetTraining_Handler,
		},
		{
			MethodName: "ListTraining",
			Handler:    _Training_ListTraining_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/training/training.proto",
}
