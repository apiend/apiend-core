/*
    fileName: mock
    author: diogoxiang@qq.com
    date: 2019/7/15
*/
package mock

import (
	"apiend-core/app/model/project"
	"github.com/globalsign/mgo/bson"
	"github.com/gogf/gf/g"
	"testing"
)

// 测试数据插入
func TestInsert(t *testing.T) {

	minfo := new(MockInfo)
	pinfo := new(project.Project)

	minfo.MockUrl = "/getname"
	minfo.MockResponse = "name"
	minfo.MockMethod ="get"
	minfo.MockResponseMode = ResJson

	pinfo.Pid = 27
	pinfo.Uid = 10027
	pinfo.ProjectUrl="aname"

	err := Insert(minfo,pinfo)

	if err != nil {
		t.Error(err)
	}

}

// 根据一个条件.查找一个项目
func TestFindBySearchOne(t *testing.T) {
	selector := bson.M{
		"_id":bson.ObjectIdHex("5d2ec5678a5edb10e8a6da17"),
	}

	Pinfo,err:=FindBySearchOne(selector)

	t.Log(Pinfo)
	t.Log(err)

}

func TestFindBySearchAll(t *testing.T) {

	selector := bson.M{
		"Uid":10027,
	}
	// 这个用来控制 页面跳转
	Skip := 0
	limter :=10


	prolist,err := FindBySearchAll(selector,nil,Skip,limter,"-CreatedAt")

	// t.Log(g.)
	t.Log(err)
	g.Dump(prolist)

}