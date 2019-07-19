/*
   fileName: project
   author: diogoxiang@qq.com
   date: 2019/7/9
*/
package project

import (
	"apiend-core/app/lib/cdb"
	"apiend-core/app/model"
	"apiend-core/app/model/user"
	"github.com/globalsign/mgo/bson"
	"github.com/gogf/gf/g/util/gconv"
)

const (
	// 用户信息 collection
	CollectionName = "c_project"
	// 自增 collection
	Counter = "counters"
	// 自增 name
	CounterName = "project_ids"
)

// 新建项目
func Insert(pro *Project, umodel *user.UserInfo) error {
	pro.SetFieldsValue()
	pro.Pid, _ = cdb.GetAutoId(Counter, CounterName)
	pro.Uid = umodel.Uid
	err := cdb.Insert(CollectionName, pro)

	if err != nil {
		// glog.Errorf("%s", err)
		model.DoLog(err)
		return err
	}
	return nil

}

/* sid  bson.id
搜索项目 bson.id 或是用户id 或是 项目ID
查找的单个数据 ,可以组合用.也可以单独用
*/
func FindByid(sid bson.M) (*Project, error) {
	// selectID := interface{}(nil)
	// fmt.Println(sid["id"])
	oid := gconv.String(sid["id"])
	selectM := bson.M{
		"_id": bson.ObjectIdHex(oid),
		// "Uid": gconv.String(sid["Uid"]),
		// "Pid":gconv.String(sid["Pid"]),
	}
	proinfo := new(Project)

	err := cdb.FindOne(CollectionName, proinfo, selectM, nil)

	if err != nil {
		model.DoLog(err)
		// panic(err)
		return nil, err
	}

	return proinfo, nil

}

// 根据条件查找 一个
func FindBySearchOne(selectM bson.M) (*Project, error) {
	// selectID := interface{}(nil)
	// fmt.Println(sid["id"])
	// oid := gconv.String(sid["id"])
	// selectM := bson.M{
	// 	"_id":bson.ObjectIdHex(oid),
	// 	// "Uid": gconv.String(sid["Uid"]),
	// 	// "Pid":gconv.String(sid["Pid"]),
	// }
	proinfo := new(Project)

	err := cdb.FindOne(CollectionName, proinfo, selectM, nil)

	if err != nil {
		model.DoLog(err)
		// panic(err)
		return nil, err
	}

	return proinfo, nil

}

/**
by搜索 信息可以自己定义 返回列表形式
*/
func FindBySearch(selector bson.M, fields bson.M, skip int, limit int, sort ...string) ([]Project, error) {
	proList := []Project{}

	// 当传参数 的时候 默认搜索10个
	if limit == 0 {
		limit = 10
	}

	err := cdb.FindAll(CollectionName, &proList, selector, fields, skip, limit, sort...)

	if err != nil {
		model.DoLog(err)
		// panic(err)
		return nil, err
	}

	return proList, nil

}

/**
Update 更新项目信息 one
*/
func UpdateOne(selector bson.M, update bson.M) error {

	err := cdb.UpdateOne(CollectionName, selector, update)

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
func UpdateAll(selector bson.M, update bson.M) (int, error) {

	Num, err := cdb.UpdateAll(CollectionName, selector, update)

	if err != nil {
		model.DoLog(err)
		// panic(err)
		return 0, err
	}

	return Num, nil
}

// 获取项目列表
func (list *ProList) GetListPage(selector bson.M, fields bson.M) error {
	Skip := (list.Page - 1) * list.PageCount
	Limitr := list.PageCount
	err := cdb.FindAll(CollectionName, &list.List, selector, fields, Skip, Limitr, "-CreatedAt", "-Uid")

	list.TotalNum, err = cdb.Count(CollectionName, selector)

	if err != nil {
		model.DoLog(err)
		// panic(err)
		return err
	}
	return nil
}

// 软 删除数据
func DelOne(selector bson.M) error {

	err := cdb.DeleteOne(CollectionName, selector)

	if err != nil {
		model.DoLog(err)
		// panic(err)
		return err
	}
	return nil

}

// 批量 软 删除
func DelAll(selector bson.M) (int, error) {

	num, err := cdb.DeleteAll(CollectionName, selector)

	if err != nil {
		model.DoLog(err)
		// panic(err)
		return 0, err
	}
	return num, nil

}
