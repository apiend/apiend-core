/*
    fileName: micache
    author: diogoxiang@qq.com
    date: 2019/8/9
*/
package micache

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"io/ioutil"
	"os"
	"path"
	"time"
)

var Dir string = "runtime/cache/"
var Error string

type Cache struct {
	Time     int64
	Value    interface{}
	Expires  int64
	FilePath string
}

/**
key get cache
*/
func Get(key string) interface{} {
	filePath := getFilePath(key)
	c := Cache{}
	f, err := os.Open(filePath)

	if err != nil {
		Error = err.Error()
		return nil
	}
	r, _ := ioutil.ReadAll(f)
	f.Close()
	json.Unmarshal(r, &c)
	if c.Time != 0 {
		if (c.Expires < time.Now().Unix()) {
			defer os.Remove(filePath)
			return nil
		}
	}

	return c.Value
}

/**
get struct decoding
*/
func GetDecoding(key string, value interface{}) bool {
	filePath := getFilePath(key)
	c := Cache{}
	f, err := os.Open(filePath)
	if err != nil {
		Error = err.Error()
		return false
	}
	r, _ := ioutil.ReadAll(f)
	f.Close()
	json.Unmarshal(r, &c)
	if c.Time != 0 {
		if (c.Expires < time.Now().Unix()) {
			defer os.Remove(filePath)
			return false
		}
	}
	v,_:=json.Marshal(c.Value)
	json.Unmarshal(v,value)
	return true
}

/**
set cache
*/
func Set(key string, value interface{}, timeNum int64) bool {
	filePath := getFilePath(key)
	dir, _ := path.Split(filePath)
	if mkdirPath(dir) == false {
		return false
	}
	c := Cache{
		Time:     timeNum,
		Value:    value,
		Expires:  time.Now().Unix() + timeNum,
		FilePath: filePath,
	}
	if setFile(c) == false {
		return false
	}
	return true
}


/**
key get file name
*/
func genFileName(name string) string {
	hash := md5.New()
	hash.Write([]byte(name))
	resu := hash.Sum(nil)
	return hex.EncodeToString(resu)
}

/**
key get file path
*/
func getFilePath(key string) string {
	fimeName := genFileName(key)
	filePath := Dir + fimeName[:2] + "/" + fimeName[2:] + ".txt"
	return filePath
}

/**
mkdir
*/
func mkdirPath(dir string) bool {
	err := os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		Error = err.Error()
		return false
	}
	return true
}

/**
set cache file
*/
func setFile(cache Cache) bool {
	c, _ := json.Marshal(cache)
	file, err := os.Create(cache.FilePath)
	defer file.Close()
	if err != nil {
		Error = err.Error()
		return false
	}
	_, err = file.Write(c)
	if err != nil {
		Error = err.Error()
		return false
	}
	return true
}

/**
key is exists
*/
func IsExist(key string) bool {
	filePath := getFilePath(key)
	f, err := os.Open(filePath)
	if err != nil && os.IsNotExist(err) {
		return false
	}
	body, _ := ioutil.ReadAll(f)
	f.Close()
	c := Cache{}
	json.Unmarshal(body, &c)
	if c.Time != 0 {
		if (c.Expires < time.Now().Unix()) {
			defer os.Remove(filePath)
			return false
		}
	}
	return true
}

/**
delete cache file
*/
func Delete(key string) bool {
	filePath := getFilePath(key)
	err := os.Remove(filePath)
	if err != nil {
		Error = err.Error()
		return false
	}

	return true
}


