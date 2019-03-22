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

type AccountController struct {
	controller.BaseController
}

//登录
func (this *AccountController) AccountLogin() {
	inputs := forms.LoginForm{}
	if err := this.ParseInput(&inputs); err.Code != 0 {
		this.ErrorResponse(err)
		return
	}
	session, err := account_models.AccountLogin(&inputs)
	if err.Code != 0 {
		this.ErrorResponse(err)
		return
	}
	this.SuccessWithData(session)
}

//登出
func (this *AccountController) AccountLogout() {
	this.User = this.GetUser(this.Ctx.Request.Header.Get("X-Access-Token"))
	if this.User.IsGuest {
		this.RequireLogin()
		return
	}
	err := account_models.AccountLogout(&this.Request)
	if err.Code != 0 {
		this.ErrorResponse(err)
		return
	}
	this.SuccessWithData(helper.JSON{})
}

//获取账号列表
func (this *AccountController) AccountList() {
	role := this.GetString("role")
	this.User = this.GetUser(this.Ctx.Request.Header.Get("X-Access-Token"))
	if this.User.IsGuest {
		this.RequireLogin()
		return
	}
	page, pageCount := this.GetPageInfo()
	admins, techs, stus, total := account_models.AccountList(role, page, pageCount)
	switch role {
	case controller.ROLE_ADMIN:
		this.SuccessWithDataList(admins, controller.PageInfoWithEndPage{
			CurrentPage: page,
			IsEndPage:   stringi.Judge(len(admins) < pageCount, "yes", "no"),
			TotalPage:   int(math.Ceil(float64(total) / float64(pageCount))),
		})
	case controller.ROLE_TEACHER:
		this.SuccessWithDataList(techs, controller.PageInfoWithEndPage{
			CurrentPage: page,
			IsEndPage:   stringi.Judge(len(techs) < pageCount, "yes", "no"),
			TotalPage:   int(math.Ceil(float64(total) / float64(pageCount))),
		})
	case controller.ROLE_STUDENT:
		this.SuccessWithDataList(stus, controller.PageInfoWithEndPage{
			CurrentPage: page,
			IsEndPage:   stringi.Judge(len(stus) < pageCount, "yes", "no"),
			TotalPage:   int(math.Ceil(float64(total) / float64(pageCount))),
		})
	default:
		this.SuccessWithDataList(helper.JSON{}, controller.PageInfoWithEndPage{})
	}
}

//修改账号信息
func (this *AccountController) AccountUpdate() {
	role := this.GetString("role")
	this.User = this.GetUser(this.Ctx.Request.Header.Get("X-Access-Token"))
	if this.User.IsGuest {
		this.RequireLogin()
		return
	}
	switch role {
	case controller.ROLE_ADMIN:
		inputs := forms.UpdateAdminForm{}
		if err := this.ParseInput(&inputs); err.Code != 0 {
			this.ErrorResponse(err)
			return
		}
		if inputs.AdminId != this.User.AdminId && this.User.Role != controller.ROLE_ADMIN {
			this.ErrorResponse(gtms_error.GetError("only_myself_or_admin"))
			return
		}
		if err := account_models.UpdateAdmin(&inputs); err.Code != 0 {
			this.ErrorResponse(err)
			return
		}
		this.SuccessWithData(helper.JSON{})
	case controller.ROLE_TEACHER:
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
	case controller.ROLE_STUDENT:
		inputs := forms.UpdateStudentForm{}
		if err := this.ParseInput(&inputs); err.Code != 0 {
			this.ErrorResponse(err)
			return
		}
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
}
