package validator

import (
	"GTMS/library/stringi"
	"reflect"
	"strconv"
	"strings"
)

var formatMapping = map[string]func(s string) bool{
	"required": IsRequired,
	"email":    IsEmail,
	"ip":       IsIP,
	"url":      IsURL,
	"numeric":  IsNumeric,
}

var valueMapping = map[string]func(s string, value interface{}) (bool, error){
	"min":       min,
	"max":       max,
	"minLength": minLength,
	"maxLength": maxLength,
	"minSize":   minSize,
	"maxSize":   maxSize,
}

func isValid(checker *Checker, name string, rule string, value interface{}) *Error {
	arr := strings.Split(rule, ":")
	if arr[0] == "switch" {
		v, ok1 := value.(string)
		if !ok1 {
			panic("switch only support string type")
		}
		ok2, err := testSwitch(rule, v)
		if err != nil {
			panic(name + err.Error())
		}
		if !ok2 {
			msg, err2 := GetParam(checker.Dicts, "tpl", "switch")
			if err2 != nil {
				panic(name + err2.Error())
			}
			attr, _ := GetParam(checker.Dicts, "dict", name)
			msg = strings.Replace(msg, ":attr", attr, 1)
			msg = strings.Replace(msg, ":limit", arr[1], 1)
			return &Error{
				Code: 1,
				Msg:  msg,
			}
		}
	} else if len(arr) == 1 {
		v, ok := value.(string)
		if !ok {
			panic(name + " must be string type")
		}
		fn, ok := formatMapping[arr[0]]
		if !ok {
			panic(arr[0] + "'s handler not exist")
		}
		if !fn(v) {
			return &Error{
				Code: 1,
				Msg:  checker.GetMessage(arr[0], name),
			}
		}
	} else {
		limit, err1 := strconv.ParseFloat(arr[1], 64)
		if err1 != nil {
			panic(name + " tag error")
		}
		fn, ok := valueMapping[arr[0]]
		if !ok {
			panic(arr[0] + "'s handler not exist")
		}

		pass, err2 := fn(rule, value)
		if err2 != nil {
			panic(name + err2.Error())
		}

		if !pass {
			return &Error{
				Code: 1,
				Msg:  checker.GetMessage(arr[0], name, limit),
			}
		}
	}
	return &Error{}
}

// 设置默认值
func setValue(field reflect.Value, v string) {
	dataType := field.Type().Name()
	if dataType == "string" && field.String() == "" {
		field.SetString(v)
	} else if (dataType == "int64" || dataType == "int") && field.Int() == 0 {
		val, _ := strconv.Atoi(v)
		field.SetInt(int64(val))
	}
}

type Error struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func Check(inputs interface{}) *Error {
	checker := NewChecker(dicts)
	t := reflect.TypeOf(inputs).Elem()
	v := reflect.ValueOf(inputs).Elem()
	for i := 0; i < t.NumField(); i++ {
		defaultValue := t.Field(i).Tag.Get("default")
		if defaultValue != "" {
			setValue(v.Field(i), defaultValue)
		}

		rule := t.Field(i).Tag.Get("valid")
		if rule == "" {
			continue
		}
		val := v.Field(i).Interface()

		name := stringi.ToCamel(t.Field(i).Name)
		alias := t.Field(i).Tag.Get("alias")
		if alias != "" {
			name = alias
		}

		fields := strings.Split(rule, "|")
		for _, field := range fields {
			e := isValid(checker, name, field, val)
			if e.Code != 0 {
				return e
			}
		}
	}
	return &Error{}
}
