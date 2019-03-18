package admin_account

import (
	"GTMS/library/controller"
	"GTMS/library/stringi"
	"GTMS/models/account/admin_account"
	"GTMS/v1/account"
	"math"
)

type AdminAccountController struct {
	controller.BaseController
}

//管理员登录
func (this *AdminAccountController) Login() {
	inputs := account.LoginForm{}
	if err := this.ParseInput(&inputs); err.Code != 0 {
		this.ErrorResponse(err)
		return
	}
	session, err := admin_account.Login(&inputs)
	if err.Code != 0 {
		this.ErrorResponse(err)
		return
	}
	this.SuccessWithData(session)
}

//获取管理员列表
func (this *AdminAccountController) AdminList() {
	this.User = this.GetUser(this.Ctx.Request.Header.Get("X-Access-Token"))
	if this.User.IsGuest {
		this.RequireLogin()
		return
	}
	page, pageCount := this.GetPageInfo()
	admins, total := admin_account.AdminList(page, pageCount)
	var pageInfo = controller.PageInfoWithEndPage{
		CurrentPage: page,
		IsEndPage:   stringi.Judge(len(admins) < pageCount, "yes", "no"),
		TotalPage:   int(math.Ceil(float64(total) / float64(pageCount))),
	}
	this.SuccessWithDataList(admins, pageInfo)
}
