// Code generated by protoc-gen-go.
// source: binlogservice.proto
// DO NOT EDIT!

/*
Package binlogservice is a generated protocol buffer package.

It is generated from these files:
	binlogservice.proto

It has these top-level messages:
*/
package binlogservice

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import binlogdata "github.com/youtube/vitess/go/vt/proto/binlogdata"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for UpdateStream service

type UpdateStreamClient interface {
	// StreamUpdate streams the binlog events, to know which objects have changed.
	StreamUpdate(ctx context.Context, in *binlogdata.StreamUpdateRequest, opts ...grpc.CallOption) (UpdateStream_StreamUpdateClient, error)
	// StreamKeyRange returns the binlog transactions related to
	// the specified Keyrange.
	StreamKeyRange(ctx context.Context, in *binlogdata.StreamKeyRangeRequest, opts ...grpc.CallOption) (UpdateStream_StreamKeyRangeClient, error)
	// StreamTables returns the binlog transactions related to
	// the specified Tables.
	StreamTables(ctx context.Context, in *binlogdata.StreamTablesRequest, opts ...grpc.CallOption) (UpdateStream_StreamTablesClient, error)
}

type updateStreamClient struct {
	cc *grpc.ClientConn
}

func NewUpdateStreamClient(cc *grpc.ClientConn) UpdateStreamClient {
	return &updateStreamClient{cc}
}

func (c *updateStreamClient) StreamUpdate(ctx context.Context, in *binlogdata.StreamUpdateRequest, opts ...grpc.CallOption) (UpdateStream_StreamUpdateClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_UpdateStream_serviceDesc.Streams[0], c.cc, "/binlogservice.UpdateStream/StreamUpdate", opts...)
	if err != nil {
		return nil, err
	}
	x := &updateStreamStreamUpdateClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type UpdateStream_StreamUpdateClient interface {
	Recv() (*binlogdata.StreamUpdateResponse, error)
	grpc.ClientStream
}

type updateStreamStreamUpdateClient struct {
	grpc.ClientStream
}

