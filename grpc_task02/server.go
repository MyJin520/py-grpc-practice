package main

import (
	"fmt"

	"grpc-develop/grpc_task02/pb"
	"grpc-develop/grpc_task02/server"
	"grpc-develop/grpc_task02/service"
)

func main() {
	rpcServer, listener, err := server.RpcServer()
	if err != nil {
		fmt.Println("rpc服务启动失败")
		return
	}

	// 注册rpc服务（使用指针）
	pb.RegisterBookServerServer(rpcServer, &service.GrpcBookServer{})

	// 启动服务器
	fmt.Println("服务器启动，监听端口8080...")
	if err := rpcServer.Serve(listener); err != nil {
		fmt.Println("服务器运行失败:", err)
		return
	}
}
