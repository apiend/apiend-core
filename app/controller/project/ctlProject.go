/*
   fileName: project
   author: diogoxiang@qq.com
   date: 2019/7/15
*/
package ctlProject

import (
	"apiend-core/app/controller"
	"apiend-core/app/lib/response"
	"apiend-core/app/model"
	"apiend-core/app/model/project"
	"apiend-core/app/model/user"
	"fmt"
	"github.com/globalsign/mgo/bson"
	"github.com/gogf/gf/g"
	"github.com/gogf/gf/g/net/ghttp"
	"github.com/gogf/gf/g/os/glog"
	"github.com/gogf/gf/g/os/gtime"
	"github.com/gogf/gf/g/util/gconv"
)

//  管理对象
type ProController struct{}

// 公用方法 可在所有方法调用之前
func (c *ProController) Init(r *ghttp.Request) {
	glog.Println("Project init ")
}

// 公用方法 方法执行完了之后 调用
func (c *ProController) Shut(r *ghttp.Request) {
	glog.Println("Project over ")
}

// 新建项目
func (c *ProController) CreateProject(r *ghttp.Request) {
	// 新加project
	uid := r.GetPostInt("uid") // 获取用户的ID
	utoken := r.GetPostString("userToken")

	// 数据校验
	// type Tuser struct {
	// 	Username string `gvalid:"username@required|length:6,30#请输入用户名称|用户名称长度非法"`
	// 	Password string `gvalid:"password@required|length:6,30#请输入密码|密码长度为6-30"`
	// }
	// tuser := new(Tuser)


	// 验证token 是否合法
	if controller.CheckToken(utoken) != nil {
		lib_res.Refail(r, 40003, "用户信息错误")
	}

	proInfo := new(project.Project)

	proInfo.ProjectName = r.GetPostString("projectName")
	proInfo.Description = r.GetPostString("projectDes")
	proInfo.ProjectUrl = r.GetPostString("projectUrl")


	// TODO: 模拟数据
	mUser := new(user.UserInfo)
	mUser.Uid = uid

	err := project.Insert(proInfo,mUser)
	/**
		新建的项目Name 与 Url  做唯一索引
	 */

	if err != nil {
		model.DoLog(err)
		lib_res.Refail(r, 40031, fmt.Sprint(err))
	}

	lib_res.Json(r, 200, "done",g.Map{})

}

// 修改及编辑项目
func (c *ProController) UpdataProject(r *ghttp.Request) {

	// 修改信息project
	pid := r.GetPostInt("pid") // 获取用户的ID
	utoken := r.GetPostString("userToken")
	// 验证token 是否合法
	if controller.CheckToken(utoken) != nil {
		lib_res.Refail(r, 40003, "用户信息错误")
	}

	selectM := bson.M{
		"Pid":pid,
	}

	proInfo, err := project.FindBySearchOne(selectM)

	if err != nil {
		model.DoLog(err)
		lib_res.Refail(r, 40032, fmt.Sprint(err))
	}

	// var setM bson.M
	var setD = bson.M{}



	 ProjectName := r.GetPostString("projectName")
	 Description := r.GetPostString("projectDes")
	 ProjectUrl := r.GetPostString("projectUrl")

	 // util.If()
	 // setD['ProjectName'] =ProjectName
	 if ProjectName != "" {
	 	// pname := bson.M{
	 	// 	"ProjectName":ProjectName,
		// }
	 	setD["ProjectName"] =ProjectName
	 	proInfo.ProjectName = ProjectName
	 	// setD=append(setD,pname)
	 }

	if Description != "" {
		// temp := bson.M{
		// 	"Description":Description,
		// }
		// setD=append(setD,temp)
		setD["Description"] = Description
		proInfo.Description = Description
	}

	if ProjectUrl != "" {
		// temp := bson.M{
		// 	"ProjectUrl":ProjectUrl,
		// }
		// setD=append(setD,temp)
		setD["ProjectUrl"] = ProjectUrl
		proInfo.ProjectUrl = ProjectUrl
	}


	changeM :=bson.M{
		"$set":setD,
	}

	 err = project.UpdateOne(selectM,changeM)


	if err != nil {
		model.DoLog(err)
		lib_res.Refail(r, 40031, fmt.Sprint(err))
	}

	 proInfo.UpdatedAt = gtime.Now().Time
	// lib_res.Json(r, 200, "done",g.Map{})
	// 重新数据
	// proInfo.ProjectName = r.GetPostString("projectName")
	// proInfo.Description = r.GetPostString("projectDes")
	// proInfo.ProjectUrl = r.GetPostString("projectUrl")

	// fmt.Println(proInfo)
	lib_res.Json(r, 200, "done",gconv.Map(proInfo))



}

// 查询项目 by id 查询项目
func (c *ProController) QueryProjectById(r *ghttp.Request) {
	// 修改信息project
	aid := r.GetPostString("id") // 获取项目的ID
	utoken := r.GetPostString("userToken")
	// 验证token 是否合法
	if controller.CheckToken(utoken) != nil {
		lib_res.Refail(r, 40003, "用户信息错误")
	}

	selectM := bson.M{
		"_id":bson.ObjectIdHex(aid),
	}

	proInfo, err := project.FindBySearchOne(selectM)

	if err != nil {
		model.DoLog(err)
		lib_res.Refail(r, 40032, fmt.Sprint(err))
	}
	lib_res.Json(r, 200, "done", gconv.Map(proInfo))
}

// 根据用户Uid 查询数据 带列表分页
func (c *ProController) QueryProjectListByUid(r *ghttp.Request) {
	// 获取参数
	Page := r.GetPostInt("page")
	PageCount := r.GetPostInt("pageCount")
	uid := r.GetPostInt("uid")
	userToken := r.GetPostString("userToken")
	// 参数的判断与获取
	if "" == userToken || Page == 0 || PageCount == 0 || uid == 0 {
		// lib_response.Refail(r, 40035, "")
		// r.ExitAll()
		lib_res.Refail(r, 40035, "")
	}

	proList := project.ProList{Page:Page,PageCount:PageCount}

	selectM :=bson.M{
		"Uid":uid,
	}
	err := proList.GetListPage(selectM,nil)

	if err != nil {
		model.DoLog(err)
		lib_res.Refail(r, 40032, fmt.Sprint(err))
	}


	// 正常返回
	lib_res.Json(r, 200, "done", gconv.Map(proList))


}