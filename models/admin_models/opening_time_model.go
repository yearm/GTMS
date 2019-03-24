package admin_models

import (
	"GTMS/boot"
	"GTMS/library/controller"
	"GTMS/library/db"
	"GTMS/library/gtms_error"
	"GTMS/library/helper"
	"GTMS/library/stringi"
	"GTMS/library/validator"
	"GTMS/v1/forms"
	"github.com/astaxie/beego/orm"
)

type OpeningTime struct {
	Year        int    `orm:"pk" json:"year"`
	StartTime   string `json:"startTime"`
	EndTime     string `json:"endTime"`
	OperateUid  string `json:"operateUid"`
	OperateName string `json:"operateName"`
	OperateTime string `json:"operateTime"`
}

func init() {
	//需要在init中注册定义的model
	orm.RegisterModel(new(OpeningTime))
}

func GetOpeningTime() *OpeningTime {
	o := boot.GetSlaveMySQL()
	openingTime := OpeningTime{Year: stringi.ToInt(helper.Date("Y"))}
	o.Read(&openingTime)
	return &openingTime
}

func UpdateOpeningTime(opt *forms.UpdateOpeningTimeForm, req *controller.Request) *validator.Error {
	_, err := db.Exec(db.ReplaceSQL("opening_time", stringi.Form{
		"year":         helper.Date("Y"),
		"start_time":   opt.StartTime,
		"end_time":     opt.EndTime,
		"operate_uid":  req.User.AdminId,
		"operate_name": req.User.AdminName,
		"operate_time": helper.Date("Y-m-d H:i:s"),
	}))
	if err != nil {
		return gtms_error.GetError("update_info_error")
	}
	return &validator.Error{}
}
