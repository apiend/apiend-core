/*
    fileName: qiniu
    author: diogoxiang@qq.com
    date: 2019/7/24
*/
package qiniu

import (
	"github.com/qiniu/api.v7/auth/qbox"
	"github.com/qiniu/api.v7/storage"
)

var (
	accessKey = "pUr2s6VcRoB8T0MKe6dmmlhDpQcqacnbJollpdL4"
	secretKey = "19DhOANFb9vAkb_9o51NPIWdCuvurwjGqDojN6Gl"
	bucket    = "apiend"
)

func CreateTokenQiniu()  string {
	// 简单上传凭证
	putPolicy := storage.PutPolicy{
		Scope: bucket,
		ReturnBody: `{"key":"$(key)","hash":"$(etag)","fsize":$(fsize),"bucket":"$(bucket)","name":"$(x:name)"}`,
	}
	putPolicy.Expires = 7200
	mac := qbox.NewMac(accessKey, secretKey)
	upToken := putPolicy.UploadToken(mac)

	return upToken
}