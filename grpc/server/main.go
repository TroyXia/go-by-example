package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	pb "go-by-example/grpc/proto/proto"
)

// 服务器实现
type server struct {
	pb.UnimplementedGreeterServer
}

// 实现 SayHello 方法
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {
	log.Printf("Received: %v", in.GetName())
	return &pb.HelloResponse{Message: "Hello " + in.GetName()}, nil
}

func main() {
	// 监听指定端口
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// 创建 gRPC 服务器
	s := grpc.NewServer()

	// 注册服务
	pb.RegisterGreeterServer(s, &server{})

	log.Println("Server listening on :50051")

	// 启动服务
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}