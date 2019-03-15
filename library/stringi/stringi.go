package stringi

import (
	"regexp"
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

// 字符串模板
func Template(tpl string, args ...string) string {
	re := regexp.MustCompile(`{{.*?}}`)
	i := -1
	return re.ReplaceAllStringFunc(tpl, func(s string) string {
		i++
		return args[i]
	})
}

func IsEmpty(str string) bool {
	str = strings.TrimSpace(str)
	return (str == "") || (str == "0")
}

func DefaultValue(s string, val string) string {
	if IsEmpty(s) {
		return val
	} else {
		return s
	}
}

func Swap(a string, b string) (string, string) {
	return b, a
}

//反转切片
func Reverse(arr []string) {
	var n int
	var length = len(arr)
	n = length / 2
	for i := 0; i < n; i++ {
		arr[i], arr[length-1-i] = Swap(arr[i], arr[length-1-i])
	}
}

//转义引号
func AddSlashes(str string) string {
	str = strings.Replace(str, "'", "\\'", -1)
	str = strings.Replace(str, "\"", "\\\"", -1)
	str = strings.Replace(str, "`", "\\`", -1)
	return str
}
