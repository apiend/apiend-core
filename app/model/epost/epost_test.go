/*
   fileName: epost
   author: diogoxiang@qq.com
   date: 2019/7/30
*/
package epost

import (
	"github.com/globalsign/mgo/bson"
	"github.com/gogf/gf/g/os/gtime"
	"testing"
)

func TestInsertType(t *testing.T) {
	doc := new(PostType)

	doc.TypeTxt = "Vue项目,time"

	err := InsertType(doc)

	if err != nil {
		t.Error(err)
	}

}

func TestInsertEpost(t *testing.T) {
	doc := new(PostDetail)
	docType := new(PostType)
	docType.TypeId = 1
	docType.TypeTxt = "Vue组件,time"

	doc.PostTypeInfo = docType

	doc.PostTypeId = 1

	doc.PostTitle = "Vue 002"
	doc.PostPic = "sss3"
	doc.PostContent = "content2"
	doc.PostInfo = "info2"
	doc.PostAuthor = "username2"
	doc.PostTag = []string{"vue", "e22s"}

	err := InsertEpost(doc)

	if err != nil {
		t.Error(err)
	}

}

func TestFindBySearchOneType(t *testing.T) {

	selector := bson.M{
		"TypeId":1,
	}
	info,err := FindBySearchOneType(selector)
	t.Log(err)
	t.Log(info.CreatedAt.UTC())
}

func TestFindBySearchAllType(t *testing.T) {
	selector := bson.M{
		"TypeId":bson.M{
			"$gt":1,
		},
	}
	// 这个用来控制 页面跳转
	Skip := 0
	limter :=10

	pinfo , err :=FindBySearchAllType(selector,nil,Skip,limter,"-CreatedAt")

	t.Log(err)
	t.Log(pinfo)
	t.Log(len(pinfo))

}



func TestFindBySearchOneEpost(t *testing.T) {
	// cstLocal, _ := time.LoadLocation("Asia/Shanghai")
	selector := bson.M{
		"PostId":2,
	}
	info,err := FindBySearchOneEpost(selector)
	t.Log(err)
	t.Log(info.CreatedAt.UTC())
}

func TestFindBySearchAllEpost(t *testing.T) {

	// 获取当前时间
	// Etime := gtime.Now().Time

	// 2019-07-30 09:18:51.215 +0000 UTC 指定时时间
	Etime := gtime.NewFromStr("2019-07-31 09:18:51").Time
	t.Log(Etime)
	selector := bson.M{
		"CreatedAt":bson.M{
			"$lt":Etime,
		},
	}
	// 这个用来控制 页面跳转
	Skip := 0
	Limter :=10
	pinfo,err := FindBySearchAllEpost(selector,nil,Skip,Limter)
	t.Log(err)

	t.Log(len(pinfo))

}

func TestUpdateOneEpost(t *testing.T) {

	selector := bson.M{
		"PostId":6,
	}
	chance := bson.M{
		"$set":bson.M{
			"PostTitle":"modfiyDio",
			"PostInfo":"diogoxiang",
		},
	}

	err := UpdateOneEpost(selector,chance)


	t.Log(err)

}

func TestUpdateAllEpost(t *testing.T) {

	selector := bson.M{
		"TypeId":1,
	}

	chance := bson.M{
		"$set":bson.M{
			"PostTitle":"modfiyDio",
			"PostInfo":"diogoxiang",
		},
	}


	fond,err := UpdateAllEpost(selector,chance)


	t.Log(fond)
	t.Log(err)

}