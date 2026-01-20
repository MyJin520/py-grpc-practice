package client

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"grpc-develop/grpc_task02/pb"
	"log"
)

const address = "localhost:8080"

func RpcClient() (pb.BookServerClient, *grpc.ClientConn) {
	conn, err := grpc.NewClient(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("客户端连接失败：%v", err)
	}
	// 创建客户端 --> 来自grpc.pb.go文件
	client := pb.NewBookServerClient(conn)
	return client, conn
}
