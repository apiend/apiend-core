/*
    fileName: apiend_core
    author: diogoxiang
    date: 2019/6/6
*/
package main

import (
	_ "apiend-core/boot"
	_ "apiend-core/router"

	"github.com/gogf/gf/g"
)

func main() {
	g.Server().SetServerAgent("Diogo")
	g.Server().Run()
}


// package main
//
// import (
// 	"time"
// 	"github.com/gogf/gf/g"
// 	"github.com/gogf/gf/g/os/gproc"
// 	"github.com/gogf/gf/g/net/ghttp"
// )
//
// func main() {
// 	s := g.Server()
// 	s.BindHandler("/", func(r *ghttp.Request){
// 		r.Response.Writeln("哈喽！")
// 	})
// 	s.BindHandler("/pid", func(r *ghttp.Request){
// 		r.Response.Writeln(gproc.Pid())
// 	})
// 	s.BindHandler("/sleep", func(r *ghttp.Request){
// 		r.Response.Writeln(gproc.Pid())
// 		time.Sleep(10*time.Second)
// 		r.Response.Writeln(gproc.Pid())
// 	})
// 	s.EnableAdmin()
// 	s.SetPort(8199)
// 	s.Run()
// }
