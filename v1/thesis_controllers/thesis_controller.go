package thesis_controllers

import (
	"GTMS/library/controller"
	"GTMS/library/gtms_error"
	"GTMS/library/helper"
	"GTMS/library/stringi"
	"GTMS/models/thesis_models"
	"GTMS/v1/forms"
	"math"
)

type ThesisController struct {
	controller.BaseController
}

//添加论文
func (this *ThesisController) AddThesis() {
	this.User = this.GetUser(this.Ctx.Request.Header.Get("X-Access-Token"))
	if this.User.IsGuest {
		this.RequireLogin()
		return
	}
	//只允许教师添加
	if this.User.Role != controller.ROLE_TEACHER {
		this.ErrorResponse(gtms_error.GetError("access_denied"))
		return
	}
	inputs := forms.AddThesisForm{}
	if err := this.ParseInput(&inputs); err.Code != 0 {
		this.ErrorResponse(err)
		return
	}
	err := thesis_models.AddThesis(&this.Request, &inputs)
	if err.Code != 0 {
		this.ErrorResponse(err)
		return
	}
	this.SuccessWithData(helper.JSON{})
}

//删除论文
func (this *ThesisController) DelThesis() {
	this.User = this.GetUser(this.Ctx.Request.Header.Get("X-Access-Token"))
	if this.User.IsGuest {
		this.RequireLogin()
		return
	}
	//只允许教师删除
	if this.User.Role != controller.ROLE_TEACHER {
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

//修改论文
func (this *ThesisController) UpdateThesis() {
	this.User = this.GetUser(this.Ctx.Request.Header.Get("X-Access-Token"))
	if this.User.IsGuest {
		this.RequireLogin()
		return
	}
	//只允许教师修改
	if this.User.Role != controller.ROLE_TEACHER {
		this.ErrorResponse(gtms_error.GetError("access_denied"))
		return
	}
	inputs := forms.UpdateThesisForm{}
	if err := this.ParseInput(&inputs); err.Code != 0 {
		this.ErrorResponse(err)
		return
	}
	err := thesis_models.UpdateThesis(&inputs, &this.Request)
	if err.Code != 0 {
		this.ErrorResponse(err)
		return
	}
	this.SuccessWithData(helper.JSON{})
}

//获取论文列表
func (this *ThesisController) ThesisList() {
	this.User = this.GetUser(this.Ctx.Request.Header.Get("X-Access-Token"))
	if this.User.IsGuest {
		this.RequireLogin()
		return
	}
	page, pageCount := this.GetPageInfo()
	thesiss, total := thesis_models.ThesisList(page, pageCount)
	pageInfo := controller.PageInfoWithEndPage{
		CurrentPage: page,
		IsEndPage:   stringi.Judge(len(thesiss) < pageCount, "yes", "no"),
		TotalPage:   int(math.Ceil(float64(total) / float64(pageCount))),
	}
	this.SuccessWithDataList(thesiss, pageInfo)
}
