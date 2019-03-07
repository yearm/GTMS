package helper

import (
	"GTMS/library/stringi"
	"crypto/rand"
	"fmt"
	"github.com/astaxie/beego/logs"
	"math/big"
	"os"
	"path/filepath"
	"strings"
)

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
