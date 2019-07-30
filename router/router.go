/*
    fileName: router
    author: diogoxiang
    date: 2019/6/6
*/
package router

import (
	"github.com/gogf/gf/g"
	"github.com/gogf/gf/g/net/ghttp"
	"github.com/gogf/gf/g/os/glog"
)

func init()  {

	glog.Print("Router Init")
	s := g.Server()

	// 公用 及错误码 初始化
	StatusCode(s)
	// 系统配置类
	SystemInitRouter(s)
	// 初始化 用户API
	UserInitRouter(s)
	// 项目API
	ProInitRouter(s)

	// mock
	MockInitRouter(s)

	// 某些浏览器直接请求favicon.ico文件，特别是产生404时
	//s.SetRewrite("/favicon.ico", "/favicon.ico")

	// 用户模块 路由注册 - 使用执行对象注册方式
	//s.BindObject("/user", new(user.UserController))
	//s.BindHookHandlerByMap(p, map[string]ghttp.HandlerFunc{
	//	ghttp.HOOK_BEFORE_SERVE:  func(r *ghttp.Request) { glog.Println(ghttp.HOOK_BEFORE_SERVE) },
	//	ghttp.HOOK_AFTER_SERVE:   func(r *ghttp.Request) { glog.Println(ghttp.HOOK_AFTER_SERVE) },
	//	ghttp.HOOK_BEFORE_OUTPUT: func(r *ghttp.Request) { glog.Println(ghttp.HOOK_BEFORE_OUTPUT) },
	//	ghttp.HOOK_AFTER_OUTPUT:  func(r *ghttp.Request) { glog.Println(ghttp.HOOK_AFTER_OUTPUT) },
	//	ghttp.HOOK_BEFORE_CLOSE:  func(r *ghttp.Request) { glog.Println(ghttp.HOOK_BEFORE_CLOSE) },
	//	ghttp.HOOK_AFTER_CLOSE:   func(r *ghttp.Request) { glog.Println(ghttp.HOOK_AFTER_CLOSE) },
	//})
	// 测试hook 功能
	p := "/check"
	s.BindHandler(p, HookHandler)
}

func HookHandler(r *ghttp.Request) {
	glog.Println("Hook Handler")
	r.Response.Writeln("Hook Handler")
}