func (x *updateStreamStreamUpdateClient) Recv() (*binlogdata.StreamUpdateResponse, error) {
	m := new(binlogdata.StreamUpdateResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *updateStreamClient) StreamKeyRange(ctx context.Context, in *binlogdata.StreamKeyRangeRequest, opts ...grpc.CallOption) (UpdateStream_StreamKeyRangeClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_UpdateStream_serviceDesc.Streams[1], c.cc, "/binlogservice.UpdateStream/StreamKeyRange", opts...)
	if err != nil {
		return nil, err
	}
	x := &updateStreamStreamKeyRangeClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type UpdateStream_StreamKeyRangeClient interface {
	Recv() (*binlogdata.StreamKeyRangeResponse, error)
	grpc.ClientStream
}

type updateStreamStreamKeyRangeClient struct {
	grpc.ClientStream
}

func (x *updateStreamStreamKeyRangeClient) Recv() (*binlogdata.StreamKeyRangeResponse, error) {
	m := new(binlogdata.StreamKeyRangeResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *updateStreamClient) StreamTables(ctx context.Context, in *binlogdata.StreamTablesRequest, opts ...grpc.CallOption) (UpdateStream_StreamTablesClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_UpdateStream_serviceDesc.Streams[2], c.cc, "/binlogservice.UpdateStream/StreamTables", opts...)
	if err != nil {
		return nil, err
	}
	x := &updateStreamStreamTablesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type UpdateStream_StreamTablesClient interface {
	Recv() (*binlogdata.StreamTablesResponse, error)
	grpc.ClientStream
}

type updateStreamStreamTablesClient struct {
	grpc.ClientStream
}

func (x *updateStreamStreamTablesClient) Recv() (*binlogdata.StreamTablesResponse, error) {
	m := new(binlogdata.StreamTablesResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for UpdateStream service

type UpdateStreamServer interface {
	// StreamUpdate streams the binlog events, to know which objects have changed.
	StreamUpdate(*binlogdata.StreamUpdateRequest, UpdateStream_StreamUpdateServer) error
	// StreamKeyRange returns the binlog transactions related to
	// the specified Keyrange.
	StreamKeyRange(*binlogdata.StreamKeyRangeRequest, UpdateStream_StreamKeyRangeServer) error
	// StreamTables returns the binlog transactions related to
	// the specified Tables.
	StreamTables(*binlogdata.StreamTablesRequest, UpdateStream_StreamTablesServer) error
}

func RegisterUpdateStreamServer(s *grpc.Server, srv UpdateStreamServer) {
	s.RegisterService(&_UpdateStream_serviceDesc, srv)
}

func _UpdateStream_StreamUpdate_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(binlogdata.StreamUpdateRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(UpdateStreamServer).StreamUpdate(m, &updateStreamStreamUpdateServer{stream})
}

type UpdateStream_StreamUpdateServer interface {
	Send(*binlogdata.StreamUpdateResponse) error
	grpc.ServerStream
}

type updateStreamStreamUpdateServer struct {
	grpc.ServerStream
}

func (x *updateStreamStreamUpdateServer) Send(m *binlogdata.StreamUpdateResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _UpdateStream_StreamKeyRange_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(binlogdata.StreamKeyRangeRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(UpdateStreamServer).StreamKeyRange(m, &updateStreamStreamKeyRangeServer{stream})
}

type UpdateStream_StreamKeyRangeServer interface {
	Send(*binlogdata.StreamKeyRangeResponse) error
	grpc.ServerStream
}

type updateStreamStreamKeyRangeServer struct {
	grpc.ServerStream
}

func (x *updateStreamStreamKeyRangeServer) Send(m *binlogdata.StreamKeyRangeResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _UpdateStream_StreamTables_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(binlogdata.StreamTablesRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(UpdateStreamServer).StreamTables(m, &updateStreamStreamTablesServer{stream})
}

type UpdateStream_StreamTablesServer interface {
	Send(*binlogdata.StreamTablesResponse) error
	grpc.ServerStream
}

type updateStreamStreamTablesServer struct {
	grpc.ServerStream
}

func (x *updateStreamStreamTablesServer) Send(m *binlogdata.StreamTablesResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _UpdateStream_serviceDesc = grpc.ServiceDesc{
	ServiceName: "binlogservice.UpdateStream",
	HandlerType: (*UpdateStreamServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamUpdate",
			Handler:       _UpdateStream_StreamUpdate_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "StreamKeyRange",
			Handler:       _UpdateStream_StreamKeyRange_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "StreamTables",
			Handler:       _UpdateStream_StreamTables_Handler,
			ServerStreams: true,
		},
	},
	Metadata: fileDescriptor0,
}

func init() { proto.RegisterFile("binlogservice.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 163 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x12, 0x4e, 0xca, 0xcc, 0xcb,
	0xc9, 0x4f, 0x2f, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17,
	0xe2, 0x45, 0x11, 0x94, 0x12, 0x80, 0x70, 0x53, 0x12, 0x4b, 0x12, 0x21, 0x0a, 0x8c, 0x66, 0x32,
	0x71, 0xf1, 0x84, 0x16, 0x00, 0x05, 0x52, 0x83, 0x4b, 0x8a, 0x52, 0x13, 0x73, 0x85, 0x42, 0xb9,
	0x78, 0x20, 0x2c, 0x88, 0xa8, 0x90, 0xbc, 0x1e, 0x92, 0x1e, 0x64, 0x99, 0xa0, 0xd4, 0xc2, 0xd2,
	0xd4, 0xe2, 0x12, 0x29, 0x05, 0xdc, 0x0a, 0x8a, 0x0b, 0xf2, 0xf3, 0x8a, 0x53, 0x95, 0x18, 0x0c,
	0x18, 0x85, 0xa2, 0xb9, 0xf8, 0x20, 0x72, 0xde, 0xa9, 0x95, 0x41, 0x89, 0x79, 0xe9, 0xa9, 0x42,
	0x8a, 0x98, 0xfa, 0x60, 0x72, 0x30, 0xa3, 0x95, 0xf0, 0x29, 0x41, 0x32, 0x1c, 0xee, 0xe6, 0x90,
	0xc4, 0xa4, 0x9c, 0xd4, 0x62, 0x6c, 0x6e, 0x86, 0xc8, 0xe0, 0x71, 0x33, 0x4c, 0x01, 0xc2, 0xd8,
	0x24, 0x36, 0x70, 0x10, 0x19, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x4c, 0x73, 0x00, 0xff, 0x5a,
	0x01, 0x00, 0x00,
}
