package admin_controllers

import (
	"GTMS/library/controller"
	"GTMS/library/gtms_error"
	"GTMS/library/helper"
	"GTMS/models/admin_models"
	"GTMS/v1/forms"
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
	//只允许管理员操作
	if this.User.Role != controller.ROLE_ADMIN {
		this.ErrorResponse(gtms_error.GetError("access_denied"))
		return
	}
	//解析参数
	inputs := forms.AddAccountForm{}
	if err := this.ParseInput(&inputs); err.Code != 0 {
		this.ErrorResponse(err)
		return
	}
	//调用model
	if err := admin_models.AddAccount(&inputs); err.Code != 0 {
		this.ErrorResponse(err)
		return
	}
	this.SuccessWithData(helper.JSON{})
}


//删除账号
func (this *AccountManageController) DelAccount() {
	this.User = this.GetUser(this.Ctx.Request.Header.Get("X-Access-Token"))
	if this.User.IsGuest {
		this.RequireLogin()
		return
	}
	if this.User.Role != controller.ROLE_ADMIN {
		this.ErrorResponse(gtms_error.GetError("access_denied"))
		return
	}
	inputs := forms.DelAccountForm{}
	if err := this.ParseInput(&inputs); err.Code != 0 {
		this.ErrorResponse(err)
		return
	}
	admin_models.DelAccount(&inputs)
	this.SuccessWithData(helper.JSON{})
}
