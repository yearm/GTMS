package notice_controllers

import (
	"GTMS/library/controller"
	"GTMS/library/gtms_error"
	"GTMS/library/helper"
	"GTMS/library/stringi"
	"GTMS/models/notice_models"
	"GTMS/v1/forms"
	"strings"
)

type NoticeController struct {
	controller.BaseController
}

//公告列表
func (this *NoticeController) NoticeList() {
	this.User = this.GetUser(this.Ctx.Request.Header.Get("X-Access-Token"))
	if this.User.IsGuest {
		this.RequireLogin()
		return
	}
	page, pageCount := this.GetPageInfo()
	notices, total := notice_models.NoticeList(page, pageCount)
	pageInfo := controller.PageInfoWithEndPage{
		CurrentPage: page,
		IsEndPage:   stringi.Judge(len(notices) < pageCount, "yes", "no"),
		Total:       total,
	}
	this.SuccessWithDataList(notices, pageInfo)
}

//添加公告
func (this *NoticeController) AddNotice() {
	this.User = this.GetUser(this.Ctx.Request.Header.Get("X-Access-Token"))
	if this.User.IsGuest {
		this.RequireLogin()
		return
	}
	//只允许管理员添加
	if this.User.Role != controller.ROLE_ADMIN {
		this.ErrorResponse(gtms_error.GetError("access_denied"))
		return
	}
	inputs := forms.AddNoticeForm{}
	if err := this.ParseInput(&inputs); err.Code != 0 {
		this.ErrorResponse(err)
		return
	}
	//获取附件
	files, _ := this.GetFiles("attach")
	var attachs []string
	for _, v := range files {
		attachs = append(attachs, v.Filename)
		this.SaveFile(controller.Notice_attach, "attach", v.Filename)
	}
	attachments := strings.Join(attachs, ",")
	err := notice_models.AddNotice(&this.Request, &inputs, attachments)
	if err.Code != 0 {
		this.ErrorResponse(err)
		return
	}
	this.SuccessWithData(helper.JSON{})
}

//删除公告
func (this *NoticeController) NoticeDel() {
	nid := this.GetString("nid")
	this.User = this.GetUser(this.Ctx.Request.Header.Get("X-Access-Token"))
	if this.User.IsGuest {
		this.RequireLogin()
		return
	}
	//只允许管理员删除
	if this.User.Role != controller.ROLE_ADMIN {
		this.ErrorResponse(gtms_error.GetError("access_denied"))
		return
	}
	err := notice_models.NoticeDel(nid)
	if err.Code != 0 {
		this.ErrorResponse(err)
		return
	}
	this.SuccessWithData(helper.JSON{})
}

//公告详情
func (this *NoticeController) NoticeDetail() {
	nid := this.GetString("nid")
	this.User = this.GetUser(this.Ctx.Request.Header.Get("X-Access-Token"))
	if this.User.IsGuest {
		this.RequireLogin()
		return
	}
	notice := notice_models.NoticeDetail(nid)
	this.SuccessWithData(notice)
}

//下载公告附件
func (this *NoticeController) DownloadAttach() {
	attachFile := this.GetString("attach")
	this.User = this.GetUser(this.Ctx.Request.Header.Get("X-Access-Token"))
	if this.User.IsGuest {
		this.RequireLogin()
		return
	}
	filePath := helper.GetRootPath() + "/upload/" + controller.Notice_attach + "/"
	if b, err := helper.FolderExists(filePath + attachFile); !b || err != nil {
		this.ErrorResponse(gtms_error.GetError("file_not_exits"))
		return
	}
	this.Ctx.Output.Download(filePath + attachFile)
}
