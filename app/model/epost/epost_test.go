/*
   fileName: epost
   author: diogoxiang@qq.com
   date: 2019/7/30
*/
package epost

import "testing"

func TestInsertType(t *testing.T) {
	doc := new(PostType)

	doc.TypeTxt = "Vue项目"

	err := InsertType(doc)

	if err != nil {
		t.Error(err)
	}

}

func TestInsertEpost(t *testing.T) {
	doc := new(PostDetail)
	docType := new(PostType)
	docType.TypeId = 1
	docType.TypeTxt = "Vue组件"

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
