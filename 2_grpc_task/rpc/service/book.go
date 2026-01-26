package service

import (
	"context"
	"grpc-develop/2_grpc_task/rpc/global"
	"grpc-develop/2_grpc_task/rpc/orm/models"
	"grpc-develop/2_grpc_task/rpc/pb"
)

type RpcBookServer struct {
	pb.UnimplementedBookServerServer
}

// 添加书籍
func (r *RpcBookServer) AddBook(c context.Context, req *pb.AddBookReq) (*pb.CommonResp, error) {
	book := models.Book{
		BookName: req.BookName,
		Author:   req.Author,
		Price:    req.Price,
		Count:    req.Count,
		Status:   req.Status,
	}
	result := global.DBConn.Create(&book)
	if result.Error != nil {
		return &pb.CommonResp{
			Success: false,
			Msg:     "新是书籍添加失败：" + result.Error.Error(),
			Code:    500,
		}, nil
	}

	return &pb.CommonResp{
		Success: true,
		Msg:     "新书籍添加成功",
		Code:    200,
	}, nil
}

func (r *RpcBookServer) GetBookByID(c context.Context, req *pb.BookIDReq) (*pb.GetBookByIDResp, error) {
	book := &models.Book{}
	if result := global.DBConn.First(&book, req.Id); result.Error != nil {
		return &pb.GetBookByIDResp{Success: false, Book: nil, Code: 500}, nil
	}

	return &pb.GetBookByIDResp{
		Success: true,
		Book:    toProtoBook(book),
		Code:    200,
	}, nil
}

// DeleteBookByID 根据ID删除书籍
func (r *RpcBookServer) DeleteBookByID(c context.Context, req *pb.BookIDReq) (*pb.CommonResp, error) {
	result := global.DBConn.DB.Delete(&models.Book{}, req.Id)
	if result.Error != nil {
		return &pb.CommonResp{
			Success: false,
			Msg:     "书籍信息删除失败：" + result.Error.Error(),
			Code:    500,
		}, nil
	}

	// 检查是否真的删除了记录
	if result.RowsAffected == 0 {
		return &pb.CommonResp{
			Success: false,
			Msg:     "未找到要删除的书籍",
			Code:    404,
		}, nil
	}

	return &pb.CommonResp{
		Success: true,
		Msg:     "书籍信息删除成功",
		Code:    200,
	}, nil
}

// FindBooks 查询书籍列表
func (r *RpcBookServer) FindBooks(c context.Context, req *pb.FindBooksRep) (*pb.FindBooksResp, error) {
	var books []models.Book

	query := global.DBConn.DB

	// 1. 如果用户传递了ids参数，则只查询包含这些ids的书籍
	if len(req.Ids) > 0 {
		query = query.Where("id IN ?", req.Ids)
	}

	if req.Status != -1 {
		query = query.Where("status = ?", req.Status)
	}

	result := query.Find(&books)
	if result.Error != nil {
		return &pb.FindBooksResp{Success: false, Book: nil, Code: 500}, result.Error
	}

	protoBooks := make([]*pb.Book, len(books))
	for i, book := range books {
		protoBooks[i] = &pb.Book{
			Id:       int32(book.ID),
			BookName: book.BookName,
			Author:   book.Author,
			Price:    book.Price,
			Count:    book.Count,
			Status:   book.Status,
		}
	}

	return &pb.FindBooksResp{Success: true, Book: protoBooks, Code: 200}, nil
}

func (r *RpcBookServer) UpdateBook(c context.Context, req *pb.UpDateBookReq) (*pb.CommonResp, error) {
	updateData := map[string]interface{}{
		"book_name": req.BookName,
		"author":    req.Author,
		"price":     req.Price,
		"count":     req.Count,
		"status":    req.Status,
	}

	result := global.DBConn.Model(&models.Book{}).Where("id=?", req.Id).Updates(updateData)
	if result.Error != nil {
		return &pb.CommonResp{
			Success: false,
			Msg:     "书籍信息更新失败",
			Code:    500,
		}, nil
	}

	if result.RowsAffected == 0 {
		return &pb.CommonResp{
			Success: false,
			Msg:     "未找到要更新的书籍",
			Code:    404,
		}, nil
	}

	return &pb.CommonResp{
		Success: true,
		Msg:     "书籍信息更新成功",
		Code:    200,
	}, nil
}

func toProtoBook(book *models.Book) *pb.Book {
	return &pb.Book{
		Id:       int32(book.ID),
		BookName: book.BookName,
		Author:   book.Author,
		Price:    book.Price,
		Count:    book.Count,
		Status:   book.Status,
	}
}
