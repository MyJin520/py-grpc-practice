package service

import (
	"context"
	"grpc-develop/grpc_task02/pb"
)

type GrpcBookServer struct {
	pb.UnimplementedBookServerServer
}

func (g *GrpcBookServer) AddBook(ctx context.Context, req *pb.AddBookReq) (*pb.CommonResp, error) {
	//req.BookName = "静夜思"
	//req.Author = "李白"
	//req.Price = 32
	//req.Count = 12
	//req.Status = true

	p := &pb.CommonResp{
		Success: true,
		Msg:     "添加成功" + " 书名：" + req.BookName,
		Code:    200,
	}
	return p, nil
}
