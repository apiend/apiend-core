/*
    fileName: gb
    author: diogoxiang@qq.com
    date: 2019/7/26
*/
package gb

import "testing"

func TestNewInmemStore(t *testing.T) {

		estore := InitStore(NewInmemStore())


		estore.Set("k1",[]byte("100"))

		t.Log(estore.Get("k1"))

}