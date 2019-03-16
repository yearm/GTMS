package db

type QueryResult map[string]*Result

type Result struct {
	v  interface{}
	s  string
	i  int
	b  bool
	i6 int64
}

func NewResult(v interface{}) *Result {
	var st = new(Result)
	if v == nil {
		v = ""
	}
	st.v = v
	return st
}
