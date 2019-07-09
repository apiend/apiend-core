/*
    fileName: project
    author: diogoxiang@qq.com
    date: 2019/7/9
*/
package project

import (
 	"apiend-core/app/model/user"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"github.com/gogf/gf/g"
	"testing"
)

func TestInsert(t *testing.T) {
	t.Log("插入项目")

	model := new(Project)
	model.ProjectName = "project2019"
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