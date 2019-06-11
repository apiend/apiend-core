/*
   fileName: util
   author: diogoxiang
   date: 2019/4/25
*/
package util

import (
	"testing"
	"unsafe"
)

func TestCreateToken(t *testing.T) {

	key, err := CreateToken([]byte("diogoxiang"))

	t.Log(key)
	t.Log(err)
	// one 81aeb256409e42a180cca3a0eb84d06c
}

func toString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

func TestCheckTokenKey(t *testing.T) {

	v, exists, err := CheckTokenKey("diogo9")

	temp := toString(v)

	//t.Log(v)
	t.Logf("%v", temp)
	t.Log(exists)
	t.Log(err)

}
