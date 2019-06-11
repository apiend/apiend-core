package conn

func init() {
	mgoPools = make(map[string]*MgoPool)
}

var mgoPools map[string]*MgoPool

// MgoSet 在连接池列表里增加mongodb连接池
func MgoSet(key string, p *MgoPool) {
	mgoPools[key] = p
}

// GetMgoPool 根据key从连接池列表里取出对应的mongodb连接池
func GetMgoPool(key string) *MgoPool {
	return mgoPools[key]
}
