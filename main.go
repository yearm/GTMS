package main

import (
	"GTMS/boot"
	"GTMS/conf"
	_ "GTMS/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	conf.LoadingConf()
	logs.SetLogger(logs.AdapterMultiFile, `{"filename":"./runtime/log/project.log","separate":["emergency", "alert", "critical", "error", "warning", "notice", "info", "debug"]}`)
	logs.Async()
	boot.ConnectMySQL()
	boot.ConnectRedis()
	beego.Run()
}
