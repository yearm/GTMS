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

func Login(opt *account.SignInForm) (*controller.Session, *validator.Error) {
	o := boot.GetSlaveMySQL()
	admin := Admin{AdminId: opt.Account}
	o.Read(&admin)
	if helper.CheckHashedPassword(admin.Pwd, opt.Password) {
		accessToken := helper.CreateToken()
		adminInfo := controller.AdminInfo{
			AdminId:   admin.AdminId,
			AdminName: admin.AdminName,
			AdminSex:  admin.AdminSex,
		}
		go func() {
			//开协程写redis、写user_session表
			s, _ := jsoniter.MarshalToString(adminInfo)
			boot.CACHE.Set(accessToken, s, time.Hour*24*30)
			db.Exec(db.ReplaceSQL("user_session", stringi.Form{
				"uid":         admin.AdminId,
				"token":       accessToken,
				"role":        "admin",
				"update_time": helper.Date("Y-m-d H:i:s"),
			}))
		}()
		return &controller.Session{
			AccessToken: accessToken,
			IsGuest:     false,
			Role:        "admin",
			UpdateTime:  time.Now().Unix(),
			AdminInfo:   adminInfo,
		}, &validator.Error{}
	} else {
		return nil, gtms_error.GetError("sign_in_error")
	}
}
