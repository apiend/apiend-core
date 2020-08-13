package router

import (
	// "github.com/gogf/gf/g"
	"github.com/gogf/gf/g/net/ghttp"
)

func StatusCode(s *ghttp.Server) {
	// s := g.Server()
	// 统一处理一些错误
	s.BindStatusHandlerByMap(map[int]ghttp.HandlerFunc{
		403: func(r *ghttp.Request) { r.Response.Writeln("403") },
		404: func(r *ghttp.Request) { r.Response.Writeln("404") },
		500: func(r *ghttp.Request) { r.Response.Writeln("500") },
	})
	// TODO: 放开 CORS 限制
	s.BindHookHandlerByMap("/*any", map[string]ghttp.HandlerFunc{
		"BeforeServe": func(r *ghttp.Request) {

			// r.Response.CORSDefault()
			r.Response.CORS(ghttp.CORSOptions{
				AllowOrigin: "*",
			})
		},
	})

	// 当没有匹配到 URL 的时候
	s.BindObjectMethod("/*any", new(common), "ErrorIndex")

}

// common
type common struct{}

// 公用类
func (o *common) ErrorIndex(r *ghttp.Request) {
	r.Response.Writeln("未找到相应的URL")
}
