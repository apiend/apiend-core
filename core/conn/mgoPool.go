package conn

import (
	"time"

	"github.com/globalsign/mgo"
	log "github.com/sirupsen/logrus"
)

// MgoPool mongodb 连接池
type MgoPool struct {
	p   pool
	opt MgoPoolOption
}

// MgoPoolOption mongodb 连接池配置项
type MgoPoolOption struct {
	Size    int
	Host    string
	DbName  string
	SlowRes time.Duration
}

// NewMgoPool 创建一个 mongodb 连接池
func NewMgoPool(opt MgoPoolOption) (*MgoPool, error) {
	p := MgoPool{}
	err := p.init(opt)
	return &p, err
}

func (p *MgoPool) init(opt MgoPoolOption) error {
	p.opt = opt
	p.p.init(opt.Size)

	session, err := mgo.Dial(opt.Host)
	if err != nil {
		return err
	}
	for i := 0; i < p.p.size; i++ {
		p.p.c <- struct{}{}
	}
	p.p.l.PushBack(session)
	return nil
}

// Get 获取一个mongo连接
func (p *MgoPool) Get() *mgo.Session {
	_ = <-p.p.c
	p.p.m.Lock()
	defer p.p.m.Unlock()
	return p.p.l.Front().Value.(*mgo.Session).Clone()
}

// Put 释放一个mongo连接
func (p *MgoPool) Put(c Conn) {
	p.p.m.Lock()
	defer p.p.m.Unlock()
	c.(*mgo.Session).Close()
	p.p.c <- struct{}{}
}

// Exec 使用连接池
func (p *MgoPool) Exec(collection string, callback func(*mgo.Collection)) {
	start := time.Now()
	_session := p.Get()
	defer func() {
		p.Put(_session)
		if err := recover(); err != nil {
			log.Errorln("mongodb exec err, ", err)
			panic(err)
		}
		t := time.Since(start)
		if t >= p.opt.SlowRes && p.opt.SlowRes != 0 {
			log.Warnln("mongodb exec ", collection, t)
		}
	}()
	c := _session.DB(p.opt.DbName).C(collection)
	callback(c)
}
// 返回 *mgo.Database
func (p *MgoPool) ExecDB(collection string, callback func(database *mgo.Database))  {
	start := time.Now()
	_session := p.Get()
	defer func() {
		p.Put(_session)
		if err := recover(); err != nil {
			log.Errorln("mongodb exec err, ", err)
			panic(err)
		}
		t := time.Since(start)
		if t >= p.opt.SlowRes && p.opt.SlowRes != 0 {
			log.Warnln("mongodb exec ", collection, t)

		}
	}()
	callback(_session.DB(p.opt.DbName))
	// return mdb
}