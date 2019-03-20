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

type TeacherAccountController struct {
	controller.BaseController
}

func (this *TeacherAccountController) TechLogin() {
	inputs := forms.LoginForm{}
	if err := this.ParseInput(&inputs); err.Code != 0 {
		this.ErrorResponse(err)
		return
	}
	session, err := account_models.TechLogin(&inputs)
	if err.Code != 0 {
		this.ErrorResponse(err)
		return
	}
	this.SuccessWithData(session)
}

//获取教师列表
func (this *TeacherAccountController) TechList() {
	this.User = this.GetUser(this.Ctx.Request.Header.Get("X-Access-Token"))
	if this.User.IsGuest {
		this.RequireLogin()
		return
	}
	page, pageCount := this.GetPageInfo()
	techs, total := account_models.TechList(page, pageCount)
	var pageInfo = controller.PageInfoWithEndPage{
		CurrentPage: page,
		IsEndPage:   stringi.Judge(len(techs) < pageCount, "yes", "no"),
		TotalPage:   int(math.Ceil(float64(total) / float64(pageCount))),
	}
	this.SuccessWithDataList(techs, pageInfo)
}


//修改教师信息
func (this *TeacherAccountController) UpdateTeacher() {
	this.User = this.GetUser(this.Ctx.Request.Header.Get("X-Access-Token"))
	if this.User.IsGuest {
		this.RequireLogin()
		return
	}
	inputs := forms.UpdateTeacherForm{}
	if err := this.ParseInput(&inputs); err.Code != 0 {
		this.ErrorResponse(err)
		return
	}
	if inputs.TechId != this.User.TechId && this.User.Role != controller.ROLE_ADMIN {
		this.ErrorResponse(gtms_error.GetError("only_myself_or_admin"))
		return
	}
	if err := account_models.UpdateTeacher(&inputs); err.Code != 0 {
		this.ErrorResponse(err)
		return
	}
	this.SuccessWithData(helper.JSON{})
}
