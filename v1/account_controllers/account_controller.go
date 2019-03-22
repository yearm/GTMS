package account_controllers

import (
	"GTMS/library/controller"
	"GTMS/library/helper"
	"GTMS/models/account_models"
)

type AccountController struct {
	controller.BaseController
}

//登出
func (this *AccountController) AccountLogout() {
	this.User = this.GetUser(this.Ctx.Request.Header.Get("X-Access-Token"))
	if this.User.IsGuest {
		this.RequireLogin()
		return
	}
	err := account_models.AccountLogout(&this.Request)
	if err.Code != 0 {
		this.ErrorResponse(err)
		return
	}
	this.SuccessWithData(helper.JSON{})
}
