/*
    fileName: gb
    author: diogoxiang@qq.com
    date: 2019/7/26
	公用内存级存储,常用参数..可用于时时获取 config.xml 最新的参数
*/
package gb

import "errors"

var ErrNotFound = errors.New("Key not found")


type Store interface {
	Get(key string) ([]byte, error)
	Set(key string, data []byte) error
	ForEach(func(key string, data []byte)) error
	Close() error
}

func InitStore(store Store) Store{



	return store


}