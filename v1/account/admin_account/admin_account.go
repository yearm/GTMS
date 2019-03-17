package admin_account

import (
	"GTMS/library/controller"
	"GTMS/models/account/admin_account"
	"GTMS/v1/account"
)

type AdminAccountController struct {
	controller.BaseController
}

//管理员登录
func (this *AdminAccountController) Login() {
	inputs := account.SignInForm{}
	if err := this.ParseInput(&inputs); err.Code != 0 {
		this.ErrorResponse(err)
		return
	}
	session, err := admin_account.Login(&inputs)
	if err.Code != 0 {
		this.ErrorResponse(err)
		return
	}
	this.SuccessWithData(session)
}
