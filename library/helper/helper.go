package helper

import (
	"GTMS/library/stringi"
	"crypto/rand"
	"fmt"
	"github.com/astaxie/beego/logs"
	"math/big"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

type JSON map[string]interface{}

//获取项目根路径
func GetRootPath() (path string) {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		logs.Error(err.Error())
	}
	path = strings.Replace(dir, "\\", "/", -1)
	return
}

//获取真随机数
func RealRandNum(maxNum int64) (i int) {
	result, _ := rand.Int(rand.Reader, big.NewInt(maxNum))
	str := fmt.Sprintf("%s", result)
	i = stringi.ToInt(str)
	return
}

//格式化时间戳
func Date(format string, timestamp ...int64) string {
	var ts = time.Now().Unix()
	if len(timestamp) > 0 {
		ts = timestamp[0]
	}
	var t = time.Unix(ts, 0)
	y := strconv.Itoa(t.Year())
	m := fmt.Sprintf("%02d", t.Month())
	d := fmt.Sprintf("%02d", t.Day())
	h := fmt.Sprintf("%02d", t.Hour())
	i := fmt.Sprintf("%02d", t.Minute())
	s := fmt.Sprintf("%02d", t.Second())
	format = strings.Replace(format, "Y", y, -1)
	format = strings.Replace(format, "m", m, -1)
	format = strings.Replace(format, "d", d, -1)
	format = strings.Replace(format, "H", h, -1)
	format = strings.Replace(format, "i", i, -1)
	format = strings.Replace(format, "s", s, -1)
	return format
}
