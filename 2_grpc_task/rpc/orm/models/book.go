package models

import (
	"gorm.io/gorm"
	"time"
)

type BaseORM struct {
	ID        uint           `gorm:"primaryKey;autoIncrement;comment:主键ID，自增" json:"id"` // 主键ID，自增
	CreatedAt time.Time      `gorm:"autoCreateTime;comment:创建时间，自动生成" json:"createdAt"`  // 创建时间，自动生成
	UpdatedAt time.Time      `gorm:"autoUpdateTime;comment:更新时间，自动更新" json:"updatedAt"`  // 更新时间，自动更新
	DeletedAt gorm.DeletedAt `gorm:"index;comment:删除时间，软删除标记" json:"-"`                  // 删除时间，软删除标记，JSON序列化时忽略
}

type Book struct {
	BaseORM          // 嵌入基础结构体，继承通用字段
	BookName string  `gorm:"size:255;not null;comment:书籍名称" json:"bookName"`        // 书籍名称，最大长度255，不能为空
	Author   string  `gorm:"size:255;not null;comment:作者名称" json:"author"`          // 作者名称，最大长度255，不能为空
	Price    float32 `gorm:"type:decimal(10,2);not null;comment:书籍价格" json:"price"` // 书籍价格，十进制类型，总长度10，小数位2，不能为空
	Count    int32   `gorm:"default:0;comment:库存数量" json:"count"`                   // 库存数量，默认值为0
	Status   int32   `gorm:"default:1;comment:书籍状态，1表示可销售，0表示不可销售" json:"status"`   // 书籍状态，1表示可销售，0表示不可销售，默认值为1
}
