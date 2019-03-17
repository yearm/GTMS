package admin_account

import (
	"GTMS/boot"
	"GTMS/library/controller"
	"GTMS/library/gtms_error"
	"GTMS/library/helper"
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

func SignIn(opt *account.SignInForm) (*controller.Session, *validator.Error) {
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
		s, _ := jsoniter.MarshalToString(adminInfo)
		boot.CACHE.Set(accessToken, s, time.Hour*24*30)
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
