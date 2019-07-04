/*
    fileName: util
    author: diogoxiang@qq.com
    date: 2019/7/2
*/
package util

import (
	"apiend-core/app/lib/buntstore"
	"github.com/gogf/gf/g"
	"github.com/gogf/gf/g/os/glog"
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



// 验证 key  值 是否正解 并返回相应的数据
func ValidTokenKey(key string) (b []byte, exists bool, err error) {
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

// ----------------另一套生成机制

// 检测用户名是否已经生成过的如果已经生成过.则删除已经生成的,并重新生成
func NewToken(uname string) (key string, err error) {
	// 先找是否已经生成过token
	v, fond, _ := ValidTokenKey(uname)
 	if fond {
 		vd := gconv.String(v)
		err = delToken(vd)
		// 检测是否删除成功
		if err != nil {
			return "", err
		}

	}
	key, err = createTokenByName(uname)

	err = createTokenByuKey(key,uname)

	if err != nil {
		glog.Debug(err)
		return "0", err
	}

	return key,nil

}

// 根据 用户名生成相应的token, 用户名是 uname
func createTokenByName(uname string) (key string, err error) {
	db := getTestDatabase()
	defer db.Close()

	bs := buntstore.NewStore(db)

	tKey := GenerateUUID()

	// eKey := gconv.Bytes(ukey)
	eValue := gconv.Bytes(tKey)
	sValue := gconv.String(tKey)

	err = bs.Save(uname, eValue, time.Now().Add(time.Duration(tokenCacheTime)*time.Minute))

	return sValue, err

}

// 根据生成的token(ukey) 存储 用户名信息
func createTokenByuKey(ukey string, uname string) (err error) {
	db := getTestDatabase()
	defer db.Close()
	bs := buntstore.NewStore(db)
	// 存储用户信息(用户名)
	sValue := gconv.Bytes(uname)

	err = bs.Save(ukey, sValue, time.Now().Add(time.Duration(tokenCacheTime)*time.Minute))

	return err

}


// 删除token
func delToken(utoken string)  error {
	db := getTestDatabase()
	defer db.Close()

	bs := buntstore.NewStore(db)

	err := bs.Delete(utoken)

	return err
}



