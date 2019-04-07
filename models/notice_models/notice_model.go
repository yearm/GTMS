package notice_models

import (
	"GTMS/boot"
	"GTMS/library/controller"
	"GTMS/library/gtms_error"
	"GTMS/library/helper"
	"GTMS/library/stringi"
	"GTMS/library/validator"
	"GTMS/v1/forms"
	"github.com/astaxie/beego/orm"
)

type Notice struct {
	Nid        int64  `orm:"pk" json:"nid"`
	Title      string `json:"title"`
	Content    string `json:"content"`
	Attachment string `json:"attachment"`
	View       int    `json:"view"`
	Uid        string `json:"uid"`
	CreateTime string `json:"createTime"`
}

func init() {
	//需要在init中注册定义的model
	orm.RegisterModel(new(Notice))
}

func NoticeList(page int, pageCount int) (notices []*Notice, total int) {
	o := boot.GetSlaveMySQL()
	qs := o.QueryTable((*Notice)(nil))
	_, err := qs.Limit(pageCount, (page-1)*pageCount).All(&notices)
	if err != nil {
		return
	}
	t, _ := qs.Count()
	total = int(t)
	return
}

func AddNotice(req *controller.Request, opt *forms.AddNoticeForm, attachment string) *validator.Error {
	o := boot.GetMasterMySQL()
	_, err := o.Insert(&Notice{
		Title:      opt.Title,
		Content:    opt.Content,
		Attachment: attachment,
		View:       0,
		Uid:        req.User.AdminId,
		CreateTime: helper.Date("Y-m-d H:i:s"),
	})
	if err != nil {
		return gtms_error.GetError("insert_error")
	}
	return &validator.Error{}
}

func NoticeDel(nid string) *validator.Error {
	o := boot.GetMasterMySQL()
	o.Delete(&Notice{Nid: stringi.ToInt64(nid)})
	return &validator.Error{}
}

func NoticeDetail(nid string) *Notice {
	o := boot.GetSlaveMySQL()
	notice := Notice{Nid: stringi.ToInt64(nid)}
	o.Read(&notice)
	notice.View = notice.View + 1
	o.Update(&notice)
	return &notice
}
