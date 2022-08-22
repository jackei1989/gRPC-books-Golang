// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.1
// source: book.proto

package grpc

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

// BookServiceClient is the client API for BookService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BookServiceClient interface {
	GetBook(ctx context.Context, in *BookRequest, opts ...grpc.CallOption) (*Book, error)
	GetBooks(ctx context.Context, in *BookRequest, opts ...grpc.CallOption) (BookService_GetBooksClient, error)
}

type bookServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBookServiceClient(cc grpc.ClientConnInterface) BookServiceClient {
	return &bookServiceClient{cc}
}

func (c *bookServiceClient) GetBook(ctx context.Context, in *BookRequest, opts ...grpc.CallOption) (*Book, error) {
	out := new(Book)
	err := c.cc.Invoke(ctx, "/grpc.BookService/GetBook", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookServiceClient) GetBooks(ctx context.Context, in *BookRequest, opts ...grpc.CallOption) (BookService_GetBooksClient, error) {
	stream, err := c.cc.NewStream(ctx, &BookService_ServiceDesc.Streams[0], "/grpc.BookService/GetBooks", opts...)
	if err != nil {
		return nil, err
	}
	x := &bookServiceGetBooksClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type BookService_GetBooksClient interface {
	Recv() (*Book, error)
	grpc.ClientStream
}

type bookServiceGetBooksClient struct {
	grpc.ClientStream
}

func (x *bookServiceGetBooksClient) Recv() (*Book, error) {
	m := new(Book)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// BookServiceServer is the server API for BookService service.
// All implementations must embed UnimplementedBookServiceServer
// for forward compatibility
type BookServiceServer interface {
	GetBook(context.Context, *BookRequest) (*Book, error)
	GetBooks(*BookRequest, BookService_GetBooksServer) error
	mustEmbedUnimplementedBookServiceServer()
}

// UnimplementedBookServiceServer must be embedded to have forward compatible implementations.
type UnimplementedBookServiceServer struct {
}

func (UnimplementedBookServiceServer) GetBook(context.Context, *BookRequest) (*Book, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBook not implemented")
}
func (UnimplementedBookServiceServer) GetBooks(*BookRequest, BookService_GetBooksServer) error {
	return status.Errorf(codes.Unimplemented, "method GetBooks not implemented")
}
func (UnimplementedBookServiceServer) mustEmbedUnimplementedBookServiceServer() {}

// UnsafeBookServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BookServiceServer will
// result in compilation errors.
type UnsafeBookServiceServer interface {
	mustEmbedUnimplementedBookServiceServer()
}

func RegisterBookServiceServer(s grpc.ServiceRegistrar, srv BookServiceServer) {
	s.RegisterService(&BookService_ServiceDesc, srv)
}

func _BookService_GetBook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BookRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookServiceServer).GetBook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.BookService/GetBook",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookServiceServer).GetBook(ctx, req.(*BookRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookService_GetBooks_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(BookRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(BookServiceServer).GetBooks(m, &bookServiceGetBooksServer{stream})
}

type BookService_GetBooksServer interface {
	Send(*Book) error
	grpc.ServerStream
}

type bookServiceGetBooksServer struct {
	grpc.ServerStream
}

func (x *bookServiceGetBooksServer) Send(m *Book) error {
	return x.ServerStream.SendMsg(m)
}

// BookService_ServiceDesc is the grpc.ServiceDesc for BookService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BookService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.BookService",
	HandlerType: (*BookServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetBook",
			Handler:    _BookService_GetBook_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetBooks",
			Handler:       _BookService_GetBooks_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "book.proto",
}