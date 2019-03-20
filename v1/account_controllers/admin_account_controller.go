package account_controllers

import (
	"GTMS/library/controller"
	"GTMS/library/gtms_error"
	"GTMS/library/helper"
	"GTMS/library/stringi"
	"GTMS/models/account_models"
	"GTMS/v1/forms"
	"math"
)

type AdminAccountController struct {
	controller.BaseController
}

//管理员登录
func (this *AdminAccountController) AdminLogin() {
	inputs := forms.LoginForm{}
	if err := this.ParseInput(&inputs); err.Code != 0 {
		this.ErrorResponse(err)
		return
	}
	session, err := account_models.AdminLogin(&inputs)
	if err.Code != 0 {
		this.ErrorResponse(err)
		return
	}
	this.SuccessWithData(session)
}

//获取管理员列表
func (this *AdminAccountController) AdminList() {
	this.User = this.GetUser(this.Ctx.Request.Header.Get("X-Access-Token"))
	if this.User.IsGuest {
		this.RequireLogin()
		return
	}
	page, pageCount := this.GetPageInfo()
	admins, total := account_models.AdminList(page, pageCount)
	var pageInfo = controller.PageInfoWithEndPage{
		CurrentPage: page,
		IsEndPage:   stringi.Judge(len(admins) < pageCount, "yes", "no"),
		TotalPage:   int(math.Ceil(float64(total) / float64(pageCount))),
	}
	this.SuccessWithDataList(admins, pageInfo)
}

//修改管理员信息
func (this *AdminAccountController) UpdateAdmin() {
	this.User = this.GetUser(this.Ctx.Request.Header.Get("X-Access-Token"))
	if this.User.IsGuest {
		this.RequireLogin()
		return
	}
	inputs := forms.UpdateAdminForm{}
	if err := this.ParseInput(&inputs); err.Code != 0 {
		this.ErrorResponse(err)
		return
	}
	if inputs.AdminId != this.User.AdminId && this.User.Role != controller.ROLE_ADMIN {
		this.ErrorResponse(gtms_error.GetError("only_myself_or_admin"))
		return
	}
	if err := account_models.UpdateAdmin(&inputs); err.Code != 0 {
		this.ErrorResponse(err)
		return
	}
	this.SuccessWithData(helper.JSON{})
}
