package boot

import (
	"GTMS/conf"
	"GTMS/library/helper"
	"GTMS/library/stringi"
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"strconv"
)

var gtms map[string]orm.Ormer

func init() {
	gtms = make(map[string]orm.Ormer, 0)
}

func ConnectMySQL() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	LoadingGtmsCluster()
}

//加载数据库集群
func LoadingGtmsCluster() {
	configs := conf.GetMySQLConfig()
	tpl := "{userName}:{password}@tcp({host}:{port})/{DbName}?charset=utf8&loc=Asia%2FShanghai"
	maxIdle := 200 //最大空闲连接
	maxConn := 200 //最大数据库连接
	for i := 0; i < len(configs); i++ {
		dataSource := stringi.Build(tpl, stringi.Form{
			"userName": configs[i].UserName,
			"password": configs[i].Password,
			"host":     configs[i].Host,
			"port":     configs[i].Port,
			"DbName":   configs[i].DbName,
		})
		if i == 0 {
			orm.RegisterDataBase("default", "mysql", dataSource, maxIdle, maxConn)
			gtms["gtms-0"] = orm.NewOrm()
		} else {
			name := "gtms-" + strconv.Itoa(i)
			orm.RegisterDataBase(name, "mysql", dataSource, maxIdle, maxConn)
			gtms[name] = orm.NewOrm()
		}
	}
}

func GetMasterMySQL() orm.Ormer {
	gtms["gtms-0"].Using("default")
	return gtms["gtms-0"]
}

func GetSlaveMySQL() orm.Ormer {
	if len(gtms) <= 1 {
		gtms["gtms-0"].Using("default")
		return gtms["gtms-0"]
	} else {
		i := helper.RealRandNum(int64(len(gtms)))
		//如果随机到0(主库)，就加1取从库
		if i == 0 {
			i++
		}
		name := "gtms-" + fmt.Sprintf("%s", i)
		gtms[name].Using(name)
		return gtms[name]
	}
}
