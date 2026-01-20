package main

import (
	"context"
	"fmt"
	"time"

	"grpc-develop/grpc_task02/client"
	"grpc-develop/grpc_task02/pb"
)

func main() {
	// 启动rpc客户端
	rpcClient, conn := client.RpcClient()
	defer conn.Close()

	book, err := rpcClient.AddBook(context.Background(), &pb.AddBookReq{
		BookName: "静夜思",
		Author:   "李白",
		Price:    21,
		Count:    1,
		Status:   false,
	})

	if err != nil {
		fmt.Println("调用失败:", err)
		return
	}

	fmt.Printf("调用成功: %+v\n", book)
	time.Sleep(5 * time.Second)
}
