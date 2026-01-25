package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"grpc-develop/grpc_task02/go_gin/router"
	"grpc-develop/grpc_task02/rpc/pb"
)

// GrpcClient gRPC客户端结构体

type GrpcClient struct {
	BookClient pb.BookServerClient
	Conn       *grpc.ClientConn
}

// 初始化gRPC客户端

func initGrpcClient() (*GrpcClient, error) {
	// gRPC服务器地址
	address := "localhost:8000"

	// 创建gRPC连接
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("无法连接到gRPC服务器: %v", err)
		return nil, err
	}

	// 创建gRPC客户端
	client := pb.NewBookServerClient(conn)

	return &GrpcClient{
		BookClient: client,
		Conn:       conn,
	}, nil
}

func main() {
	// 初始化gRPC客户端
	grpcClient, err := initGrpcClient()
	if err != nil {
		log.Fatalf("初始化gRPC客户端失败: %v", err)
	}
	defer grpcClient.Conn.Close()

	// 初始化Gin引擎
	engine := gin.Default()

	// 注册路由
	router.SetupBookRouter(engine, grpcClient.BookClient)

	// 启动服务器
	log.Println("Gin服务器启动，监听端口: 8080")
	if err := engine.Run(":8080"); err != nil {
		log.Fatalf("启动Gin服务器失败: %v", err)
	}
}
