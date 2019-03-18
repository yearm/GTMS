package account

//登录Form
type LoginForm struct {
	Account  string `form:"account" valid:"required"`
	Password string `form:"password" valid:"required"`
}

//更新管理员Form
type UpdateAdminForm struct {
	AdminId   string `form:"adminId"`
	Pwd       string `form:"pwd"`
	AdminName string `form:"adminName"`
	AdminSex  string `form:"AdminSex"`
}

//更新学生Form
type UpdateStudentForm struct {
	StuNo        string `form:"stuNo"`
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

//更新教师Form
type UpdateTeacherForm struct {
	TechId            string `form:"techId"`
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
