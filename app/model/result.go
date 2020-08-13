/*
   fileName: model
   author: diogoxiang@qq.com
   date: 2019/7/1
*/
package model

import (
	"apiend-core/app/lib/util"
	"github.com/globalsign/mgo/bson"
	"time"
)

/**
	DBref 数组应当是断言为[]interface{}，[]map[string]interface就应该还往里面循环
	这里面有坑  待更加熟悉golang后来优化
*/
func FilterResult(m bson.M) map[string]interface{} {
	var mapInfo = map[string]interface{}{}
	for key, value := range m {
		mapInfo[key] = value
		if util.IsBsonM(value) {
			//判断是否为DBRef
			var isDBRef = false
			for refKey, _ := range value.(bson.M) {
				if refKey == "$id" || refKey == "$ref" {
					isDBRef = true
					break
				}
			}
			if isDBRef {
				mapInfo[key] = map[string]interface{}{
					"objectId":  value.(bson.M)["$id"].(bson.ObjectId).Hex(),
					"className": value.(bson.M)["$ref"].(string),
					"__type":    "Pointer",
				}
			} else {
				mapInfo[key] = FilterResult(value.(bson.M))
			}
		}
		if util.IsAnyArray(value) {
			//是否maps数组
			if util.IsMapArray(value) {
				var datas = []bson.M{}
				for _, item := range value.([]bson.M) {
					data := FilterResult(item)
					datas = append(datas, data)
				}
				mapInfo[key] = datas
			} else if util.IsInterfaceArray(value) {
				//判断是否为DBRef 数组
				var isDBRef = false
				for _, item := range value.([]interface{}) {
					if util.IsBsonM(item) {
						for refKey, _ := range item.(bson.M) {
							if refKey == "$id" || refKey == "$ref" {
								isDBRef = true
								break
							}
						}
					}
					//默认为数组里面都是同一类型
					break
				}
				if isDBRef {
					var refs = []map[string]interface{}{}
					for _, item := range value.([]interface{}) {
						ref := map[string]interface{}{
							"objectId":  item.(bson.M)["$id"].(bson.ObjectId).Hex(),
							"className": item.(bson.M)["$ref"].(string),
							"__type":    "Pointer",
						}
						refs = append(refs, ref)
					}
					mapInfo[key] = refs
				}
			}
		}
	}

	delete(mapInfo, "ACL")
	delete(mapInfo, "_r")
	delete(mapInfo, "_w")
	objectId := mapInfo["_id"]
	delete(mapInfo, "_id")
	if objectId != nil {
		mapInfo["objectId"] = objectId
	}
	if mapInfo["createdAt"] != nil && util.IsTime(mapInfo["createdAt"]) {
		mapInfo["createdAt"] = mapInfo["createdAt"].(time.Time).UTC()
	}
	if mapInfo["updatedAt"] != nil && util.IsTime(mapInfo["updatedAt"]) {
		mapInfo["updatedAt"] = mapInfo["updatedAt"].(time.Time).UTC()
	}
	return mapInfo
}

func FilterResults(ms []bson.M) []bson.M {
	datas := []bson.M{}
	for _, m := range ms {
		datas = append(datas, FilterResult(m))
	}
	return datas
}
