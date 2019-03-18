package admin

import (
	"GTMS/library/controller"
	"GTMS/library/db"
	"GTMS/library/gtms_error"
	"GTMS/library/helper"
	"GTMS/library/stringi"
	"GTMS/library/validator"
	"GTMS/v1/admin"
	"github.com/astaxie/beego/logs"
)

const (
	table_admin      = "admin"
	table_teacher    = "teacher"
	table_student    = "student"
	table_session    = "user_session"
	default_password = "123456"
)

func AddAccount(opt *admin.AddAccountForm) *validator.Error {
	if opt.Role == controller.ROLE_ADMIN {
		userVal := stringi.Forms{}
		for i := 0; i < len(opt.Uids); i++ {
			pwd, _ := helper.HashedPassword(default_password)
			userVal = append(userVal, stringi.Form{
				"admin_id":   opt.Uids[i],
				"pwd":        pwd,
				"admin_name": opt.Names[i],
			})
		}
		_, err := db.Exec(db.InsertAllSQL(table_admin, userVal))
		if err != nil {
			logs.Error(err)
			return gtms_error.GetError("insert_error")
		}
	} else if opt.Role == controller.ROLE_STUDENT {
		userVal := stringi.Forms{}
		for i := 0; i < len(opt.Uids); i++ {
			pwd, _ := helper.HashedPassword(default_password)
			userVal = append(userVal, stringi.Form{
				"stu_no":   opt.Uids[i],
				"pwd":      pwd,
				"stu_name": opt.Names[i],
			})
		}
		_, err := db.Exec(db.InsertAllSQL(table_student, userVal))
		if err != nil {
			logs.Error(err)
			return gtms_error.GetError("insert_error")
		}
	} else if opt.Role == controller.ROLE_TEACHER {
		userVal := stringi.Forms{}
		for i := 0; i < len(opt.Uids); i++ {
			pwd, _ := helper.HashedPassword(default_password)
			userVal = append(userVal, stringi.Form{
				"tech_id":   opt.Uids[i],
				"pwd":       pwd,
				"tech_name": opt.Names[i],
			})
		}
		_, err := db.Exec(db.InsertAllSQL(table_teacher, userVal))
		if err != nil {
			logs.Error(err)
			return gtms_error.GetError("insert_error")
		}
	}
	return &validator.Error{}
}

func DelAccount(opt *admin.DelAccountForm) {
	if opt.Role == controller.ROLE_ADMIN {
		db.Exec(db.DeleteSQL(table_admin, "admin_id", opt.Uids))
		go func() {
			for _, v := range opt.Uids {
				controller.DelRedisToken(v)
			}
		}()
	} else if opt.Role == controller.ROLE_TEACHER {
		db.Exec(db.DeleteSQL(table_teacher, "tech_id", opt.Uids))
		go func() {
			for _, v := range opt.Uids {
				controller.DelRedisToken(v)
			}
		}()
	} else if opt.Role == controller.ROLE_STUDENT {
		db.Exec(db.DeleteSQL(table_student, "stu_no", opt.Uids))
		go func() {
			for _, v := range opt.Uids {
				controller.DelRedisToken(v)
			}
		}()
	}
	go db.Exec(db.DeleteSQL(table_session, "uid", opt.Uids))
}
