package routers

import (
	"GTMS/tests"
	"GTMS/v1/account/admin_account"
	"GTMS/v1/account/student_account"
	"GTMS/v1/account/teacher_account"
	"GTMS/v1/admin/account_manage"
	"GTMS/v1/index"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/test", &test.TestController{}, "*:Test")
	beego.Router("/", &index.MainController{}, "get:Index")

	//登录
	beego.Router("/v1/admin/login", &admin_account.AdminAccountController{}, "post:Login")
	beego.Router("/v1/teacher/login", &teacher_account.TeacherAccountController{}, "post:Login")
	beego.Router("/v1/student/login", &student_account.StudentAccountController{}, "post:Login")

	//管理员操作
	beego.Router("/v1/admin/account", &account_manage.AccountManageController{}, "post:AddAccount")
}
