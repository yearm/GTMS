package routers

import (
	"GTMS/tests"
	"GTMS/v1/account/admin_account"
	"GTMS/v1/index"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/test", &test.TestController{}, "*:Test")
	beego.Router("/", &index.MainController{}, "get:Index")
	beego.Router("/v1/admin/signIn", &admin_account.AdminAccountController{}, "post:SignIn")
}
