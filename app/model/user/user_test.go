/*
   fileName: user
   author: diogoxiang
   date: 2019/6/10
*/
package user

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/globalsign/mgo/bson"
)

func TestUserInfo_Create(t *testing.T) {
	model := new(UserInfo)
	model.NickName = "diogoxiang201901"
	model.SetFieldsValue()
	err := model.Create(nil)
	// etime := time.Now()
	// fmt.Println(etime)
	if err != nil {
		t.Error(err)
	}

}

// 测试一下并发写入5000 个用户
func TestBenchInsert(t *testing.T) {
	var successCount int32
	var wg sync.WaitGroup
	var start = time.Now()

	for i := 0; i < 5000; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()

			td := &UserInfo{Username: fmt.Sprintf("diogoxiang_%d", i)}
			td.SetFieldsValue()
			err := td.Create(nil)
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

func TestInsertUser(t *testing.T) {
	model := new(UserInfo)
	model.NickName = "dio201901"
	err := InsertUser(model)
	if err != nil {
		t.Error(err)
	}
}

func BenchmarkInsertUser(b *testing.B) {
	model := new(UserInfo)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		model.NickName = fmt.Sprintf("%v_%s", i, "diogoxiang")

		_ = InsertUser(model)

	}

}

func TestUpdateUser(t *testing.T) {

	selector := bson.M{
		"Uid": 5010,
	}
	chance := bson.M{
		"$set": bson.M{
			"NickName": "modfiyDio",
			"Username": "diogoxiang",
		},
	}

	err := UpdateUser(selector, chance)

	if err != nil {
		t.Error(err)
	}

}

func TestFindById(t *testing.T) {

	oid := "5d1da8908a5edb09980a7f29"
	puser, err := FindById(oid)

	if err != nil {
		t.Error(err)
	}

	t.Log(puser)

}
