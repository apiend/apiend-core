/*
    fileName: user
    author: diogoxiang
    date: 2019/6/10
*/
package user

import "apiend-core/app/model"

//noinspection GoStructTag
type UserInfo struct {
	model.PublicFields 	   `bson:",inline"` // 公共字段，id和时间
	Uid      int           `bson:"Uid" json:"Uid"`
	Username string        `bson:"Username" json:"Username"`
	Password string        `bson:"Password" json:"Password"`
	NickName string        `bson:"NickName" json:"NickName"`
	HeadImg  string        `bson:"HeadImg" json:"HeadImg"`
	// Role     Role          `bson:"role" json:"role"`
	// Status   UserStatus    `bson:"status" json:"status"`
}