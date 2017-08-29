package common

import (
	"beego-common-utils/utils/logger"
	"beego-common-utils/utils/comutil"
)

func (this *CommonController) ResResult(status int, info string, data interface{}) {
	var result map[string]interface{} = make(map[string]interface{})
	result["status"] = status
	result["info"] = info
	result["data"] = data
	logger.Info(" ResData:", comutil.MapToJsonString(result))
	this.Data["json"] = result
	this.ServeJSON()
	this.StopRun()
}

//返回通用成功数据
func (this *CommonController) ResSuccess(data interface{}) {
	var success map[string]interface{} = make(map[string]interface{})
	success["status"] = 200
	success["info"] = ""
	success["data"] = data
	logger.Info(" ResData:", comutil.MapToJsonString(success))
	this.Data["json"] = success
	this.ServeJSON()
	this.StopRun()
}

//失败
func (this *CommonController) ResFail(desc string) {
	var fail map[string]interface{} = make(map[string]interface{})
	fail["status"] = 0
	fail["info"] = ""
	fail["data"] = desc
	logger.Info(" ResData:", comutil.MapToJsonString(desc))
	this.Data["json"] = fail
	this.ServeJSON()
	this.StopRun()
}
