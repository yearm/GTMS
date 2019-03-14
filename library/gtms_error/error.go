package gtms_error

import (
	"GTMS/library/helper"
	"GTMS/library/stringi"
	"GTMS/library/validator"
	"github.com/go-ini/ini"
)

var (
	errFile *ini.File
	err     error
)

func init() {
	rootPath := helper.GetRootPath()
	errFile, err = ini.Load(rootPath + "/conf/err/error.ini")
	if err != nil {
		panic("load error.ini failed")
	}
}

func getDict(cfg *ini.File, section string, key string) string {
	sec, err1 := cfg.GetSection(section)
	if err1 != nil {
		panic(err1.Error())
	}
	k, err2 := sec.GetKey(key)
	if err2 != nil {
		panic(err2.Error())
	}
	return k.String()
}

func getErrorTemplate(tpl string) *validator.Error {
	cfg := errFile
	section, _ := cfg.GetSection(tpl)
	tmp, _ := section.GetKey("code")
	code, _ := tmp.Int()
	msg, _ := section.GetKey("msg")
	return &validator.Error{
		Code: code,
		Msg:  msg.String(),
	}
}

func GetError(tpl string, args ...string) *validator.Error {
	cfg := validator.GetDicts()
	var errorTpl = getErrorTemplate(tpl)
	var arr = make([]string, 0)
	for _, attr := range args {
		trans := getDict(cfg, "dict", attr)
		arr = append(arr, trans)
	}
	return &validator.Error{
		Code: errorTpl.Code,
		Msg:  stringi.Template(errorTpl.Msg, arr...),
	}
}
