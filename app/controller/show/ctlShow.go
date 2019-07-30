/*
    fileName: show
    author: diogoxiang@qq.com
    date: 2019/7/23
	//  不做限制模块.主要只是用来做 show

*/
package show

import (
	"apiend-core/app/lib/response"
	"apiend-core/app/model/mock"
	"apiend-core/app/model/project"
	"encoding/hex"
	"github.com/globalsign/mgo/bson"
	"github.com/gogf/gf/g/encoding/gjson"
	"github.com/gogf/gf/g/net/ghttp"
)

type ShowMock struct{}

// 主返回的 路由控制  /项目id/项目路径前缀/mock路径
func (sw *ShowMock) Show(r *ghttp.Request) {

	pid := r.GetString("pid")

	// _id 是否为真判断
	d, err := hex.DecodeString(pid)
	if err != nil || len(d) != 12 {
		// panic(fmt.Sprintf("invalid input to ObjectIdHex: %q", s))
		r.Response.Write(`未找到指定的项目[30001]`)
		r.Exit()
	}
	// 优先处理

	selector :=bson.M{
		"_id":bson.ObjectId(d),
	}

	pinfo, err := project.FindBySearchOne(selector)

	ppath := r.GetRouterString("ppath")

	if pinfo.ProjectUrl != ppath || err !=nil {
		// lib_res.Refail(r,)
		r.Response.Write(`未找到指定的Mock[30001]`)
		r.Exit()
	}

	// 校验是否为禁用项目 0 和1 为正常  -1为禁止
	if pinfo.Status == project.ProBanned {
		r.Response.Write(lib_res.ReturnCode["48004"])
		r.Exit()
	}


	cRouter := r.GetRouterString("any")

	selector = bson.M{
		"MockUrl":cRouter,
	}

	mockinfo,err := mock.FindBySearchOne(selector)
	// r.Response.Write(pid)

	if err != nil {
		// lib_res.Refail(r, 40032, err.Error())
		r.Response.Write(`未找到指定的Mock[30001]`)
		r.Exit()
	}

	if gjson.Valid(mockinfo.MockResponse) == true {


		// resjson, _ := (gconv.Bytes(mockinfo.MockResponse))
		temp,_ := gjson.DecodeToJson(mockinfo.MockResponse)

		r.Response.WriteJson(temp.ToMap())
		r.Exit()

	} else {
		r.Response.Write(mockinfo.MockResponse)
		r.Exit()
	}


}