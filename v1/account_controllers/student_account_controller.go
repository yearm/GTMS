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

type StudentAccountController struct {
	controller.BaseController
}

//登录
func (this *StudentAccountController) StuLogin() {
	inputs := forms.LoginForm{}
	if err := this.ParseInput(&inputs); err.Code != 0 {
		this.ErrorResponse(err)
		return
	}
	session, err := account_models.StuLogin(&inputs)
	if err.Code != 0 {
		this.ErrorResponse(err)
		return
	}
	this.SuccessWithData(session)
}

//获取学生列表
func (this *StudentAccountController) StuList() {
	this.User = this.GetUser(this.Ctx.Request.Header.Get("X-Access-Token"))
	if this.User.IsGuest {
		this.RequireLogin()
		return
	}
	page, pageCount := this.GetPageInfo()
	stus, total := account_models.StuList(page, pageCount)
	var pageInfo = controller.PageInfoWithEndPage{
		CurrentPage: page,
		IsEndPage:   stringi.Judge(len(stus) < pageCount, "yes", "no"),
		TotalPage:   int(math.Ceil(float64(total) / float64(pageCount))),
	}
	this.SuccessWithDataList(stus, pageInfo)
}

//修改学生信息
func (this *StudentAccountController) UpdateStudent() {
	this.User = this.GetUser(this.Ctx.Request.Header.Get("X-Access-Token"))
	if this.User.IsGuest {
		this.RequireLogin()
		return
	}
	inputs := forms.UpdateStudentForm{}
	if err := this.ParseInput(&inputs); err.Code != 0 {
		this.ErrorResponse(err)
		return
	}
	//仅管理员和自己能修改
	if inputs.StuNo != this.User.StuNo && this.User.Role != controller.ROLE_ADMIN {
		this.ErrorResponse(gtms_error.GetError("only_myself_or_admin"))
		return
	}
	if err := account_models.UpdateStudent(&inputs); err.Code != 0 {
		this.ErrorResponse(err)
		return
	}
	this.SuccessWithData(helper.JSON{})
}
