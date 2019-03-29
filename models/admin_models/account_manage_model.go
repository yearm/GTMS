package admin_models

import (
	"GTMS/boot"
	"GTMS/library/controller"
	"GTMS/library/db"
	"GTMS/library/gtms_error"
	"GTMS/library/helper"
	"GTMS/library/stringi"
	"GTMS/library/validator"
	"GTMS/models/account_models"
	"GTMS/v1/forms"
)

const (
	table_admin      = "admin"
	table_teacher    = "teacher"
	table_student    = "student"
	table_session    = "user_session"
	default_password = "123456"
)

func AddAdminAccount(opt *forms.AddAdminAccountForm) *validator.Error {
	o := boot.GetMasterMySQL()
	pwd, _ := helper.HashedPassword(default_password)
	opt.AdminSex = stringi.DefaultValue(opt.AdminSex, "男")
	admin := account_models.Admin{AdminId: opt.AdminId, Pwd: pwd, AdminName: opt.AdminName, AdminSex: opt.AdminSex, Phone: opt.Phone, Email: opt.Email}
	_, err := o.Insert(&admin)
	if err != nil {
		return gtms_error.GetError("insert_error")
	}
	return &validator.Error{}
}

func AddTechAccount(opt *forms.AddTechAccountForm) *validator.Error {
	o := boot.GetMasterMySQL()
	pwd, _ := helper.HashedPassword(default_password)
	opt.TechSex = stringi.DefaultValue(opt.TechSex, "男")
	tech := account_models.Teacher{
		TechId:            opt.TechId,
		Pwd:               pwd,
		TechName:          opt.TechName,
		TechSex:           opt.TechSex,
		Education:         opt.Education,
		Degree:            opt.Degree,
		ResearchDirection: opt.ResearchDirection,
		JobTitle:          opt.JobTitle,
		Job:               opt.Job,
		InstructNums:      opt.InstructNums,
		InstructMajor:     opt.InstructMajor,
		Email:             opt.Email,
		Phone:             opt.Phone,
		Qq:                opt.Qq,
		WeChat:            opt.WeChat,
	}
	_, err := o.Insert(&tech)
	if err != nil {
		return gtms_error.GetError("insert_error")
	}
	return &validator.Error{}
}

func AddStuAccount(opt *forms.AddStuAccountForm) *validator.Error {
	o := boot.GetMasterMySQL()
	pwd, _ := helper.HashedPassword(default_password)
	opt.StuSex = stringi.DefaultValue(opt.StuSex, "男")
	stu := account_models.Student{
		StuNo:        opt.StuNo,
		Pwd:          pwd,
		StuName:      opt.StuName,
		StuSex:       opt.StuSex,
		IdCard:       opt.IdCard,
		Birthplace:   opt.Birthplace,
		Department:   opt.Department,
		Major:        opt.Major,
		Class:        opt.Class,
		Phone:        opt.Phone,
		Qq:           opt.Qq,
		Email:        opt.Email,
		WeChat:       opt.WeChat,
		SchoolSystem: opt.SchoolSystem,
		EntryDate:    opt.EntryDate,
		Education:    opt.Education,
	}
	_, err := o.Insert(&stu)
	if err != nil {
		return gtms_error.GetError("insert_error")
	}
	return &validator.Error{}
}

func DelAccount(uid string, role string) {
	if role == controller.ROLE_ADMIN {
		db.Exec(db.DeleteSQL(table_admin, "admin_id", uid))
	} else if role == controller.ROLE_TEACHER {
		db.Exec(db.DeleteSQL(table_teacher, "tech_id", uid))
	} else if role == controller.ROLE_STUDENT {
		db.Exec(db.DeleteSQL(table_student, "stu_no", uid))
	}
	go func() {
		controller.DelRedisToken(uid)
		db.Exec(db.DeleteSQL(table_session, "uid", uid))
	}()

}

func ResetPwd(uid string, role string) {
	sql := `UPDATE @table SET @value WHERE @ukey = :uid`
	pwd, _ := helper.HashedPassword(default_password)
	value := db.Set(stringi.Form{
		"pwd": pwd,
	})
	if role == controller.ROLE_ADMIN {
		db.Exec(sql, stringi.Form{
			"table": table_admin,
			"value": value,
			"ukey":  "admin_id",
			"uid":   uid,
		})
	} else if role == controller.ROLE_TEACHER {
		db.Exec(sql, stringi.Form{
			"table": table_teacher,
			"value": value,
			"ukey":  "tech_id",
			"uid":   uid,
		})
	} else if role == controller.ROLE_STUDENT {
		db.Exec(sql, stringi.Form{
			"table": table_student,
			"value": value,
			"ukey":  "stu_no",
			"uid":   uid,
		})
	}
	go func() {
		controller.DelRedisToken(uid)
		db.Exec(db.DeleteSQL(table_session, "uid", uid))
	}()
}
