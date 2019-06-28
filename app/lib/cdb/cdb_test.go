/*
   fileName: cdb
   author: diogoxiang@qq.com
   date: 2019/6/19
*/
package cdb

import (
	"apiend-core/app/model"
	"fmt"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"math/rand"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

const testColl = "c_test"

// 这个主要用来测试用的
type UserInfo struct {
	model.PublicFields `bson:",inline"` // 公共字段，id和时间
	Uid                int              `bson:"Uid" json:"Uid"`
	Username           string           `bson:"Username" json:"Username"`
	Password           string           `bson:"Password,omitempty" json:"Password,omitempty"`
	NickName           string           `bson:"NickName,omitempty" json:"NickName,omitempty"` // 在tag里面加上omitempy，可以在序列化的时候忽略0值或者空值
	HeadImg            string           `bson:"HeadImg,omitempty" json:"HeadImg,omitempty"`
	// Role     Role          `bson:"role" json:"role"`
	// Status   UserStatus    `bson:"status" json:"status"`
}

// 当创建有唯一索引的时候。插入数据失败
func TestInsert(t *testing.T) {
	doc := new(UserInfo)
	doc.Username = "detoxing1"
	doc.SetFieldsValue()

	err := Insert(testColl, doc)

	if err != nil {
		t.Error(err)
	}
	// t.Log(doc)
}

func TestFindOne(t *testing.T) {
	doc := new(UserInfo)
	findDoc := bson.M{"Uid": 0}

	err := FindOne(testColl, doc, findDoc, nil)
	if err != nil {
		t.Error(err)
	}
	t.Log(doc.Username)
}

func TestFindAll(t *testing.T) {
	// 查询某个时间段的参数
	timeString := "2019-06-26 10:00:37.344 +0800 CST"
	ttime, _ := time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", timeString)
	// 查找所有数据
	tds := []UserInfo{}
	findDoc := bson.M{"Uid": bson.M{"$gt": 100}, "createdAt": bson.M{
		"$gt": ttime,
	}}
	err := FindAll(testColl, &tds, findDoc, nil, 0, 100, "-createdAt")
	if err != nil {
		t.Error(err)
	}
	t.Log(len(tds))
	t.Log(tds[0].CreatedAt.Local())
	t.Log(tds[0].NickName)

}

func TestUpdateOne(t *testing.T) {

	selector := bson.M{"Uid": 0}
	update := bson.M{"$set": bson.M{"Username": "diogoxiang"}}

	err := UpdateOne(testColl, selector, update)

	if err != nil {
		t.Error(err)
	}

}

func TestUpdateAll(t *testing.T) {
	selector := bson.M{"Uid": 0}
	update := bson.M{"$set": bson.M{"Username": "diogoxiang"}}

	info, err := UpdateAll(testColl, selector, update)

	if err != nil {
		t.Error(err)
	}

	t.Log(info)
}

func TestDeleteOne(t *testing.T) {

	selector := bson.M{"Uid": 0}
	err := DeleteOne(testColl, selector)

	if err != nil {
		t.Error(err)
	}

}

func TestDeleteAll(t *testing.T) {

	selector := bson.M{"Uid": 0}
	info, err := DeleteAll(testColl, selector)

	if err != nil {
		t.Error(err)
	}
	t.Log(info)
}

func TestDeleteOneReal(t *testing.T) {
	selector := bson.M{"Uid": 0}
	err := DeleteOneReal(testColl, selector)

	if err != nil {
		t.Error(err)
	}

}

func TestDeleteAllReal(t *testing.T) {
	selector := bson.M{"Uid": 0}
	info, err := DeleteAllReal(testColl, selector)

	if err != nil {
		t.Error(err)
	}
	t.Log(info)
}

func TestCount(t *testing.T) {
	selector := bson.M{"Uid": 0}

	info, err := Count(testColl, selector)

	if err != nil {
		t.Error(err)
	}
	t.Log(info)

}

func TestCountAll(t *testing.T) {
	selector := bson.M{"Uid": 0}

	info, err := CountAll(testColl, selector)

	if err != nil {
		t.Error(err)
	}
	t.Log(info)
}

func TestFindAndModify(t *testing.T) {

	selector := bson.M{"Uid": 0}
	update := bson.M{"$set": bson.M{"Username": "FindAndModify"}}
	doc := new(UserInfo)

	info, err := FindAndModify(testColl, doc, selector, update)

	if err != nil {
		t.Error(err)
	}
	t.Log(info)
	t.Log(doc.Username)

}

func TestEnsureIndexKey(t *testing.T) {
	err := EnsureIndexKey(testColl, "Username")
	if err != nil {
		t.Error(err)
	}

	doc := new(UserInfo)
	doc.NickName = "detoxing1"
	doc.SetFieldsValue()

	cerr := Insert(testColl, doc)

	if cerr == nil {
		// t.Error(cerr)
		t.Log(cerr)
	}
	t.Log(doc)

}

// 这个 没整明白 .干啥用的.
func TestEnsureIndex(t *testing.T) {
	// 创建索引
	index := mgo.Index{
		Key:        []string{"Uid", "Username"}, // 索引字段， 默认升序,若需降序在字段前加-
		Unique:     true,                        // 唯一索引 同mysql唯一索引
		DropDups:   true,                        // 索引重复替换旧文档,Unique为true时失效
		Background: true,                        // 后台创建索引
	}

	err := EnsureIndex(testColl, index)
	if err != nil {
		t.Error(err)
	}

	doc := new(UserInfo)
	doc.Username = "detoxing1"
	doc.SetFieldsValue()

	cerr := Insert(testColl, doc)

	if cerr != nil {
		t.Error("index name unique failed")
		return

	}
	t.Log(doc)

}

// 测试并发插入数据
func TestBenchInsert(t *testing.T) {
	var successCount int32
	var wg sync.WaitGroup
	var start = time.Now()

	for i := 0; i < 5000; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()

			td := UserInfo{Username: fmt.Sprintf("zhansan_%d", i), Uid: randAge()}
			td.SetFieldsValue()
			err := Insert(testColl, td)
			if err != nil {
				t.Error(err)
				return
			}

			atomic.AddInt32(&successCount, 1)
		}(i)
	}

	wg.Wait()

	fmt.Printf("\nwrite success count = %d, time = %s\n", successCount, time.Now().Sub(start))
}

// 测试并发读取数据，建立索引前后并发查询速度相差10倍左右
func TestBenchRead(t *testing.T) {
	var successCount int32
	var wg sync.WaitGroup
	var start = time.Now()

	for i := 0; i < 5000; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()

			selector := bson.M{"Username": fmt.Sprintf("zhansan_%d", i)}
			td := &UserInfo{}
			err := FindOne(testColl, td, selector, nil)
			if err != nil {
				t.Error(err)
				return
			}
			if td.Uid < 1 {
				t.Errorf("got %d, expected >0", td.Uid)
				return
			}

			atomic.AddInt32(&successCount, 1)
		}(i)
	}

	wg.Wait()

	fmt.Printf("\nfind success count = %d, time = %s\n", successCount, time.Now().Sub(start))
	// 10万 数据库
	// 没有索引的数据 find success count = 5000, time = 3m9.8989002s
	// 建立索引后的  find success count = 5000, time = 2.8361137s

	// 100万数据
	// find success count = 5000, time = 3.2426464s

}

func randAge() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(99) + 1
}
