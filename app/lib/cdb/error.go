/*
   fileName: cdb
   author: diogoxiang@qq.com
   date: 2019/6/27
*/
package cdb

type OrmError string

func (e OrmError) Error() string {
	return string(e)
}

// 检查错误
// 如果err等于nil，会调用panic函数，在调用的函数中需要用recover处理错误信息
func CheckErr(err error) {
	if err != nil {
		panic(OrmError(err.Error()))
	}
}
