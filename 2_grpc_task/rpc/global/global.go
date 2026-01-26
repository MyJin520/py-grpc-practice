package global

import "gorm.io/gorm"

// DB 全局数据库连接
type DB struct {
	*gorm.DB
}

// DBConn 全局数据库连接实例
var DBConn DB
