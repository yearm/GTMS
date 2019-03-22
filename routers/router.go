package routers

import (
	"GTMS/tests"
	"GTMS/v1/account_controllers"
	"GTMS/v1/admin_controllers"
	"GTMS/v1/index"
	"GTMS/v1/thesis_controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/test", &test.TestController{}, "*:Test")
	beego.Router("/", &index.MainController{}, "get:Index")

	//登录
	beego.Router("v1/account/login", &account_controllers.AccountController{}, "post:AccountLogin")

	//登出
	beego.Router("/v1/account/logout", &account_controllers.AccountController{}, "delete:AccountLogout")

	//管理员添加删除账号
	beego.Router("/v1/account", &admin_controllers.AccountManageController{}, "post:AddAccount;delete:DelAccount")

	//获取账号列表、修改账号信息
	beego.Router("/v1/account", &account_controllers.AccountController{}, "get:AccountList;put:AccountUpdate")

	//发送邮件重置密码
	beego.Router("v1/account/sendEmail", &account_controllers.ResetPwdController{}, "post:SendEmailToResetPwd")
	beego.Router("v1/account/resetPwd", &account_controllers.ResetPwdController{}, "put:ResetPwd")

	//论文
	beego.Router("/v1/thesis", &thesis_controllers.ThesisController{}, "get:ThesisList;post:AddThesis;put:UpdateThesis;delete:DelThesis")
}
