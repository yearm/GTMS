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
	AdminId   string `orm:"pk" json:"adminId"`
	Pwd       string `json:"-"`
	AdminName string `json:"adminName"`
	AdminSex  string `json:"adminSex"`
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
			controller.DelRedisToken(admin.AdminId)
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

func AdminList(page int, pageCount int) (admins []*Admin, total int) {
	o := boot.GetSlaveMySQL()
	qs := o.QueryTable((*Admin)(nil))
	_, err := qs.Limit(pageCount, (page-1)*pageCount).All(&admins)
	if err != nil {
		return
	}
	t, _ := qs.Count()
	total = int(t)
	return
}
