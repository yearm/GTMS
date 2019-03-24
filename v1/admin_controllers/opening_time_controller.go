package admin_controllers

import (
	"GTMS/library/controller"
	"GTMS/library/gtms_error"
	"GTMS/library/helper"
	"GTMS/models/admin_models"
	"GTMS/v1/forms"
)

type OpeningTimeController struct {
	controller.BaseController
}

//获取开放时间
func (this *OpeningTimeController) GetOpeningTime() {
	this.User = this.GetUser(this.Ctx.Request.Header.Get("X-Access-Token"))
	if this.User.IsGuest {
		this.RequireLogin()
		return
	}
	//只允许管理员操作
	if this.User.Role != controller.ROLE_ADMIN {
		this.ErrorResponse(gtms_error.GetError("access_denied"))
		return
	}
	openingTime := admin_models.GetOpeningTime()
	this.SuccessWithData(openingTime)
}

//修改开放时间
func (this *OpeningTimeController) UpdateOpeningTime() {
	this.User = this.GetUser(this.Ctx.Request.Header.Get("X-Access-Token"))
	if this.User.IsGuest {
		this.RequireLogin()
		return
	}
	//只允许管理员操作
	if this.User.Role != controller.ROLE_ADMIN {
		this.ErrorResponse(gtms_error.GetError("access_denied"))
		return
	}
	inputs := forms.UpdateOpeningTimeForm{}
	if err := this.ParseInput(&inputs); err.Code != 0 {
		this.ErrorResponse(err)
		return
	}
	err := admin_models.UpdateOpeningTime(&inputs, &this.Request)
	if err.Code != 0 {
		this.ErrorResponse(err)
		return
	}
	this.SuccessWithData(helper.JSON{})
}
