package orm

import (
	"gorm.io/gorm"
	"time"
)

// BaseTabel 所有数据库模型的基础结构体，包含通用字段
// gorm:"comment:基础表结构，包含所有模型的通用字段"
type BaseTabel struct {
	ID        uint           `gorm:"primaryKey;autoIncrement;comment:主键ID，自增" json:"id"` // 主键ID，自增
	CreatedAt time.Time      `gorm:"autoCreateTime;comment:创建时间，自动生成" json:"createdAt"`  // 创建时间，自动生成
	UpdatedAt time.Time      `gorm:"autoUpdateTime;comment:更新时间，自动更新" json:"updatedAt"`  // 更新时间，自动更新
	DeletedAt gorm.DeletedAt `gorm:"index;comment:删除时间，软删除标记" json:"-"`                  // 删除时间，软删除标记，JSON序列化时忽略
}

// Book 定义书籍模型
// gorm:"comment:书籍表，存储书籍基本信息"
type Book struct {
	BaseTabel         // 嵌入基础结构体，继承通用字段
	BookName  string  `gorm:"size:255;not null;comment:书籍名称" json:"bookName"`                // 书籍名称，最大长度255，不能为空
	Author    string  `gorm:"size:255;not null;comment:作者名称" json:"author"`                  // 作者名称，最大长度255，不能为空
	Price     float32 `gorm:"type:decimal(10,2);not null;comment:书籍价格" json:"price"`         // 书籍价格，十进制类型，总长度10，小数位2，不能为空
	Count     int32   `gorm:"default:0;comment:库存数量" json:"count"`                           // 库存数量，默认值为0
	Status    bool    `gorm:"default:true;comment:书籍状态，true表示可销售，false表示不可销售" json:"status"` // 书籍状态，true表示可销售，false表示不可销售，默认值为true
}
