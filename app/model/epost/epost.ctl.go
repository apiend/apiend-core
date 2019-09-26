/*
    fileName: epost
    author: diogoxiang@qq.com
    date: 2019/7/30
*/
package epost

import (
	"apiend-core/app/lib/cdb"
	"apiend-core/app/model"
	"github.com/globalsign/mgo/bson"
)

const (
	// 文章分类 collection
	CollectionNameType = "c_epost_type"
	// 文章 collection
	CollectionName = "c_epost"
	// 自增 collection
	Counter = "counters"
	// 自增 name
	CounterNameType = "epost_type_ids"
	CounterName = "epost_ids"
)

// 新建分类
func InsertType(doc *PostType) error {
	doc.SetFieldsValue()
	doc.TypeId, _ = cdb.GetAutoId(Counter, CounterNameType)

	err := cdb.Insert(CollectionNameType, doc)
	if err != nil {
		// glog.Errorf("%s", err)
		model.DoLog(err)
		return err
	}
	return nil

}

//根据条件查询一个
func FindBySearchOneType(selector bson.M) (*PostType,error)  {

	newPostType := new(PostType)

	err := cdb.FindOne(CollectionNameType, newPostType, selector, nil)

	if err != nil {
		model.DoLog(err)
		// panic(err)
		return nil, err
	}

	return newPostType, nil

}

// 根据条件查询所有数据
func FindBySearchAllType(selector bson.M, fields bson.M, skip int, limit int, sort ...string) ([]PostType,error)  {

	proList := []PostType{}

	// 当传参数 的时候 默认搜索10个
	if limit == 0 {
		limit = 10
	}

	err := cdb.FindAll(CollectionNameType, &proList, selector, fields, skip, limit, sort...)

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
func UpdateOneType(selector bson.M, update bson.M) error {

	err := cdb.UpdateOne(CollectionNameType,selector,update)
	if err != nil {
		model.DoLog(err)
		// panic(err)
		return err
	}

	return nil
}

// ---------------------------------------------- Epost

// 插入文档
func InsertEpost(doc *PostDetail)  error{

	doc.SetFieldsValue()
	doc.PostId, _ = cdb.GetAutoId(Counter, CounterName)
	err := cdb.Insert(CollectionName, doc)
	if err != nil {
		// glog.Errorf("%s", err)
		model.DoLog(err)
		return err
	}
	return nil
}


//根据条件查询一个
func FindBySearchOneEpost(selector bson.M) (*PostDetail,error)  {

	newPostDetail := new(PostDetail)

	err := cdb.FindOne(CollectionName, newPostDetail, selector, nil)

	if err != nil {
		model.DoLog(err)
		// panic(err)
		return nil, err
	}

	return newPostDetail, nil

}

// 根据条件查询所有数据
func FindBySearchAllEpost(selector bson.M, fields bson.M, skip int, limit int, sort ...string) ([]PostDetail,error)  {

	proList := []PostDetail{}

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
func UpdateOneEpost(selector bson.M, update bson.M) error {

	err := cdb.UpdateOne(CollectionName,selector,update)
	if err != nil {
		model.DoLog(err)
		// panic(err)
		return err
	}

	return nil
}

/**
	update all epost
	新项目信息 All
 */
func UpdateAllEpost(selector bson.M, update bson.M) (int,error) {

	Num,err := cdb.UpdateAll(CollectionName,selector,update)

	if err != nil {
		model.DoLog(err)
		// panic(err)
		return 0,err
	}
	return Num,nil
}