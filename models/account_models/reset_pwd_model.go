package account_models

import (
	"GTMS/boot"
	"GTMS/library/controller"
	"GTMS/library/db"
	"GTMS/library/gtms_error"
	"GTMS/library/helper"
	"GTMS/library/stringi"
	"GTMS/library/validator"
	"GTMS/v1/forms"
	"github.com/json-iterator/go"
	"time"
)

const (
	subject = "黄冈师范学院毕业论文管理系统重置密码"
	url     = "https://www.upcoder.cn/account/resetPwd"
)

func SendEmailToResetPwd(opt *forms.SendEmailToResetPwd) *validator.Error {
	sql := `SELECT * FROM @table WHERE @ukey = :uid`
	switch opt.Role {
	case controller.ROLE_ADMIN:
		admin := Admin{}
		db.QueryRow(sql, stringi.Form{
			"table": "admin",
			"ukey":  "admin_id",
			"uid":   opt.Uid,
		}, &admin)
		if admin.AdminId == "" {
			return gtms_error.GetError("account_not_exist")
		}
		if admin.Email == "" {
			return gtms_error.GetError("not_bind_email")
		}
		//创建一个临时的token
		accessToken := helper.CreateToken()
		session := controller.Session{
			AccessToken: accessToken,
			Role:        controller.ROLE_ADMIN,
			AdminInfo: controller.AdminInfo{
				AdminId: admin.AdminId,
			},
		}
		s, _ := jsoniter.MarshalToString(session)
		boot.CACHE.Set(accessToken, s, time.Minute*5) //设置5分钟缓存
		body := `<h4>` + admin.AdminName + `:<br>
			你好！<br>
			您在本系统中的用户名是` + admin.AdminId + `，请点击以下超链接，在浏览器中重置自己的密码(5分钟内有效):<br><br>
			<a href="` + url + `?token=` + accessToken + `">` + url + `?token=` + accessToken + `</a><br><br>
			黄冈师范学院毕业论文管理系统<br><br>
			注:此邮件为本系统所发，非对方邮箱所发，所以请勿回邮。
		</h4>`
		//发送邮件
		if err := boot.SendEmail(admin.Email, admin.AdminName, subject, body); err != nil {
			return gtms_error.GetError("send_email_error")
		}
	case controller.ROLE_TEACHER:
		tech := Teacher{}
		db.QueryRow(sql, stringi.Form{
			"table": "teacher",
			"ukey":  "tech_id",
			"uid":   opt.Uid,
		}, &tech)
		if tech.TechId == "" {
			return gtms_error.GetError("account_not_exist")
		}
		if tech.Email == "" {
			return gtms_error.GetError("not_bind_email")
		}
		//创建一个临时的token
		accessToken := helper.CreateToken()
		session := controller.Session{
			AccessToken: accessToken,
			Role:        controller.ROLE_TEACHER,
			TechInfo: controller.TechInfo{
				TechId: tech.TechId,
			},
		}
		s, _ := jsoniter.MarshalToString(session)
		boot.CACHE.Set(accessToken, s, time.Minute*5) //设置5分钟缓存
		body := `<h4>` + tech.TechName + `:<br>
			你好！<br>
			您在本系统中的用户名是` + tech.TechId + `，请点击以下超链接，在浏览器中重置自己的密码(5分钟内有效):<br><br>
			<a href="` + url + `?token=` + accessToken + `">` + url + `?token=` + accessToken + `</a><br><br>
			黄冈师范学院毕业论文管理系统<br><br>
			注:此邮件为本系统所发，非对方邮箱所发，所以请勿回邮。
		</h4>`
		if err := boot.SendEmail(tech.Email, tech.TechName, subject, body); err != nil {
			return gtms_error.GetError("send_email_error")
		}
	case controller.ROLE_STUDENT:
		stu := Student{}
		db.QueryRow(sql, stringi.Form{
			"table": "student",
			"ukey":  "stu_no",
			"uid":   opt.Uid,
		}, &stu)
		if stu.StuNo == "" {
			return gtms_error.GetError("account_not_exist")
		}
		if stu.Email == "" {
			return gtms_error.GetError("not_bind_email")
		}
		//创建一个临时的token
		accessToken := helper.CreateToken()
		session := controller.Session{
			AccessToken: accessToken,
			Role:        controller.ROLE_STUDENT,
			StuInfo: controller.StuInfo{
				StuNo: stu.StuNo,
			},
		}
		s, _ := jsoniter.MarshalToString(session)
		boot.CACHE.Set(accessToken, s, time.Minute*5) //设置5分钟缓存
		body := `<h4>` + stu.StuName + `:<br>
			你好！<br>
			您在本系统中的用户名是` + stu.StuNo + `，请点击以下超链接，在浏览器中重置自己的密码(5分钟内有效):<br><br>
			<a href="` + url + `?token=` + accessToken + `">` + url + `?token=` + accessToken + `</a><br><br>
			黄冈师范学院毕业论文管理系统<br><br>
			注:此邮件为本系统所发，非对方邮箱所发，所以请勿回邮。
		</h4>`
		if err := boot.SendEmail(stu.Email, stu.StuName, subject, body); err != nil {
			return gtms_error.GetError("send_email_error")
		}
	}
	return &validator.Error{}
}

func ResetPwd(opt *forms.ResetPwdForm) *validator.Error {
	user := &controller.Session{}
	str, redisError := boot.CACHE.Get(opt.Token).Result()
	if len(str) > 0 && redisError == nil {
		jsoniter.UnmarshalFromString(str, user)
	} else {
		return gtms_error.GetError("invalid_token")
	}
	sql := `UPDATE @table SET pwd = :newPwd WHERE @ukey = :uid`
	var table string
	var ukey string
	var uid string
	if user.Role == controller.ROLE_ADMIN {
		table = "admin"
		ukey = "admin_id"
		uid = user.AdminId
	} else if user.Role == controller.ROLE_TEACHER {
		table = "teacher"
		ukey = "tech_id"
		uid = user.TechId
	} else {
		table = "student"
		ukey = "stu_no"
		uid = user.StuNo
	}
	newPwd, _ := helper.HashedPassword(opt.NewPwd)
	_, err := db.Exec(sql, stringi.Form{
		"table":  table,
		"newPwd": newPwd,
		"ukey":   ukey,
		"uid":    uid,
	})
	if err != nil {
		gtms_error.GetError("update_info_error")
	} else {
		//重置密码成功后删除临时token
		boot.CACHE.Del(opt.Token).Result()
	}
	return &validator.Error{}
}
