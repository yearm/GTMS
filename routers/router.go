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
	beego.Router("/v1/account/login", &account_controllers.AccountController{}, "post:AccountLogin")

	//登出
	beego.Router("/v1/account/logout", &account_controllers.AccountController{}, "delete:AccountLogout")

	//管理员添加删除账号
	beego.Router("/v1/account", &admin_controllers.AccountManageController{}, "post:AddAccount;delete:DelAccount")

	//获取账号列表、修改账号信息
	beego.Router("/v1/account", &account_controllers.AccountController{}, "get:AccountList;put:AccountUpdate")

	//发送邮件重置密码
	beego.Router("/v1/account/sendEmail", &account_controllers.ResetPwdController{}, "post:SendEmailToResetPwd")
	beego.Router("/v1/account/resetPwd", &account_controllers.ResetPwdController{}, "put:ResetPwd")

	//论文信息的增删改查、论文的上传下载
	beego.Router("/v1/thesis", &thesis_controllers.ThesisController{}, "get:ThesisList;post:AddThesis;put:UpdateThesis;delete:DelThesis;post:UploadThesis")

	//学生选题
	beego.Router("/v1/selectThesis/student", &thesis_controllers.SelectThesisController{}, "post:SelectThesis")

	//获取教师确认、未确认的选题、教师确认学生选题(双向选择)
	beego.Router("/v1/confirmThesis/teacher", &thesis_controllers.SelectThesisController{}, "get:GetNotOrConfirmThesis;put:ConfirmSelectedThesis")

	//获取所有已选题目
	beego.Router("/v1/selectedThesis", &thesis_controllers.SelectThesisController{}, "get:SelectedThesisList")

	//获取学生自己已选题目
	beego.Router("/v1/thesis/myself", &thesis_controllers.SelectThesisController{}, "get:GetThesis")

	//获取开放时间
	beego.Router("/v1/openingTime", &admin_controllers.OpeningTimeController{}, "get:GetOpeningTime;put:UpdateOpeningTime")
}
