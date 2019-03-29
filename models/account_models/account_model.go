package account_models

import (
	"GTMS/boot"
	"GTMS/library/controller"
	"GTMS/library/db"
	"GTMS/library/gtms_error"
	"GTMS/library/helper"
	"GTMS/library/stringi"
	"GTMS/library/validator"
	"GTMS/v1/forms"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"github.com/json-iterator/go"
	"time"
)

type Admin struct {
	AdminId   string `orm:"pk" json:"adminId"`
	Pwd       string `json:"-"`
	AdminName string `json:"adminName"`
	AdminSex  string `json:"adminSex"`
	Phone     string `json:"phone"`
	Email     string `json:"email"`
}

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
	orm.RegisterModel(new(Admin))
	orm.RegisterModel(new(Teacher))
	orm.RegisterModel(new(Student))
}

func AccountLogin(opt *forms.LoginForm) (*controller.Session, *validator.Error) {
	o := boot.GetSlaveMySQL()
	accessToken := helper.CreateToken()
	var redisStr string
	var uid string
	var role string
	session := controller.Session{}
	switch opt.Role {
	case controller.ROLE_ADMIN:
		admin := Admin{AdminId: opt.Account}
		o.Read(&admin)
		if helper.CheckHashedPassword(admin.Pwd, opt.Password) {
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
					Phone:     admin.Phone,
					Email:     admin.Email,
				},
			}
			redisStr, _ = jsoniter.MarshalToString(user)
			uid = admin.AdminId
			role = controller.ROLE_ADMIN
		} else {
			return nil, gtms_error.GetError("sign_in_error")
		}
	case controller.ROLE_TEACHER:
		tech := Teacher{TechId: opt.Account}
		o.Read(&tech)
		if helper.CheckHashedPassword(tech.Pwd, opt.Password) {
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
			redisStr, _ = jsoniter.MarshalToString(user)
			uid = tech.TechId
			role = controller.ROLE_TEACHER
		} else {
			return nil, gtms_error.GetError("sign_in_error")
		}
	case controller.ROLE_STUDENT:
		stu := Student{StuNo: opt.Account}
		o.Read(&stu)
		if helper.CheckHashedPassword(stu.Pwd, opt.Password) {
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
			redisStr, _ = jsoniter.MarshalToString(user)
			uid = stu.StuNo
			role = controller.ROLE_STUDENT
		} else {
			return nil, gtms_error.GetError("sign_in_error")
		}
	}
	//设置新token，一个月缓存
	boot.CACHE.Set(accessToken, redisStr, time.Hour*24*30)
	go func() {
		//删除旧token
		controller.DelRedisToken(uid)
		//更新user_session表
		db.Exec(db.ReplaceSQL("user_session", stringi.Form{
			"uid":         uid,
			"token":       accessToken,
			"role":        role,
			"update_time": helper.Date("Y-m-d H:i:s"),
		}))
	}()
	jsoniter.UnmarshalFromString(redisStr, &session)
	return &session, &validator.Error{}
}

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

func AccountList(role string, page int, pageCount int) (admins []*Admin, techs []*Teacher, stus []*Student, total int) {
	o := boot.GetSlaveMySQL()
	switch role {
	case controller.ROLE_ADMIN:
		qs := o.QueryTable((*Admin)(nil))
		_, err := qs.Limit(pageCount, (page-1)*pageCount).All(&admins)
		if err != nil {
			return
		}
		t, _ := qs.Count()
		total = int(t)
		return
	case controller.ROLE_TEACHER:
		qs := o.QueryTable((*Teacher)(nil))
		_, err := qs.Limit(pageCount, (page-1)*pageCount).All(&techs)
		if err != nil {
			return
		}
		t, _ := qs.Count()
		total = int(t)
		return
	case controller.ROLE_STUDENT:
		qs := o.QueryTable((*Student)(nil))
		_, err := qs.Limit(pageCount, (page-1)*pageCount).All(&stus)
		if err != nil {
			return
		}
		t, _ := qs.Count()
		total = int(t)
		return
	}
	return
}

func UpdateAdmin(opt *forms.UpdateAdminForm) *validator.Error {
	sql := `UPDATE @table SET @value WHERE admin_id = :admin_id`
	form := helper.StructToFormWithClearNilField(*opt, controller.FormatAdmin)
	if form["pwd"] != "" {
		form["pwd"], _ = helper.HashedPassword(opt.Pwd)
		//更新密码会删除token
		controller.DelRedisToken(opt.AdminId)
	}
	values := db.Set(form)
	_, err := db.Exec(sql, stringi.Form{
		"table":    "admin",
		"value":    values,
		"admin_id": opt.AdminId,
	})
	if err != nil {
		logs.Error(err)
		return gtms_error.GetError("update_info_error")
	}
	return &validator.Error{}
}

func UpdateStudent(opt *forms.UpdateStudentForm) *validator.Error {
	sql := `UPDATE @table SET @value WHERE stu_no = :stu_no`
	form := helper.StructToFormWithClearNilField(*opt, controller.FormatStudent)
	if form["pwd"] != "" {
		form["pwd"], _ = helper.HashedPassword(opt.Pwd)
		//更新密码会删除token
		controller.DelRedisToken(opt.StuNo)
	}
	values := db.Set(form)
	_, err := db.Exec(sql, stringi.Form{
		"table":  "student",
		"value":  values,
		"stu_no": opt.StuNo,
	})
	if err != nil {
		logs.Error(err)
		return gtms_error.GetError("update_info_error")
	}
	return &validator.Error{}
}

func UpdateTeacher(opt *forms.UpdateTeacherForm) *validator.Error {
	sql := `UPDATE @table SET @value WHERE tech_id = :tech_id`
	form := helper.StructToFormWithClearNilField(*opt, controller.FormatTeacher)
	if form["pwd"] != "" {
		form["pwd"], _ = helper.HashedPassword(opt.Pwd)
		//更新密码会删除token
		controller.DelRedisToken(opt.TechId)
	}
	values := db.Set(form)
	_, err := db.Exec(sql, stringi.Form{
		"table":   "teacher",
		"value":   values,
		"tech_id": opt.TechId,
	})
	if err != nil {
		logs.Error(err)
		return gtms_error.GetError("update_info_error")
	}
	return &validator.Error{}
}
