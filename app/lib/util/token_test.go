/*
    fileName: util
    author: diogoxiang@qq.com
    date: 2019/7/2
*/
package util

import (
	"encoding/json"
	"github.com/gogf/gf/g/encoding/gjson"
	"testing"
	"unsafe"
)


func toString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}


func TestCreateToken(t *testing.T) {

	key, err := CreateToken([]byte(`{"name":"diogoxiang"}`))

	t.Log(key)
	t.Log(err)
	// one 81aeb256409e42a180cca3a0eb84d06c
}


func TestValidTokenKey(t *testing.T) {

	v, exists, err := ValidTokenKey("4c7c7204404353b7802ba3a0f4bc66f2")
	temp,err := gjson.Decode(v)

	// temp := gconv.String(v)
	// jsonStr, err := json.Marshal(temp)

	// if err != nil {
	// 	t.Fatal(err)
	// }
	//t.Log(v)
	t.Logf("%v", temp)

	// t.Logf("jsonStr %s", jsonStr)
	t.Log(exists)
	t.Log(err)

}


func TestCreateTokenByName(t *testing.T) {
	ukey, err := createTokenByName("diogoxiang")
	t.Log(ukey)
	t.Log(err)
	err = createTokenByuKey(ukey,"diogoxiang")
	t.Log(err)
}

func TestNewToken(t *testing.T) {
	ukey, err := NewToken("diogoxiang1")
	t.Log(ukey)
	t.Log(err)
}



func TestMap2Json(t *testing.T) {
	mapInstance := make(map[string]interface{})
	mapInstance["Name"] = "liang637210"
	mapInstance["Age"] = 28
	mapInstance["Address"] = "北京昌平区"

	t.Log(mapInstance)
	jsonStr, err := json.Marshal(mapInstance)

	if err != nil {
		t.Fatal(err)
	}

	t.Logf("Map2Json 得到 json 字符串内容:%s", jsonStr)
}

// part 3 项目
func TestNewCreateToken(t *testing.T) {

	uname := "diogoxiang9"
	uinfo := ` {
            "Uid": 10026,
            "Username": "diogoxiang9",
            "id": "5d1da8fd8a5edb09980a7f2a"
        }`

	t.Log(uname)
	t.Log(uinfo)
	temp,err := gjson.Encode(uinfo)
	utoken, err := NewCreateToken(uname,temp)

	t.Log(utoken)
	t.Log(err)
	t.Log(`---------`)

	v, _, err := ValidTokenKey(utoken)
	tempe,errs := gjson.DecodeToJson(v)

	t.Log(tempe.Get("Uid"))
	t.Log(errs)
}
