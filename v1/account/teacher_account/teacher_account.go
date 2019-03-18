package teacher_account

import (
	"GTMS/library/controller"
	"GTMS/library/stringi"
	"GTMS/models/account/teacher_account"
	"GTMS/v1/account"
	"math"
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

//获取教师列表
func (this *TeacherAccountController) TechList() {
	this.User = this.GetUser(this.Ctx.Request.Header.Get("X-Access-Token"))
	if this.User.IsGuest {
		this.RequireLogin()
		return
	}
	page, pageCount := this.GetPageInfo()
	techs, total := teacher_account.TechList(page, pageCount)
	var pageInfo = controller.PageInfoWithEndPage{
		CurrentPage: page,
		IsEndPage:   stringi.Judge(len(techs) < pageCount, "yes", "no"),
		TotalPage:   int(math.Ceil(float64(total) / float64(pageCount))),
	}
	this.SuccessWithDataList(techs, pageInfo)
}
