package conn

import (
	"time"

	"github.com/eclipse/paho.mqtt.golang"
)

// MqttPool mqtt 连接池
type MqttPool struct {
	p   pool
	opt MqttPoolOption
}

// MqttPoolOption mqtt 连接池配置项
type MqttPoolOption struct {
	Size   int
	Broker string
}

// NewMqttPool 创建一个 mqtt 连接池
func NewMqttPool(opt MqttPoolOption) *MqttPool {
	var p MqttPool
	p.init(opt)
	return &p
}

func (p *MqttPool) init(opt MqttPoolOption) error {
	p.opt = opt
	p.p.init(opt.Size)

	mqttOpt := mqtt.NewClientOptions()
	mqttOpt.AddBroker(opt.Broker)
	mqttOpt.SetKeepAlive(5 * time.Second)

	for i := 0; i < p.p.size; i++ {
		client := mqtt.NewClient(mqttOpt)
		for {
			token := client.Connect()
			if token.Wait() && token.Error() != nil {
				time.Sleep(3 * time.Second)
				continue
			}
			break
		}
		p.p.c <- struct{}{}
		p.p.l.PushBack(client)
	}
	return nil
}

// Get 获取连接
func (p *MqttPool) Get() Conn {
	_ = <-p.p.c
	p.p.m.Lock()
	defer p.p.m.Unlock()
	element := p.p.l.Front()
	p.p.l.Remove(element)

	return element.Value.(mqtt.Client)
}

// Put 释放连接
func (p *MqttPool) Put(c Conn) {
	p.p.m.Lock()
	defer p.p.m.Unlock()
	p.p.l.PushBack(c)
	p.p.c <- struct{}{}
}
