package orm

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitPgSql() *gorm.DB {
	dsn := "host=localhost user=root password=123456 dbname=myapp_db port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("连接数据库失败")
	}
	return db
}

func Migration(db *gorm.DB) {
	err := db.AutoMigrate(&Book{})
	if err != nil {
		fmt.Println("数据库迁移失败:", err)
		return
	}
}
