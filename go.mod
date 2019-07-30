module apiend-core

go 1.12

replace (
	golang.org/x/crypto => github.com/golang/crypto v0.0.0-20190701094942-4def268fd1a4
	golang.org/x/net => github.com/golang/net v0.0.0-20180821023952-922f4815f713
	golang.org/x/sync => github.com/golang/sync v0.0.0-20181108010431-42b317875d0f
	golang.org/x/sys => github.com/golang/sys v0.0.0-20180905080454-ebe1bf3edb33
	golang.org/x/text => github.com/golang/text v0.0.0-20170915032832-14c0d48ead0c
	golang.org/x/tools => github.com/golang/tools v0.0.0-20190706070813-72ffa07ba3db
	google.golang.org/appengine => github.com/golang/appengine v1.6.1
)

require (
	github.com/dgraph-io/badger v1.6.0
	github.com/eclipse/paho.mqtt.golang v1.2.0
	github.com/fortytw2/leaktest v1.3.0 // indirect
	github.com/garyburd/redigo v1.6.0 // indirect
	github.com/globalsign/mgo v0.0.0-20181015135952-eeefdecb41b8
	github.com/glycerine/goconvey v0.0.0-20190410193231-58a59202ab31
	github.com/go-stomp/stomp v2.0.3+incompatible // indirect
	github.com/gogf/gf v1.7.0
	github.com/gopherjs/gopherjs v0.0.0-20190430165422-3e4dfb77656c // indirect
	github.com/jolestar/go-commons-pool v2.0.0+incompatible // indirect
	github.com/jtolds/gls v4.20.0+incompatible // indirect
	github.com/kr/beanstalk v0.0.0-20180818045031-cae1762e4858 // indirect
	github.com/kr/pretty v0.1.0 // indirect
	github.com/lib/pq v1.1.1
	github.com/nickalie/go-queue v0.0.0-20180806121409-9f0a137abb9c
	github.com/onsi/ginkgo v1.8.0 // indirect
	github.com/onsi/gomega v1.5.0 // indirect
	github.com/pkg/errors v0.8.1
	github.com/qiniu/api.v7 v7.2.5+incompatible
	github.com/qiniu/x v7.0.8+incompatible // indirect
	github.com/sirupsen/logrus v1.4.2
	github.com/streadway/amqp v0.0.0-20190404075320-75d898a42a94 // indirect
	github.com/stretchr/testify v1.3.0
	github.com/tidwall/btree v0.0.0-20170113224114-9876f1454cf0 // indirect
	github.com/tidwall/buntdb v1.1.0
	github.com/tidwall/gjson v1.3.0 // indirect
	github.com/tidwall/grect v0.0.0-20161006141115-ba9a043346eb // indirect
	github.com/tidwall/rtree v0.0.0-20180113144539-6cd427091e0e // indirect
	github.com/tidwall/tinyqueue v0.0.0-20180302190814-1e39f5511563 // indirect
	github.com/xujiajun/gorouter v1.2.0
	github.com/xujiajun/nutsdb v0.4.0
	github.com/xujiajun/utils v0.0.0-20190123093513-8bf096c4f53b
	google.golang.org/appengine v0.0.0-00010101000000-000000000000 // indirect
	gopkg.in/check.v1 v1.0.0-20180628173108-788fd7840127 // indirect
	gopkg.in/redis.v5 v5.2.9
	qiniupkg.com/x v7.0.8+incompatible // indirect

)
