package router

import (
	"github.com/gin-gonic/gin"
	"grpc-develop/2_grpc_task/go_gin/api"
	"grpc-develop/2_grpc_task/rpc/pb"
)

// SetupBookRouter 设置书籍路由

func SetupBookRouter(engine *gin.Engine, client pb.BookServerClient) {
	bookHandler := api.NewBookHandler(client)

	// 书籍相关路由
	bookGroup := engine.Group("/book")
	{
		bookGroup.POST("/add", bookHandler.AddBook)
		bookGroup.GET("/get", bookHandler.GetBookByID)
		bookGroup.DELETE("/delete", bookHandler.DeleteBookByID)
		bookGroup.POST("/find", bookHandler.FindBooks)
		bookGroup.PUT("/update", bookHandler.UpdateBook)
	}
}
