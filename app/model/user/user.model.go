/*
   fileName: user
   author: diogoxiang
   date: 2019/6/10
*/
package user

import "apiend-core/app/model"

//noinspection GoStructTag
type UserInfo struct {
	model.PublicFields `bson:",inline"` // 公共字段，id和时间
	Uid                int              `bson:"Uid" json:"Uid"`
	Username           string           `bson:"Username" json:"Username"`
	Password           string           `bson:"Password,omitempty" json:"Password,omitempty"`
	NickName           string           `bson:"NickName,omitempty" json:"NickName,omitempty"`
	HeadImg            string           `bson:"HeadImg,omitempty" json:"HeadImg,omitempty"` // 在tag里面加上omitempy，可以在序列化的时候忽略0值或者空值
	Salt               string           `bson:"Salt,omitempty"  json:"Salt,omitempty"`
	// Role     Role          `bson:"role" json:"role"`
	Status UserStatus `bson:"status" json:"status"` // 用户状态 0,1 为正常  -1 为禁止
}

// 用户状态
type UserStatus int

const (
	UserBanned  UserStatus = -1
	UserActived UserStatus = 1
)

func (s UserStatus) Defined() bool {
	switch s {
	case UserBanned, UserActived:
		return true
	}
	return false
}

// 用户列表信息 分页参数
type UserList struct {
	Page      int // 当前页码
	PageCount int // 当前页面的数量
	TotalNum  int // 数量总数
	List      []UserInfo
}
