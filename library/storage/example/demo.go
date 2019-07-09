/*
    fileName: example
    author: diogoxiang@qq.com
    date: 2019/7/7
*/
package main

import (
	"apiend-core/library/storage"
	"fmt"
)

func main()  {

	setDB()
}

func setDB()  {

	db:=storage.NewBadgerStore("/tmp/badger")

	db.Set([]byte("10"),[]byte("100"))

	info,err:=db.Get([]byte("10"))

	if err{
		fmt.Println(err)
	}

	fmt.Println(info)

}