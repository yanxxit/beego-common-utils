package comutil

import (
	"encoding/json"
	"time"
	"beego-common-utils/utils/timeutil"
	"bytes"
	"strconv"
	"github.com/shawflying/beego-common-utils/utils/comutil"
)

//interface 转换为string类型
func TransInterfaceToString(e interface{}) string {
	var str string
	switch v := e.(type) {
	case int:
		str = string(v)
	case string:
		str = v
	}
	return str
}

//将字符串转换为json
func TransStrToJSON(str string) (error, interface{}) {
	var params interface{}
	err := json.Unmarshal([]byte(str), &params)
	return err, params
}

func TransBytesToJSON(str []byte) (error, interface{}) {
	var params interface{}
	err := json.Unmarshal(str, &params)
	return err, params
}

func GetPayid() string {
	time.Sleep(1)
	timeUnix := time.Now().Unix() //已知的时间戳
	formatTimeStr := time.Unix(timeUnix, 0).Format("20060102150405")
	return formatTimeStr + timeutil.RandomString(10)
}

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

//interface 转换为json 任意类型转换为json
func InterfaceToJSON(data interface{}) string {
	content, _ := json.Marshal(data)
	return StringsToJSON(string(content))
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

//map 转换为 json string  JSON.stringify
//{"data":{"age":12,"name":"inter"},"detail":{"Admin":"admin","Pwd":"12312312"},"info":"","status":0}
func MapToJsonString(data interface{}) string {
	b, _ := json.Marshal(data)
	return string(b)
}

//map 转换为json format
//{
//	"data": {
//		"age": 12,
//		"name": "inter"
//	},
//	"detail": {
//		"Admin": "admin",
//		"Pwd": "12312312"
//	},
//	"info": "",
//	"status": 0
//}
func MapToJsonFormat(data interface{}) string {
	content, _ := json.MarshalIndent(data, "", "  ")
	return string(content)
}
