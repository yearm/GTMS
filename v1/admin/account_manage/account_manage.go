package account_manage

import "GTMS/library/controller"

type AccountManageController struct {
	controller.BaseController
}

func (this *AccountManageController) AddStuAccount() {
	if this.User.IsGuest {
		this.RequireSignin()
		return
	}
}
