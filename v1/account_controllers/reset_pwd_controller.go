package account_controllers

import (
	"GTMS/library/controller"
	"GTMS/library/helper"
	"GTMS/models/account_models"
	"GTMS/v1/forms"
)

type ResetPwdController struct {
	controller.BaseController
}

//发送重置密码邮件
func (this *ResetPwdController) SendEmailToResetPwd() {
	inputs := forms.SendEmailToResetPwd{}
	if err := this.ParseInput(&inputs); err.Code != 0 {
		this.ErrorResponse(err)
		return
	}
	err := account_models.SendEmailToResetPwd(&inputs)
	if err.Code != 0 {
		this.ErrorResponse(err)
		return
	}
	this.SuccessWithData(helper.JSON{})
}

//重置密码
func (this *ResetPwdController) ResetPwd() {
	inputs := forms.ResetPwdForm{}
	if err := this.ParseInput(&inputs); err.Code != 0 {
		this.ErrorResponse(err)
		return
	}
	err := account_models.ResetPwd(&inputs)
	if err.Code != 0 {
		this.ErrorResponse(err)
		return
	}
	this.SuccessWithData(helper.JSON{})
}
