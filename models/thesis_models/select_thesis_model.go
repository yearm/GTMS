package thesis_models

import (
	"GTMS/boot"
	"GTMS/library/controller"
	"GTMS/library/db"
	"GTMS/library/gtms_error"
	"GTMS/library/helper"
	"GTMS/library/stringi"
	"GTMS/library/validator"
	"GTMS/models/account_models"
	"GTMS/models/admin_models"
	"GTMS/v1/forms"
	"github.com/astaxie/beego/orm"
	"strconv"
)

type SelectedThesis struct {
	Sid         int64  `orm:"pk" json:"sid"`
	Uid         string `json:"uid"`
	Tid         string `json:"tid"`
	TechId      string `json:"techId"`
	Confirm     string `json:"confirm"`
	CreateTime  string `json:"createTime"`
	ConfirmTime string `json:"confirmTime"`
}

//教师未确定的论文
type NotConfirmThesis struct {
	SelectedThesis         `json:"selectedThesis"`
	account_models.Student `json:"student"`
	Thesis                 `json:"thesis"`
}

//已选题目
type ConfirmThesis struct {
	SelectedThesis         `json:"selectedThesis"`
	account_models.Student `json:"student"`
	Thesis                 `json:"thesis"`
	account_models.Teacher `json:"teacher"`
}

const (
	subject = "黄冈师范学院毕业论文管理系统学生选题"
)

func init() {
	//需要在init中注册定义的model
	orm.RegisterModel(new(SelectedThesis))
}

func SelectThesis(opt *forms.SelectThesisForm, req *controller.Request) *validator.Error {
	o := boot.GetMasterMySQL()
	//获取选题开放时间
	openingTime := admin_models.OpeningTime{Year: stringi.ToInt(helper.Date("Y"))}
	o.Read(&openingTime)
	if openingTime.StartTime > helper.Date("Y-m-d H:i:s") || openingTime.EndTime < helper.Date("Y-m-d H:i:s") {
		return gtms_error.GetError("not_openingTime")
	}
	//判断有没有选题(教师未同意可以重选)
	if o.QueryTable((*SelectedThesis)(nil)).Filter("uid", req.User.StuNo).Filter("confirm__in", 0, 1).Exist() {
		return gtms_error.GetError("only_select_one")
	}
	//根据tid查询论文信息
	qs := o.QueryTable((*Thesis)(nil))
	thesis := Thesis{Tid: stringi.ToInt64(opt.Tid)}
	o.Read(&thesis)
	if thesis.Status == NotOptional_Status {
		return gtms_error.GetError("other_selected")
	}
	//开启事物
	o.Begin()
	_, err1 := o.Insert(&SelectedThesis{
		Uid:         req.User.StuNo,
		Tid:         opt.Tid,
		TechId:      thesis.UpdateUid,
		Confirm:     Optional_Status,
		CreateTime:  helper.Date("Y-m-d H:i:s"),
		ConfirmTime: helper.Date("Y-m-d H:i:s"),
	})
	_, err2 := qs.Filter("tid", opt.Tid).Update(orm.Params{
		"status": NotOptional_Status,
	})
	if err1 != nil || err2 != nil {
		o.Rollback()
		return gtms_error.GetError("select_thesis_failed")
	}
	o.Commit()
	//发送邮件给教师...
	tech := account_models.Teacher{TechId: thesis.UpdateUid}
	o.Read(&tech)
	body := `<h4>` + thesis.UpdateUser + `:<br>
			您好！<br>
			学生(` + req.User.StuName + `)已选取您的毕业论文题目(` + thesis.Subject + `),请尽快处理!<br><br>
			黄冈师范学院毕业论文管理系统<br><br>
			注:此邮件为本系统所发，非对方邮箱所发，所以请勿回邮。
		</h4>`
	boot.SendEmail(tech.Email, thesis.UpdateUser, subject, body)
	return &validator.Error{}
}

