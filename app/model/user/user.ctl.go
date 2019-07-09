/*
   fileName: user
   author: diogoxiang
   date: 2019/6/10
*/
package user

import (
	"apiend-core/app/lib/cdb"
	"apiend-core/app/model"
	"github.com/globalsign/mgo/bson"
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

/**
InsertUser adds the user to the db
*/
func InsertUser(user *UserInfo) error {

	user.SetFieldsValue()
	user.Uid, _ = cdb.GetAutoId(Counter, CounterName)

	err := cdb.Insert(CollectionName, user)

	if err != nil {
		// glog.Errorf("%s", err)
		model.DoLog(err)
		return err
	}
	return nil
}

//UpdateUser updates a user with the given id and handler made struct
func UpdateUser(selector bson.M, update bson.M) error {
	err := cdb.UpdateOne(CollectionName, selector, update)
	if err != nil {
		// glog.Errorf("%s", err)
		// panic(err)
		model.DoLog(err)
		return err
	}
	return nil
}

//  Find by id hex
func FindById(id string) (*UserInfo, error) {
	oid := bson.ObjectIdHex(id)
	person := new(UserInfo)

	selector := bson.M{
		"_id": oid,
	}
	fielder := bson.M{
		"Password": 0,
	}
	err := cdb.FindOne(CollectionName, person, selector, fielder)


	if err != nil {
		model.DoLog(err)
		// panic(err)
		return nil, err
	}

	return person, nil
}

//  Find by Username
func FindByName(name string) (*UserInfo, error) {

	person := new(UserInfo)

	selector := bson.M{
		"Username": name,
	}
	err := cdb.FindOne(CollectionName, person, selector, nil)

	if err != nil {
		model.DoLog(err)
		// panic(err)
		return nil, err
	}

	return person, nil
}

//  Find by search txt 公用搜索 只搜1个
func FindBySearch(selector bson.M, fields bson.M) (*UserInfo, error) {

	person := new(UserInfo)

	err := cdb.FindOne(CollectionName, person, selector, fields)

	if err != nil {
		model.DoLog(err)
		// panic(err)
		return nil, err
	}
	return person, nil
}

// 获取用户列表
func (list *UserList) GetListPage(selector bson.M, fields bson.M) error{
	Skip := (list.Page - 1) * list.PageCount
	Limitr := list.PageCount
	err := cdb.FindAll(CollectionName,&list.List,selector,fields, Skip, Limitr,"-CreatedAt","-Uid")

	list.TotalNum, err = cdb.Count(CollectionName,selector)

	if err != nil {
		model.DoLog(err)
		// panic(err)
		return   err
	}
	return nil
}