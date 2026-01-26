package request

// AddBookRequest 添加书籍请求

type AddBookRequest struct {
	BookName string  `json:"bookName" `
	Author   string  `json:"author" `
	Price    float32 `json:"price" `
	Count    int32   `json:"count" `
	Status   int32   `json:"status" `
}

// BookIDRequest 书籍ID请求

type BookIDRequest struct {
	ID int32 `form:"id" `
}

// UpdateBookRequest 更新书籍请求

type UpdateBookRequest struct {
	ID       int32   `json:"id" `
	BookName string  `json:"bookName" `
	Author   string  `json:"author" `
	Price    float32 `json:"price" `
	Count    int32   `json:"count" `
	Status   int32   `json:"status" `
}

// FindBooKsRequest 查询书籍请求
type FindBooKsRequest struct {
	Ids    []int32 `json:"ids" form:"ids"`
	Status int32   `json:"status" form:"status"`
}
