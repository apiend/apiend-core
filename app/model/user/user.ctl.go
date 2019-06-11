/*
   fileName: mdb
   author: diogoxiang
   date: 2019/6/10
*/
package user

import (
	"apiend-core/core/conn"
	"github.com/globalsign/mgo"
	"github.com/gogf/gf/g"
)

var (
	c = g.Config()
	// dataBase   = c.GetString("setting.mgoDbName")
	mongoURL      = c.GetString("mongo.mongoUrl")
	mongoPoolSize = c.GetInt("mongo.mgoPoolSize")
	dbName        = c.GetString("mongo.mgoDbName")
)

const (
	// 用户信息
	CollectionName = "c_user"
	// 自增
	Counter = "counters"
)

//  创建用户
func (user *UserInfo) Create(doc interface{}) (err error) {

	expected := &UserInfo{}
	expected.SetFieldsValue()

	conn.GetMgoPool(dbName).Exec(CollectionName, func(c *mgo.Collection) {
		err = c.Insert(expected)
	})

	return err
}
