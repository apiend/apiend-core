/*
    fileName: role
    author: diogoxiang@qq.com
    date: 2019/6/30
	用户权限的拆分
	1. 用户所署team  所在组 拥用的权限  Role

*/
package role

import (
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"time"
)


// 团队
type Team struct {
	Id       bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Users    []*mgo.DBRef
	TeamName string
	Motto    string
	RoleList []Role  // 权限集合
}

// 权限
type Role struct {
	Id          bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Title       string
	ImageUrl    string
	TimeStamp   time.Time
	Deadline    time.Time

}