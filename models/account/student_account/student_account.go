package student_account

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

type Student struct {
	StuId        string `orm:"pk"`
	Pwd          string
	StuNo        string
	StuName      string
	StuSex       string
	IdCard       string
	Birthplace   string
	Department   string
	Major        string
	Class        string
	Phone        string
	QQ           string
	Email        string
	WeChat       string
	SchoolSystem string
	EntryDate    string
	Education    string
}

func init() {
	//需要在init中注册定义的model
	orm.RegisterModel(new(Student))
}

func SignIn(opt *account.SignInForm) (*controller.Session, *validator.Error) {
	o := boot.GetSlaveMySQL()
	stu := Student{StuId: opt.Account}
	o.Read(&stu)
	if helper.CheckHashedPassword(stu.Pwd, opt.Password) {
		accessToken := helper.CreateToken()
		return &controller.Session{
			AccessToken: accessToken,
			IsGuest:     false,
			Role:        "student",
			UpdateTime:  time.Now().Unix(),
			StuInfo: controller.StuInfo{
				StuId:        stu.StuId,
				StuNo:        stu.StuNo,
				StuName:      stu.StuName,
				StuSex:       stu.StuSex,
				IdCard:       stu.IdCard,
				Birthplace:   stu.Birthplace,
				Department:   stu.Department,
				Major:        stu.Major,
				Class:        stu.Class,
				Phone:        stu.Phone,
				QQ:           stu.QQ,
				Email:        stu.Email,
				WeChat:       stu.WeChat,
				SchoolSystem: stu.SchoolSystem,
				EntryDate:    stu.EntryDate,
				Education:    stu.Education,
			},
		}, &validator.Error{}
	} else {
		return nil, gtms_error.GetError("sign_in_error")
	}
}
