// Code generated by protoc-gen-go. DO NOT EDIT.
// source: trillian_log_sequencer_api.proto

package trillian

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

func init() { proto.RegisterFile("trillian_log_sequencer_api.proto", fileDescriptor_f32c68ea33658ef4) }

var fileDescriptor_f32c68ea33658ef4 = []byte{
	// 129 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x28, 0x29, 0xca, 0xcc,
	0xc9, 0xc9, 0x4c, 0xcc, 0x8b, 0xcf, 0xc9, 0x4f, 0x8f, 0x2f, 0x4e, 0x2d, 0x2c, 0x4d, 0xcd, 0x4b,
	0x4e, 0x2d, 0x8a, 0x4f, 0x2c, 0xc8, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x80, 0xa9,
	0x30, 0x12, 0xe3, 0x12, 0x09, 0x81, 0xb2, 0x7d, 0xf2, 0xd3, 0x83, 0x61, 0x6a, 0x9d, 0xc2, 0xb9,
	0x24, 0x93, 0xf3, 0x73, 0xf5, 0xd2, 0xf3, 0xf3, 0xd3, 0x73, 0x52, 0xf5, 0x60, 0xca, 0x21, 0xda,
	0x9d, 0x64, 0xb0, 0x69, 0x71, 0x2c, 0xc8, 0x0c, 0x00, 0xc9, 0x06, 0x30, 0x46, 0x49, 0xa5, 0x67,
	0x96, 0x64, 0x94, 0x26, 0xe9, 0x25, 0xe7, 0xe7, 0xea, 0x43, 0x4c, 0xd0, 0x87, 0x99, 0x90, 0xc4,
	0x06, 0x36, 0xc2, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x86, 0xa5, 0x47, 0x41, 0xa5, 0x00, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// TrillianLogSequencerClient is the client API for TrillianLogSequencer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TrillianLogSequencerClient interface {
}

type trillianLogSequencerClient struct {
	cc *grpc.ClientConn
}

func NewTrillianLogSequencerClient(cc *grpc.ClientConn) TrillianLogSequencerClient {
	return &trillianLogSequencerClient{cc}
}

// TrillianLogSequencerServer is the server API for TrillianLogSequencer service.
type TrillianLogSequencerServer interface {
}

func RegisterTrillianLogSequencerServer(s *grpc.Server, srv TrillianLogSequencerServer) {
	s.RegisterService(&_TrillianLogSequencer_serviceDesc, srv)
}

var _TrillianLogSequencer_serviceDesc = grpc.ServiceDesc{
	ServiceName: "trillian.TrillianLogSequencer",
	HandlerType: (*TrillianLogSequencerServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams:     []grpc.StreamDesc{},
	Metadata:    "trillian_log_sequencer_api.proto",
}
