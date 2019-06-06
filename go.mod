module apiend-core

go 1.12

replace (
	golang.org/x/net => github.com/golang/net v0.0.0-20180821023952-922f4815f713
	golang.org/x/sync => github.com/golang/sync v0.0.0-20181108010431-42b317875d0f
	golang.org/x/sys => github.com/golang/sys v0.0.0-20180905080454-ebe1bf3edb33
	golang.org/x/text => github.com/golang/text v0.0.0-20170915032832-14c0d48ead0c
)

require (
	github.com/globalsign/mgo v0.0.0-20181015135952-eeefdecb41b8
	github.com/gogf/gf v1.6.17
)
