package db

import (
	"GTMS/boot"
	"GTMS/library/stringi"
	"database/sql"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
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

func QueryRows(sql string, bind stringi.Form, st interface{}) (int64, error) {
	var client = boot.GetSlaveMySQL()
	sql = BuildSQL(sql, bind)
	return client.Raw(sql).QueryRows(st)
}

func Fetch(sql string, bind ...stringi.Form) QueryResult {
	sql = BuildSQL(sql, bind...)
	res := FetchAll(sql)
	if len(res) > 0 {
		return res[0]
	}
	return nil
}

func FetchAll(sql string, bind ...stringi.Form) []QueryResult {
	var client = boot.GetSlaveMySQL()
	sql = BuildSQL(sql, bind...)
	var temp = make([]orm.Params, 0)
	var results = make([]QueryResult, 0)
	_, err := client.Raw(sql).Values(&temp)
	if err != nil {
		println(err.Error())
	}
	for _, row := range temp {
		var item = QueryResult{}
		for k, v := range row {
			item[k] = NewResult(v)
		}
		results = append(results, item)
	}
	return results
}

func Exec(sql string, bind ...stringi.Form) (sql.Result, error) {
	var client = boot.GetMasterMySQL()
	sql = BuildSQL(sql, bind...)
	res, err := client.Raw(sql).Exec()
	if err != nil {
		beego.BeeLogger.Error("SQL: %s, Message: %s", sql, err.Error())
	}
	return res, err
}

func InsertSQL(tableName string, data stringi.Form) string {
	var keys []string
	var values []string
	for key, _ := range data {
		keys = append(keys, key)
		values = append(values, ":"+key)
	}
	var keyString = strings.Join(keys, ", ")
	var valueString = strings.Join(values, ", ")
	var sql = "INSERT INTO {tableName} ({keys}) VALUES ({values})"
	sql = stringi.Build(sql, stringi.Form{
		"tableName": tableName,
		"keys":      keyString,
		"values":    valueString,
	})
	return BuildSQL(sql, data)
}

func InsertAllSQL(tableName string, data []stringi.Form) string {
	var keys []string
	for key, _ := range data[0] {
		keys = append(keys, key)
	}
	var keyString = strings.Join(keys, ", ")
	var vals []string
	for _, item := range data {
		var str = "('" + strings.Join(stringi.ArrayValues(item, keys), "', '") + "')"
		vals = append(vals, str)
	}
	var values = strings.Join(vals, ", ")
	var sql = "INSERT INTO {tableName} ({keys}) VALUES {values}"
	sql = stringi.Build(sql, stringi.Form{
		"tableName": tableName,
		"keys":      keyString,
		"values":    values,
	})
	return sql
}

func ReplaceSQL(tableName string, data stringi.Form) string {
	var keys []string
	var values []string
	for key, _ := range data {
		keys = append(keys, key)
		values = append(values, ":"+key)
	}
	var keyString = strings.Join(keys, ", ")
	var valueString = strings.Join(values, ", ")
	var sql = "REPLACE INTO {tableName} ({keys}) VALUES ({values})"
	sql = stringi.Build(sql, stringi.Form{
		"tableName": tableName,
		"keys":      keyString,
		"values":    valueString,
	})
	return BuildSQL(sql, data)
}

func DeleteSQL(tableName string, field string, data string) string {
	sql := `DELETE FROM {tableName} WHERE {field} IN ({value})`
	sql = stringi.Build(sql, stringi.Form{
		"tableName": tableName,
		"field":     field,
		"value":     data,
	})
	return sql
}

// 设置新数据
func Set(data stringi.Form) string {
	var arr = make([]string, 0)
	for k, v := range data {
		var item = stringi.Build("{k} = '{v}'", stringi.Form{
			"k": k,
			"v": stringi.AddSlashes(v),
		})
		arr = append(arr, item)
	}
	return strings.Join(arr, ", ")
}
