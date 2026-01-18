import grpc
from basic_learning.server import hello_pb2, hello_pb2_grpc


# 开启客户端【连接客户端】【参数：ip:端口号】
def run_client():
    with grpc.insecure_channel('localhost:50051') as channel:
        # 创建 stub「存根」 【stub 是 gRPC 客户端的一个代理，用于调用服务端的方法】
        stub = hello_pb2_grpc.HelloServiceStub(channel)

        # 调用服务端的方法【参数：请求字段... ，超时时间】
        response = stub.SayHello(hello_pb2.HelloRequest(name='张三', message='你好，很高兴认识你', age=18), timeout=10)
        # 打印响应
        resp = {
            'message': response.message,
            'error_code': response.error_code,
        }
        print(resp)

if __name__ == '__main__':
    run_client()