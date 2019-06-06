/*
    fileName: apiend_core
    author: diogoxiang
    date: 2019/6/6
*/
package main

import (
	_ "apiend-core/boot"
	_ "apiend-core/router"

	"github.com/gogf/gf/g"
)

func main() {
	g.Server().SetServerAgent("Diogo")
	g.Server().Run()
}

