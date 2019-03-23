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
	CreateTime       string `json:"createTime"`
}

func init() {
	//需要在init中注册定义的model
	orm.RegisterModel(new(Thesis))
}

func AddThesis(opt *forms.AddThesisForm) *validator.Error {
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
		CreateTime:       helper.Date("Y-m-d H:i:s"),
	})
	if err != nil {
		return gtms_error.GetError("insert_error")
	}
	return &validator.Error{}
}

func DelThesis(opt *forms.DelThesisForm) *validator.Error {
	o := boot.GetMasterMySQL()
	o.Delete(&Thesis{Tid: stringi.ToInt64(opt.Tid)})
	return &validator.Error{}
}

func UpdateThesis(opt *forms.UpdateThesisForm) *validator.Error {
	sql := `UPDATE @table SET @value WHERE tid = :tid`
	value := db.Set(helper.StructToFormWithClearNilField(*opt, controller.FormatThesis))
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
