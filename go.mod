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
	github.com/gogf/gf v1.7.0
	github.com/kr/pretty v0.1.0 // indirect
	github.com/lib/pq v1.1.1
	github.com/onsi/ginkgo v1.8.0 // indirect
	github.com/onsi/gomega v1.5.0 // indirect
	github.com/sirupsen/logrus v1.4.2
	github.com/tidwall/btree v0.0.0-20170113224114-9876f1454cf0 // indirect
	github.com/tidwall/buntdb v1.1.0
	github.com/tidwall/gjson v1.3.0 // indirect
	github.com/tidwall/grect v0.0.0-20161006141115-ba9a043346eb // indirect
	github.com/tidwall/rtree v0.0.0-20180113144539-6cd427091e0e // indirect
	github.com/tidwall/tinyqueue v0.0.0-20180302190814-1e39f5511563 // indirect
	github.com/xujiajun/gorouter v1.2.0
	github.com/xujiajun/nutsdb v0.4.0
	github.com/xujiajun/utils v0.0.0-20190123093513-8bf096c4f53b
	gopkg.in/check.v1 v1.0.0-20180628173108-788fd7840127 // indirect
	gopkg.in/redis.v5 v5.2.9
)
