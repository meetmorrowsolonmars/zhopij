// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.19.4
// source: api/v1/quiz/quiz.proto

package quiz

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
	QuizService_CreateQuiz_FullMethodName = "/zhopij.quiz.v1.quiz.QuizService/CreateQuiz"
)

// QuizServiceClient is the client API for QuizService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type QuizServiceClient interface {
	CreateQuiz(ctx context.Context, in *CreateQuizRequest, opts ...grpc.CallOption) (*CreateQuizResponse, error)
}

type quizServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewQuizServiceClient(cc grpc.ClientConnInterface) QuizServiceClient {
	return &quizServiceClient{cc}
}

func (c *quizServiceClient) CreateQuiz(ctx context.Context, in *CreateQuizRequest, opts ...grpc.CallOption) (*CreateQuizResponse, error) {
	out := new(CreateQuizResponse)
	err := c.cc.Invoke(ctx, QuizService_CreateQuiz_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QuizServiceServer is the server API for QuizService service.
// All implementations must embed UnimplementedQuizServiceServer
// for forward compatibility
type QuizServiceServer interface {
	CreateQuiz(context.Context, *CreateQuizRequest) (*CreateQuizResponse, error)
	mustEmbedUnimplementedQuizServiceServer()
}

// UnimplementedQuizServiceServer must be embedded to have forward compatible implementations.
type UnimplementedQuizServiceServer struct {
}

func (UnimplementedQuizServiceServer) CreateQuiz(context.Context, *CreateQuizRequest) (*CreateQuizResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateQuiz not implemented")
}
func (UnimplementedQuizServiceServer) mustEmbedUnimplementedQuizServiceServer() {}

// UnsafeQuizServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to QuizServiceServer will
// result in compilation errors.
type UnsafeQuizServiceServer interface {
	mustEmbedUnimplementedQuizServiceServer()
}

func RegisterQuizServiceServer(s grpc.ServiceRegistrar, srv QuizServiceServer) {
	s.RegisterService(&QuizService_ServiceDesc, srv)
}

func _QuizService_CreateQuiz_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateQuizRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QuizServiceServer).CreateQuiz(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: QuizService_CreateQuiz_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QuizServiceServer).CreateQuiz(ctx, req.(*CreateQuizRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// QuizService_ServiceDesc is the grpc.ServiceDesc for QuizService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var QuizService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "zhopij.quiz.v1.quiz.QuizService",
	HandlerType: (*QuizServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateQuiz",
			Handler:    _QuizService_CreateQuiz_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/v1/quiz/quiz.proto",
}