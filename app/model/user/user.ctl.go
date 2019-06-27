/*
   fileName: mdb
   author: diogoxiang
   date: 2019/6/10
*/
package user

import (
	"apiend-core/app/lib/cdb"
)

const (
	// 用户信息 collection
	CollectionName = "c_user"
	// 自增 collection
	Counter = "counters"
	// 自增 name
	CounterName = "user_ids"
)

var comErr error

//  创建用户
// func NewUser(Uname, Upassword string) *UserInfo {
// 	return &UserInfo{Username: Uname, Password: Upassword}
// }

func (user *UserInfo) Create(doc interface{}) error {

	expected := &UserInfo{
		Username: user.Username,
		Password: user.Password,
		NickName: user.NickName,
		HeadImg:  user.HeadImg,
	}
	expected.SetFieldsValue()
	expected.Uid, _ = cdb.GetAutoId(Counter, CounterName)
	// conn.GetMgoPool(dbName).Exec(CollectionName, func(c *mgo.Collection) {
	// 	comErr = c.Insert(expected)
	// })

	comErr = cdb.Insert(CollectionName, expected)

	return comErr
}
