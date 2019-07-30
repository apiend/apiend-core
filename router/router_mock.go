/*
    fileName: router
    author: diogoxiang@qq.com
    date: 2019/7/19
*/
package router

import (
	"apiend-core/app/controller/mock"
	"apiend-core/app/controller/show"
	"github.com/gogf/gf/g/net/ghttp"
)

// 用户API路由
func MockInitRouter(s *ghttp.Server) {
	// 分组路由
	e := s.Group("/api")
	e.POST("/cm", new(mock.MockController))

	// 主返回的 路由控制  /项目id/项目路径前缀/mock路径
	s.BindObjectMethod("/m/:pid/:ppath/*any", new(show.ShowMock), "Show")
}
