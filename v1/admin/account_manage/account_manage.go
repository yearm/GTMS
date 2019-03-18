package account_manage

import (
	"GTMS/library/controller"
	"GTMS/library/gtms_error"
	"GTMS/library/helper"
	models "GTMS/models/admin"
	"GTMS/v1/admin"
)

type AccountManageController struct {
	controller.BaseController
}

//添加账号
func (this *AccountManageController) AddAccount() {
	this.User = this.GetUser(this.Ctx.Request.Header.Get("X-Access-Token"))
	//限制游客
	if this.User.IsGuest {
		this.RequireLogin()
		return
	}
	//只允许管理员添加
	if this.User.Role != controller.ROLE_ADMIN {
		this.ErrorResponse(gtms_error.GetError("access_denied"))
		return
	}
	//解析参数
	inputs := admin.AddAccountForm{}
	if err := this.ParseInput(&inputs); err.Code != 0 {
		this.ErrorResponse(err)
		return
	}
	//调用model
	if err := models.AddAccount(&inputs); err.Code != 0 {
		this.ErrorResponse(err)
		return
	}
	this.SuccessWithData(helper.JSON{})
}
