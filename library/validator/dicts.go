package validator

import (
	"GTMS/library/helper"
	"errors"
	"github.com/go-ini/ini"
)

var (
	dicts *ini.File
	err   error
)

func init() {
	rootPath := helper.GetRootPath()
	dicts, err = ini.Load(rootPath + "conf/err/validator.ini")
	if err != nil {
		panic("load validator.ini failed")
	}
}

func GetParam(cfg *ini.File, section string, key string) (string, error) {
	s, err1 := cfg.GetSection(section)
	if err1 != nil {
		return "", errors.New("tpl " + section + " not exist")
	}
	k, err2 := s.GetKey(key)
	if err2 != nil {
		return "", errors.New("dict " + key + " not exist")
	}
	return k.String(), nil
}

func GetDicts() *ini.File {
	return dicts
}
