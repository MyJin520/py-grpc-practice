package api

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"grpc-develop/2_grpc_task/go_gin/request"
	"grpc-develop/2_grpc_task/rpc/pb"
)

type BookHandler struct {
	Client pb.BookServerClient
}

// NewBookHandler 创建书籍处理器

func NewBookHandler(client pb.BookServerClient) *BookHandler {
	return &BookHandler{
		Client: client,
	}
}

// AddBook 添加书籍

func (h *BookHandler) AddBook(c *gin.Context) {
	var req request.AddBookRequest
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"msg":     "请求参数错误",
			"code":    http.StatusBadRequest,
		})
		return
	}

	// 转换为gRPC请求
	grpcReq := &pb.AddBookReq{
		BookName: req.BookName,
		Author:   req.Author,
		Price:    req.Price,
		Count:    req.Count,
		Status:   req.Status,
	}

	// 调用gRPC服务
	resp, err := h.Client.AddBook(context.Background(), grpcReq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"msg":     "添加书籍失败",
			"code":    http.StatusInternalServerError,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": resp.Success,
		"msg":     resp.Msg,
		"code":    resp.Code,
	})
}

// GetBookByID 根据ID获取书籍

func (h *BookHandler) GetBookByID(c *gin.Context) {
	var req request.BookIDRequest
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"msg":     "请求参数错误",
			"code":    http.StatusBadRequest,
		})
		return
	}

	// 转换为gRPC请求
	grpcReq := &pb.BookIDReq{
		Id: req.ID,
	}

	// 调用gRPC服务
	resp, err := h.Client.GetBookByID(context.Background(), grpcReq)
	if err != nil || resp.Success != true {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"msg":     "获取书籍失败",
			"code":    http.StatusInternalServerError,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": resp.Success,
		"book":    resp.Book,
		"code":    resp.Code,
	})
}

func (h *BookHandler) DeleteBookByID(c *gin.Context) {
	var req request.BookIDRequest
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"msg":     "请求参数错误",
			"code":    http.StatusBadRequest,
		})
		return
	}

	// 转换为gRPC请求
	grpcReq := &pb.BookIDReq{
		Id: req.ID,
	}

	// 调用gRPC服务
	resp, err := h.Client.DeleteBookByID(context.Background(), grpcReq)
	if err != nil || resp.Success != true {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"msg":     "删除书籍失败",
			"code":    http.StatusInternalServerError,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": resp.Success,
		"msg":     resp.Msg,
		"code":    resp.Code,
	})
}

// FindBooks 查询所有书籍

func (h *BookHandler) FindBooks(c *gin.Context) {
	var req = request.FindBooKsRequest{}
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"msg":     "请求参数错误",
			"code":    http.StatusBadRequest,
		})
		return
	}

	// 转换为gRPC请求
	grpcReq := &pb.FindBooksRep{
		Ids:    req.Ids,
		Status: req.Status,
	}

	// 调用gRPC服务
	resp, err := h.Client.FindBooks(context.Background(), grpcReq)
	if err != nil || resp.Success != true {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"msg":     "查询书籍失败",
			"code":    http.StatusInternalServerError,
		})
		return
	}

	fmt.Println("resp.Book", resp.Book)
	c.JSON(http.StatusOK, gin.H{
		"success": resp.Success,
		"book":    resp.Book,
		"code":    resp.Code,
	})
}

// UpdateBook 更新书籍

func (h *BookHandler) UpdateBook(c *gin.Context) {
	var req request.UpdateBookRequest
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"msg":     "请求参数错误",
			"code":    http.StatusBadRequest,
		})
		return
	}

	// 转换为gRPC请求
	grpcReq := &pb.UpDateBookReq{
		Id:       req.ID,
		BookName: req.BookName,
		Author:   req.Author,
		Price:    req.Price,
		Count:    req.Count,
		Status:   req.Status,
	}

	// 调用gRPC服务
	resp, err := h.Client.UpdateBook(context.Background(), grpcReq)
	if err != nil || resp.Success != true {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"msg":     "更新书籍失败",
			"code":    http.StatusInternalServerError,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": resp.Success,
		"msg":     resp.Msg,
		"code":    resp.Code,
	})
}
