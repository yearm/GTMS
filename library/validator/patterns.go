package validator

import (
	"GTMS/library/stringi"
	"github.com/pkg/errors"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

// Basic regular expressions for validating strings
const (
	Email             string = "^(((([a-zA-Z]|\\d|[!#\\$%&'\\*\\+\\-\\/=\\?\\^_`{\\|}~]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])+(\\.([a-zA-Z]|\\d|[!#\\$%&'\\*\\+\\-\\/=\\?\\^_`{\\|}~]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])+)*)|((\\x22)((((\\x20|\\x09)*(\\x0d\\x0a))?(\\x20|\\x09)+)?(([\\x01-\\x08\\x0b\\x0c\\x0e-\\x1f\\x7f]|\\x21|[\\x23-\\x5b]|[\\x5d-\\x7e]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])|(\\([\\x01-\\x09\\x0b\\x0c\\x0d-\\x7f]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}]))))*(((\\x20|\\x09)*(\\x0d\\x0a))?(\\x20|\\x09)+)?(\\x22)))@((([a-zA-Z]|\\d|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])|(([a-zA-Z]|\\d|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])([a-zA-Z]|\\d|-|\\.|_|~|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])*([a-zA-Z]|\\d|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])))\\.)+(([a-zA-Z]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])|(([a-zA-Z]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])([a-zA-Z]|\\d|-|_|~|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])*([a-zA-Z]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])))\\.?$"
	CreditCard        string = "^(?:4[0-9]{12}(?:[0-9]{3})?|5[1-5][0-9]{14}|6(?:011|5[0-9][0-9])[0-9]{12}|3[47][0-9]{13}|3(?:0[0-5]|[68][0-9])[0-9]{11}|(?:2131|1800|35\\d{3})\\d{11})$"
	ISBN10            string = "^(?:[0-9]{9}X|[0-9]{10})$"
	ISBN13            string = "^(?:[0-9]{13})$"
	UUID3             string = "^[0-9a-f]{8}-[0-9a-f]{4}-3[0-9a-f]{3}-[0-9a-f]{4}-[0-9a-f]{12}$"
	UUID4             string = "^[0-9a-f]{8}-[0-9a-f]{4}-4[0-9a-f]{3}-[89ab][0-9a-f]{3}-[0-9a-f]{12}$"
	UUID5             string = "^[0-9a-f]{8}-[0-9a-f]{4}-5[0-9a-f]{3}-[89ab][0-9a-f]{3}-[0-9a-f]{12}$"
	UUID              string = "^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$"
	Alpha             string = "^[a-zA-Z]+$"
	Alphanumeric      string = "^[a-zA-Z0-9]+$"
	Numeric           string = "^[0-9]+$"
	Int               string = "^(?:[-+]?(?:0|[1-9][0-9]*))$"
	Float             string = "^(?:[-+]?(?:[0-9]+))?(?:\\.[0-9]*)?(?:[eE][\\+\\-]?(?:[0-9]+))?$"
	Hexadecimal       string = "^[0-9a-fA-F]+$"
	Hexcolor          string = "^#?([0-9a-fA-F]{3}|[0-9a-fA-F]{6})$"
	RGBcolor          string = "^rgb\\(\\s*(0|[1-9]\\d?|1\\d\\d?|2[0-4]\\d|25[0-5])\\s*,\\s*(0|[1-9]\\d?|1\\d\\d?|2[0-4]\\d|25[0-5])\\s*,\\s*(0|[1-9]\\d?|1\\d\\d?|2[0-4]\\d|25[0-5])\\s*\\)$"
	ASCII             string = "^[\x00-\x7F]+$"
	Multibyte         string = "[^\x00-\x7F]"
	FullWidth         string = "[^\u0020-\u007E\uFF61-\uFF9F\uFFA0-\uFFDC\uFFE8-\uFFEE0-9a-zA-Z]"
	HalfWidth         string = "[\u0020-\u007E\uFF61-\uFF9F\uFFA0-\uFFDC\uFFE8-\uFFEE0-9a-zA-Z]"
	Base64            string = "^(?:[A-Za-z0-9+\\/]{4})*(?:[A-Za-z0-9+\\/]{2}==|[A-Za-z0-9+\\/]{3}=|[A-Za-z0-9+\\/]{4})$"
	PrintableASCII    string = "^[\x20-\x7E]+$"
	DataURI           string = "^data:.+\\/(.+);base64$"
	Latitude          string = "^[-+]?([1-8]?\\d(\\.\\d+)?|90(\\.0+)?)$"
	Longitude         string = "^[-+]?(180(\\.0+)?|((1[0-7]\\d)|([1-9]?\\d))(\\.\\d+)?)$"
	DNSName           string = `^([a-zA-Z0-9_]{1}[a-zA-Z0-9_-]{0,62}){1}(\.[a-zA-Z0-9_]{1}[a-zA-Z0-9_-]{0,62})*[\._]?$`
	URLSchema         string = `((ftp|tcp|udp|wss?|https?):\/\/)`
	URLUsername       string = `(\S+(:\S*)?@)`
	URLPath           string = `((\/|\?|#)[^\s]*)`
	URLPort           string = `(:(\d{1,5}))`
	URLIP             string = `([1-9]\d?|1\d\d|2[01]\d|22[0-3])(\.(1?\d{1,2}|2[0-4]\d|25[0-5])){2}(?:\.([0-9]\d?|1\d\d|2[0-4]\d|25[0-4]))`
	URLSubdomain      string = `((www\.)|([a-zA-Z0-9]+([-_\.]?[a-zA-Z0-9])*[a-zA-Z0-9]\.[a-zA-Z0-9]+))`
	URL               string = `(http|https):\/\/[\w\-_]+(\.[\w\-_]+)+([\w\-\.,@?^=%&:/~\+#]*[\w\-\@?^=%&/~\+#])?`
	SSN               string = `^\d{3}[- ]?\d{2}[- ]?\d{4}$`
	WinPath           string = `^[a-zA-Z]:\\(?:[^\\/:*?"<>|\r\n]+\\)*[^\\/:*?"<>|\r\n]*$`
	UnixPath          string = `^(/[^/\x00]*)+/?$`
	Semver            string = "^v?(?:0|[1-9]\\d*)\\.(?:0|[1-9]\\d*)\\.(?:0|[1-9]\\d*)(-(0|[1-9]\\d*|\\d*[a-zA-Z-][0-9a-zA-Z-]*)(\\.(0|[1-9]\\d*|\\d*[a-zA-Z-][0-9a-zA-Z-]*))*)?(\\+[0-9a-zA-Z-]+(\\.[0-9a-zA-Z-]+)*)?$"
	tagName           string = "valid"
	hasLowerCase      string = ".*[[:lower:]]"
	hasUpperCase      string = ".*[[:upper:]]"
	hasWhitespace     string = ".*[[:space:]]"
	hasWhitespaceOnly string = "^[[:space:]]+$"
	Phone             string = "1[0-9]{10}"
)

// Used by IsFilePath func
const (
	// Unknown is unresolved OS type
	Unknown = iota
	// Win is Windows type
	Win
	// Unix is *nix OS types
	Unix
)

var (
	userRegexp          = regexp.MustCompile("^[a-zA-Z0-9!#$%&'*+/=?^_`{|}~.-]+$")
	hostRegexp          = regexp.MustCompile("^[^\\s]+\\.[^\\s]+$")
	userDotRegexp       = regexp.MustCompile("(^[.]{1})|([.]{1}$)|([.]{2,})")
	rxEmail             = regexp.MustCompile(Email)
	rxCreditCard        = regexp.MustCompile(CreditCard)
	rxISBN10            = regexp.MustCompile(ISBN10)
	rxISBN13            = regexp.MustCompile(ISBN13)
	rxUUID3             = regexp.MustCompile(UUID3)
	rxUUID4             = regexp.MustCompile(UUID4)
	rxUUID5             = regexp.MustCompile(UUID5)
	rxUUID              = regexp.MustCompile(UUID)
	rxAlpha             = regexp.MustCompile(Alpha)
	rxAlphanumeric      = regexp.MustCompile(Alphanumeric)
	rxNumeric           = regexp.MustCompile(Numeric)
	rxInt               = regexp.MustCompile(Int)
	rxFloat             = regexp.MustCompile(Float)
	rxHexadecimal       = regexp.MustCompile(Hexadecimal)
	rxHexcolor          = regexp.MustCompile(Hexcolor)
	rxRGBcolor          = regexp.MustCompile(RGBcolor)
	rxASCII             = regexp.MustCompile(ASCII)
	rxPrintableASCII    = regexp.MustCompile(PrintableASCII)
	rxMultibyte         = regexp.MustCompile(Multibyte)
	rxFullWidth         = regexp.MustCompile(FullWidth)
	rxHalfWidth         = regexp.MustCompile(HalfWidth)
	rxBase64            = regexp.MustCompile(Base64)
	rxDataURI           = regexp.MustCompile(DataURI)
	rxLatitude          = regexp.MustCompile(Latitude)
	rxLongitude         = regexp.MustCompile(Longitude)
	rxDNSName           = regexp.MustCompile(DNSName)
	rxURL               = regexp.MustCompile(URL)
	rxSSN               = regexp.MustCompile(SSN)
	rxWinPath           = regexp.MustCompile(WinPath)
	rxUnixPath          = regexp.MustCompile(UnixPath)
	rxSemver            = regexp.MustCompile(Semver)
	rxHasLowerCase      = regexp.MustCompile(hasLowerCase)
	rxHasUpperCase      = regexp.MustCompile(hasUpperCase)
	rxHasWhitespace     = regexp.MustCompile(hasWhitespace)
	rxHasWhitespaceOnly = regexp.MustCompile(hasWhitespaceOnly)
)

func IsEmail(email string) bool {
	return rxEmail.MatchString(email)
}

func IsNumeric(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}

func IsURL(url string) bool {
	return rxURL.MatchString(url)
}

func IsIP(ip string) bool {
	arr := strings.Split(ip, ".")
	if len(arr) != 4 {
		return false
	}
	for _, item := range arr {
		num, err := strconv.Atoi(item)
		if err != nil {
			return false
		}
		if num < 0 || num > 255 {
			return false
		}
	}
	return true
}

func IsRequired(s string) bool {
	return strings.TrimSpace(s) != ""
}

func min(s string, value interface{}) (bool, error) {
	arr := strings.Split(s, ":")
	limit, err1 := strconv.ParseFloat(arr[1], 64)
	if err1 != nil {
		return false, errors.New(" tag error")
	}

	v, err2 := stringi.ToFloat64(value)
	if err2 != nil {
		return false, errors.New("min" + err2.Error())
	}
	return v >= limit, nil
}

func max(s string, value interface{}) (bool, error) {
	arr := strings.Split(s, ":")
	limit, err1 := strconv.ParseFloat(arr[1], 64)
	if err1 != nil {
		return false, errors.New(" tag error")
	}

	v, err2 := stringi.ToFloat64(value)
	if err2 != nil {
		return false, errors.New("max" + err2.Error())
	}
	return v <= limit, nil
}

func maxLength(s string, value interface{}) (bool, error) {
	arr := strings.Split(s, ":")
	limit, err1 := strconv.Atoi(arr[1])
	if err1 != nil {
		return false, errors.New(" tag error")
	}

	v, ok := value.(string)
	if !ok {
		return false, errors.New("maxLength only support string type")
	}
	return len(v) <= limit, nil
}

func minLength(s string, value interface{}) (bool, error) {
	arr := strings.Split(s, ":")
	limit, err1 := strconv.Atoi(arr[1])
	if err1 != nil {
		return false, errors.New(" tag error")
	}

	v, ok := value.(string)
	if !ok {
		return false, errors.New("minLength only support string type")
	}
	return len(v) >= limit, nil
}

func testSwitch(rule string, value string) (bool, error) {
	arr := strings.Split(rule, ":")
	if len(arr) != 2 {
		return false, errors.New(" tag error")
	}

	items := strings.Split(arr[1], ",")
	for _, item := range items {
		if item == value {
			return true, nil
		}
	}
	return false, nil
}

func minSize(s string, value interface{}) (bool, error) {
	arr := strings.Split(s, ":")
	limit, err1 := strconv.Atoi(arr[1])
	if err1 != nil {
		return false, errors.New(" tag error")
	}

	tp := reflect.TypeOf(value).String()
	if tp != "[]string" && tp != "[]int" && tp != "[]int64" {
		return false, errors.New(" must be slice type")
	}

	if tp == "[]string" {
		sli1, ok1 := value.([]string)
		if ok1 {
			return len(sli1) >= limit, nil
		}
	} else if tp == "[]int" {
		sli2, ok2 := value.([]int)
		if ok2 {
			return len(sli2) >= limit, nil
		}
	}
	sli3, _ := value.([]int64)
	return len(sli3) >= limit, nil
}

func maxSize(s string, value interface{}) (bool, error) {
	arr := strings.Split(s, ":")
	limit, err1 := strconv.Atoi(arr[1])
	if err1 != nil {
		return false, errors.New(" tag error")
	}

	tp := reflect.TypeOf(value).String()
	if tp != "[]string" && tp != "[]int" && tp != "[]int64" {
		return false, errors.New(" must be slice type")
	}

	if tp == "[]string" {
		sli1, ok1 := value.([]string)
		if ok1 {
			return len(sli1) <= limit, nil
		}
	} else if tp == "[]int" {
		sli2, ok2 := value.([]int)
		if ok2 {
			return len(sli2) <= limit, nil
		}
	}
	sli3, _ := value.([]int64)
	return len(sli3) <= limit, nil
}
