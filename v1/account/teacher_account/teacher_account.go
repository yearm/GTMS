package teacher_account

import (
	"GTMS/library/controller"
	"GTMS/models/account/teacher_account"
	"GTMS/v1/account"
)

type TeacherAccountController struct {
	controller.BaseController
}

func (this *TeacherAccountController) SignIn() {
	inputs := account.SignInForm{}
	if err := this.ParseInput(&inputs); err.Code != 0 {
		this.ErrorResponse(err)
		return
	}
	session, err := teacher_account.SignIn(&inputs)
	if err.Code != 0 {
		this.ErrorResponse(err)
		return
	}
	this.SuccessWithData(session)
}
