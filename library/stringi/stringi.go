package stringi

import (
	"strconv"
	"strings"
)

type Form map[string]string
type Forms []Form
type JSON map[string]interface{}

//字符串模板
func Build(s string, bind Form) (str string) {
	for k, v := range bind {
		str = strings.Replace(s, "{"+k+"}", v, -1)
		s = str
	}
	return
}

func ToInt(str string) int {
	i, _ := strconv.Atoi(str)
	return i
}
