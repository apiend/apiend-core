/*
    fileName: util
    author: diogoxiang@qq.com
    date: 2019/6/28
*/
package util

import (
	"testing"
	"errors"
)

func TestErrors(t *testing.T) {
	errs := []error{
		errors.New("error_text1"),
		errors.New("error_text2"),
		nil,
		errors.New("error_text4"),
		errors.New("error_text5"),
		nil,
		errors.New("error_text7"),
	}
	err := Merge(errs...)
	t.Log(err)
	err = Append(err, nil)
	t.Log(err)
	err = Append(err, errs...)
	t.Log(err)
}