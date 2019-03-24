package forms

//修改开放时间Form
type UpdateOpeningTimeForm struct {
	StartTime string `form:"startTime"`
	EndTime   string `form:"endTime"`
}
