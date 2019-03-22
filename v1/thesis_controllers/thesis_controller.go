package thesis_controllers

import (
	"GTMS/library/controller"
	"GTMS/library/gtms_error"
	"GTMS/library/helper"
	"GTMS/models/thesis_models"
	"GTMS/v1/forms"
)

type ThesisController struct {
	controller.BaseController
}

func (this *ThesisController) AddThesis() {
	this.User = this.GetUser(this.Ctx.Request.Header.Get("X-Access-Token"))
	if this.User.IsGuest {
		this.RequireLogin()
		return
	}
	//只允许管理员和教师添加
	if this.User.Role != controller.ROLE_ADMIN && this.User.Role != controller.ROLE_TEACHER {
		this.ErrorResponse(gtms_error.GetError("access_denied"))
		return
	}
	inputs := forms.AddThesisForm{}
	if err := this.ParseInput(&inputs); err.Code != 0 {
		this.ErrorResponse(err)
		return
	}
	err := thesis_models.AddThesis(&inputs)
	if err.Code != 0 {
		this.ErrorResponse(err)
		return
	}
	this.SuccessWithData(helper.JSON{})
}

func (this *ThesisController) DelThesis() {
	this.User = this.GetUser(this.Ctx.Request.Header.Get("X-Access-Token"))
	if this.User.IsGuest {
		this.RequireLogin()
		return
	}
	//只允许管理员和教师删除
	if this.User.Role != controller.ROLE_ADMIN && this.User.Role != controller.ROLE_TEACHER {
		this.ErrorResponse(gtms_error.GetError("access_denied"))
		return
	}
	inputs := forms.DelThesisForm{}
	if err := this.ParseInput(&inputs); err.Code != 0 {
		this.ErrorResponse(err)
		return
	}
	err := thesis_models.DelThesis(&inputs)
	if err.Code != 0 {
		this.ErrorResponse(err)
		return
	}
	this.SuccessWithData(helper.JSON{})
}

func (this *ThesisController) UpdateThesis() {
	this.User = this.GetUser(this.Ctx.Request.Header.Get("X-Access-Token"))
	if this.User.IsGuest {
		this.RequireLogin()
		return
	}
	//只允许管理员和教师修改
	if this.User.Role != controller.ROLE_ADMIN && this.User.Role != controller.ROLE_TEACHER {
		this.ErrorResponse(gtms_error.GetError("access_denied"))
		return
	}
	inputs := forms.UpdateThesisForm{}
	if err := this.ParseInput(&inputs); err.Code != 0 {
		this.ErrorResponse(err)
		return
	}
	err := thesis_models.UpdateThesis(&inputs)
	if err.Code != 0 {
		this.ErrorResponse(err)
		return
	}
	this.SuccessWithData(helper.JSON{})
}

func (this *ThesisController) ThesisList() {
	this.User = this.GetUser(this.Ctx.Request.Header.Get("X-Access-Token"))
	if this.User.IsGuest {
		this.RequireLogin()
		return
	}
}
