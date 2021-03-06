/*
    fileName: boot
    author: diogoxiang
    date: 2019/6/6
*/
package boot

import (
	_ "apiend-core/app/lib/cdb" // 初始化链接mongo
	"github.com/gogf/gf/g"
	"github.com/gogf/gf/g/net/ghttp"
	"github.com/gogf/gf/g/os/glog"
)

// 用于应用初始化。
func init() {
	//v := g.View()
	c := g.Config()
	s := g.Server()

	// 配置对象及视图对象配置
	c.AddPath("config")
	// TODO 模板. 没用默认的模板解析
	//v.AddPath("template")
	//v.SetDelimiters("${", "}")

	// glog配置
	logpath := c.GetString("setting.logpath")
	glog.SetPath(logpath)
	// glog.SetStdPrint(true)
	glog.SetDebug(true)

	s.SetLogPath(logpath)

	// Web Server配置 后台不提共 静态目录
	// publicPath := c.GetString("setting.publicPath")
	// s.SetServerRoot(publicPath)
	// TODO 关闭 静态文件服务
	s.SetFileServerEnabled(false)



	// 接口地址的 URI方式
	// s.SetNameToUriType(ghttp.NAME_TO_URI_TYPE_ALLLOWER)
	s.SetNameToUriType(ghttp.NAME_TO_URI_TYPE_CAMEL)


	s.SetErrorLogEnabled(true)
	s.SetAccessLogEnabled(true)
	// 后台端口
	apiPort := c.GetInt("setting.apiport")
	s.SetPort(apiPort)
}


