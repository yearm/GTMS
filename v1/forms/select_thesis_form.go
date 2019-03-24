package forms

//学生选题Form
type SelectThesisForm struct {
	Tid string `form:"tid" valid:"required"`
}

//教师确认选题Form
type ConfirmSelectedlThesisForm struct {
	Sid     string `form:"sid" valid:"required"`
	Confirm string `form:"confirm" valid:"required|switch:1,2"`
}
