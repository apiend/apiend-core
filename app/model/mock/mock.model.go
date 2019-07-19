/*
   fileName: mock
   author: diogoxiang@qq.com
   date: 2019/7/15
*/
package mock

import "apiend-core/app/model"

// 接口类型
type MockInfo struct {
	model.PublicFields `bson:",inline"` // 公共字段，id和时间
	Pid                int              `bson:"Pid" json:"Pid"`                           // 当前 项目的ID
	Uid                int              `bson:"Uid" json:"Uid"`                           // 用户ID
	MockPreUrl		   string			`bson:"MockPreUrl" json:"MockPreUrl"`			  // 项目的路径
	MockId             int              `bson:"MockId" json:"MockId"`                     // mock ID
	MockMode           string           `bson:"MockMode" json:"MockMode"`                 // mode
	MockUrl            string           `bson:"MockUrl" json:"MockUrl"`                   // 请求的路径
	MockDescription    string           `bson:"MockDescription" json:"MockDescription"`   // 接口说明
	MockMethod         string           `bson:"MockMethod" json:"MockMethod"`             // 请求的方式
	MockResponse       string           `bson:"MockResponse" json:"MockResponse"`         // 返回的主体
	MockResponseMode   ResMode          `bson:"MockResponseMode" json:"MockResponseMode"` // 返回的格式类型
	Status             MockStatus       `bson:"Status" json:"Status"`                     // MOCK状态

}

// 返回的类型
type ResMode int

const (
	ResJson ResMode = -1
	ResTxt  ResMode = 1
)

func (s ResMode) Defined() bool {
	switch s {
	case ResJson, ResTxt:
		return true
	}
	return false
}

// 项目状态 0 为正常 1为封禁 -1为删除
type MockStatus int

const (
	MockBanned  MockStatus = -1
	MockActived MockStatus = 1
)

func (s MockStatus) Defined() bool {
	switch s {
	case MockBanned, MockActived:
		return true
	}
	return false
}

// mock列表信息 分页参数
type MockList struct {
	Page      int // 当前页码
	PageCount int // 当前页面的数量
	TotalNum  int // 数量总数
	List      []MockInfo
}
