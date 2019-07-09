/*
   fileName: model
   author: diogoxiang
   date: 2019/6/10
*/
package model

import (
	"errors"
	"github.com/globalsign/mgo/bson"
	"github.com/gogf/gf/g/os/glog"
	"time"
)

// 公用的 Fields
type PublicFields struct {
	ID        bson.ObjectId `bson:"_id,omitempty" json:"id,omitempty"`
	CreatedAt time.Time     `bson:"CreatedAt,omitempty" json:"CreatedAt,omitempty"` // 创建时间
	UpdatedAt time.Time     `bson:"UpdatedAt,omitempty" json:"UpdatedAt,omitempty"` // 修改时间
	DeletedAt *time.Time    `bson:"DeletedAt,omitempty" json:"DeletedAt,omitempty"` // 删除时间

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

// PublicFieldsInt 设置公共字段值，自定id，在插入数据时使用
type PublicFieldsInt struct {
	ID        int64      `bson:"_id" json:"id,string"`                           // 唯一ID
	CreatedAt time.Time  `bson:"createdAt" json:"createdAt"`                     // 创建时间
	UpdatedAt time.Time  `bson:"updatedAt" json:"updatedAt"`                     // 修改时间
	DeletedAt *time.Time `bson:"deletedAt,omitempty" json:"deletedAt,omitempty"` // 删除时间
}

// Init 初始化
func (m *PublicFieldsInt) SetFieldsValue(newID int64) {
	t := time.Now()
	if m.ID == 0 {
		m.ID = newID
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = t
	}
}

// ----------------------------------- 公共函数 ----------------------------------------

// 执行更新操作之前先判断有没有$操作符
func CheckUpdateContent(update bson.M) error {
	for k := range update {
		if k[0] != '$' {
			return errors.New("update content must start with '$'")
		}
	}
	return nil
}

// excludeDeleted 不包含已删除的
func ExcludeDeleted(selector bson.M) bson.M {
	selector["deletedAt"] = bson.M{"$exists": false}
	return selector
}

// updatedTime 更新updatedAt时间
func UpdatedTime(update bson.M) bson.M {
	if v, ok := update["$set"]; ok {
		v.(bson.M)["updatedAt"] = time.Now()
	} else {
		update["$set"] = bson.M{"updatedAt": time.Now()}
	}
	return update
}

// deletedTime 更新deletedAt时间
func DeletedTime(update bson.M) bson.M {
	if v, ok := update["$set"]; ok {
		v.(bson.M)["deletedAt"] = time.Now()
	} else {
		update["$set"] = bson.M{"deletedAt": time.Now()}
	}
	return update
}

// model 写log
func DoLog(err error) {
	if err != nil {
		glog.Printf("%s", err)
		glog.Warning(err)
	}
}

// Query Model
type QueryModel struct {
	Where   string `form:"where"`
	Include string `form:"include"`
	Skip    int    `form:"skip"`
	Limit   int    `form:"limit"`
	Count   int    `form:"count"`
	Order   string `form:"order"`
}
