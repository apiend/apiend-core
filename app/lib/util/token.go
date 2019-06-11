/*
   fileName: util
   author: diogoxiang
   date: 2019/4/25
*/
package util

import (
	"apiend-back/app/library/buntstore"
	"github.com/gogf/gf/g"
	"github.com/gogf/gf/g/util/gconv"
	"github.com/tidwall/buntdb"
	"time"
)

var (
	c       = g.Config()
	tokenDB = c.GetString("setting.localCacheDB")
	// 缓存有效时间
	tokenCacheTime = c.GetInt("setting.cacheTime")
)

// remove old test DB if it exists and create a new one
func getTestDatabase() *buntdb.DB {
	// 不删除 数据
	//err := os.Remove(tokenDB)
	//if err != nil {
	//	panic(err)
	//}
	db, err := buntdb.Open(tokenDB)
	if err != nil {
		panic(err)
	}
	return db
}

// 创建Token 并保存 相应的信息 ,并返回 生成的token key
func CreateToken(value []byte) (key string, err error) {
	db := getTestDatabase()
	defer db.Close()

	bs := buntstore.NewStore(db)

	key = GenerateUUID()

	err = bs.Save(key, value, time.Now().Add(time.Duration(tokenCacheTime)*time.Minute))

	return key, err

}

// 根据 用户名生成相应的token, 用户名是 key
func CreateTokenByName(value []byte) (key string, err error) {
	db := getTestDatabase()
	defer db.Close()

	bs := buntstore.NewStore(db)

	key = GenerateUUID()

	eKey := gconv.Bytes(key)
	eValue := gconv.String(value)

	err = bs.Save(eValue, eKey, time.Now().Add(time.Duration(tokenCacheTime)*time.Minute))

	return key, err

}

// 验证 key  值 是否正解 并返回相应的数据
func CheckTokenKey(key string) (b []byte, exists bool, err error) {
	db := getTestDatabase()
	defer db.Close()

	bs := buntstore.NewStore(db)

	b, exists, err = bs.Find(key)
	return
}

// 只验证是否正常
func ValidToken(key string) (fond bool) {

	db := getTestDatabase()
	defer db.Close()

	bs := buntstore.NewStore(db)

	_, fond, _ = bs.Find(key)
	return fond
}
