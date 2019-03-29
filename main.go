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
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "X-Access-Token", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		AllowCredentials: true,
	}))
	boot.ConnectMySQL()
	boot.ConnectRedis()
	beego.Run()
}
