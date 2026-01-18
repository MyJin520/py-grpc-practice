package go_grpc_client

// UserRegistrationRequest 用户注册请求结构体
type UserRegistrationRequest struct {
	Name     string   `json:"name"`     // 用户名
	Email    string   `json:"email"`    // 邮箱
	Password string   `json:"password"` // 密码
	Age      int32    `json:"age"`      // 年龄
	Hobbies  []string `json:"hobbies"`  // 爱好列表
}

// UserLoginRequest 用户登录请求结构体
type UserLoginRequest struct {
	Email    string `json:"email"`    // 邮箱
	Password string `json:"password"` // 密码
}

// UserQueryRequest 用户查询请求结构体
type UserQueryRequest struct {
	Email string `json:"email"` // 邮箱
}

// UserUpdateRequest 用户更新请求结构体
type UserUpdateRequest struct {
	Email   string   `json:"email"`   // 邮箱
	Age     int32    `json:"age"`     // 年龄
	Hobbies []string `json:"hobbies"` // 爱好列表
}
