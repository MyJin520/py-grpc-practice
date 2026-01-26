package main

import (
	"fmt"
	"google.golang.org/grpc"
	"grpc-develop/2_grpc_task/rpc/orm/init_sql"
	"grpc-develop/2_grpc_task/rpc/pb"
	"grpc-develop/2_grpc_task/rpc/service"

	"net"
)

func main() {
	// 初始化全局DB
	init_sql.InitPgSql()
	fmt.Println("数据库初始化完成")

	listen, err := net.Listen("tcp", ":8000")
	if err != nil {
		fmt.Printf("rpc服务启动失败：%v \n", err)
		return
	}
	fmt.Printf("rpc服务启动成功，监听端口：%v \n", listen.Addr())
	rpcServer := grpc.NewServer()
	// 注册服务
	pb.RegisterBookServerServer(rpcServer, &service.RpcBookServer{})

	err = rpcServer.Serve(listen)
	if err != nil {
		fmt.Printf("rpc服务启动失败：%v \n", err)
		return
	}
}
