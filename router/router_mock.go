/*
    fileName: router
    author: diogoxiang@qq.com
    date: 2019/7/19
*/
package router

import (
	"apiend-core/app/controller/mock"
	"github.com/gogf/gf/g/net/ghttp"
)

// 用户API路由
func MockInitRouter(s *ghttp.Server) {
	// 分组路由
	e := s.Group("/api")
	e.POST("/cm", new(mock.MockController))
}