func ConfirmSelectedThesis(opt *forms.ConfirmSelectedlThesisForm) *validator.Error {
	o := boot.GetMasterMySQL()
	//获取选题开放时间
	openingTime := admin_models.OpeningTime{Year: stringi.ToInt(helper.Date("Y"))}
	o.Read(&openingTime)
	if openingTime.StartTime > helper.Date("Y-m-d H:i:s") || openingTime.EndTime < helper.Date("Y-m-d H:i:s") {
		return gtms_error.GetError("not_openingTime")
	}
	qs := o.QueryTable((*SelectedThesis)(nil))
	_, err := qs.Filter("sid", opt.Sid).Update(orm.Params{
		"confirm": opt.Confirm,
	})
	//如果教师不同意，则该论文重新进入可选状态
	var err1 error
	if opt.Confirm == "2" {
		sql := `UPDATE thesis SET status = '0' WHERE tid = (SELECT tid FROM selected_thesis WHERE sid = :sid)`
		_, err1 = db.Exec(db.BuildSQL(sql, stringi.Form{
			"sid": opt.Sid,
		}))
	}
	if err != nil || err1 != nil {
		return gtms_error.GetError("confirm_error")
	}
	return &validator.Error{}

}

func GetNotOrConfirmThesis(req *controller.Request, page int, pageCount int, confirm string) (ncThesis []*NotConfirmThesis, total int) {
	if confirm == "" { //默认获取已确认的选题
		confirm = "1"
	}
	sql := `SELECT tmp.*, s.*
FROM (SELECT st.*,
             t.subject,
             t.subtopic,
             t.keyword,
             t.type,
             t.source,
             t.workload,
             t.degree_difficulty,
             t.research_direc,
             t.content,
             t.update_uid,
             t.update_user,
             t.update_time,
             t.status
      FROM selected_thesis AS st
             INNER JOIN thesis AS t ON st.tid = t.tid
      WHERE st.tech_id = :tech_id
        AND st.confirm = :confirm) AS tmp
       INNER JOIN student AS s
ON tmp.uid = s.stu_no
LIMIT @start, @pageCount`
	db.QueryRows(sql, stringi.Form{
		"tech_id":   req.User.TechId,
		"confirm":   confirm,
		"start":     strconv.Itoa((page - 1) * pageCount),
		"pageCount": strconv.Itoa(pageCount),
	}, &ncThesis)
	db.QueryRow(`SELECT COUNT(*) FROM selected_thesis WHERE tech_id = :tech_id AND confirm = :confirm`, stringi.Form{
		"tech_id": req.User.TechId,
		"confirm": confirm,
	}, &total)
	return
}

func SelectedThesisList(page int, pageCount int) (confirmThesis []*ConfirmThesis, total int) {
	sql := `SELECT temp.*, te.*
FROM (SELECT tmp.*, s.*
      FROM (SELECT st.*,
                   t.SUBJECT,
                   t.subtopic,
                   t.keyword,
                   t.type,
                   t.source,
                   t.workload,
                   t.degree_difficulty,
                   t.research_direc,
                   t.content,
                   t.update_time,
                   t.STATUS
            FROM selected_thesis AS st
                   INNER JOIN thesis AS t ON st.tid = t.tid
            WHERE st.confirm = '1') AS tmp
             INNER JOIN student AS s
      ON tmp.uid = s.stu_no) temp
       INNER JOIN teacher te ON temp.tech_id = te.tech_id
LIMIT @start, @pageCount`
	db.QueryRows(sql, stringi.Form{
		"start":     strconv.Itoa((page - 1) * pageCount),
		"pageCount": strconv.Itoa(pageCount),
	}, &confirmThesis)
	db.QueryRow(`SELECT COUNT(*) FROM selected_thesis WHERE confirm = '1'`, stringi.Form{
	}, &total)
	return
}

func GetThesis(req *controller.Request) (thesis []*ConfirmThesis) {
	sql := `SELECT temp.*, te.*
FROM (SELECT tmp.*, s.*
      FROM (SELECT st.*,
                   t.SUBJECT,
                   t.subtopic,
                   t.keyword,
                   t.type,
                   t.source,
                   t.workload,
                   t.degree_difficulty,
                   t.research_direc,
                   t.content,
                   t.update_time,
                   t.STATUS
            FROM selected_thesis AS st
                   INNER JOIN thesis AS t ON st.tid = t.tid
            WHERE st.confirm = '1') AS tmp
             INNER JOIN student AS s ON tmp.uid = s.stu_no) temp
       INNER JOIN teacher te ON temp.tech_id = te.tech_id
WHERE temp.stu_no = :stu_no
LIMIT 1`
	db.QueryRows(sql, stringi.Form{
		"stu_no": req.User.StuNo,
	}, &thesis)
	return
}
