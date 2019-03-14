package routers

import (
	"GTMS/v1/index"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &index.MainController{}, "get:Index")
}
