package router

import (
	"apiend-core/app/controller/user"
	"github.com/gogf/gf/g"
	"github.com/gogf/gf/g/net/ghttp"
	"github.com/gogf/gf/g/os/glog"
)

func init() {
	s := g.Server()

	// 某些浏览器直接请求favicon.ico文件，特别是产生404时
	//s.SetRewrite("/favicon.ico", "/favicon.ico")

	// 用户模块 路由注册 - 使用执行对象注册方式
	//s.BindObject("/user", new(user.UserController))

	// 分组路由
	e := s.Group("/api")
	e.POST("/user", new(ctlUser.UserController))

	// 测试hook 功能
	p := "/check"
	//s.BindHookHandlerByMap(p, map[string]ghttp.HandlerFunc{
	//	ghttp.HOOK_BEFORE_SERVE:  func(r *ghttp.Request) { glog.Println(ghttp.HOOK_BEFORE_SERVE) },
	//	ghttp.HOOK_AFTER_SERVE:   func(r *ghttp.Request) { glog.Println(ghttp.HOOK_AFTER_SERVE) },
	//	ghttp.HOOK_BEFORE_OUTPUT: func(r *ghttp.Request) { glog.Println(ghttp.HOOK_BEFORE_OUTPUT) },
	//	ghttp.HOOK_AFTER_OUTPUT:  func(r *ghttp.Request) { glog.Println(ghttp.HOOK_AFTER_OUTPUT) },
	//	ghttp.HOOK_BEFORE_CLOSE:  func(r *ghttp.Request) { glog.Println(ghttp.HOOK_BEFORE_CLOSE) },
	//	ghttp.HOOK_AFTER_CLOSE:   func(r *ghttp.Request) { glog.Println(ghttp.HOOK_AFTER_CLOSE) },
	//})
	s.BindHandler(p, HookHandler)

}

func HookHandler(r *ghttp.Request) {
	glog.Println("Hook Handler")
	r.Response.Writeln("Hook Handler")
}
