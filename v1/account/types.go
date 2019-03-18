package account

type LoginForm struct {
	Account  string `form:"account" valid:"required"`
	Password string `form:"password" valid:"required"`
}
