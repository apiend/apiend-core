/*
   fileName: user
   author: diogoxiang
   date: 2019/6/10
*/
package ctlUser

import (
	"apiend-core/app/lib/cdb"
	"apiend-core/app/lib/response"
	"apiend-core/app/lib/util"
	"apiend-core/app/model"
	"apiend-core/app/model/user"
	"fmt"
	"github.com/globalsign/mgo/bson"
	"github.com/gogf/gf/g"
	"github.com/gogf/gf/g/net/ghttp"
	"github.com/gogf/gf/g/os/glog"
	"github.com/gogf/gf/g/text/gstr"
	"github.com/gogf/gf/g/util/gconv"
	"github.com/gogf/gf/g/util/gvalid"
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

	// var pd map[string]string
	// postData := r.GetPostMap(pd)
	type Tuser struct {
		Username string `gvalid:"username@required|length:6,30#请输入用户名称|用户名称长度非法"`
		Password string `gvalid:"password@required|length:6,30#请输入密码|密码长度为6-30"`
	}
	tuser := new(Tuser)
	r.GetToStruct(tuser)

	// 验证数据
	if err := gvalid.CheckStruct(tuser, nil); err != nil {
		// r.Response.WriteJson(err.String())
		lib_res.Refail(r, 40030, err.String())
		// r.ExitAll()
	}

	// uinfo,err :=user.FindByName(tuser.Username)

	glog.Println("---100---")

	// 如找到用户则直接返回
	// if uinfo != nil {
	// 	lib_res.Refail(r, 40001,fmt.Sprint(err))
	// 	// r.ExitAll()
	// }

	// 添加用户

	Inuser := new(user.UserInfo)

	Inuser.Username = tuser.Username
	// Inuser.Password = tuser.Password

	// encrypt 参数
	salt := cdb.Salt(24, false)
	// secretPassword := cdb.Encrypt(interface{}(tuser.Password).(string), salt)
	secretPassword := cdb.Encrypt(tuser.Password, salt)
	Inuser.Password = secretPassword
	Inuser.Salt = salt

	err := user.InsertUser(Inuser)

	if err != nil {
		model.DoLog(err)

		if gstr.Contains(err.Error(), "duplicate") {
			lib_res.Refail(r, 40001, "用户名已经注册,请重新输入")
		} else {
			lib_res.Refail(r, 40030, fmt.Sprint(err))
		}

		// lib_res.Refail(r, 40030, fmt.Sprint(err))
	}

	lib_res.Json(r, 200, "done")

}

// 用户登录 /user/SignUp
func (c *UserController) SignUp(r *ghttp.Request) {
	glog.Println("SignUp ")
	type Tuser struct {
		Username string `gvalid:"username@required|length:6,30#请输入用户名称|用户名称长度非法"`
		Password string `gvalid:"password@required|length:6,30#请输入密码|密码长度为6-30"`
	}
	tuser := new(Tuser)
	r.GetToStruct(tuser)

	// 验证数据
	if err := gvalid.CheckStruct(tuser, nil); err != nil {
		// r.Response.WriteJson(err.String())
		lib_res.Refail(r, 40030, err.String())
		// r.ExitAll()
	}

	uinfo, err := user.FindByName(tuser.Username)

	if err != nil {
		lib_res.Refail(r, 40030, fmt.Sprint(err))
		// r.ExitAll()
	}

	// encrypt 参数
	// salt := cdb.Salt(48, false)
	// secretPassword := cdb.Encrypt(interface{}(tuser.Password).(string), salt)
	secretPassword := cdb.Encrypt(tuser.Password, uinfo.Salt)

	// 密码错误
	if secretPassword != uinfo.Password {
		lib_res.Refail(r, 40002, "")
	}

	// 转换参数
	euser := gconv.MapDeep(uinfo)
	delete(euser,"Password")
	delete(euser,"Salt")


	// 生成token
	token, err := util.CreateTokenByName(uinfo.Username)


	// 生成 token 失败
	if err != nil {
		model.DoLog(err)
	}


	//登录成功
	lib_res.Json(r, 200,"done",g.Map{
		"token":token,
		"userinfo":euser,
	})


}

// 修改用户信息 /user/UpdateUser
func (c *UserController) UpdateUser(r *ghttp.Request) {
	glog.Println("UpdateUser ")


}

// 获取用户列表 /user/GetList
func (c *UserController) GetList(r *ghttp.Request) {
	glog.Println("GetList ")

	Page := r.GetPostInt("Page")
	PageCount := r.GetPostInt("PageCount")

	ulist := new(user.UserList)

	ulist.Page = Page
	ulist.PageCount = PageCount

	err := ulist.GetListPage(bson.M{},bson.M{"Password":0,"Salt":0,})

	if err != nil {
		model.DoLog(err)
	}

	lib_res.Json(r,200, "done", gconv.Map(ulist))

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
