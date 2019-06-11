package conn

import (
	"time"

	log "github.com/sirupsen/logrus"
	"gopkg.in/redis.v5"
)

// RedisPool redis 连接池
type RedisPool struct {
	p   pool
	opt RedisPoolOption
}

// RedisPoolOption redis 连接池配置项
type RedisPoolOption struct {
	Size     int
	Host     string
	Password string
	DB       int
	SlowRes  time.Duration
}

// NewRedisPool 创建一个 redis 连接池
func NewRedisPool(opt RedisPoolOption) (*RedisPool, error) {
	var p RedisPool
	if opt.SlowRes == 0 {
		opt.SlowRes = time.Millisecond * 100
	}
	err := p.init(opt)
	return &p, err
}

func (p *RedisPool) init(opt RedisPoolOption) error {
	p.opt = opt
	p.p.init(opt.Size)

	redisOpt := redis.Options{
		Addr:     opt.Host,
		Password: opt.Password,
		PoolSize: opt.Size,
		DB:       opt.DB,
	}
	c := redis.NewClient(&redisOpt)
	err := c.Ping().Err()
	if err != nil {
		return err
	}

	for i := 0; i < p.p.size; i++ {
		p.p.c <- struct{}{}
	}
	p.p.l.PushBack(c)
	return nil
}

// Close 关闭连接
func (p *RedisPool) Close() {
	p.p.m.Lock()
	defer p.p.m.Unlock()

	client := p.p.l.Front()
	client.Value.(*redis.Client).Close()
}

// Get 获取一个redis连接
func (p *RedisPool) Get() *redis.Client {
	_ = <-p.p.c
	p.p.m.Lock()
	defer p.p.m.Unlock()
	return p.p.l.Front().Value.(*redis.Client)
}

// Put 释放一个redis连接
func (p *RedisPool) Put(c *redis.Client) {
	p.p.m.Lock()
	defer p.p.m.Unlock()
	p.p.c <- struct{}{}
}

// Exec 使用连接池
func (p *RedisPool) Exec(callback func(*redis.Client)) {
	start := time.Now()
	client := p.Get()
	defer func() {
		p.Put(client)
		if err := recover(); err != nil {
			log.Errorln("redis exec err, ", err)
			panic(err)
		}
		t := time.Since(start)
		if t >= p.opt.SlowRes && p.opt.SlowRes != 0 {
			log.Warnln("redis exec db:", p.opt.DB, t)
		}
	}()
	callback(client)
}
