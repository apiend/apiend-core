/*
   fileName: eError
   author: diogoxiang
   date: 2019/4/30
*/
package eError

import "testing"

func TestError_Error(t *testing.T) {

	err2 := myErr()

	t.Logf("------my err:%v, %v", err2.Error(), err2.(*Error).ErrCode)
}

func myErr() error {
	err := NewError(2, "err test")
	return err
}
