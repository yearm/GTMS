package conf

import (
	"GTMS/library/helper"
	"fmt"
	"github.com/astaxie/beego/orm"
	"github.com/go-ini/ini"
	"github.com/json-iterator/go"
)

var cfg *ini.File
var env *ini.File

func getKey(conf *ini.File, sec string, key string) (k *ini.Key) {
	k, err := conf.Section(sec).GetKey(key)
	if err != nil {
		panic(fmt.Sprintf("Config error! sec: %s, key: %s", sec, key))
	}
	return
}

//加载配置文件
func LoadingConf() {
	rootPath := helper.GetRootPath()
	jsoniter.ConfigDefault = jsoniter.Config{EscapeHTML: false}.Froze() // 禁止HTML转义
	env, _ := ini.Load(rootPath + "/conf/app.conf")
	runMode := getKey(env, "", "runmode").String()
	if runMode == "prod" {
		cfg, _ = ini.Load(rootPath + "/conf/app_prod.ini")
	} else if runMode == "dev" {
		cfg, _ = ini.Load(rootPath + "/conf/app_dev.ini")
		orm.Debug = true //打印SQL
	}
}

//获取mysql的配置
func GetMySQLConfig() (configs []MySQlConfig) {
	hosts := cfg.Section("mysql").Key("host").Strings(",")
	port := cfg.Section("mysql").Key("port").String()
	userName := cfg.Section("mysql").Key("userName").String()
	password := cfg.Section("mysql").Key("password").String()
	dbName := cfg.Section("mysql").Key("dbName").String()
	for _, v := range hosts {
		configs = append(configs, MySQlConfig{
			Host:     v,
			Port:     port,
			UserName: userName,
			Password: password,
			DbName:   dbName,
		})
	}
	return
}

//获取redis的配置
func GetRedisConfig() RedisConfig {
	host := cfg.Section("redis").Key("host").String()
	port := cfg.Section("redis").Key("port").String()
	password := cfg.Section("redis").Key("password").String()
	db, _ := cfg.Section("redis").Key("db").Int()
	return RedisConfig{
		Host:     host,
		Port:     port,
		Password: password,
		Db:       db,
	}
}
