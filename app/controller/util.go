/*
    fileName: controller
    author: diogoxiang@qq.com
    date: 2019/6/30
*/
package controller

import (
	"github.com/globalsign/mgo/bson"
	"github.com/gogf/gf/g/net/ghttp"
	"io/ioutil"
)

func TranspilePostParams(c *ghttp.Request) bson.M {
	data, _ := ioutil.ReadAll(c.Request.Body)
	var params = bson.M{}
	bson.UnmarshalJSON(data, &params)
	return params
}