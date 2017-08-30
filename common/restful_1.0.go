package common

import (
	"github.com/shawflying/beego-common-utils/utils/logger"
	"github.com/shawflying/beego-common-utils/utils/comutil"
	"github.com/shawflying/beego-common-utils/open"
)

func (this *CommonController) ResResult(status int, info string, data interface{}) {
	var result map[string]interface{} = make(map[string]interface{})
	result["status"] = status
	result["info"] = info
	result["data"] = data
	if open.IsOpenResponseLog {
		logger.Info(" ResData:", comutil.MapToJsonString(result))
	}
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

	if open.IsOpenResponseLog {
		logger.Info(" ResData:", comutil.MapToJsonString(success))
	}
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

	if open.IsOpenResponseLog {
		logger.Info(" ResData:", comutil.MapToJsonString(fail))
	}
	this.Data["json"] = fail
	this.ServeJSON()
	this.StopRun()
}
