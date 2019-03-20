package routers

import (
	"GTMS/tests"
	"GTMS/v1/account_controllers"
	"GTMS/v1/admin_controllers"
	"GTMS/v1/index"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/test", &test.TestController{}, "*:Test")
	beego.Router("/", &index.MainController{}, "get:Index")

	//登录
	beego.Router("/v1/admin/login", &account_controllers.AdminAccountController{}, "post:AdminLogin")
	beego.Router("/v1/teacher/login", &account_controllers.TeacherAccountController{}, "post:TechLogin")
	beego.Router("/v1/student/login", &account_controllers.StudentAccountController{}, "post:StuLogin")

	//管理员添加删除账号
	beego.Router("/v1/account", &admin_controllers.AccountManageController{}, "post:AddAccount;delete:DelAccount")

	//修改信息
	beego.Router("/v1/admin", &account_controllers.AdminAccountController{}, "get:AdminList;put:UpdateAdmin")
	beego.Router("/v1/teacher", &account_controllers.TeacherAccountController{}, "get:TechList;put:UpdateTeacher")
	beego.Router("/v1/student", &account_controllers.StudentAccountController{}, "get:StuList;put:UpdateStudent")

	//重置密码
	beego.Router("v1/account/sendEmail", &account_controllers.ResetPwdController{}, "post:SendEmailToResetPwd")
	beego.Router("v1/account/resetPwd", &account_controllers.ResetPwdController{}, "put:ResetPwd")
}
