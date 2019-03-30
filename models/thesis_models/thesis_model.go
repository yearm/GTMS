package thesis_models

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
	"strconv"
)

type Thesis struct {
	Tid              int64  `orm:"pk" json:"tid"`
	Subject          string `json:"subject"`
	Subtopic         string `json:"subtopic"`
	Keyword          string `json:"keyword"`
	Type             string `json:"type"`
	Source           string `json:"source"`
	Workload         string `json:"workload"`
	DegreeDifficulty string `json:"degreeDifficulty"`
	ResearchDirec    string `json:"researchDirec"`
	Content          string `json:"content"`
	UpdateUid        string `json:"updateUid"`
	UpdateUser       string `json:"updateUser"`
	UpdateTime       string `json:"updateTime"`
	Status           string `json:"status"`
}

type ThesisFile struct {
	Uid           int64  `orm:"pk" json:"uid"`
	Tid           string `json:"tid"`
	OpeningReport string `json:"openingReport"`
	ReportTime    string `json:"reportTime"`
	Thesis        string `json:"thesis"`
	ThesisTime    string `json:"thesisTime"`
}

const (
	Optional_Status    = "0" //论文可选
	NotOptional_Status = "1" //不可选
	File_Type          = "pdf"
)

func init() {
	//需要在init中注册定义的model
	orm.RegisterModel(new(Thesis))
	orm.RegisterModel(new(ThesisFile))

}

func AddThesis(req *controller.Request, opt *forms.AddThesisForm) *validator.Error {
	o := boot.GetMasterMySQL()
	_, err := o.Insert(&Thesis{
		Subject:          opt.Subject,
		Subtopic:         opt.Subtopic,
		Keyword:          opt.Keyword,
		Type:             opt.Type,
		Source:           opt.Source,
		Workload:         opt.Workload,
		DegreeDifficulty: opt.DegreeDifficulty,
		ResearchDirec:    opt.ResearchDirec,
		Content:          opt.Content,
		UpdateUid:        req.User.TechId,
		UpdateUser:       req.User.TechName,
		UpdateTime:       helper.Date("Y-m-d H:i:s"),
		Status:           Optional_Status,
	})
	if err != nil {
		return gtms_error.GetError("insert_error")
	}
	return &validator.Error{}
}

func DelThesis(tid string) *validator.Error {
	o := boot.GetMasterMySQL()
	o.Delete(&Thesis{Tid: stringi.ToInt64(tid)})
	return &validator.Error{}
}

func UpdateThesis(opt *forms.UpdateThesisForm, req *controller.Request) *validator.Error {
	sql := `UPDATE @table SET @value WHERE tid = :tid`
	form := helper.StructToFormWithClearNilField(*opt, controller.FormatThesis)
	form["update_uid"] = req.User.TechId
	form["update_user"] = req.User.TechName
	form["update_time"] = helper.Date("Y-m-d H:i:s")
	value := db.Set(form)
	_, err := db.Exec(sql, stringi.Form{
		"table": "thesis",
		"value": value,
		"tid":   opt.Tid,
	})
	if err != nil {
		return gtms_error.GetError("update_info_error")
	}
	return &validator.Error{}
}

func ThesisList(page int, pageCount int) (thesiss []*Thesis, total int) {
	o := boot.GetSlaveMySQL()
	qs := o.QueryTable((*Thesis)(nil))
	_, err := qs.Limit(pageCount, (page-1)*pageCount).All(&thesiss)
	if err != nil {
		return
	}
	t, _ := qs.Count()
	total = int(t)
	return
}

func UploadThesis(opt *forms.UploadThesisForm, req *controller.Request) (*validator.Error, string) {
	var fileName string
	sql := `SELECT * FROM thesis WHERE tid = (SELECT tid FROM selected_thesis WHERE uid = :uid AND confirm = '1') LIMIT 1`
	thesis := Thesis{}
	db.QueryRow(sql, stringi.Form{
		"uid": req.User.StuNo,
	}, &thesis)
	if thesis.Tid == 0 {
		return gtms_error.GetError("thesis_unconfirmed"), ""
	}
	o := boot.GetSlaveMySQL()
	thesisFile := ThesisFile{Uid: stringi.ToInt64(req.User.StuNo)}
	o.Read(&thesisFile)
	if opt.ThesisType == controller.Opening_report {
		_, err := db.Exec(db.ReplaceSQL("thesis_file", stringi.Form{
			"uid":            req.User.StuNo,
			"tid":            strconv.FormatInt(thesis.Tid, 10),
			"opening_report": thesis.Subject + "_开题报告_" + req.User.StuName + "." + File_Type,
			"report_time":    helper.Date("Y-m-d H:i:s"),
			"thesis":         thesisFile.Thesis,
			"thesis_time":    thesisFile.ThesisTime,
		}))
		fileName = thesis.Subject + "_开题报告_" + req.User.StuName + "." + File_Type
		if err != nil {
			return gtms_error.GetError("upload_error"), ""
		}
	} else if opt.ThesisType == controller.Thesis {
		_, err := db.Exec(db.ReplaceSQL("thesis_file", stringi.Form{
			"uid":            req.User.StuNo,
			"tid":            strconv.FormatInt(thesis.Tid, 10),
			"opening_report": thesisFile.OpeningReport,
			"report_time":    thesisFile.ReportTime,
			"thesis":         thesis.Subject + "_" + req.User.StuName + "." + File_Type,
			"thesis_time":    helper.Date("Y-m-d H:i:s"),
		}))
		fileName = thesis.Subject + "_" + req.User.StuName + "." + File_Type
		if err != nil {
			return gtms_error.GetError("upload_error"), ""
		}
	}
	return &validator.Error{}, fileName
}
