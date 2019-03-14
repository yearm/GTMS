package controller

import (
	"GTMS/library/helper"
	"GTMS/library/validator"
	"github.com/astaxie/beego"
	"net/http"
)

type BaseController struct {
	beego.Controller
	Request
}

func (this *BaseController) JSON(code int, v interface{}) {
	this.Ctx.ResponseWriter.Header().Set("Content-Type", "application/json; charset=UTF-8")
	this.Ctx.ResponseWriter.WriteHeader(code)
	this.Ctx.ResponseWriter.Write(helper.MustMarshal(v))
}

func (this *BaseController) SuccessWithData(data interface{}) {
	this.JSON(http.StatusOK, helper.JSON{
		"status": helper.JSON{
			"code":              0,
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
			"code":              0,
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

//解析参数
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

func (this *BaseController) ErrorResponse(err *validator.Error, datas ...helper.JSON) {
	var data = helper.JSON{}
	if len(datas) > 0 {
		data = datas[0]
	}

	this.JSON(http.StatusOK, helper.JSON{
		"status": helper.JSON{
			"code":             err.Code,
			"message":          err.Msg,
			"time":             helper.Date("Y-m-d H:i:s"),
			"accessTokenState": "keep",
		},
		"data": data,
	})
}
