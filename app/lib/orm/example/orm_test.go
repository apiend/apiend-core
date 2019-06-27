/*
    fileName: example
    author: diogoxiang@qq.com
    date: 2019/6/17
*/
package example



import (
	"apiend-core/app/lib/orm"
	"encoding/json"
	"fmt"

	"testing"
	"time"
)

type User struct {
	orm.Model `bson:",inline"`
	Name        string
	Password    string
	DeletedAt   time.Time
}

func (this *User) BeforeSave() (err error) {
	fmt.Println(this)
	this.Name = fmt.Sprint("aaa", time.Now().Unix())
	return
}

func (this *User) AfterSave() (err error) {
	fmt.Println(this)
	return
}

func init() {
	orm.Init("127.0.0.1:27017", "test", false, time.Second*30)
	orm.UseSoftDelete(User{})
}

func TestSave(t *testing.T) {

	user := &User{
		Name:     fmt.Sprint("你好", time.Now().Unix()),
		Password: "密码",
		//DeletedAt: time.Now().UTC(),
	}
	//user.SetDoc(user)
	//err := user.Save()
	err := orm.Save(user)
	if err != nil {
		t.Error(err)
	}
}

func TestSelect(t *testing.T) {
	var users []User
	err := orm.FindAll(orm.Query{
		SortFields: []string{"-name"},
		Limit:      3,
		Skip:       1,
	}, &users)
	if err != nil {
		t.Error(err)
	}
	bs, _ := json.Marshal(users)
	fmt.Println(string(bs))
}

func TestCount(t *testing.T) {
	var user User
	user.SetDoc(user)

	query := orm.Query{}

	// 统计所有
	query.Contain = orm.All
	n, err := user.Count(query)
	if err != nil {
		t.Error(err)
	}
	fmt.Println("全部:", n)

	// 默认查询不包含被软删除的
	n, err = user.Count(query)
	if err != nil {
		t.Error(err)
	}
	fmt.Println("默认查询:", n)

	// 统计被软删除的
	query.Contain = orm.DeletedOnly
	n, err = user.Count(query)
	if err != nil {
		t.Error(err)
	}
	fmt.Println("已被删除:", n)
}

func TestSoftDelete(t *testing.T) {
	var users []User
	err := orm.FindAll(orm.Query{
		Contain: orm.DeletedOnly,
	}, &users)
	if err != nil {
		t.Error(err)
	}
	for _, v := range users {
		v.SetDoc(v)
		err := v.Remove(orm.M{"name": v.Name})
		if err != nil {
			t.Error(err)
		}
	}

	fmt.Println(users)
	bs, _ := json.Marshal(users)
	fmt.Println(string(bs))
}

func TestTrueDelete(t *testing.T) {
	var users []User
	err := orm.FindAll(orm.Query{
		Contain: orm.DeletedOnly,
	}, &users)
	if err != nil {
		t.Error(err)
	}
	for _, v := range users {
		v.SetDoc(v)
		err := v.RemoveTrue(orm.M{"name": v.Name})
		if err != nil {
			t.Error(err)
		}
	}
	fmt.Println(users)
	bs, _ := json.Marshal(users)
	fmt.Println(string(bs))
}

func TestUpdate(t *testing.T) {
	var users []User
	err := orm.FindAll(orm.Query{
		SortFields: []string{"-name"},
		Limit:      3,
		Skip:       1,
	}, &users)
	if err != nil {
		t.Error(err)
	}
	for _, v := range users {
		v.SetDoc(v)
		err = v.Update(orm.M{"name": v.Name}, orm.M{"name": v.Name + "	哈哈哈"})
		if err == nil {
			continue
		}
		t.Error(err)
	}
}

func TestUpdateId(t *testing.T) {
	u := User{}
	u.SetDoc(u)
	err := u.UpdateId("5c24f058d66e067ba2e80f0e", orm.M{
		"name": "中国我爱你",
	})
	if err != nil {
		t.Error(err)
	}
}

