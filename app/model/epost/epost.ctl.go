/*
    fileName: epost
    author: diogoxiang@qq.com
    date: 2019/7/30
*/
package epost

import (
	"apiend-core/app/lib/cdb"
	"apiend-core/app/model"
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
