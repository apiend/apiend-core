/*
    fileName: example
    author: diogoxiang@qq.com
    date: 2019/7/24
*/
package main

import (
	// "encoding/base64"
	"fmt"
	"github.com/qiniu/api.v7/auth/qbox"
	"github.com/qiniu/api.v7/storage"
	"os"
	// "strings"


)

var (
	accessKey = os.Getenv("pUr2s6VcRoB8T0MKe6dmmlhDpQcqacnbJollpdL4")
	secretKey = os.Getenv("19DhOANFb9vAkb_9o51NPIWdCuvurwjGqDojN6Gl")
	bucket    = os.Getenv("apiend")
)

func main() {

	// 简单上传凭证
	putPolicy := storage.PutPolicy{
		Scope: bucket,
	}
	mac := qbox.NewMac(accessKey, secretKey)
	upToken := putPolicy.UploadToken(mac)
	fmt.Println(upToken)



}

