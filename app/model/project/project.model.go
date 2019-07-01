/*
    fileName: project
    author: diogoxiang@qq.com
    date: 2019/6/30
*/
package project

import (
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"time"
)

//Cast struct
type Cast struct {
	Id bson.ObjectId `json:"id" bson:"_id,omitempty"`
	User *mgo.DBRef
	Role string
}


// 项目model
type Project struct {
	Id bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Title string
	URL string
	ShortDescription string
	Description string
	Cast []*mgo.DBRef
	PostedDate time.Time
	User *mgo.DBRef
	Contest *mgo.DBRef
}

