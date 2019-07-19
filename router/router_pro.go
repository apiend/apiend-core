/*
    fileName: router
    author: diogoxiang@qq.com
    date: 2019/7/15

	项目Router 初始化
	方便以后多接口可以共存的情况
*/
package router

import (
	"apiend-core/app/controller/project"
	"github.com/gogf/gf/g/net/ghttp"
)

// 项目初始化
func ProInitRouter(s *ghttp.Server)  {

	e := s.Group("/api")
	e.POST("/p", new(ctlProject.ProController))
}