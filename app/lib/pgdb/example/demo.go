/*
    fileName: example
    author: diogoxiang
    date: 2019/6/13
*/
package main

import (
	_ "github.com/lib/pq"
	"github.com/gogf/gf/g/database/gdb"
)

var db gdb.DB

// 初始化配置及创建数据库
func init () {
	gdb.AddDefaultConfigNode(gdb.ConfigNode {
		Host    : "10.0.75.1",
		Port    : "5432",
		User    : "apiend",
		Pass    : "123456",
		Name    : "apiend_com",
		Type    : "pgsql",
		LinkInfo : "user=apiend password=123456 host=10.0.75.1 port=5432 dbname=apiend_com sslmode=disable",
	})
	db, _ = gdb.New()
}

func main() {
	db.SetDebug(true)

	// 执行3条SQL查询
	for i := 1; i <= 3; i++ {
		db.Table("user").Where("uid=?", i).One()
	}
	// 构造一条错误查询
	db.Table("user").Where("no_such_field=?", "just_test").One()
}
