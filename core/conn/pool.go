package conn

import (
	"container/list"
	"sync"
)

// Conn 连接接口定义
type Conn interface{}

// Pool 连接池接口定义
type Pool interface {
	Get() Conn
	Put(Conn)
	Exec(...interface{})
}

type pool struct {
	size int
	l    *list.List
	c    chan struct{}
	m    sync.Mutex
}

func (p *pool) init(size int) {
	if size == 0 {
		size = 100
	}
	p.size = size
	p.l = list.New()
	p.c = make(chan struct{}, size)
}
