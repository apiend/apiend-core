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
// 注册用户 /user/Register
func (c *UserController) Register(r *ghttp.Request) {
	glog.Println("Register ")
	//简单的注册只需要用户名以及密码就行了
	// postPayload := controller.TranspilePostParams(r)
	postPayload := r.GetString("username")

	var pd map[string]string

	postData := r.GetPostMap(pd)
	glog.Println(postData)
	glog.Println(postPayload)


}

// 用户登录 /user/SignUp
func (c *UserController) SignUp(r *ghttp.Request) {
	glog.Println("SignUp ")
}

// 修改用户信息 /user/UpdateUser
func (c *UserController) UpdateUser(r *ghttp.Request) {
	glog.Println("UpdateUser ")
}

// 获取用户列表 /user/GetList
func (c *UserController) GetList(r *ghttp.Request) {
	glog.Println("GetList ")
}

// 根据用户名 查询用户 信息 /user/GetSearchName
func (c *UserController) GetSearchName(r *ghttp.Request) {
	// log
	glog.Println("GetSearchName ")
}

// 根据用户Uid  获取数据 /user/GetByUid
func (c *UserController) GetByUid(r *ghttp.Request) {
	glog.Println("GetByUid ")
}