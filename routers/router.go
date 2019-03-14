package routers

import (
	"GTMS/tests"
	"GTMS/v1/index"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/test", &test.TestController{}, "post:Test")
	beego.Router("/", &index.MainController{}, "get:Index")
}
