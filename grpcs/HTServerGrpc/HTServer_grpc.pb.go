// Copyright 2015 gRPC authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.3
// source: HTServer.proto

package HTServerGrpc

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
	HTServerService_HandleHTRequest_FullMethodName = "/HTServerGrpc.HTServerService/HandleHTRequest"
)

// HTServerServiceClient is the client API for HTServerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HTServerServiceClient interface {
	HandleHTRequest(ctx context.Context, in *HTRequest, opts ...grpc.CallOption) (*HTReply, error)
}

type hTServerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewHTServerServiceClient(cc grpc.ClientConnInterface) HTServerServiceClient {
	return &hTServerServiceClient{cc}
}

func (c *hTServerServiceClient) HandleHTRequest(ctx context.Context, in *HTRequest, opts ...grpc.CallOption) (*HTReply, error) {
	out := new(HTReply)
	err := c.cc.Invoke(ctx, HTServerService_HandleHTRequest_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HTServerServiceServer is the server API for HTServerService service.
// All implementations must embed UnimplementedHTServerServiceServer
// for forward compatibility
type HTServerServiceServer interface {
	HandleHTRequest(context.Context, *HTRequest) (*HTReply, error)
	mustEmbedUnimplementedHTServerServiceServer()
}

// UnimplementedHTServerServiceServer must be embedded to have forward compatible implementations.
type UnimplementedHTServerServiceServer struct {
}

func (UnimplementedHTServerServiceServer) HandleHTRequest(context.Context, *HTRequest) (*HTReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HandleHTRequest not implemented")
}
func (UnimplementedHTServerServiceServer) mustEmbedUnimplementedHTServerServiceServer() {}

// UnsafeHTServerServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HTServerServiceServer will
// result in compilation errors.
type UnsafeHTServerServiceServer interface {
	mustEmbedUnimplementedHTServerServiceServer()
}

func RegisterHTServerServiceServer(s grpc.ServiceRegistrar, srv HTServerServiceServer) {
	s.RegisterService(&HTServerService_ServiceDesc, srv)
}

func _HTServerService_HandleHTRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HTRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HTServerServiceServer).HandleHTRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HTServerService_HandleHTRequest_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HTServerServiceServer).HandleHTRequest(ctx, req.(*HTRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// HTServerService_ServiceDesc is the grpc.ServiceDesc for HTServerService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var HTServerService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "HTServerGrpc.HTServerService",
	HandlerType: (*HTServerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "HandleHTRequest",
			Handler:    _HTServerService_HandleHTRequest_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "HTServer.proto",
}
