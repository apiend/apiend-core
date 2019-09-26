/*
    fileName: router
    author: diogoxiang@qq.com
    date: 2019/8/13
*/
package router

import (
	"apiend-core/app/controller/epost"
	"github.com/gogf/gf/g/net/ghttp"
)

// 初始化项目的
func EpostInitRouter(s *ghttp.Server)  {
	e := s.Group("/api")
	e.POST("/e", new(epost.EpostController))
}