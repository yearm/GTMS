package forms

//添加账号form
type AddAccountForm struct {
	Uids  []string `form:"uids" minSize:"1"`
	Names []string `form:"names" minSize:"1"`
	Role  string   `form:"role" valid:"required|switch:admin,teacher,student"`
}

//删除账号form
type DelAccountForm struct {
	Uids []string `form:"uids" minSize:"1"`
	Role string   `form:"role" valid:"required|switch:admin,teacher,student"`
}
