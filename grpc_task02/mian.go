package main

import (
	"fmt"
	"grpc-develop/grpc_task02/orm"
)

func main() {
	db := orm.InitPgSql()
	if db == nil {
		println("连接数据库失败")
	}
	fmt.Println("连接数据库成功")
	orm.Migration(db)
	fmt.Println("数据库迁移成功")
}
