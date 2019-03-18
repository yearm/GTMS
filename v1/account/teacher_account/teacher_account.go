package teacher_account

import (
	"GTMS/library/controller"
	"GTMS/models/account/teacher_account"
	"GTMS/v1/account"
)

type TeacherAccountController struct {
	controller.BaseController
}

func (this *TeacherAccountController) Login() {
	inputs := account.LoginForm{}
	if err := this.ParseInput(&inputs); err.Code != 0 {
		this.ErrorResponse(err)
		return
	}
	session, err := teacher_account.Login(&inputs)
	if err.Code != 0 {
		this.ErrorResponse(err)
		return
	}
	this.SuccessWithData(session)
}
