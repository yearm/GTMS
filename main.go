package main

import (
	"GTMS/boot"
	"GTMS/conf"
	_ "GTMS/routers"
	"github.com/astaxie/beego"
)

func main() {
	conf.LoadingConf()
	boot.ConnectMySQL()
	boot.ConnectRedis()
	beego.Run()
}
