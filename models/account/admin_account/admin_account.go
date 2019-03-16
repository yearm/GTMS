package admin_account

import (
	"GTMS/library/controller"
	"GTMS/library/db"
	"GTMS/library/gtms_error"
	"GTMS/library/helper"
	"GTMS/library/stringi"
	"GTMS/library/validator"
	"GTMS/v1/account"
)

func SignIn(opt *account.SignInForm) (*controller.Session, *validator.Error) {
	sal := `SELECT pwd FROM admin WHERE admin_id = :account`
	qs := struct {
		Pwd string
	}{}
	db.QueryRow(sal, stringi.Form{
		"account": opt.Account,
	}, &qs)
	b := helper.CheckHashedPassword(qs.Pwd, opt.Password)
	if b {
		return &controller.Session{}, &validator.Error{}
	} else {
		return nil, gtms_error.GetError("sign_in_error")
	}
}
