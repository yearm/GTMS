package teacher_account

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

func Login(opt *account.LoginForm) (*controller.Session, *validator.Error) {
	o := boot.GetSlaveMySQL()
	tech := Teacher{TechId: opt.Account}
	o.Read(&tech)
	if helper.CheckHashedPassword(tech.Pwd, opt.Password) {
		accessToken := helper.CreateToken()
		user := controller.Session{
			AccessToken: accessToken,
			IsGuest:     false,
			Role:        controller.ROLE_TEACHER,
			ErrorKey:    "",
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
		}
		s, _ := jsoniter.MarshalToString(user)
		boot.CACHE.Set(accessToken, s, time.Hour*24*30)
		go func() {
			//删除旧token
			sql := `SELECT token FROM user_session WHERE uid = :uid`
			var token string
			db.QueryRow(sql, stringi.Form{
				"uid": tech.TechId,
			}, &token)
			boot.CACHE.Del(token).Result()
			//更新user_session表
			db.Exec(db.ReplaceSQL("user_session", stringi.Form{
				"uid":         tech.TechId,
				"token":       accessToken,
				"role":        "admin",
				"update_time": helper.Date("Y-m-d H:i:s"),
			}))
		}()
		return &user, &validator.Error{}
	} else {
		return nil, gtms_error.GetError("sign_in_error")
	}
}
