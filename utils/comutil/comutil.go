package comutil

import (
	"encoding/json"
	"time"
	"github.com/shawflying/beego-common-utils/utils/timeutil"
	"bytes"
	"strconv"
	"crypto/md5"
	"fmt"
)

//interface 转换为string类型
func TransInterfaceToString(data interface{}) string {
	var str string
	switch v := data.(type) {
	case int64:
		str = strconv.FormatInt(v, 10)
	case int:
		str = strconv.Itoa(v)
	case string:
		str = v
	case []byte:
		str = string(v)
	default:
		t, err := json.Marshal(data)
		if err != nil {
			str = "数据类型未识别"
		}
		str = string(t)
	}
	return str
}

//生成流水号
func GetPayid() string {
	time.Sleep(1)
	timeUnix := time.Now().Unix() //已知的时间戳
	formatTimeStr := time.Unix(timeUnix, 0).Format("20060102150405")
	return formatTimeStr + timeutil.RandomString(10)
}

//string 转json
func StringsToJSON(str string) string {
	var jsons bytes.Buffer
	for _, r := range str {
		rint := int(r)
		if rint < 128 {
			jsons.WriteRune(r)
		} else {
			jsons.WriteString("\\u")
			jsons.WriteString(strconv.FormatInt(int64(rint), 16))
		}
	}
	return jsons.String()
}

//将 string []byte interface
func InterfaceTo2(data interface{}) (interface{}) {
	var params interface{}
	switch v := data.(type) {
	case int:
		json.Unmarshal([]byte(string(v)), &params)
	case string:
		json.Unmarshal([]byte(v), &params)
	case []byte:
		json.Unmarshal(v, &params)
	default:
		params = data
	}
	return params
}

//map 转换为json format
func MapToJsonFormat(data interface{}) string {
	content, _ := json.MarshalIndent(data, "", "  ")
	return string(content)
}

//md5 string 转换为MD5
func MD5Encode(value string) string {
	data := []byte(value)
	has := md5.Sum(data)
	md5str1 := fmt.Sprintf("%x", has) //将[]byte转成16进制
	return md5str1
}

//判断数据是否为空
//1. 为空的口径是什么
//2. 判断那些类型
//1. string 
//2. []byte
//3. map
//
//注：数字类型，bool 都为非空
func IsEmpty(data interface{}) bool {
	empty := false
	str := "";
	switch v := data.(type) {

	case bool:
		empty = false

	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, uintptr:
		empty = false
	case float32, float64, complex64, complex128:
		empty = false
	case string:
		if len(v) == 0 || v == "null" || v == "" || v == "undefined" {
			empty = true
		}
	case []byte:
		str = string(v)
		if len(str) == 0 || str == "null" || str == "" || str == "undefined" {
			empty = true
		}
	case map[interface{}]interface{}:
		if len(v) == 0 {
			empty = true
		}
	default:
		t, err := json.Marshal(data)
		if err != nil {
			str = ""
		}
		str = string(t)
		if len(str) == 0 || str == "null" || str == "" || str == "undefined" {
			empty = true
		}
	}
	return empty
}
