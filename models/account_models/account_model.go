package account_models

import (
	"GTMS/boot"
	"GTMS/library/controller"
	"GTMS/library/db"
	"GTMS/library/stringi"
	"GTMS/library/validator"
)

func AccountLogout(req *controller.Request) *validator.Error {
	var token string
	var uid string
	sql := `DELETE FROM user_session WHERE uid = :uid`
	if req.User.Role == controller.ROLE_ADMIN {
		token = req.User.AccessToken
		uid = req.User.AdminId
	} else if req.User.Role == controller.ROLE_TEACHER {
		token = req.User.AccessToken
		uid = req.User.TechId
	} else if req.User.Role == controller.ROLE_STUDENT {
		token = req.User.AccessToken
		uid = req.User.StuNo
	}
	boot.CACHE.Del(token).Result()
	db.Exec(sql, stringi.Form{
		"uid": uid,
	})
	return &validator.Error{}
}
