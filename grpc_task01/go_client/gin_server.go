package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"grpc-develop/grpc_task01"
	"grpc-develop/grpc_task01/go_grpc_client"
	"log"
)

// APIResponse 统一API响应结构
type APIResponse struct {
	Code    int         `json:"code"`    // 状态码
	Message string      `json:"message"` // 响应消息
	Data    interface{} `json:"data"`    // 响应数据
}

// gRPC客户端连接
type GrpcClient struct {
	UserClient grpc_task01.UserServiceClient
	Conn       *grpc.ClientConn
}

// 初始化gRPC客户端
func initGrpcClient() (*GrpcClient, error) {
	address := "localhost:8000"
	conn, err := grpc.NewClient(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	return &GrpcClient{
		UserClient: grpc_task01.NewUserServiceClient(conn),
		Conn:       conn,
	}, nil
}

// 错误响应处理
func sendErrorResponse(c *gin.Context, statusCode int, message string, err error) {
	response := APIResponse{
		Code:    statusCode,
		Message: message,
		Data:    nil,
	}
	if err != nil {
		response.Data = err.Error()
	}
	c.JSON(statusCode, response)
}

// 成功响应处理
func sendSuccessResponse(c *gin.Context, data interface{}) {
	response := APIResponse{
		Code:    200,
		Message: "success",
		Data:    data,
	}
	c.JSON(200, response)
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
	setupRoutes(engine, grpcClient)

	// 启动服务器
	log.Println("Gin服务器启动，监听端口: 8080")
	if err := engine.Run(); err != nil {
		log.Fatalf("启动Gin服务器失败: %v", err)
	}
}

// 设置路由
func setupRoutes(engine *gin.Engine, grpcClient *GrpcClient) {
	// 用户相关路由组
	userGroup := engine.Group("/user")
	{
		userGroup.POST("/register", handleRegister(grpcClient))
		userGroup.POST("/login", handleLogin(grpcClient))
		userGroup.POST("/query", handleQuery(grpcClient))
		userGroup.POST("/update", handleUpdate(grpcClient))
	}
}

// 处理用户注册请求
func handleRegister(grpcClient *GrpcClient) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 绑定请求参数
		var req go_grpc_client.UserRegistrationRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			sendErrorResponse(c, 400, "请求参数错误", err)
			return
		}

		// 转换为gRPC请求
		grpcReq := &grpc_task01.UserRegistration{
			Name:     req.Name,
			Email:    req.Email,
			Password: req.Password,
			Age:      req.Age,
			Hobbies:  req.Hobbies,
		}

		// 调用gRPC服务
		resp, err := grpcClient.UserClient.Register(context.Background(), grpcReq)
		if err != nil {
			sendErrorResponse(c, 500, "调用gRPC服务失败", err)
			return
		}

		if resp.Code != 200 {
			sendErrorResponse(c, 500, resp.Message, nil)
			return
		}

		sendSuccessResponse(c, resp.Message)
	}
}

// 处理用户登录请求
func handleLogin(grpcClient *GrpcClient) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 绑定请求参数
		var req go_grpc_client.UserLoginRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			sendErrorResponse(c, 400, "请求参数错误", err)
			return
		}

		// 转换为gRPC请求
		grpcReq := &grpc_task01.UserLoginRequest{
			Email:    req.Email,
			Password: req.Password,
		}

		// 调用gRPC服务
		resp, err := grpcClient.UserClient.Login(context.Background(), grpcReq)
		if err != nil {
			sendErrorResponse(c, 500, "调用gRPC服务失败", err)
			return
		}

		if resp.Code != 200 {
			sendErrorResponse(c, 401, resp.Message, nil)
			return
		}

		sendSuccessResponse(c, gin.H{"token": "mock-token", "message": resp.Message})
	}
}

// 处理用户查询请求
func handleQuery(grpcClient *GrpcClient) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 绑定请求参数
		var req go_grpc_client.UserQueryRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			sendErrorResponse(c, 400, "请求参数错误", err)
			return
		}

		// 转换为gRPC请求
		grpcReq := &grpc_task01.UserQueryRequest{
			Email: req.Email,
		}

		// 调用gRPC服务
		resp, err := grpcClient.UserClient.Query(context.Background(), grpcReq)
		if err != nil {
			sendErrorResponse(c, 500, "调用gRPC服务失败", err)
			return
		}

		if resp.Code != 200 {
			sendErrorResponse(c, 404, resp.Message, nil)
			return
		}

		sendSuccessResponse(c, resp.Message)
	}
}

// 处理用户更新请求
func handleUpdate(grpcClient *GrpcClient) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 绑定请求参数
		var req go_grpc_client.UserUpdateRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			sendErrorResponse(c, 400, "请求参数错误", err)
			return
		}
		fmt.Print("age", req.Age, " email:", req.Hobbies, " hobbies:", req.Hobbies)
		// 转换为gRPC请求
		grpcReq := &grpc_task01.UserUpdateRequest{
			Email:   req.Email,
			Age:     req.Age,
			Hobbies: req.Hobbies,
		}

		// 调用gRPC服务
		resp, err := grpcClient.UserClient.Update(context.Background(), grpcReq)
		if err != nil {
			sendErrorResponse(c, 500, "调用gRPC服务失败", err)
			return
		}

		if resp.Code != 200 {
			sendErrorResponse(c, 500, resp.Message, nil)
			return
		}

		sendSuccessResponse(c, resp.Message)
	}
}
