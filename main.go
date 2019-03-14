package main

import (
	"GTMS/boot"
	"GTMS/conf"
	_ "GTMS/routers"
	"github.com/astaxie/beego"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	conf.LoadingConf()
	boot.ConnectMySQL()
	boot.ConnectRedis()
	beego.Run()
}
