/*
   fileName: project
   author: diogoxiang@qq.com
   date: 2019/6/30
*/
package project

import (
	"apiend-core/app/model"
	"github.com/globalsign/mgo"
)

// type Cast struct {
// 	Id bson.ObjectId `json:"id" bson:"_id,omitempty"`
// 	User *mgo.DBRef
// 	Role string
// }

// 项目model
type Project struct {
	model.PublicFields `bson:",inline"` // 公共字段，id和时间
	Pid         int           `bson:"Pid" json:"Pid"` // 当前 项目的ID
	Uid         int           `bson:"Uid" json:"Uid"` // 创建 者用户ID
	ProjectName string        `bson:"ProjectName" json:"ProjectName"`
	ProjectUrl  string        `bson:"ProjectUrl" json:"ProjectUrl"`
	Description string        `bson:"Description" json:"Description"`
	SwaggerUrl  string        `bson:"SwaggerUrl" json:"SwaggerUrl"`
	Status      ProStatus     `bson:"Status" json:"Status"` // 项目状态
	Users       []*mgo.DBRef  `bson:"Users" json:"Users"`   // 当前项目组可编辑人员ID

}

// 项目状态 0 为正常 1为封禁 -1为删除
type ProStatus int

const (
	UserBanned  ProStatus = -1
	UserActived ProStatus = 1
)

func (s ProStatus) Defined() bool {
	switch s {
	case UserBanned, UserActived:
		return true
	}
	return false
}
