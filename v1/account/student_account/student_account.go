package student_account

import (
	"GTMS/library/controller"
	"GTMS/models/account/student_account"
	"GTMS/v1/account"
)

type StudentAccountController struct {
	controller.BaseController
}

func (this *StudentAccountController) SignIn() {
	inputs := account.SignInForm{}
	if err := this.ParseInput(&inputs); err.Code != 0 {
		this.ErrorResponse(err)
		return
	}
	session, err := student_account.SignIn(&inputs)
	if err.Code != 0 {
		this.ErrorResponse(err)
		return
	}
	this.SuccessWithData(session)
}
