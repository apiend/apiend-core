/*
    fileName: epost
    author: diogoxiang@qq.com
    date: 2019/8/9
*/
package epost

import (
	"apiend-core/app/controller"
	"apiend-core/app/lib/response"
	"apiend-core/app/model"
	"apiend-core/app/model/epost"
	"fmt"
	"github.com/globalsign/mgo/bson"
	"github.com/gogf/gf/g"
	"github.com/gogf/gf/g/net/ghttp"
	"github.com/gogf/gf/g/os/glog"
)

type EpostController struct{}

// 公用方法 可在所有方法调用之前
func (m *EpostController) Init(r *ghttp.Request) {
	glog.Println("Epost init ")
}

// 公用方法 方法执行完了之后 调用
func (m *EpostController) Shut(r *ghttp.Request) {
	glog.Println("Epost over ")
}

// 模块内部
func checkToken(r *ghttp.Request){
	utoken := r.GetPostString("userToken")
	// 验证token 是否合法
	if controller.CheckToken(utoken) != nil {
		lib_res.Refail(r, 40003, "用户信息错误")
	}
}

// ----
// 新建文档
func (m *EpostController) Create(r *ghttp.Request)  {

}

// 新建文档类型
func (m *EpostController) CreateType(r *ghttp.Request)  {
	checkToken(r)
	// 获取参数
	TypeTxt := r.GetPostString("TypeTxt") // type

	doc := new(epost.PostType)

	doc.TypeTxt=TypeTxt

	err := epost.InsertType(doc)

	if err != nil {
		model.DoLog(err)
		lib_res.Refail(r, 40030, fmt.Sprint(err))
	}

	// 返回
	lib_res.Json(r, 200, "ok","")


}


// 查找一个
func (m *EpostController) FindOne(r *ghttp.Request)  {

}


// 查找 all
func (m *EpostController) FindAll(r *ghttp.Request)  {

}


// 查找 all type
func (m *EpostController) FindAllType(r *ghttp.Request)  {

	// 检测
	checkToken(r)

	// 条件
	selector := bson.M{}
	fields :=bson.M{
		"UpdatedAt":0,
		"CreatedAt":0,
	}

	skip :=0
	// 默认是取100
	limit := 100


	TypeList,err := epost.FindBySearchAllType(selector,fields,skip,limit, "-CreatedAt")

	if err != nil {
		model.DoLog(err)
		lib_res.Refail(r, 40030, fmt.Sprint(err))
	}

	// 删除 CreatedAt 为空的. 或是指字定KEY的一些参数
	// var Ulist []map[string]interface{}
	// for _, value := range TypeList{
	// 	vmap := gconv.MapDeep(value)
	// 	Ulist=append(Ulist,vmap)
	// }

	// glog.Print(TypeList)
	lib_res.Json(r, 200, "ok", g.Map{
		"TypeList": TypeList,
	})

}
