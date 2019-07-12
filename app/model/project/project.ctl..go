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
func Insert(pro *Project, user *user.UserInfo) error {
	pro.SetFieldsValue()
	pro.Pid,_ = cdb.GetAutoId(Counter, CounterName)
	pro.Uid = user.Uid
	err := cdb.Insert(CollectionName, pro)

	if err != nil {
		// glog.Errorf("%s", err)
		model.DoLog(err)
		return err
	}
	return nil

}
// 搜索项目 bson.id 或是用户id 或是 项目ID
/* sid  bson.id

	查找的单个数据 ,可以组合用.也可以单独用
*/
func FindByid(sid bson.M) (*Project, error){
	// selectID := interface{}(nil)
	// fmt.Println(sid["id"])
	oid := gconv.String(sid["id"])
	selectM := bson.M{
		"_id":bson.ObjectIdHex(oid),
		// "Uid": gconv.String(sid["Uid"]),
		// "Pid":gconv.String(sid["Pid"]),
	}
	proinfo := new(Project)

	err := cdb.FindOne(CollectionName,proinfo,selectM,nil)

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
func FindBySearch(selector bson.M, fields bson.M, skip int, limit int, sort ...string)  ([]Project, error){
	proList := []Project{}


	// 当传参数 的时候 默认搜索10个
	if limit ==0{
		limit =10
	}


	err := cdb.FindAll(CollectionName,&proList,selector,fields,skip,limit,sort...)

	if err != nil {
		model.DoLog(err)
		// panic(err)
		return nil, err
	}

	return proList, nil


}