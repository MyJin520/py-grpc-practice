import grpc
import hello_pb2
import hello_pb2_grpc
from concurrent import futures


# 定义服务类，继承自动生成的HelloServiceServicer类
class HelloServiceServicer(hello_pb2_grpc.HelloServiceServicer):
    # 实现方法
    def SayHello(self, request, context):
        # 从请求中获取数据
        name = request.name
        message = request.message
        age = request.age

        # 处理业务逻辑
        response_message = f"你好，我叫{name}，今年{age}岁，你说：{message}"

        # 返回响应
        return hello_pb2.HelloResponse(message=response_message, error_code=200)


# 启动服务端
def run_server():
    # 创建服务器【使用线程池执行器，最大线程数为10】
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    # 注册方法到服务器【参数1：服务类的实例，参数2：服务器的实例】
    hello_pb2_grpc.add_HelloServiceServicer_to_server(HelloServiceServicer(), server)
    # 绑定端口号【ip:端口号】
    server.add_insecure_port('[::]:50051')

    # 启动服务器
    server.start()
    # 打印服务端地址
    print(f"服务端启动成功")
    # 等待服务器关闭【阻塞当前线程，直到服务器关闭】
    server.wait_for_termination()
    print("服务端已关闭")


if __name__ == '__main__':
    print("服务端启动中...")
    run_server()
