/*
   fileName: mock
   author: diogoxiang@qq.com
   date: 2019/7/18
*/
package mock

import (
	"apiend-core/app/controller"
	"apiend-core/app/lib/response"
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
func checkRequest(data interface{}) bool{

	return false
}


// 新建 mock
func (m *MockController) CreateMock(r *ghttp.Request) {

	// upid := r.GetPostInt("pid")
	userToken := r.GetPostString("userToken")
	uinfo, err := controller.GetUinfoToken(userToken)
	if err != nil {
		lib_res.Refail(r, 40003, err.Error())
	}

	glog.Println(err)
	glog.Println(uinfo.Value())
	lib_res.Json(r,200,"",uinfo.Value())

}
