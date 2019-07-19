/*
    fileName: mock
    author: diogoxiang@qq.com
    date: 2019/7/15
	控制器

*/
package mock

import (
	"apiend-core/app/lib/cdb"
	"apiend-core/app/model"
	"apiend-core/app/model/project"
	"github.com/globalsign/mgo/bson"
)

const (
	// 用户信息 collection
	CollectionName = "c_mock"
	// 自增 collection
	Counter = "counters"
	// 自增 name
	CounterName = "mock_ids"
)


// 新建 mock
func Insert(m *MockInfo, p *project.Project) error {

	m.SetFieldsValue()
	m.Uid = p.Uid
	m.Pid  = p.Pid
	m.MockId,_ = cdb.GetAutoId(Counter, CounterName)

	// 把项目的根路径
	m.MockPreUrl = p.ProjectUrl

	err := cdb.Insert(CollectionName, m)

	if err != nil {
		// glog.Errorf("%s", err)
		model.DoLog(err)
		return err
	}
	return nil

}

//根据条件查询一个
func FindBySearchOne(selector bson.M) (*MockInfo,error)  {

	return nil,nil
}

// 根据条件查询所有数据
func FindBySearchAll(selector bson.M, fields bson.M, skip int, limit int, sort ...string) ([]MockInfo,error)  {

	return nil, nil
}

/**
	Update 更新项目信息 one
 */
func UpdateOne(selector bson.M, update bson.M) error {

	err := cdb.UpdateOne(CollectionName,selector,update)
	if err != nil {
		model.DoLog(err)
		// panic(err)
		return err
	}

	return nil
}

/**
	Update 更新项目信息 all
	返回 修改的数量   Num   错误 error
 */
func UpdateAll(selector bson.M, update bson.M) (int,error) {

	Num,err := cdb.UpdateAll(CollectionName,selector,update)

	if err != nil {
		model.DoLog(err)
		// panic(err)
		return 0,err
	}
	return Num,nil
}

// 导入 mock 数据集合
func ImportList()  {

}