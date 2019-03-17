package teacher_account

import (
	"GTMS/boot"
	"GTMS/library/controller"
	"GTMS/library/gtms_error"
	"GTMS/library/helper"
	"GTMS/library/validator"
	"GTMS/v1/account"
	"github.com/astaxie/beego/orm"
	"time"
)

type Teacher struct {
	TechId            string `orm:"pk"`
	Pwd               string
	TechName          string
	TechSex           string
	Education         string
	Degree            string
	ResearchDirection string
	JobTitle          string
	Job               string
	InstructNums      string
	InstructMajor     string
	Email             string
	Phone             string
	QQ                string
	WeChat            string
}

func init() {
	//需要在init中注册定义的model
	orm.RegisterModel(new(Teacher))
}

func SignIn(opt *account.SignInForm) (*controller.Session, *validator.Error) {
	o := boot.GetSlaveMySQL()
	tech := Teacher{TechId: opt.Account}
	o.Read(&tech)
	if helper.CheckHashedPassword(tech.Pwd, opt.Password) {
		accessToken := helper.CreateToken()
		return &controller.Session{
			AccessToken: accessToken,
			IsGuest:     false,
			Role:        "teacher",
			UpdateTime:  time.Now().Unix(),
			TechInfo: controller.TechInfo{
				TechId:            tech.TechId,
				TechName:          tech.TechName,
				TechSex:           tech.TechSex,
				Education:         tech.Education,
				Degree:            tech.Degree,
				ResearchDirection: tech.ResearchDirection,
				JobTitle:          tech.JobTitle,
				Job:               tech.Job,
				InstructNums:      tech.InstructNums,
				InstructMajor:     tech.InstructMajor,
				Email:             tech.Email,
				Phone:             tech.Phone,
				QQ:                tech.QQ,
				WeChat:            tech.WeChat,
			},
		}, &validator.Error{}
	} else {
		return nil, gtms_error.GetError("sign_in_error")
	}
}
