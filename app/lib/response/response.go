package lib_res

import (
	"github.com/gogf/gf/g"
	"github.com/gogf/gf/g/net/ghttp"
	"github.com/gogf/gf/g/util/gconv"
)

// 标准返回结果数据结构封装。
// 返回固定数据结构的JSON:
// err:  错误码(200:成功, 201:失败, >200:错误码);
// msg:  请求结果信息;
// data: 请求结果,根据不同接口返回结果的数据结构不同;
func Json(r *ghttp.Request, err int, msg string, data ...interface{}) {
	responseData := interface{}(nil)
	if len(data) > 0 {
		responseData = data[0]
	}
	r.Response.WriteJson(g.Map{
		"code": err,
		"msg":  msg,
		"data": responseData,
	})
	r.Exit()
}

// 返回错误 信息
func Refail(r *ghttp.Request, Code int, msg string, data ...interface{}) {
	responseData := interface{}(nil)
	if len(data) > 0 {
		responseData = data[0]
	}
	// 统一 返回码
	errCode := gconv.String(Code)
	msgInfo := gconv.String(ReturnCode[errCode])
	r.Response.WriteJson(g.Map{
		"code": Code,
		"msg":  msgInfo+";" + msg,
		"data": responseData,
	})
	r.Exit()

}
