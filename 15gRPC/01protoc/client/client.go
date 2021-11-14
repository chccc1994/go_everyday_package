package main

import (
	"context"
	pb "go_everyday_package/15gRPC/01protoc/pb"
	"log"

	"google.golang.org/grpc"
)

const PORT = "9001"

func main() {
	conn, err := grpc.Dial(":"+PORT, grpc.WithInsecure()) // 创建与给定目标（服务端）的连接交互；Dial 建立连接targe->server地址  opts->DialOption
	if err != nil {
		log.Fatalf("grpc.Dial err: %v", err)
	}
	defer conn.Close()

	client := pb.NewSearchServiceClient(conn) // prod.pb.go 文件中产生的；创建 SearchService 的客户端对象
	//发送 RPC 请求，等待同步响应，得到回调后返回响应结果
	resp, err := client.Search(context.Background(), &pb.SearchRequest{
		Request: "gRPC",
	})
	if err != nil {
		log.Fatalf("client.Search err: %v", err)
	}

	log.Printf("resp: %s", resp.GetResponse())
}
