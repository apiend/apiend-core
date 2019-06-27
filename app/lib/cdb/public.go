/*
    fileName: cdb
    author: diogoxiang@qq.com
    date: 2019/6/19
*/
package cdb

import (
	"errors"
	"github.com/globalsign/mgo/bson"
	"time"
)

// ----------------------------------- 公共函数 ----------------------------------------

// 执行更新操作之前先判断有没有$操作符
func checkUpdateContent(update bson.M) error {
	for k := range update {
		if k[0] != '$' {
			return errors.New("update content must start with '$'")
		}
	}
	return nil
}

// excludeDeleted 不包含已删除的
func excludeDeleted(selector bson.M) bson.M {
	selector["deletedAt"] = bson.M{"$exists": false}
	return selector
}

// updatedTime 更新updatedAt时间
func updatedTime(update bson.M) bson.M {
	if v, ok := update["$set"]; ok {
		v.(bson.M)["updatedAt"] = time.Now()
	} else {
		update["$set"] = bson.M{"updatedAt": time.Now()}
	}
	return update
}

// deletedTime 更新deletedAt时间
func deletedTime(update bson.M) bson.M {
	if v, ok := update["$set"]; ok {
		v.(bson.M)["deletedAt"] = time.Now()
	} else {
		update["$set"] = bson.M{"deletedAt": time.Now()}
	}
	return update
}
