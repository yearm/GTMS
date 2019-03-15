package db

import (
	"GTMS/boot"
	"GTMS/library/stringi"
	"sort"
	"strings"
)

// :转义, @不转义
func BuildSQL(sql string, bind ...stringi.Form) string {
	var params = stringi.Form{}
	if len(bind) > 0 {
		params = bind[0]
	}
	var keys []string
	for key, _ := range params {
		keys = append(keys, key)
	}
	if len(keys) == 0 {
		return sql
	}
	sort.Strings(keys)
	stringi.Reverse(keys)
	for _, key := range keys {
		sql = strings.Replace(sql, "@"+key, params[key], -1)
		val := stringi.AddSlashes(params[key])
		sql = strings.Replace(sql, ":"+key, "'"+val+"'", -1)
	}
	return sql
}

func QueryRow(sql string, bind stringi.Form, st interface{}) error {
	var client = boot.GetSlaveMySQL()
	sql = BuildSQL(sql, bind)
	return client.Raw(sql).QueryRow(st)
}
