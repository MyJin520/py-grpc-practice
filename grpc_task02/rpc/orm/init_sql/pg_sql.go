package init_sql

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"grpc-develop/grpc_task02/rpc/global"
	"grpc-develop/grpc_task02/rpc/orm/models"
)

func InitPgSql() {
	dsn := "host=localhost user=root password=123456 dbname=myapp_db port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("连接数据库失败" + err.Error())
	}

	// 显式创建public schema（如果不存在）
	if err := db.Exec("CREATE SCHEMA IF NOT EXISTS public").Error; err != nil {
		fmt.Println("创建schema失败:", err)
		return
	}

	// 设置默认schema
	if err := db.Exec("SET search_path TO public").Error; err != nil {
		fmt.Println("设置默认schema失败:", err)
		return
	}

	global.DBConn = global.DB{DB: db}
	Migration(global.DBConn)
}

func Migration(db global.DB) {
	err := db.AutoMigrate(&models.Book{})
	if err != nil {
		fmt.Println("数据库迁移失败:", err)
		return
	}
}
