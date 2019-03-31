package thesis_controllers

import (
	"GTMS/library/controller"
	"GTMS/library/gtms_error"
	"GTMS/library/helper"
	"GTMS/library/stringi"
	"GTMS/models/thesis_models"
	"GTMS/v1/forms"
)

type SelectThesisController struct {
	controller.BaseController
}

//学生选题
func (this *SelectThesisController) SelectThesis() {
	this.User = this.GetUser(this.Ctx.Request.Header.Get("X-Access-Token"))
	if this.User.IsGuest {
		this.RequireLogin()
		return
	}
	//只允许学生选题
	if this.User.Role != controller.ROLE_STUDENT {
		this.ErrorResponse(gtms_error.GetError("access_denied"))
		return
	}
	inputs := forms.SelectThesisForm{}
	if err := this.ParseInput(&inputs); err.Code != 0 {
		this.ErrorResponse(err)
		return
	}
	err := thesis_models.SelectThesis(&inputs, &this.Request)
	if err.Code != 0 {
		this.ErrorResponse(err)
		return
	}
	this.SuccessWithData(helper.JSON{})
}

//教师确认学生选题(双向选择)
func (this *SelectThesisController) ConfirmSelectedThesis() {
	this.User = this.GetUser(this.Ctx.Request.Header.Get("X-Access-Token"))
	if this.User.IsGuest {
		this.RequireLogin()
		return
	}
	//只允许教师操作
	if this.User.Role != controller.ROLE_TEACHER {
		this.ErrorResponse(gtms_error.GetError("access_denied"))
		return
	}
	inputs := forms.ConfirmSelectedlThesisForm{}
	if err := this.ParseInput(&inputs); err.Code != 0 {
		this.ErrorResponse(err)
		return
	}
	err := thesis_models.ConfirmSelectedThesis(&inputs)
	if err.Code != 0 {
		this.ErrorResponse(err)
		return
	}
	this.SuccessWithData(helper.JSON{})
}

//获取教师未确认、确认的选题
func (this *SelectThesisController) GetNotOrConfirmThesis() {
	confirm := this.GetString("confirm")
	this.User = this.GetUser(this.Ctx.Request.Header.Get("X-Access-Token"))
	if this.User.IsGuest {
		this.RequireLogin()
		return
	}
	//只允许教师访问
	if this.User.Role != controller.ROLE_TEACHER {
		this.ErrorResponse(gtms_error.GetError("access_denied"))
		return
	}
	page, pageCount := this.GetPageInfo()
	ncThesis, total := thesis_models.GetNotOrConfirmThesis(&this.Request, page, pageCount, confirm)
	pageInfo := controller.PageInfoWithEndPage{
		CurrentPage: page,
		IsEndPage:   stringi.Judge(len(ncThesis) < pageCount, "yes", "no"),
		Total:       total,
	}
	this.SuccessWithDataList(ncThesis, pageInfo)
}

//获取所有已选题目
func (this *SelectThesisController) SelectedThesisList() {
	this.User = this.GetUser(this.Ctx.Request.Header.Get("X-Access-Token"))
	if this.User.IsGuest {
		this.RequireLogin()
		return
	}
	page, pageCount := this.GetPageInfo()
	confirmThesis, total := thesis_models.SelectedThesisList(page, pageCount)
	pageInfo := controller.PageInfoWithEndPage{
		CurrentPage: page,
		IsEndPage:   stringi.Judge(len(confirmThesis) < pageCount, "yes", "no"),
		Total:       total,
	}
	this.SuccessWithDataList(confirmThesis, pageInfo)
}

//获取学生自己已选题目(教师确认同意的)
func (this *SelectThesisController) GetThesis() {
	this.User = this.GetUser(this.Ctx.Request.Header.Get("X-Access-Token"))
	if this.User.IsGuest {
		this.RequireLogin()
		return
	}
	//只允许学生访问
	if this.User.Role != controller.ROLE_STUDENT {
		this.ErrorResponse(gtms_error.GetError("access_denied"))
		return
	}
	thesis := thesis_models.GetThesis(&this.Request)
	this.SuccessWithData(thesis)
}
