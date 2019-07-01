/*
   fileName: cdb
   author: diogoxiang@qq.com
   date: 2019/6/27
*/
package cdb

import (
	"errors"
	"fmt"
	"github.com/globalsign/mgo/bson"
	"math/rand"
	"reflect"
	"strings"
)

// 获取类名字
func getCName(class interface{}) string {
	if class == nil {
		return ""
	}
	name := reflect.TypeOf(class).String()
	arr := strings.Split(name, ".")
	return arr[len(arr)-1]
}

// 检查指定的值是否时nil
// 如果时nil，执行panic
// 调用者需要用recover处理错误
func isNil(doc interface{}) {
	if doc == nil {
		CheckErr(errors.New("doc不能为nil"))
	}
}

var (
	zeroVal  reflect.Value
	zeroArgs []reflect.Value
)

// 调用模型的指定方法
func callToDoc(method string, doc interface{}) error {
	docV := reflect.ValueOf(doc)
	if docV.Kind() != reflect.Ptr {
		e := fmt.Sprintf("ormgo: Passed non-pointer: %v (kind=%v), method:%s", doc, docV.Kind(), method)
		return errors.New(e)
	}
	fn := docV.Elem().Addr().MethodByName(method)
	if fn != zeroVal {
		ret := fn.Call(zeroArgs)
		if len(ret) > 0 && !ret[0].IsNil() {
			return ret[0].Interface().(error)
		}
	}
	return nil
}

// Salt 生成一个盐值
func Salt(size int, needUpper bool) string {
	// 按需要生成字符串
	var result string
	var funcLength = 2
	if needUpper {
		funcLength = 3
	}
	for i := 0; i < size; i++ {
		randNumber := rand.Intn(funcLength)
		switch randNumber {
		case 0:
			result += string(Number())
			break
		case 1:
			result += string(Lower())
			break
		case 2:
			result += string(Upper())
			break
		default:
			break
		}
	}
	return result
}

// Lower 随机生成小写字母
func Lower() byte {
	lowerhouse := []int{97, 122}
	result := uint8(lowerhouse[0] + rand.Intn(26))
	return result
}

// Number 随机生成数字
func Number() byte {
	numberhouse := []int{48, 57}
	result := byte(numberhouse[0] + rand.Intn(10))
	return result
}
// Lower 随机生成大写字母
func Upper() byte {
	upperhouse := []int{65, 90}
	result := uint8(upperhouse[0] + rand.Intn(26))
	return result
}


func DealWithSort(sort string) bson.M {
	if sort == "" {return  bson.M{}}
	if sort[0: 1] == "-"{
		var key = sort[1 : len(sort)]
		return  bson.M{
			key: -1,
		}
	}
	return bson.M{
		sort: 1,
	}
}

func IsStringInArray(array []string, str string) bool {
	isIn := false
	for _, value := range array {
		if value == str {
			isIn = true
			break
		}
	}
	return isIn
}
