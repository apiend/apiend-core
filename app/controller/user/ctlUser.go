/*
    fileName: user
    author: diogoxiang
    date: 2019/6/10
*/
package ctlUser

import (
	"github.com/gogf/gf/g"
	"github.com/gogf/gf/g/net/ghttp"
	"github.com/gogf/gf/g/os/glog"
)

var (
	c        = g.Config()
	gravatar = c.GetArray("system.gravatar")
)

// 用户API管理对象
type UserController struct{}

// 公用方法 可在所有方法调用之前
func (c *UserController) Init(r *ghttp.Request) {

	glog.Println("init ")

}

// 公用方法 方法执行完了之后 调用
func (c *UserController) Shut(r *ghttp.Request) {
	glog.Println("over ")

}