package account_models

import (
	"GTMS/library/controller"
	"GTMS/library/db"
	"GTMS/library/gtms_error"
	"GTMS/library/stringi"
	"GTMS/library/validator"
	"GTMS/v1/forms"
)

func SendEmailToResetPwd(opt *forms.SendEmailToResetPwd) *validator.Error {
	sql := `SELECT @ukey,email FROM @table WHERE @ukey = :uid`
	if opt.Role == controller.ROLE_ADMIN {
		admin := Admin{}
		db.QueryRow(sql, stringi.Form{
			"table": "admin",
			"ukey":  "admin_id",
			"uid":   opt.Uid,
		}, &admin)
		if admin.AdminId == "" {
			return gtms_error.GetError("account_not_exist")
		}
	} else if opt.Role == controller.ROLE_TEACHER {
		tech := Teacher{}
		db.QueryRow(sql, stringi.Form{
			"table": "teacher",
			"ukey":  "tech_id",
			"uid":   opt.Uid,
		}, &tech)
		if tech.TechId == "" {
			return gtms_error.GetError("account_not_exist")
		}
	} else if opt.Role == controller.ROLE_STUDENT {
		stu := Student{}
		db.QueryRow(sql, stringi.Form{
			"table": "student",
			"ukey":  "stu_id",
			"uid":   opt.Uid,
		}, &stu)
		if stu.StuNo == "" {
			return gtms_error.GetError("account_not_exist")
		}
	}
	return &validator.Error{}
}

func ResetPwd(opt *forms.ResetPwdForm) *validator.Error {
	return &validator.Error{}
}
