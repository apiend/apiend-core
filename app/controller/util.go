/*
    fileName: controller
    author: diogoxiang@qq.com
    date: 2019/6/30
*/
package controller

import (
	"apiend-core/app/lib/eError"
	"apiend-core/app/lib/util"
	"github.com/gogf/gf/g/encoding/gjson"
	"github.com/gogf/gf/g/os/glog"
)

// func TranspilePostParams(c *ghttp.Request) bson.M {
// 	data, _ := ioutil.ReadAll(c.Request.Body)
// 	var params = bson.M{}
// 	bson.UnmarshalJSON(data, &params)
// 	return params
// }

// 只检测Token 是否合法
func CheckToken(u string) error {
	if "" == u {
		err := eError.NewError(40035, "不合法的参数")
		return err
	}
	fond := util.ValidToken(u)
	if !fond {
		glog.Print("userToken 错误")
		err := eError.NewError(40003, "用户信息错误")
		// lib_response.Refail(r, 40003, "用户信息错误")
		return err
	}
	return nil
}

// 根据 token 获取缓存数据 用户信息
func GetUinfoToken(u string) (info *gjson.Json,  err error) {
	if "" == u {
		err = eError.NewError(40035, "不合法的参数")
		return nil,err
	}
	v, fond, err := util.ValidTokenKey(u)

	if !fond {
		err = eError.NewError(40035, "用户信息错误")
		return nil,err
	}

	temp,err := gjson.DecodeToJson(v)

	// glog.Println(temp.Dump())

	return temp,err
}