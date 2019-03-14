package controller

import (
	"GTMS/library/helper"
	"GTMS/library/validator"
	"github.com/astaxie/beego"
	"net/http"
)

type BaseController struct {
	beego.Controller
}

func (this *BaseController) JSON(code int, v interface{}) {
	this.Ctx.ResponseWriter.Header().Set("Content-Type", "application/json; charset=UTF-8")
	this.Ctx.ResponseWriter.WriteHeader(code)
	this.Ctx.ResponseWriter.Write(helper.MustMarshal(v))
}

func (this *BaseController) SuccessWithData(data interface{}) {
	this.JSON(http.StatusOK, helper.JSON{
		"status": helper.JSON{
			"code":              "0",
			"message":           "success",
			"time":              helper.Date("Y-m-d H:i:s"),
			"accessTokenStatus": "keep",
		},
		"data": data,
	})
}
func (this *BaseController) SuccessWithDataList(datalist interface{}, pageInfo interface{}) {
	this.JSON(http.StatusOK, helper.JSON{
		"status": helper.JSON{
			"code":              "0",
			"message":           "success",
			"time":              helper.Date("Y-m-d H:i:s"),
			"accessTokenStatus": "keep",
		},
		"data": helper.JSON{
			"dataList": datalist,
			"pageInfo": pageInfo,
		},
	})
}

func (this *BaseController) Getstring(k string) string {
	query := this.Ctx.Request.URL.Query()
	if v := query.Get(k); v != "" {
		return v
	}
	return this.Ctx.Request.PostForm.Get(k)
}

func (this *BaseController) ParseInput(obj interface{}) *validator.Error {
	err := beego.ParseForm(this.Input(), obj)
	if err != nil {
		return &validator.Error{
			Code: 1,
			Msg:  err.Error(),
		}
	}
	return validator.Check(obj)
}
