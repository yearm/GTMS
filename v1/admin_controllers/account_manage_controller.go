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
	role := this.GetString("role")
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
	switch role {
	case controller.ROLE_ADMIN:
		inputs := forms.AddAdminAccountForm{}
		if err := this.ParseInput(&inputs); err.Code != 0 {
			this.ErrorResponse(err)
			return
		}
		//调用model
		if err := admin_models.AddAdminAccount(&inputs); err.Code != 0 {
			this.ErrorResponse(err)
			return
		}
	case controller.ROLE_TEACHER:
		inputs := forms.AddTechAccountForm{}
		if err := this.ParseInput(&inputs); err.Code != 0 {
			this.ErrorResponse(err)
			return
		}
		//调用model
		if err := admin_models.AddTechAccount(&inputs); err.Code != 0 {
			this.ErrorResponse(err)
			return
		}
	case controller.ROLE_STUDENT:
		inputs := forms.AddStuAccountForm{}
		if err := this.ParseInput(&inputs); err.Code != 0 {
			this.ErrorResponse(err)
			return
		}
		//调用model
		if err := admin_models.AddStuAccount(&inputs); err.Code != 0 {
			this.ErrorResponse(err)
			return
		}
	}
	this.SuccessWithData(helper.JSON{})
}

//删除账号
func (this *AccountManageController) DelAccount() {
	role := this.GetString("role")
	uid := this.GetString("uid")
	this.User = this.GetUser(this.Ctx.Request.Header.Get("X-Access-Token"))
	if this.User.IsGuest {
		this.RequireLogin()
		return
	}
	if this.User.Role != controller.ROLE_ADMIN || this.User.AdminId == uid {
		this.ErrorResponse(gtms_error.GetError("access_denied"))
		return
	}
	admin_models.DelAccount(uid, role)
	this.SuccessWithData(helper.JSON{})
}

//重置密码
func (this *AccountManageController) ResetPwd() {
	role := this.GetString("role")
	uid := this.GetString("uid")
	this.User = this.GetUser(this.Ctx.Request.Header.Get("X-Access-Token"))
	if this.User.IsGuest {
		this.RequireLogin()
		return
	}
	if this.User.Role != controller.ROLE_ADMIN || this.User.AdminId == uid {
		this.ErrorResponse(gtms_error.GetError("access_denied"))
		return
	}
	admin_models.ResetPwd(uid, role)
	this.SuccessWithData(helper.JSON{})
}
