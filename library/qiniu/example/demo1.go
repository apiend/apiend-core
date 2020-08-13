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
	// "strings"
)

var (
	// accessKey = "pUr2s6VcRoB8T0MKe6dmmlhDpQcqacnbJollpdL4"
	// secretKey = "19DhOANFb9vAkb_9o51NPIWdCuvurwjGqDojN6Gl"

	accessKey = "U-Fn1-4RmkrzBVs-CtzxFFBQst_-_rQpBD2Yalq8"
	secretKey = "TU6-C0kJX5pSJncunKhWqpEWApVtk_fMI_IUTe13"
	bucket    = "apiend"
)

func main() {

	// 简单上传凭证
	putPolicy := storage.PutPolicy{
		Scope: bucket,
		// ReturnBody: `{"key":"$(key)","hash":"$(etag)","fsize":$(fsize),"bucket":"$(bucket)","name":"$(x:name)"}`,
	}
	putPolicy.Expires = 7200
	mac := qbox.NewMac(accessKey, secretKey)
	upToken := putPolicy.UploadToken(mac)
	fmt.Println(upToken)

}
