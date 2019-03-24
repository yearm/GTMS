package admin_controllers

import (
	"GTMS/library/controller"
	"GTMS/library/gtms_error"
)

type OpeningTimeController struct {
	controller.BaseController
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

}
