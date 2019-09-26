/*
   fileName: micache
   author: diogoxiang@qq.com
   date: 2019/8/9
*/
package micache

import (
	"fmt"
	"testing"
)

type User struct {
	Uid      int
	UserName string
}

func TestSet(t *testing.T) {
	getUser := User{20,"diogose"}
	Set("goe", getUser, 3600)

	vstr := Get("goe")

	fmt.Println(vstr)
}

func TestGetDecoding(t *testing.T) {

	getUser := User{}

	vt :=GetDecoding("goe", &getUser)

	fmt.Println(getUser.UserName)

	// vstr := Get("gode")

	fmt.Println(vt)

}
