/*
   fileName: cdb
   author: diogoxiang@qq.com
   date: 2019/6/27
*/
package cdb

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
)

// 获取类名字
func getCName(class interface{}) string {
	if class == nil {
		return ""
	}
	name := reflect.TypeOf(class).String()
	arr := strings.Split(name, ".")
	return arr[len(arr)-1]
}

// 检查指定的值是否时nil
// 如果时nil，执行panic
// 调用者需要用recover处理错误
func isNil(doc interface{}) {
	if doc == nil {
		CheckErr(errors.New("doc不能为nil"))
	}
}

var (
	zeroVal  reflect.Value
	zeroArgs []reflect.Value
)

// 调用模型的指定方法
func callToDoc(method string, doc interface{}) error {
	docV := reflect.ValueOf(doc)
	if docV.Kind() != reflect.Ptr {
		e := fmt.Sprintf("ormgo: Passed non-pointer: %v (kind=%v), method:%s", doc, docV.Kind(), method)
		return errors.New(e)
	}
	fn := docV.Elem().Addr().MethodByName(method)
	if fn != zeroVal {
		ret := fn.Call(zeroArgs)
		if len(ret) > 0 && !ret[0].IsNil() {
			return ret[0].Interface().(error)
		}
	}
	return nil
}
