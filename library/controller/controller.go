package controller

import (
	"GTMS/boot"
	"GTMS/library/db"
	"GTMS/library/gtms_error"
	"GTMS/library/helper"
	"GTMS/library/stringi"
	"GTMS/library/validator"
	"github.com/astaxie/beego"
	"github.com/json-iterator/go"
	"mime/multipart"
	"net/http"
	"os"
	"strconv"
)

const (
	default_page      = "1"
	default_pageCount = "30"
	upload_folder     = "/upload/"
	Opening_report    = "openingReport" //开题报告
	Thesis            = "thesis"        // 毕业论文
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

//错误Response
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

//重新登录
func (this *BaseController) RequireLogin() {
	user := this.Request.User
	errKey := user.ErrorKey
	if stringi.IsEmpty(errKey) {
		errKey = "required_login"
	}
	errorMsg := gtms_error.GetError(errKey)
	this.JSON(http.StatusOK, helper.JSON{
		"status": helper.JSON{
			"code":             errorMsg.Code,
			"message":          errorMsg.Msg,
			"time":             helper.Date("Y-m-d H:i:s"),
			"accessTokenState": "refresh",
		},
	})
}

//删除redisToken
func DelRedisToken(uid string) {
	sql := `SELECT token FROM user_session WHERE uid = :uid`
	var token string
	db.QueryRow(sql, stringi.Form{
		"uid": uid,
	}, &token)
	boot.CACHE.Del(token).Result()
}

//根据token获取用户
func (this *BaseController) GetUser(accessToken string) *Session {
	var user = &Session{IsGuest: true}
	if stringi.IsEmpty(accessToken) {
		return user
	}
	str, redisError := boot.CACHE.Get(accessToken).Result()
	if len(str) > 0 && redisError == nil {
		jsoniter.UnmarshalFromString(str, user)
	}
	return user
}

//获取分页参数
func (this *BaseController) GetPageInfo() (page int, pageCount int) {
	pageStr := this.GetString("page")
	pageCountStr := this.GetString("pageCount")
	pageStr = stringi.DefaultValue(pageStr, default_page)
	pageCountStr = stringi.DefaultValue(pageCountStr, default_pageCount)
	page, err1 := strconv.Atoi(pageStr)
	pageCount, err2 := strconv.Atoi(pageCountStr)
	if err1 != nil || err2 != nil || page < 0 || pageCount < 0 {
		return stringi.ToInt(default_page), stringi.ToInt(default_pageCount)
	}
	return
}

//获取上传文件
func (this *BaseController) GetFile(k string) (*multipart.FileHeader, bool) {
	files, exist := this.Ctx.Request.MultipartForm.File[k]
	if !exist || len(files) == 0 {
		return nil, false
	}
	return files[0], true
}

//保存上传文件
func (this *BaseController) SaveFile(fileType string, fromFile string, fileName string) (err error) {
	openingReportPath := helper.GetRootPath() + upload_folder + Opening_report
	thesisPath := helper.GetRootPath() + upload_folder + Thesis
	if fileType == Opening_report {
		if b, _ := helper.FolderExists(openingReportPath); !b {
			os.Mkdir(openingReportPath, os.ModePerm)
		}
		err = this.SaveToFile(fromFile, openingReportPath+"/"+fileName)
	} else {
		if b, _ := helper.FolderExists(thesisPath); !b {
			os.Mkdir(thesisPath, os.ModePerm)
		}
		err = this.SaveToFile(fromFile, thesisPath+"/"+fileName)
	}
	return
}
