package stringi

import (
	"errors"
	"fmt"
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

func ToFloat64(value interface{}) (float64, error) {
	v1, ok := value.(int)
	if ok {
		return float64(v1), nil
	}

	v2, ok := value.(int64)
	if ok {
		return float64(v2), nil
	}

	v3, ok := value.(float32)
	if ok {
		return float64(v3), nil
	}

	v4, ok := value.(float64)
	if ok {
		return float64(v4), nil
	}
	return 0, errors.New(" only support int, int64, float32, float64")
}

func ToString(v interface{}) string {
	var s = ""
	if v == nil {
		return s
	} else if f, ok := v.(string); ok {
		s = f
	} else if f, ok := v.(int64); ok {
		s = ToString(f)
	} else if f, ok := v.(float64); ok {
		s = fmt.Sprintf("%.6f", f)
		re, _ := regexp.Compile(`\.*[0]+$`)
		s = re.ReplaceAllString(s, "")
	} else {
		panic("val only supports int, int64, float, string type")
	}
	return s
}

func ArrayValues(form Form, keys []string) []string {
	var result []string
	for _, key := range keys {
		s := ToString(form[key])
		result = append(result, AddSlashes(s))
	}
	return result
}

// 三元操作
func Judge(flag bool, v1 string, v2 string) string {
	if flag == true {
		return v1
	}
	return v2
}
