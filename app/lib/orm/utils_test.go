/*
    fileName: orm
    author: diogoxiang@qq.com
    date: 2019/6/17
*/
package orm



import (
	"testing"
)

type MyUser struct {

}

func TestGetCName(t *testing.T) {
	var a interface{}
	a=&[]MyUser{}
	if getCName(a) != "MyUser" {
		t.Error("GetCName测试失败")
	}
}
