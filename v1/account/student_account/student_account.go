package student_account

import (
	"GTMS/library/controller"
	"GTMS/models/account/student_account"
	"GTMS/v1/account"
)

type StudentAccountController struct {
	controller.BaseController
}

func (this *StudentAccountController) Login() {
	inputs := account.LoginForm{}
	if err := this.ParseInput(&inputs); err.Code != 0 {
		this.ErrorResponse(err)
		return
	}
	session, err := student_account.Login(&inputs)
	if err.Code != 0 {
		this.ErrorResponse(err)
		return
	}
	this.SuccessWithData(session)
}
