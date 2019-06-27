module apiend-core

go 1.12

replace (
	golang.org/x/net => github.com/golang/net v0.0.0-20180821023952-922f4815f713
	golang.org/x/sync => github.com/golang/sync v0.0.0-20181108010431-42b317875d0f
	golang.org/x/sys => github.com/golang/sys v0.0.0-20180905080454-ebe1bf3edb33
	golang.org/x/text => github.com/golang/text v0.0.0-20170915032832-14c0d48ead0c
)

require (
	github.com/eclipse/paho.mqtt.golang v1.2.0
	github.com/globalsign/mgo v0.0.0-20181015135952-eeefdecb41b8
	github.com/gogf/gf v1.6.17
	github.com/kr/pretty v0.1.0 // indirect
	github.com/lib/pq v1.1.1
	github.com/onsi/ginkgo v1.8.0 // indirect
	github.com/onsi/gomega v1.5.0 // indirect
	github.com/pkg/errors v0.8.1 // indirect
	github.com/sirupsen/logrus v1.4.2
	github.com/xujiajun/nutsdb v0.4.0
	github.com/zhufuyi/logger v0.0.0-20180910035350-cfb136113a27
	go.uber.org/atomic v1.4.0 // indirect
	go.uber.org/multierr v1.1.0 // indirect
	go.uber.org/zap v1.10.0 // indirect
	gopkg.in/check.v1 v1.0.0-20180628173108-788fd7840127 // indirect
	gopkg.in/redis.v5 v5.2.9
)
