package main

import (
	"GTMS/boot"
	"GTMS/conf"
	_ "GTMS/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/plugins/cors"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU()) //多核运行
	conf.LoadingConf()
	//添加日志
	logs.SetLogger(logs.AdapterMultiFile, `{"filename":"./runtime/log/project.log","separate":["emergency", "alert", "critical", "error", "warning", "notice", "info", "debug"]}`)
	logs.Async()
	//处理跨域问题
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowAllOrigins:  true,                                                                                                                                 //设置为true则所有域名都可以访问本网站接口
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},                                                                                  //允许的请求类型
		AllowHeaders:     []string{"Origin", "Authorization", "X-Access-Token", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"}, //允许的头部信息
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},                            //允许暴露的头信息
		AllowCredentials: true,                                                                                                                                 //允许共享AuthTuffic证书
	}))
	boot.ConnectMySQL()
	boot.ConnectRedis()
	beego.Run()
}
