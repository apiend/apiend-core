/*
   fileName: epost
   author: diogoxiang@qq.com
   date: 2019/7/30
*/
package epost

import "apiend-core/app/model"

// 文档分类
type PostType struct {
	model.PublicFields `bson:",inline"` // 公共字段，id和时间
	TypeId             int              `bson:"TypeId" json:"TypeId"`   // 类id
	TypeTxt            string           `bson:"TypeTxt" json:"TypeTxt"` // 分类名称
}

// 文档内容细节
type PostDetail struct {
	model.PublicFields `bson:",inline"` // 公共字段，id和时间
	PostTypeInfo       *PostType `bson:"PostTypeInfo" json:"PostTypeInfo"`
	PostTypeId         int      `bson:"TypeId" json:"TypeId"`
	PostId             int      `bson:"PostId" json:"PostId"`
	PostTitle          string   `bson:"PostTitle" json:"PostTitle"`
	PostPic            string   `bson:"PostPic" json:"PostPic"`
	PostInfo           string   `bson:"PostInfo" json:"PostInfo"`
	PostContent        string   `bson:"PostContent" json:"PostContent"`
	PostAuthor         string   `bson:"PostAuthor" json:"PostAuthor"`
	PostTag            []string `bson:"PostTag" json:"PostTag"`
}
