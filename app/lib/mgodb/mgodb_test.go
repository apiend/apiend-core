/*
    fileName: mgodb
    author: diogoxiang
    date: 2019/6/10
*/
package mgodb

import (
	"apiend-core/app/model"
	"github.com/gogf/gf/g"
	"github.com/gogf/gf/g/os/glog"
	"testing"

	"time"
)

var (
	c          = g.Config()
	dataBase   = c.GetString("setting.mgoDbName")
	dbm *Dbm
)


func setDB()  {

	var err error

	connectUrl := c.GetString("setting.mongoUrl")
	timeout, _ := time.ParseDuration("5s")

	err = dbm.Init(connectUrl, dataBase, timeout)
	if err != nil {
		glog.Error("Expected error to be nil %s",err.Error())
	}
}

func TestDbm_Init(t *testing.T)  {
	var dbm *Dbm
	var err error

	connectUrl := c.GetString("setting.mongoUrl")
	timeout, _ := time.ParseDuration("5s")

	err = dbm.Init(connectUrl, dataBase, timeout)
	if err != nil {
		t.Error("Expected error to be nil")
	}
}

func TestDbm_Find(t *testing.T) {



}

func TestDbm_Insert(t *testing.T) {

	setDB()
	doc := new(model.PublicFields)
	err:=dbm.Insert("c_users",doc)

	if err != nil {
		t.Error(err.Error())
	}

}
