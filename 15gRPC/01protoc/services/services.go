package main

import (
	"context"
	pb "go_everyday_package/15gRPC/01protoc/pb"
	"log"
	"net"

	"google.golang.org/grpc"
)

type SearchService struct{}

func (s *SearchService) Search(ctx context.Context, r *pb.SearchRequest) (*pb.SearchResponse, error) {
	return &pb.SearchResponse{Response: r.GetRequest() + " Server"}, nil
}

const PORT = "9001"

func main() {
	server := grpc.NewServer()                               //创建 gRPC Server 对象，你可以理解为它是 Server 端的抽象对象
	pb.RegisterSearchServiceServer(server, &SearchService{}) // prod.pb.go 文件中产生的，注册到 gRPC Server 的内部注册中心

	lis, err := net.Listen("tcp", ":"+PORT) //net 监听 创建 Listen，监听 TCP 端口
	if err != nil {
		log.Fatalf("net.Listen err: %v", err)
	}

	server.Serve(lis)
}
