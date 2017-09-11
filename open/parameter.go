package open

import "time"

var Cookie string //cookie

var Header map[string](string) //header

var Token string //token

var Uuid = ""

var SessionName = ""

var IsOpenResponseLog = false //是否开启响应日志

//http 超时时间
var HttpCT time.Duration = 5 * time.Second
var HttpRWT time.Duration = 5 * time.Second
