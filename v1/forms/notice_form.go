package forms

//添加通知Form
type AddNoticeForm struct {
	Title   string `form:"title"`
	Content string `form:"content"`
}
