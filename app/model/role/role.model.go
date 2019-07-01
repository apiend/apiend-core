/*
    fileName: role
    author: diogoxiang@qq.com
    date: 2019/6/30
*/
package role

import (
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"time"
)

//Contest struct
type Contest struct {
	Id                 bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Title              string        `bson:"Title" json:"Title"`
	Description        string        `bson:"Description" json:"Description"`
	ParticipatingTeams []*mgo.DBRef  `bson:"ParticipatingTeams" json:"ParticipatingTeams"`
	ImageUrl           string        `bson:"ImageUrl" json:"ImageUrl"`
	StartDate          time.Time     `bson:"StartDate" json:"StartDate"`
	EndDate            time.Time     `bson:"EndDate" json:"EndDate"`
}

// 团队
type Team struct {
	Id       bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Users    []*mgo.DBRef
	TeamName string
	Motto    string
}

// 权限
type Role struct {
	Id          bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Title       string
	ImageUrl    string
	User        *mgo.DBRef
	Traits      []string
	Description string
	Script      string
	Gender      string
	Age         int
	TimeStamp   time.Time
	Deadline    time.Time
	Comment     []*mgo.DBRef
	Audition    []*mgo.DBRef
}