/*
   fileName: user
   author: diogoxiang
   date: 2019/6/10
*/
package user

import (
	"apiend-core/core/conn"
	"github.com/gogf/gf/g/os/glog"
	"os"
	"testing"
)

// 设置 db 初始化
func setDB() {
	// Mongodb init
	mgoOption := conn.MgoPoolOption{
		Host:   mongoURL,
		Size:   mongoPoolSize,
		DbName: dbName,
	}
	mgoPool, err := conn.NewMgoPool(mgoOption)
	if err != nil {
		glog.Debugf("connect mongodb: " + dbName + "  fail")
 		os.Exit(1)
	}
	conn.MgoSet(mgoOption.DbName, mgoPool)

}

func TestUserInfo_Create(t *testing.T) {
	setDB()
	model := new(UserInfo)

	err := model.Create(nil)

	if err != nil {
		t.Error(err)
	}

}
