package conn

func init() {
	redisPools = make(map[string]*RedisPool)
}

var redisPools map[string]*RedisPool

// RedisSet 在连接池列表里增加 redis 连接池
func RedisSet(key string, p *RedisPool) {
	redisPools[key] = p
}

// GetRedisPool 根据key从连接池列表里取出对应的 redis 连接池
func GetRedisPool(key string) *RedisPool {
	return redisPools[key]
}
