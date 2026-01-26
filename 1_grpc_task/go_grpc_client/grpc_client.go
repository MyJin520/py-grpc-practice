package go_grpc_client // todo 如果只进行grpc服务端+客户端测试请修改包名为main，当前修改为当前包名是为结合gin服务

import (
	"context"
	"fmt"
	"grpc-develop/1_grpc_task"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// 服务端地址（和Python服务端完全一致）
const address = "localhost:8000"

func main() {
	// 1. 连接gRPC服务端：因为Python服务端是 无证书的不安全连接 add_insecure_port
	// 所以这里用 insecure.NewCredentials() 跳过tls验证
	conn, err := grpc.NewClient(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("客户端连接服务端失败: %v", err)
	}
	// 程序退出时关闭连接
	defer conn.Close()

	// 2. 创建UserService的gRPC客户端对象（核心：自动生成的桩代码）
	client := __grpc_task.NewUserServiceClient(conn)

	// 3. 调用各个业务方法：注册 -> 登录 -> 查询 -> 更新
	ctx := context.Background()
	TestRegister(ctx, client)
	fmt.Println("------------------------")
	TestLogin(ctx, client)
	fmt.Println("------------------------")
	TestQuery(ctx, client)
	fmt.Println("------------------------")
	TestUpdate(ctx, client)
	fmt.Println("------------------------")
	TestQuery(ctx, client) // 更新后再次查询，验证结果
}

// 测试注册接口
func TestRegister(ctx context.Context, client __grpc_task.UserServiceClient) {
	// 构造注册请求参数
	req := &__grpc_task.UserRegistration{
		Name:     "张三",
		Email:    "zhangsan@test.com",
		Password: "123456",
		Age:      25,
		Hobbies:  []string{"篮球", "编程"},
	}
	// 调用gRPC的Register方法
	res, err := client.Register(ctx, req)
	if err != nil {
		log.Printf("注册失败: %v", err)
		return
	}
	fmt.Printf("注册结果 -> code: %d, message: %s\n", res.Code, res.Message)
}

// 测试登录接口
func TestLogin(ctx context.Context, client __grpc_task.UserServiceClient) {
	req := &__grpc_task.UserLoginRequest{
		Email:    "zhangsan@test.com",
		Password: "123456",
	}
	res, err := client.Login(ctx, req)
	if err != nil {
		log.Printf("登录失败: %v", err)
		return
	}
	fmt.Printf("登录结果 -> code: %d, message: %s\n", res.Code, res.Message)
}

// 测试查询接口
func TestQuery(ctx context.Context, client __grpc_task.UserServiceClient) {
	req := &__grpc_task.UserQueryRequest{
		Email: "zhangsan@test.com",
	}
	res, err := client.Query(ctx, req)
	if err != nil {
		log.Printf("查询失败: %v", err)
		return
	}
	fmt.Printf("查询结果 -> code: %d, message: %s\n", res.Code, res.Message)
}

// 测试更新接口
func TestUpdate(ctx context.Context, client __grpc_task.UserServiceClient) {
	req := &__grpc_task.UserUpdateRequest{
		Email:   "zhangsan@test.com",
		Age:     26,
		Hobbies: []string{"羽毛球", "编程", "读书"},
	}
	res, err := client.Update(ctx, req)
	if err != nil {
		log.Printf("更新失败: %v", err)
		return
	}
	fmt.Printf("更新结果 -> code: %d, message: %s\n", res.Code, res.Message)
}
