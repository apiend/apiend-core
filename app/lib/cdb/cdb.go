/*
   fileName: cdb
   author: diogoxiang
   date: 2019/6/14
*/
package cdb

import (
	"apiend-core/app/lib/cdb/mid"
	"apiend-core/core/conn"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"github.com/gogf/gf/g"
	"github.com/gogf/gf/g/os/glog"
	"time"
)

var (
	c = g.Config()
	// dataBase   = c.GetString("setting.mgoDbName")
	mongoURL      = c.GetString("mongo.mongoUrl")
	mongoPoolSize = c.GetInt("mongo.mgoPoolSize")
	dbName        = c.GetString("mongo.mgoDbName")
	SlowRes       = c.GetInt("mongo.SlowRes")
	// ErrNotFound   = mgo.ErrNotFound
)

// 项目初始化的时候  链接 mongo
func init() {
	setDB()
	time.Local = time.UTC
}

// 设置 db 初始化
func setDB() {
	// Mongodb init
	mgoOption := conn.MgoPoolOption{
		Host:   mongoURL,
		Size:   mongoPoolSize,
		DbName: dbName,
		// SlowRes:time.Duration(SlowRes),
	}
	mgoPool, err := conn.NewMgoPool(mgoOption)
	if err != nil {
		glog.Debugf("connect mongodb: " + dbName + "  fail")
		panic("connect mongodb: " + dbName + "  fail")
		// os.Exit(1)
	}
	conn.MgoSet(mgoOption.DbName, mgoPool)
	glog.Println("cdb connect mongodb done")
}

// type Dbm struct {
// 	nSession *mgo.Session
// 	printLog   bool
// }

/**
单个插入数据
*/
func Insert(collection string, doc interface{}) (err error) {
	conn.GetMgoPool(dbName).Exec(collection, func(c *mgo.Collection) {
		err = c.Insert(doc)
	})
	if err != nil {
		return err
	}
	return nil
}

/**
查找一个
*/
func FindOne(collection string, result interface{}, selector bson.M, fields bson.M) (err error) {
	conn.GetMgoPool(dbName).Exec(collection, func(c *mgo.Collection) {
		// err = c.Insert(doc)
		err = c.Find(excludeDeleted(selector)).Select(fields).One(result)
	})

	if err != nil {
		return err
	}
	return nil
}

/**
查找全部数据 通用模式
查找符合条件的多条记录，这里参数skip表示页码，limit表示每页多少行数据
*/
func FindAll(collection string, result interface{}, selector bson.M, fields bson.M, skip int, limit int, sort ...string) (err error) {

	conn.GetMgoPool(dbName).Exec(collection, func(c *mgo.Collection) {
		// err = c.Insert(doc)
		err = c.Find(excludeDeleted(selector)).Select(fields).Sort(sort...).Skip(skip * limit).Limit(limit).All(result)
	})

	if err != nil {
		return err
	}
	return nil
}

/**
Update 更新一条记录
*/
func UpdateOne(collection string, selector bson.M, update bson.M) (err error) {
	err = checkUpdateContent(update)
	if err != nil {
		return err
	}

	conn.GetMgoPool(dbName).Exec(collection, func(c *mgo.Collection) {

		err = c.Update(excludeDeleted(selector), updatedTime(update))

	})

	if err != nil {
		return err
	}
	return nil

}

/**
更新数据
返回 int 影响多少doc
error
*/
func UpdateAll(collection string, selector bson.M, update bson.M) (int, error) {
	var cerr error

	info := new(mgo.ChangeInfo)

	err := checkUpdateContent(update)
	if err != nil {
		return 0, err
	}

	conn.GetMgoPool(dbName).Exec(collection, func(c *mgo.Collection) {

		info, cerr = c.UpdateAll(excludeDeleted(selector), updatedTime(update))

	})

	if cerr != nil {
		return 0, cerr
	}
	return info.Updated, cerr

}

/**
DeleteOne 标记性删除一条记录
*/
func DeleteOne(collection string, selector bson.M) error {
	var cerr error

	conn.GetMgoPool(dbName).Exec(collection, func(c *mgo.Collection) {

		cerr = c.Update(excludeDeleted(selector), deletedTime(bson.M{}))

	})

	if cerr != nil {
		return cerr
	}

	return nil
}

