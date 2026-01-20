package server

import (
	"fmt"
	"google.golang.org/grpc"
	"net"
)

func RpcServer() (*grpc.Server, net.Listener, error) {
	listen, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Printf("服务启动失败：%v \n", err)
		return nil, nil, err
	}
	server := grpc.NewServer()

	return server, listen, nil
}
