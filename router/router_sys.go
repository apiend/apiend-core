/*
    fileName: router
    author: diogoxiang@qq.com
    date: 2019/7/24
*/
package router

import (

	"apiend-core/app/controller/uptoken"
	"github.com/gogf/gf/g/net/ghttp"
)

func SystemInitRouter(s *ghttp.Server)  {

	// 分组路由
	e := s.Group("/api")
	e.GET("/up", new(uptoken.Uptoken))

}