// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: b.proto

package pb

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/gogo/protobuf/proto"
	types "github.com/gogo/protobuf/types"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

func init() { proto.RegisterFile("b.proto", fileDescriptor_1848ba9bd6296c7e) }

var fileDescriptor_1848ba9bd6296c7e = []byte{
	// 149 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x62, 0x4f, 0xd2, 0x2b, 0x28,
	0xca, 0x2f, 0xc9, 0x17, 0x62, 0x29, 0x49, 0x2d, 0x2e, 0x91, 0x52, 0x2f, 0xc9, 0xc8, 0x2c, 0x4a,
	0x89, 0x2f, 0x48, 0x2c, 0x2a, 0xa9, 0xd4, 0x4f, 0xcf, 0xcf, 0x4f, 0xcf, 0x49, 0xd5, 0x07, 0xcb,
	0x27, 0x95, 0xa6, 0xe9, 0xa7, 0xe6, 0x16, 0x94, 0x54, 0x42, 0x94, 0x4b, 0x69, 0x60, 0x51, 0x98,
	0x58, 0x90, 0xa9, 0x9f, 0x98, 0x97, 0x97, 0x5f, 0x92, 0x58, 0x92, 0x99, 0x9f, 0x57, 0x0c, 0x51,
	0x69, 0xe4, 0xc5, 0xc5, 0xe8, 0x24, 0xe4, 0xca, 0xc5, 0xee, 0x9b, 0x5a, 0x92, 0x91, 0x9f, 0xe2,
	0x24, 0x24, 0xa6, 0x07, 0x51, 0xae, 0x07, 0x33, 0x57, 0xcf, 0x15, 0x64, 0xae, 0x14, 0x0e, 0x71,
	0x25, 0xae, 0xa6, 0xcb, 0x4f, 0x26, 0x33, 0xb1, 0x08, 0x31, 0xe9, 0x27, 0x39, 0xb1, 0x44, 0x31,
	0x15, 0x24, 0x25, 0xb1, 0x81, 0x55, 0x18, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0xf9, 0x0f, 0xdd,
	0x9d, 0xbc, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// BClient is the client API for B service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type BClient interface {
	// here is MethodB comment
	MethodB(ctx context.Context, in *types.Empty, opts ...grpc.CallOption) (*types.Empty, error)
}

type bClient struct {
	cc *grpc.ClientConn
}

func NewBClient(cc *grpc.ClientConn) BClient {
	return &bClient{cc}
}

func (c *bClient) MethodB(ctx context.Context, in *types.Empty, opts ...grpc.CallOption) (*types.Empty, error) {
	out := new(types.Empty)
	err := c.cc.Invoke(ctx, "/test.B/MethodB", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BServer is the server API for B service.
type BServer interface {
	// here is MethodB comment
	MethodB(context.Context, *types.Empty) (*types.Empty, error)
}

// UnimplementedBServer can be embedded to have forward compatible implementations.
type UnimplementedBServer struct {
}

func (*UnimplementedBServer) MethodB(ctx context.Context, req *types.Empty) (*types.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MethodB not implemented")
}

func RegisterBServer(s *grpc.Server, srv BServer) {
	s.RegisterService(&_B_serviceDesc, srv)
}

func _B_MethodB_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(types.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BServer).MethodB(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/test.B/MethodB",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BServer).MethodB(ctx, req.(*types.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _B_serviceDesc = grpc.ServiceDesc{
	ServiceName: "test.B",
	HandlerType: (*BServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "MethodB",
			Handler:    _B_MethodB_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "b.proto",
}
