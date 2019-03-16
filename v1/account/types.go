package account

type SignInForm struct {
	Account  string `form:"account" valid:"required"`
	Password string `form:"password" valid:"required"`
}
