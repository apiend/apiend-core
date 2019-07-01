/*
    fileName: cdb
    author: diogoxiang@qq.com
    date: 2019/6/30

	include功能 包含 include('creator')
	include('creator.detail')  include('feed.creator.detail')以及更多层次
	也涵盖了include('articles') 这样的数组形式的引用
	//后续需要花时间异步请求
*/
package cdb

import (
	"apiend-core/app/lib/util"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"strings"
)

func IncludeObjects(m []bson.M, includes []string, db *mgo.Database) []bson.M {
	var objects = []bson.M{}
	objectChans := make([]chan bson.M, len(m))
	for i, obj := range m {
		objectChans[i] = make(chan bson.M)
		go func(k int, obj bson.M) {
			data :=  IncludeObject(obj, includes, db)
			objectChans[k] <- data
		}(i, obj)
	}
	for _, ch := range objectChans {
		data := <-ch
		objects = append(objects, data)
	}
	return objects
}


/**
	获取某个对象的各种引用
	@params m 对象本身
	@params includes 需要引用的字段
	@params db 数据库连接对象
	@return 返回一个已经fetch过值的对象
 */
func IncludeObject(m bson.M, includes []string, db *mgo.Database) bson.M{
	//first filter includes  due to include may be "creator.detail"
	includeArray, includeMap := filterIncludes(includes)
	for _, include := range includeArray{

		if m[include] == nil {
			continue
		}
		var refs = []interface{}{}
		var isArrayParameters = util.IsAnyArray(m[include])

		//include 可能会是一个数组，那就当做全是数组
		if isArrayParameters {
			refs = m[include].([]interface{})
		} else {
			refs = append(refs, m[include])
		}
		var defaultResults = scanAndFetchRef(refs, db, include, includeMap)
		//如果是数组，那么直接复制 非数组则取第一个值
		if isArrayParameters {
			m[include] = defaultResults
		}else {
			m[include] = defaultResults[0]
		}
	}
	return m
}

func scanAndFetchRef(refs []interface{}, db *mgo.Database, include string, includeMap map[string][]string) []bson.M{
	//遍历ref获取值
	var defaultResults = []bson.M{}
	for _, _r := range refs {
		var obj = bson.M{}
		var ref mgo.DBRef
		data, _:= bson.Marshal(_r)
		bson.Unmarshal(data, &ref)
		db.FindRef(&ref).One(&obj)
		obj["__type"] = "Pointer"
		obj["className"] = ref.Collection
		for key, value := range includeMap{
			if key == include {
				obj = IncludeObject(obj, value, db)
			}
		}
		defaultResults = append(defaultResults, obj)
	}
	return defaultResults
}

/**
	example: ["creator", "creator.detail", "feed.creator.detail"]
 */
func filterIncludes(includes []string) ([]string, map[string][]string) {
	var filterArray = []string{}
	var filterMap = map[string][]string{}
	for _, include:= range includes {
		array := strings.SplitN(include, ".", 2)
		filterArray = append(filterArray, array[0])
		if len(array) > 1 {
			if _, ok := filterMap[array[0]]; ok {
				valueArray := filterMap[array[0]]
				valueArray = append(valueArray, array[1])
				filterMap[array[0]] = removeDuplicatesAndEmpty(valueArray)
			}else {
				filterMap[array[0]] = []string{array[1]}
			}
		}
	}
	return removeDuplicatesAndEmpty(filterArray), filterMap
}

/**
	数组去重 去空
 */
func removeDuplicatesAndEmpty(a []string) []string{
	a_len := len(a)
	ret := []string{}
	for i:=0; i < a_len; i++{
		if (i > 0 && a[i-1] == a[i]) || len(a[i])==0{
			continue;
		}
		ret = append(ret, a[i])
	}
	return ret
}

