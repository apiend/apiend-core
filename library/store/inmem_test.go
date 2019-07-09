/*
    fileName: store
    author: diogoxiang@qq.com
    date: 2019/7/8
*/
package store


import (
	. "github.com/glycerine/goconvey/convey"
	"testing"
)

func TestInMemStore(t *testing.T) {
	Convey("Having inmemory store", t, func() {
		testStore(NewInmemStore())
	})
}