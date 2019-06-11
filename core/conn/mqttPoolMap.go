package conn

func init() {
	mqttPools = make(map[string]*MqttPool)
}

var mqttPools map[string]*MqttPool

// MqttSet 在连接池列表里增加 mqtt 连接池
func MqttSet(key string, p *MqttPool) {
	mqttPools[key] = p
}

// GetMqttPool 根据key从连接池列表里取出对应的 mqtt 连接池
func GetMqttPool(key string) *MqttPool {
	return mqttPools[key]
}
