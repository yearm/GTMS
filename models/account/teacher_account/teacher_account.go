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
	TechId            string `orm:"pk" json:"tech_id"`
	Pwd               string `json:"-"`
	TechName          string `json:"techName"`
	TechSex           string `json:"techSex"`
	Education         string `json:"education"`
	Degree            string `json:"degree"`
	ResearchDirection string `json:"researchDirection"`
	JobTitle          string `json:"jobTitle"`
	Job               string `json:"job"`
	InstructNums      string `json:"instructNums"`
	InstructMajor     string `json:"instructMajor"`
	Email             string `json:"email"`
	Phone             string `json:"phone"`
	Qq                string `json:"qq"`
	WeChat            string `json:"weChat"`
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
				QQ:                tech.Qq,
				WeChat:            tech.WeChat,
			},
		}
		s, _ := jsoniter.MarshalToString(user)
		boot.CACHE.Set(accessToken, s, time.Hour*24*30)
		go func() {
			controller.DelRedisToken(tech.TechId)
			//更新user_session表
			db.Exec(db.ReplaceSQL("user_session", stringi.Form{
				"uid":         tech.TechId,
				"token":       accessToken,
				"role":        controller.ROLE_TEACHER,
				"update_time": helper.Date("Y-m-d H:i:s"),
			}))
		}()
		return &user, &validator.Error{}
	} else {
		return nil, gtms_error.GetError("sign_in_error")
	}
}

func TechList(page int, pageCount int) (techs []*Teacher, total int) {
	o := boot.GetSlaveMySQL()
	qs := o.QueryTable((*Teacher)(nil))
	_, err := qs.Limit(pageCount, (page-1)*pageCount).All(&techs)
	if err != nil {
		return
	}
	t, _ := qs.Count()
	total = int(t)
	return
}
