/*
    fileName: project
    author: diogoxiang@qq.com
    date: 2019/7/9
*/
package project

import (
	"apiend-core/app/model/user"
	"fmt"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"github.com/gogf/gf/g"
	"github.com/gogf/gf/g/os/gtime"
	"testing"
	"time"
)

// 测试插入项目
func TestInsert(t *testing.T) {
	t.Log("插入项目")

	model := new(Project)
	model.ProjectName = "project2019aa"
	umodel :=new(user.UserInfo)
	umodel.Uid =10028
	err := Insert(model,umodel)
	if err != nil {
		t.Error(err)
	}

}

// 新建项目并添加其他组员
func TestInsertAddUser(t *testing.T) {

	// ObjectId("5d1da8fd8a5edb09980a7f2a")
	model := new(Project)
	model.ProjectName = "project2019Q"

	usersArr := []*mgo.DBRef{}

	eid := "5d1da8fd8a5edb09980a7f2a"
	var one =  &mgo.DBRef{
		Collection: user.CollectionName,
		Id:bson.ObjectIdHex(eid),
	}

	usersArr=append(usersArr,one)

	model.Users = usersArr


	umodel :=new(user.UserInfo)
	umodel.Uid =10028

	err := Insert(model,umodel)
	if err != nil {
		t.Error(err)
	}

	g.Dump(model)
}

func TestInsertAddUserList(t *testing.T)  {
	model := new(Project)
	model.ProjectName = "project2019Q712"


	usersArr := []*UserArray{}

	// 模拟添加10个用户
	for i := 0; i < 10; i++ {
		userA := &UserArray{}
		userA.Uid = 10027
		userA.Uname = fmt.Sprintf("diogoxiang_%d",i)
		usersArr = append(usersArr,userA)
	}


	model.UserList = usersArr

	umodel :=new(user.UserInfo)
	umodel.Uid =10027





	err := Insert(model,umodel)
	if err != nil {
		t.Error(err)
	}

	g.Dump(model)

}


// 单个查询
func TestFindByid(t *testing.T) {
	// 模拟ID ObjectId("5d2460b58a5edb0f00e81c58") ObjectId("5d2460b58a5edb0f00e81c58")
	var sid = map[string]interface{}{
		"id":"5d2460b58a5edb0f00e81c58",
		// "Pid":5,
	}
	Pinfo,err:=FindByid(sid)

	t.Log(Pinfo)
	t.Log(err)
}

// 批量条件查询
func TestFindBySearch(t *testing.T) {
	selector := bson.M{
		"Uid":10028,
	}
	// 这个用来控制 页面跳转
	Skip := 0
	limter :=10


	prolist,err := FindBySearch(selector,nil,Skip,limter,"-CreatedAt")

	// t.Log(g.)
	t.Log(err)
	g.Dump(prolist)

}

// 根据用户列表中的某一个参数.查找数据
func TestFindbyUserList(t *testing.T)  {
	var orr []bson.M

	orr1 :=bson.M{
		"Uid":1,
	}

	orr = append(orr,orr1)

	orr2 :=bson.M{
		"UserList.uname":"diogoxiang_0",
	}
	orr = append(orr,orr2)
	selector := bson.M{
		"$or": orr,
	}


	Skip := 0
	limter :=10


	prolist,err := FindBySearch(selector,nil,Skip,limter,"-CreatedAt")

	// t.Log(g.)
	t.Log(err)
	g.Dump(prolist)

}


// 软 删除
func TestDelOne(t *testing.T) {

	// 根据 Pid 删除
	selector := bson.M{
			"Pid":11,
	}

	err := DelOne(selector)

	t.Log(err)

}

// 批量删除 根据日间搜索
func TestDelAll(t *testing.T) {
	cstLocal, _ := time.LoadLocation("Asia/Shanghai")
	etime,_ :=gtime.StrToTime("2019-07-16T11:24:33.642")
	t.Log(etime.In(cstLocal))
	// t.Log(time.Now().UTC())
	selector := bson.M{
			"CreatedAt":bson.M{
				"$gt":etime.In(cstLocal),
			},
	}


	num,err := DelAll(selector)

	t.Log(err)
	t.Log(num)
}