/**
DeleteOne 标记性删除所有数据
*/
func DeleteAll(collection string, selector bson.M) (int, error) {

	var cerr error

	info := new(mgo.ChangeInfo)

	conn.GetMgoPool(dbName).Exec(collection, func(c *mgo.Collection) {

		info, cerr = c.UpdateAll(excludeDeleted(selector), deletedTime(bson.M{}))

	})

	if cerr != nil {
		return 0, cerr
	}
	return info.Updated, cerr

}

/**
DeleteOneReal 真实删除
*/
func DeleteOneReal(collection string, selector bson.M) error {
	var cerr error

	conn.GetMgoPool(dbName).Exec(collection, func(c *mgo.Collection) {

		cerr = c.Remove(selector)

	})

	if cerr != nil {
		return cerr
	}

	return nil
}

/**
DeleteAllReal 删除所有匹配的记录，包括标记性删除的记录
*/
func DeleteAllReal(collection string, selector bson.M) (int, error) {
	var cerr error
	info := new(mgo.ChangeInfo)

	conn.GetMgoPool(dbName).Exec(collection, func(c *mgo.Collection) {

		info, cerr = c.RemoveAll(selector)

	})

	if cerr != nil {
		return 0, cerr
	}
	return info.Removed, cerr

}

/**
Count 统计匹配的数量，不包括标记性删除的记录
*/
func Count(collection string, selector bson.M) (int, error) {
	var cerr error
	var count int
	conn.GetMgoPool(dbName).Exec(collection, func(c *mgo.Collection) {

		count, cerr = c.Find(excludeDeleted(selector)).Count()

	})

	if cerr != nil {
		return 0, cerr
	}

	return count, nil

}

/**
CountAll 统计匹配的数量，包括标记性删除的记录
*/
func CountAll(collection string, selector bson.M) (int, error) {
	var cerr error
	var count int

	conn.GetMgoPool(dbName).Exec(collection, func(c *mgo.Collection) {

		count, cerr = c.Find(selector).Count()

	})

	if cerr != nil {
		return 0, cerr
	}

	return count, nil

}

/**
FindAndModify 更新并返回最新记录 找到一个 ,并更新updateAt 时间
*/
func FindAndModify(collection string, result interface{}, selector bson.M, update bson.M) (int, error) {
	var cerr error
	err := checkUpdateContent(update)
	if err != nil {
		return 0, err
	}
	info := new(mgo.ChangeInfo)

	change := mgo.Change{ReturnNew: true, Update: updatedTime(update)}

	conn.GetMgoPool(dbName).Exec(collection, func(c *mgo.Collection) {

		info, cerr = c.Find(excludeDeleted(selector)).Apply(change, result)

	})

	if cerr != nil {
		return 0, cerr
	}

	return info.Matched, nil

}

/**
IndexSet 设置索引key
*/
func EnsureIndexKey(collection string, indexKeys ...string) error {
	var cerr error
	conn.GetMgoPool(dbName).Exec(collection, func(c *mgo.Collection) {

		cerr = c.EnsureIndexKey(indexKeys...)

	})

	if cerr != nil {
		return cerr
	}

	return nil

}

/**
IndexSet 设置索引
*/
func EnsureIndex(collection string, index mgo.Index) error {

	var cerr error
	conn.GetMgoPool(dbName).Exec(collection, func(c *mgo.Collection) {

		cerr = c.EnsureIndex(index)

	})

	if cerr != nil {
		return cerr
	}

	return nil

}

// 公共方法, 获取 name 的自增ID
func GetAutoId(collection string, name string) (id int, err error) {
	conn.GetMgoPool(dbName).Exec(collection, func(c *mgo.Collection) {
		id, err = mid.AutoInc(c, name)
	})
	return
}

// 公用类获取
func FetchRef(ref mgo.DBRef) bson.M {
	var obj = bson.M{}
	conn.GetMgoPool(dbName).ExecDB("", func(db *mgo.Database) {
			db.FindRef(&ref).One(&obj)
	})
	return obj
}
