package admin_account

import (
	"GTMS/boot"
	"GTMS/library/controller"
	"GTMS/library/db"
	"GTMS/library/gtms_error"
	"GTMS/library/helper"
	"GTMS/library/stringi"
	"GTMS/library/validator"
	"GTMS/v1/account"
	"github.com/astaxie/beego/orm"
	"github.com/json-iterator/go"
	"time"
)

type Admin struct {
	AdminId   string `orm:"pk"`
	Pwd       string
	AdminName string
	AdminSex  string
}

func init() {
	//需要在init中注册定义的model
	orm.RegisterModel(new(Admin))
}

func Login(opt *account.LoginForm) (*controller.Session, *validator.Error) {
	o := boot.GetSlaveMySQL()
	admin := Admin{AdminId: opt.Account}
	o.Read(&admin)
	if helper.CheckHashedPassword(admin.Pwd, opt.Password) {
		accessToken := helper.CreateToken()
		user := controller.Session{
			AccessToken: accessToken,
			IsGuest:     false,
			Role:        controller.ROLE_ADMIN,
			ErrorKey:    "",
			UpdateTime:  time.Now().Unix(),
			AdminInfo: controller.AdminInfo{
				AdminId:   admin.AdminId,
				AdminName: admin.AdminName,
				AdminSex:  admin.AdminSex,
			},
		}
		s, _ := jsoniter.MarshalToString(user)
		boot.CACHE.Set(accessToken, s, time.Hour*24*30)
		go func() {
			//删除旧token
			sql := `SELECT token FROM user_session WHERE uid = :uid`
			var token string
			db.QueryRow(sql, stringi.Form{
				"uid": admin.AdminId,
			}, &token)
			boot.CACHE.Del(token).Result()
			//更新user_session表
			db.Exec(db.ReplaceSQL("user_session", stringi.Form{
				"uid":         admin.AdminId,
				"token":       accessToken,
				"role":        controller.ROLE_ADMIN,
				"update_time": helper.Date("Y-m-d H:i:s"),
			}))
		}()
		return &user, &validator.Error{}
	} else {
		return nil, gtms_error.GetError("sign_in_error")
	}
}
