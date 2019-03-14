package controller

import (
	"GTMS/library/helper"
	"github.com/astaxie/beego"
	"github.com/json-iterator/go"
	"net/http"
)

type BaseController struct {
	beego.Controller
}

func (this *BaseController) JSON(v interface{}) {
	this.Ctx.ResponseWriter.Header().Set("Content-Type", "application/json; charset=UTF-8")
	this.Ctx.ResponseWriter.WriteHeader(http.StatusOK)
	data, _ := jsoniter.Marshal(v)
	this.Ctx.ResponseWriter.Write(data)
}

func (this *BaseController) SuccessWithData(data interface{}) {
	this.JSON(helper.JSON{
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
	this.JSON(helper.JSON{
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
