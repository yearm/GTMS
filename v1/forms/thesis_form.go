package forms

//添加论文Form
type AddThesisForm struct {
	Subject          string `form:"subject"`
	Subtopic         string `form:"subtopic"`
	Keyword          string `form:"keyword"`
	Type             string `form:"type"`
	Source           string `form:"source"`
	Workload         string `form:"workload"`
	DegreeDifficulty string `form:"degreeDifficulty"`
	ResearchDirec    string `form:"researchDirec"`
	Content          string `form:"content"`
}

//删除论文Form
type DelThesisForm struct {
	Tid string `form:"tid" valid:"required"`
}

//更新论文信息Form
type UpdateThesisForm struct {
	Tid              string `form:"tid" valid:"required"`
	Subject          string `form:"subject"`
	Subtopic         string `form:"subtopic"`
	Keyword          string `form:"keyword"`
	Type             string `form:"type"`
	Source           string `form:"source"`
	Workload         string `form:"workload"`
	DegreeDifficulty string `form:"degreeDifficulty"`
	ResearchDirec    string `form:"researchDirec"`
	Content          string `form:"content"`
	Status           string `form:"status"`
}
