package student_account

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

type Student struct {
	StuNo        string `orm:"pk" json:"stuNo"`
	Pwd          string `json:"-"`
	StuName      string `json:"stuName"`
	StuSex       string `json:"stuSex"`
	IdCard       string `json:"idCard"`
	Birthplace   string `json:"birthplace"`
	Department   string `json:"department"`
	Major        string `json:"major"`
	Class        string `json:"class"`
	Phone        string `json:"phone"`
	Qq           string `json:"qq"`
	Email        string `json:"email"`
	WeChat       string `json:"weChat"`
	SchoolSystem string `json:"schoolSystem"`
	EntryDate    string `json:"entryDate"`
	Education    string `json:"education"`
}

func init() {
	//需要在init中注册定义的model
	orm.RegisterModel(new(Student))
}

func Login(opt *account.LoginForm) (*controller.Session, *validator.Error) {
	o := boot.GetSlaveMySQL()
	stu := Student{StuNo: opt.Account}
	o.Read(&stu)
	if helper.CheckHashedPassword(stu.Pwd, opt.Password) {
		accessToken := helper.CreateToken()
		user := controller.Session{
			AccessToken: accessToken,
			IsGuest:     false,
			Role:        controller.ROLE_STUDENT,
			ErrorKey:    "",
			UpdateTime:  time.Now().Unix(),
			StuInfo: controller.StuInfo{
				StuNo:        stu.StuNo,
				StuName:      stu.StuName,
				StuSex:       stu.StuSex,
				IdCard:       stu.IdCard,
				Birthplace:   stu.Birthplace,
				Department:   stu.Department,
				Major:        stu.Major,
				Class:        stu.Class,
				Phone:        stu.Phone,
				QQ:           stu.Qq,
				Email:        stu.Email,
				WeChat:       stu.WeChat,
				SchoolSystem: stu.SchoolSystem,
				EntryDate:    stu.EntryDate,
				Education:    stu.Education,
			},
		}
		s, _ := jsoniter.MarshalToString(user)
		boot.CACHE.Set(accessToken, s, time.Hour*24*30)
		go func() {
			controller.DelRedisToken(stu.StuNo)
			//更新user_session表
			db.Exec(db.ReplaceSQL("user_session", stringi.Form{
				"uid":         stu.StuNo,
				"token":       accessToken,
				"role":        controller.ROLE_STUDENT,
				"update_time": helper.Date("Y-m-d H:i:s"),
			}))
		}()
		return &user, &validator.Error{}
	} else {
		return nil, gtms_error.GetError("sign_in_error")
	}
}

func StuList(page int, pageCount int) (stus []*Student, total int) {
	o := boot.GetSlaveMySQL()
	qs := o.QueryTable((*Student)(nil))
	_, err := qs.Limit(pageCount, (page-1)*pageCount).All(&stus)
	if err != nil {
		return
	}
	t, _ := qs.Count()
	total = int(t)
	return
}

func UpdateStudent(opt *account.UpdateStudentForm) *validator.Error {
	sql := `UPDATE @table SET @value WHERE stu_no = :stu_no`
	values := db.Set(helper.StructToFormWithClearNilField(*opt))
	_, err := db.Exec(sql, stringi.Form{
		"table":  "student",
		"value":  values,
		"stu_no": opt.StuNo,
	})
	if err != nil {
		return gtms_error.GetError("update_info_error")
	}
	return &validator.Error{}
}
