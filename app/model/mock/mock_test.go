/*
    fileName: mock
    author: diogoxiang@qq.com
    date: 2019/7/15
*/
package mock

import (
	"apiend-core/app/model/project"
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