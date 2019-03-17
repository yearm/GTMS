package account_manage

import (
	"GTMS/library/controller"
	"GTMS/library/gtms_error"
	"GTMS/library/helper"
)

type AccountManageController struct {
	controller.BaseController
}

//添加学生
func (this *AccountManageController) AddStuAccount() {
	this.User = this.GetUser(this.Ctx.Request.Header.Get("X-Access-Token"))
	if this.User.IsGuest {
		this.RequireLogin()
		return
	}
	if this.User.Role != controller.ROLE_ADMIN {
		this.ErrorResponse(gtms_error.GetError("access_denied"))
		return
	}
	this.SuccessWithData(helper.JSON{})
}
