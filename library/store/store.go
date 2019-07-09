/*
    fileName: store
    author: diogoxiang@qq.com
    date: 2019/7/8
*/
package store

import (
	"github.com/pkg/errors"
)

var ErrNotFound = errors.New("Key not found")

type Store interface {
	Get(key string) ([]byte, error)
	Set(key string, data []byte) error
	ForEach(func(key string, data []byte)) error
	Close() error
}
