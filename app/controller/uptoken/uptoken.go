/*
    fileName: uptoken
    author: diogoxiang@qq.com
    date: 2019/7/24
*/
package uptoken

import (
	"apiend-core/app/lib/response"
	"apiend-core/library/qiniu"
	"github.com/gogf/gf/g"
	"github.com/gogf/gf/g/net/ghttp"
)

type Uptoken struct {}

func (up *Uptoken) Uptoken(r *ghttp.Request) {

		utoken := qiniu.CreateTokenQiniu()
		lib_res.Json(r,200,"done",g.Map{
			"uptoken":utoken,
		})

}