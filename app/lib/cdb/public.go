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

// 公用的 Fields
type PublicFields struct {
	ID        bson.ObjectId `bson:"_id,omitempty" json:"id,omitempty"`
	CreatedAt time.Time     `bson:"createdAt,omitempty" json:"createdAt,omitempty"` // 创建时间
	UpdatedAt time.Time     `bson:"updatedAt,omitempty" json:"updatedAt,omitempty"` // 修改时间
	DeletedAt *time.Time    `bson:"deletedAt,omitempty" json:"deletedAt,omitempty"` // 删除时间

}

// SetFieldsValue 设置公共字段值，在插入数据时使用
func (p *PublicFields) SetFieldsValue() {
	now := time.Now().UTC()
	if !p.ID.Valid() {
		p.ID = bson.NewObjectId()
	}
	if p.CreatedAt.IsZero() {
		p.CreatedAt = now
	}
}

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
	selector["DeletedAt"] = bson.M{"$exists": false}
	return selector
}

// updatedTime 更新updatedAt时间
func updatedTime(update bson.M) bson.M {
	if v, ok := update["$set"]; ok {
		v.(bson.M)["UpdatedAt"] = time.Now()

	} else {
		update["$set"] = bson.M{"UpdatedAt": time.Now()}
	}
	return update
}

// deletedTime 更新deletedAt时间
func deletedTime(update bson.M) bson.M {
	if v, ok := update["$set"]; ok {
		v.(bson.M)["DeletedAt"] = time.Now()
	} else {
		update["$set"] = bson.M{"DeletedAt": time.Now()}
	}
	return update
}
