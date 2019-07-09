/*
    fileName: system
    author: diogoxiang@qq.com
    date: 2019/7/9
	系统配置项
*/
package system

import "apiend-core/app/model"

// 系统的基本配置项
type SystemInfo struct {
	model.PublicFields `bson:",inline"` // 公共字段，id和时间
	WebName            string           `bson:"WebName" json:"WebName"`     // 网站名称
	WebDes             string           `bson:"WebDes" json:"WebDes"`       // 网站简介
	WebNotice          string           `bson:"WebNotice" json:"WebNotice"` // 网站公告
}
