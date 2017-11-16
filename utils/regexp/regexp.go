package regexp

import (
	"regexp"
	"fmt"
)

func IsMobile(mobile string) bool {
	var validID = regexp.MustCompile(`^[1][123456789]\d{9}$`)
	return validID.MatchString(mobile)
}

func IsOpenid(openid string) bool {
	var validID = regexp.MustCompile(`^[\x00-\x7F]{28}$`)
	return validID.MatchString(openid)
}

func IsEmail(str string) bool {
	var validID = regexp.MustCompile(`^[-_A-Za-z0-9]+@([_A-Za-z0-9]+\.)+[A-Za-z0-9]{2,3}$`)
	return validID.MatchString(str)
}

func FormatNumber(str string) string {
	return str[0:3] + "****" + str[7:len(str)]
}

//还有些问题
func FormatName(str string) string {
	a := str[0:1]
	b := len(str)
	c := ""
	for i := 1; i < b; i ++ {
		fmt.Println("11")
		c += "*"
	}
	return a + c
}

//判断设备号运营商
func Operators(mobile string) string {
	if IsMobile(mobile) {
		return "非手机号码"
	}
	if OperatorsMap(mobile[0:3]) != "" {
		return OperatorsMap(mobile[0:3])
	} else if (OperatorsMap(mobile[0:4]) != "") {
		return OperatorsMap(mobile[0:4])
	}
	return "其他"
}

//设备运营商列表
func OperatorsMap(id string) string {
	mymap := make(map[string]string)
	//中国电信
	mymap["133"] = "中国电信"
	mymap["149"] = "中国电信"
	mymap["153"] = "中国电信"
	mymap["173"] = "中国电信"
	mymap["177"] = "中国电信"
	mymap["180"] = "中国电信"
	mymap["181"] = "中国电信"
	mymap["189"] = "中国电信"
	mymap["199"] = "中国电信"

	//中国联通
	mymap["130"] = "中国联通"
	mymap["131"] = "中国联通"
	mymap["132"] = "中国联通"
	mymap["155"] = "中国联通"
	mymap["156"] = "中国联通"
	mymap["145"] = "中国联通"
	mymap["175"] = "中国联通"
	mymap["176"] = "中国联通"
	mymap["185"] = "中国联通"
	mymap["186"] = "中国联通"
	mymap["166"] = "中国联通"

	//中国移动
	mymap["135"] = "中国移动"
	mymap["136"] = "中国移动"
	mymap["137"] = "中国移动"
	mymap["138"] = "中国移动"
	mymap["139"] = "中国移动"
	mymap["147"] = "中国移动"
	mymap["150"] = "中国移动"
	mymap["151"] = "中国移动"
	mymap["152"] = "中国移动"
	mymap["157"] = "中国移动"
	mymap["158"] = "中国移动"
	mymap["159"] = "中国移动"
	mymap["178"] = "中国移动"
	mymap["182"] = "中国移动"
	mymap["183"] = "中国移动"
	mymap["184"] = "中国移动"
	mymap["187"] = "中国移动"
	mymap["188"] = "中国移动"
	mymap["198"] = "中国移动"

	//134(0-8)
	mymap["1340"] = "中国移动"
	mymap["1341"] = "中国移动"
	mymap["1342"] = "中国移动"
	mymap["1343"] = "中国移动"
	mymap["1344"] = "中国移动"
	mymap["1345"] = "中国移动"
	mymap["1346"] = "中国移动"
	mymap["1347"] = "中国移动"
	mymap["1348"] = "中国移动"

	//虚拟运营商
	mymap["1700"] = "中国电信"
	mymap["1701"] = "中国电信"
	mymap["1702"] = "中国电信"

	mymap["1703"] = "中国移动"
	mymap["1705"] = "中国移动"
	mymap["1706"] = "中国移动"

	mymap["1704"] = "中国联通"
	mymap["1707"] = "中国联通"
	mymap["1708"] = "中国联通"
	mymap["1709"] = "中国联通"
	mymap["171"] = "中国联通"

	//其他
	mymap["1349"] = "中国移动"
	//物联网网号
	mymap["1410"] = "中国电信"

	mymap["148"] = "中国移动"
	mymap["1440"] = "中国移动"

	mymap["146"] = "中国联通"
	//卫星中国移动通信业务号
	mymap["1740"] = "中国电信"
	mymap["1741"] = "中国电信"
	mymap["1742"] = "中国电信"
	mymap["1743"] = "中国电信"
	mymap["1744"] = "中国电信"
	mymap["1745"] = "中国电信"
	return mymap[id]
}
