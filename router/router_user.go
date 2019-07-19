package router

import (
	"apiend-core/app/controller/user"
	"github.com/gogf/gf/g/net/ghttp"


)

// 用户API路由
func UserInitRouter(s *ghttp.Server) {
 	// 分组路由
	e := s.Group("/api")
	e.POST("/user", new(ctlUser.UserController))
}

