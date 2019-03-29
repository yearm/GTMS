package forms

//添加管理员form
type AddAdminAccountForm struct {
	AdminId   string `form:"adminId" valid:"required"`
	AdminName string `form:"adminName"`
	AdminSex  string `form:"adminSex"`
	Phone     string `form:"phone"`
	Email     string `form:"email"`
}

//添加教师form
type AddTechAccountForm struct {
	TechId            string `form:"techId" valid:"required"`
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

//添加学生form
type AddStuAccountForm struct {
	StuNo        string `form:"stuNo" valid:"required"`
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

//删除账号form
type DelAccountForm struct {
	Uid string `form:"uid" valid:"required"`
}
