package forms

//登录Form
type LoginForm struct {
	Account  string `form:"account" valid:"required"`
	Password string `form:"password" valid:"required"`
}

//更新管理员信息Form
type UpdateAdminForm struct {
	AdminId   string `form:"adminId" valid:"required"`
	Pwd       string `form:"pwd"`
	AdminName string `form:"adminName"`
	AdminSex  string `form:"AdminSex"`
	Phone     string `form:"phone"`
	Email     string `form:"email"`
}

//更新学生信息Form
type UpdateStudentForm struct {
	StuNo        string `form:"stuNo" valid:"required"`
	Pwd          string `form:"pwd"`
	StuName      string `form:"stuName"`
	StuSex       string `form:"stuSex"`
	IdCard       string `form:"idCard"`
	Birthplace   string `form:"birthplace"`
	Department   string `form:"department"`
	Major        string `form:"major"`
	Class        string `form:"class"`
	Phone        string `form:"phone"`
	Qq           string `form:"qq"`
	Email        string `form:"email"`
	WeChat       string `form:"weChat"`
	SchoolSystem string `form:"schoolSystem"`
	EntryDate    string `form:"entryDate"`
	Education    string `form:"education"`
}

//更新教师信息Form
type UpdateTeacherForm struct {
	TechId            string `form:"techId" valid:"required"`
	Pwd               string `form:"pwd"`
	TechName          string `form:"techName"`
	TechSex           string `form:"techSex"`
	Education         string `form:"education"`
	Degree            string `form:"degree"`
	ResearchDirection string `form:"researchDirection"`
	JobTitle          string `form:"jobTitle"`
	Job               string `form:"job"`
	InstructNums      string `form:"instructNums"`
	InstructMajor     string `form:"instructMajor"`
	Email             string `form:"email"`
	Phone             string `form:"phone"`
	Qq                string `form:"qq"`
	WeChat            string `form:"weChat"`
}

//发送重置密码邮件Form
type SendEmailToResetPwd struct {
	Role string `form:"role" valid:"required|switch:admin,teacher,student"`
	Uid  string `form:"uid" valid:"required"`
}

//重置密码Form
type ResetPwdForm struct {
	Token  string `form:"token" valid:"required"`
	NewPwd string `form:"newPwd" valid:"required"`
}
