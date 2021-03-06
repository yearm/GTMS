package helper

import (
	"GTMS/library/stringi"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/json-iterator/go"
	"github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
	"math/big"
	"os"
	"path/filepath"
	"reflect"
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

//创建token
func CreateToken() (token string) {
	token = uuid.NewV4().String()
	return
}

// 结构体转form,去掉为空的字段,格式化key
func StructToFormWithClearNilField(obj interface{}, keys stringi.Form) stringi.Form {
	data := stringi.Form{}
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)
	for i := 0; i < t.NumField(); i++ {
		key := t.Field(i).Name
		val := v.Field(i).String()
		if val != "" {
			data[keys[key]] = v.Field(i).String()
		}
	}
	return data
}

// 判断文件夹、文件是否存在
func FolderExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
