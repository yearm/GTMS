package helper

import (
	"GTMS/library/stringi"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/json-iterator/go"
	"golang.org/x/crypto/bcrypt"
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

func MustMarshal(v interface{}) []byte {
	b, _ := jsoniter.Marshal(v)
	return b
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

//转驼峰
func ToCamel(s string) string {
	b := []byte(s)
	if b[0] >= 'A' && b[0] <= 'Z' {
		b[0] += 32
	}
	return string(b)
}

//base64加密
func Base64Encode(str string) string {
	var in = []byte(str)
	return base64.StdEncoding.EncodeToString(in)
}

//base64解密
func Base64Decode(str string) string {
	bytes, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		return ""
	}
	return string(bytes)
}

//bcrypt加密密码
func HashedPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

//bcrypt校验密码
func CheckHashedPassword(hashedPassword string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
