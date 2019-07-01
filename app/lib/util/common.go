package util

import (
	"github.com/gogf/gf/g/net/ghttp"
	"github.com/gogf/gf/g/os/glog"
)

// 主要用来做一些判断用
//a, b := 2, 3
//max := If(a > b, a, b).(int)
//println(max)
func If(condition bool, trueVal, falseVal interface{}) interface{} {
	if condition {
		return trueVal
	}
	return falseVal
}

// RawQuery 将map转换为url的query字符串
func RawQuery(q map[string]string) string {
	var str string
	first := true
	for k, v := range q {
		if first {
			str += k + "=" + v
			first = false
			continue
		}
		str += "&" + k + "=" + v
	}
	return str
}

// 公用的错误处理方案
func Fill(err error, r *ghttp.Request) {
	// 打印错误
	glog.Error(err)
	r.Exit()
}

// 批量判断是否有相应的参数 如果为空则 返回过 nil
//func QueryFieldsString(p ...string) error {
//
//	r := make(map[string]interface{},len(p))
//
//	for _, s := range q  {
//		r[s] = 1
//	}
//
//	return nil
//}
