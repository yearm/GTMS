package controller

import (
	"github.com/astaxie/beego/context"
)

const (
	ROLE_ADMIN   = "admin"
	ROLE_TEACHER = "teacher"
	ROLE_STUDENT = "student"
)

type Request struct {
	User    *Session
	Context *context.Context
}

type Session struct {
	AccessToken string `json:"accessToken"`
	IsGuest     bool   `json:"isGuest"`
	Role        string `json:"role"`
	ErrorKey    string `json:"errorKey"`
	UpdateTime  int64  `json:"updateTime"`
	AdminInfo   `json:"adminInfo"`
	TechInfo    `json:"techInfo"`
	StuInfo     `json:"stuInfo"`
}

type AdminInfo struct {
	AdminId   string `json:"adminId"`
	AdminName string `json:"adminName"`
	AdminSex  string `json:"adminSex"`
}

type TechInfo struct {
	TechId            string `json:"techId"`            //教师id
	TechName          string `json:"techName"`          //姓名
	TechSex           string `json:"techSex"`           //性别
	Education         string `json:"education"`         //学历
	Degree            string `json:"degree"`            //学位
	ResearchDirection string `json:"researchDirection"` //研究方向
	JobTitle          string `json:"jobTitle"`          //职称
	Job               string `json:"job"`               //职务
	InstructNums      string `json:"instructNums"`      //指导人数
	InstructMajor     string `json:"instructMajor"`     //指导专业
	Email             string `json:"email"`             //邮箱
	Phone             string `json:"phone"`             //手机
	QQ                string `json:"qq"`                //qq
	WeChat            string `json:"weChat"`            //微信
}

type StuInfo struct {
	StuNo        string `json:"stuNo"`        //学号
	StuName      string `json:"stuName"`      //学生姓名
	StuSex       string `json:"stuSex"`       //性别
	IdCard       string `json:"idCard"`       //身份证
	Birthplace   string `json:"birthplace"`   //籍贯
	Department   string `json:"department"`   //院系
	Major        string `json:"major"`        //专业
	Class        string `json:"class"`        //班级
	Phone        string `json:"phone"`        //手机号
	QQ           string `json:"qq"`           //QQ号
	Email        string `json:"email"`        //邮箱
	WeChat       string `json:"weChat"`       //微信
	SchoolSystem string `json:"schoolSystem"` //学制
	EntryDate    string `json:"entryDate"`    //入学日期
	Education    string `json:"education"`    //学历
}
