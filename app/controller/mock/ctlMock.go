/*
   fileName: mock
   author: diogoxiang@qq.com
   date: 2019/7/18
*/
package mock

import (
	"apiend-core/app/controller"
	"apiend-core/app/lib/response"
	"apiend-core/app/model"
	"apiend-core/app/model/mock"
	"apiend-core/app/model/project"
	"fmt"
	"github.com/globalsign/mgo/bson"
	"github.com/gogf/gf/g"
	"github.com/gogf/gf/g/net/ghttp"
	"github.com/gogf/gf/g/os/glog"
)

//  管理对象
type MockController struct{}

// 公用方法 可在所有方法调用之前
func (m *MockController) Init(r *ghttp.Request) {
	glog.Println("mock init ")
}

// 公用方法 方法执行完了之后 调用
func (m *MockController) Shut(r *ghttp.Request) {
	glog.Println("mock over ")
}

// 验证 request 的参数据或是数据
func checkRequest(data... interface{}) bool{

	for key, value := range  data {

		switch value {

		case "userToken":
			fmt.Println(value)

		default:

			fmt.Println(key)

		}

	}

	return false
}


// 新建 mock
func (m *MockController) CreateMock(r *ghttp.Request) {

	// 项目ID
	upid := r.GetPostInt("pid")
	userToken := r.GetPostString("userToken")

	if "" == userToken || upid == 0 {
		lib_res.Refail(r, 40035, "")

	}

	// 可获取 用户的信息
	err := controller.CheckToken(userToken)
	if err != nil {
		lib_res.Refail(r, 40003, err.Error())
	}

	// 获取项目信息
	pInfo,err  := project.FindBySearchOne(bson.M{"Pid":upid})
	if err != nil {
		lib_res.Refail(r, 40032, err.Error())
	}

	mockinfo := new(mock.MockInfo)

	// 基础参数
	// mockinfo.Pid = pInfo.Pid
	// mockinfo.Uid = uinfo.GetInt("Uid")
	// mockinfo.MockPreUrl = pInfo.ProjectUrl


	// 提交的参数
	mockinfo.MockMode = r.GetPostString("MockMode")
	mockinfo.MockUrl = r.GetPostString("MockUrl")
	mockinfo.MockDescription = r.GetPostString("MockDescription")
	mockinfo.MockResponse = r.GetPostString("MockResponse")
	mockinfo.MockResponseMode = -1

	err = mock.Insert(mockinfo,pInfo)


	if err !=nil {
		model.DoLog(err)
		lib_res.Refail(r, 40031, fmt.Sprint(err))
	}

	// 新建成功
	lib_res.Json(r, 200, "done",g.Map{})




	// glog.Println(err)
	// glog.Println(uinfo.Value())
	// lib_res.Json(r,200,"",uinfo.Value())



}

// 根据条件 查询 mock , 默认是根据 用户Uid
func (m *MockController) FindOne(r *ghttp.Request)  {




}

// 根据 path 路径查询数据
func (m *MockController) FindOnePath(r *ghttp.Request) {

}

// 获取列表数据
func (m *MockController) FindAllList(r *ghttp.Request)  {

	// 获取参数
	Page := r.GetPostInt("page")
	PageCount := r.GetPostInt("pageCount")
	pid := r.GetPostInt("pid")
	userToken := r.GetPostString("userToken")
	// 参数的判断与获取
	if "" == userToken || Page == 0 || PageCount == 0 || pid == 0 {
		lib_res.Refail(r, 40035, "")
		// r.ExitAll()
	}



	//初始化
	// mocklist := mock.MockList{Page:Page,PageCount:PageCount}

	// 根据项目PID选择
	selector := bson.M{
		"Pid":pid,
	}

	// 这个用来控制 页面跳转
	// Skip := 0
	// limter :=10


	mockList,err  := mock.FindBySearchAll(selector,nil,Page,PageCount,"-CreatedAt")


	if err !=nil {
		model.DoLog(err)
		lib_res.Refail(r, 40032, fmt.Sprint(err))
	}

	lib_res.Json(r, 200, "ok", g.Map{
		"mockList": mockList,
	})

}

// 根据mid 来修改数据
func (m *MockController) UpdataOne(r *ghttp.Request)  {

	mid := r.GetPostInt("mid")
	userToken := r.GetPostString("userToken")
	// 参数的判断与获取
	if "" == userToken || mid == 0 {
		lib_res.Refail(r, 40035, "")
	}

	selector := bson.M{
		"MockId":mid,
	}
	// 查询数据
	mockInfo,err := mock.FindBySearchOne(selector)

	if err !=nil || mockInfo.MockId ==0 {
		model.DoLog(err)
		lib_res.Refail(r, 40010, fmt.Sprint(err))
	}


	// 可以修改的参数
	MockMode := r.GetPostString("MockMode")
	MockUrl := r.GetPostString("MockUrl")
	MockDescription := r.GetPostString("MockDescription")
	MockResponse := r.GetPostString("MockResponse")

	// var setM bson.M
	var setD = bson.M{}

	if MockMode != "" {
		setD["MockMode"] = MockMode
	}

	if MockUrl != "" {
		setD["MockUrl"] = MockUrl
	}

	if MockDescription != ""{
		setD["MockDescription"] = MockDescription
	}

	if MockResponse != "" {
		setD["MockResponse"] = MockResponse
	}

	changeM := bson.M{
		"$set":setD,
	}

	err = mock.UpdateOne(selector,changeM)

	if err !=nil {
		model.DoLog(err)
		lib_res.Refail(r, 40032, fmt.Sprint(err))
	}

	lib_res.Json(r, 200, "ok","")


}

// 软删除 数据
func (m  *MockController) DelOne(r *ghttp.Request)  {

}

// 软删除 列表
func (m *MockController) DelAll(r *ghttp.Request)  {

}


