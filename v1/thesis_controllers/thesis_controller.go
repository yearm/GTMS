package thesis_controllers

import (
	"GTMS/conf"
	"GTMS/library/controller"
	"GTMS/library/gtms_error"
	"GTMS/library/helper"
	"GTMS/library/stringi"
	"GTMS/models/thesis_models"
	"GTMS/v1/forms"
	"os/exec"
	"strings"
)

type ThesisController struct {
	controller.BaseController
}

//添加论文
func (this *ThesisController) AddThesis() {
	this.User = this.GetUser(this.Ctx.Request.Header.Get("X-Access-Token"))
	if this.User.IsGuest {
		this.RequireLogin()
		return
	}
	//只允许教师添加
	if this.User.Role != controller.ROLE_TEACHER {
		this.ErrorResponse(gtms_error.GetError("access_denied"))
		return
	}
	inputs := forms.AddThesisForm{}
	if err := this.ParseInput(&inputs); err.Code != 0 {
		this.ErrorResponse(err)
		return
	}
	err := thesis_models.AddThesis(&this.Request, &inputs)
	if err.Code != 0 {
		this.ErrorResponse(err)
		return
	}
	this.SuccessWithData(helper.JSON{})
}

//删除论文
func (this *ThesisController) DelThesis() {
	tid := this.GetString("tid")
	this.User = this.GetUser(this.Ctx.Request.Header.Get("X-Access-Token"))
	if this.User.IsGuest {
		this.RequireLogin()
		return
	}
	//只允许教师删除
	if this.User.Role != controller.ROLE_TEACHER {
		this.ErrorResponse(gtms_error.GetError("access_denied"))
		return
	}
	err := thesis_models.DelThesis(tid)
	if err.Code != 0 {
		this.ErrorResponse(err)
		return
	}
	this.SuccessWithData(helper.JSON{})
}

//修改论文
func (this *ThesisController) UpdateThesis() {
	this.User = this.GetUser(this.Ctx.Request.Header.Get("X-Access-Token"))
	if this.User.IsGuest {
		this.RequireLogin()
		return
	}
	//只允许教师修改
	if this.User.Role != controller.ROLE_TEACHER {
		this.ErrorResponse(gtms_error.GetError("access_denied"))
		return
	}
	inputs := forms.UpdateThesisForm{}
	if err := this.ParseInput(&inputs); err.Code != 0 {
		this.ErrorResponse(err)
		return
	}
	err := thesis_models.UpdateThesis(&inputs, &this.Request)
	if err.Code != 0 {
		this.ErrorResponse(err)
		return
	}
	this.SuccessWithData(helper.JSON{})
}

//获取论文列表
func (this *ThesisController) ThesisList() {
	this.User = this.GetUser(this.Ctx.Request.Header.Get("X-Access-Token"))
	if this.User.IsGuest {
		this.RequireLogin()
		return
	}
	page, pageCount := this.GetPageInfo()
	thesiss, total := thesis_models.ThesisList(page, pageCount)
	pageInfo := controller.PageInfoWithEndPage{
		CurrentPage: page,
		IsEndPage:   stringi.Judge(len(thesiss) < pageCount, "yes", "no"),
		Total:       total,
	}
	this.SuccessWithDataList(thesiss, pageInfo)
}

//论文上传
func (this *ThesisController) UploadThesis() {
	this.User = this.GetUser(this.Ctx.Request.Header.Get("X-Access-Token"))
	if this.User.IsGuest {
		this.RequireLogin()
		return
	}
	//只允许学生上传
	if this.User.Role != controller.ROLE_STUDENT {
		this.ErrorResponse(gtms_error.GetError("access_denied"))
		return
	}
	inputs := forms.UploadThesisForm{}
	if err := this.ParseInput(&inputs); err.Code != 0 {
		this.ErrorResponse(err)
		return
	}
	file, _ := this.GetFile("file")
	fileStr := strings.Split(file.Filename, ".")
	//限制上传格式
	if fileStr[len(fileStr)-1] != thesis_models.File_Type {
		this.ErrorResponse(gtms_error.GetError("only_pdf"))
		return
	}
	err, fileName := thesis_models.UploadThesis(&inputs, &this.Request)
	if err.Code != 0 {
		this.ErrorResponse(err)
		return
	} else {
		switch inputs.ThesisType {
		case controller.Opening_report:
			this.SaveFile(controller.Opening_report, "file", fileName)
		case controller.Thesis:
			this.SaveFile(controller.Thesis, "file", fileName)
		}
	}
	this.SuccessWithData(helper.JSON{})
}

//论文下载
func (this *ThesisController) DownloadThesis() {
	fileName := this.GetString("fileName")
	this.User = this.GetUser(this.Ctx.Request.Header.Get("X-Access-Token"))
	if this.User.IsGuest {
		this.RequireLogin()
		return
	}
	filePath := helper.GetRootPath() + "/upload/"
	fileStr := strings.Split(fileName, "_")
	if fileStr[1] == "开题报告" {
		filePath = filePath + controller.Opening_report + "/"
	} else {
		filePath = filePath + controller.Thesis + "/"
	}
	if b, err := helper.FolderExists(filePath + fileName); !b || err != nil {
		this.ErrorResponse(gtms_error.GetError("file_not_exits"))
		return
	}
	this.Ctx.Output.Download(filePath + fileName)
}

//论文预览
func (this *ThesisController) PreviewThesis() {
	fileName := this.GetString("fileName")
	this.User = this.GetUser(this.Ctx.Request.Header.Get("X-Access-Token"))
	if this.User.IsGuest {
		this.RequireLogin()
		return
	}
	filePath := helper.GetRootPath() + "/upload/"
	fileStr := strings.Split(fileName, "_")
	if fileStr[1] == "开题报告" {
		filePath = filePath + controller.Opening_report + "/"
	} else {
		filePath = filePath + controller.Thesis + "/"
	}
	if b, err := helper.FolderExists(filePath + fileName); !b || err != nil {
		this.ErrorResponse(gtms_error.GetError("file_not_exits"))
		return
	}
	//复制文件到gtms_web/pdfjs/web
	if conf.GetRunMode() == "dev" {
		cmd := exec.Command("/bin/sh", "-c", `cp `+filePath+fileName+` ~/WebstormProjects/gtms_web/pdfjs/web/`)
		if _, err := cmd.Output(); err != nil {
			this.ErrorResponse(gtms_error.GetError("preview_failed"))
			return
		}
	} else {
		cmd := exec.Command("/bin/sh", "-c", `cp `+filePath+fileName+` /www/gtms_web/pdfjs/web/`)
		if _, err := cmd.Output(); err != nil {
			this.ErrorResponse(gtms_error.GetError("preview_failed"))
			return
		}
	}
	this.SuccessWithData(helper.JSON{})
}
