import grpc
from concurrent import futures

from grpc_task01 import user_pb2_grpc, user_pb2

# 全局用户存储字典
user_map = {}


class UserService(user_pb2_grpc.UserServiceServicer):

    def Register(self, request, context):
        if request.age <= 0 or request.age >= 120:
            resp = "用户注册失败，年龄异常"
            code = 400
            return user_pb2.UserRegistrationResponse(message=resp, code=code)

        user_info = {
            "Name": request.name,
            "Email": request.email,
            "Password": request.password,
            "Age": request.age,
            "Hobbies": request.hobbies,
        }
        user_map[request.email] = user_info
        resp = f"用户注册成功，用户信息：{user_info}"
        code = 200
        return user_pb2.UserRegistrationResponse(message=resp, code=code)

    def Login(self, request, context):
        if request.email not in user_map:
            resp = "用户登录失败，邮箱不存在"
            code = 400
            return user_pb2.UserLoginResponse(message=resp, code=code)
        if user_map[request.email]["Password"] != request.password:
            resp = "用户登录失败，密码错误"
            code = 400
            return user_pb2.UserLoginResponse(message=resp, code=code)

        resp = f"用户登录成功，用户信息：{user_map[request.email]}"
        code = 200
        return user_pb2.UserLoginResponse(message=resp, code=code)

    def Query(self, request, context):
        if request.email not in user_map:
            resp = "查询的用户不存在"
            code = 400
            return user_pb2.UserQueryResponse(message=resp, code=code)

        resp = str(user_map[request.email])
        code = 200
        return user_pb2.UserQueryResponse(message=resp, code=code)

    def Update(self, request, context):

        if request.email not in user_map:
            resp = "更新失败，用户不存在"
            code = 400
            return user_pb2.UserUpdateResponse(message=resp, code=code)

        user_map[request.email]["Age"] = request.age
        user_map[request.email]["Hobbies"] = request.hobbies

        new_info = user_map[request.email]
        resp = f"用户信息更新成功，最新信息：{new_info}"
        code = 200
        return user_pb2.UserUpdateResponse(message=resp, code=code)


def run_server():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=5))
    user_pb2_grpc.add_UserServiceServicer_to_server(UserService(), server)
    server.add_insecure_port("localhost:8000")

    server.start()
    print("服务启动成功，监听端口 localhost:8000")
    server.wait_for_termination()
    print("服务已关闭")


if __name__ == '__main__':
    run_server()