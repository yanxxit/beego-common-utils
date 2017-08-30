package common

import (
	"github.com/shawflying/beego-common-utils/utils/logger"
	"github.com/shawflying/beego-common-utils/utils/comutil"
	"github.com/shawflying/beego-common-utils/open"
)

func (this *CommonController) ResV2Result(error_code int, error_msg string, data interface{}) {
	var result map[string]interface{} = make(map[string]interface{})
	result["error_code"] = error_code
	result["error_msg"] = error_msg
	result["data"] = data
	if open.IsOpenResponseLog {
		logger.Info(" ResData:", comutil.MapToJsonString(result))
	}

	this.Data["json"] = result
	this.ServeJSON()
	this.StopRun()
}

//返回通用成功数据
func (this *CommonController) ResV2Success(data interface{}) {
	var success map[string]interface{} = make(map[string]interface{})
	success["error_code"] = 200
	success["error_msg"] = ""
	success["data"] = data
	if open.IsOpenResponseLog {
		logger.Info(" ResData:", comutil.MapToJsonString(success))
	}

	this.Data["json"] = success
	this.ServeJSON()
	this.StopRun()
}

//失败
func (this *CommonController) ResV2Fail(desc string) {
	var fail map[string]interface{} = make(map[string]interface{})
	fail["error_code"] = 0
	fail["error_msg"] = ""
	fail["data"] = desc
	if open.IsOpenResponseLog {
		logger.Info(" ResData:", comutil.MapToJsonString(desc))
	}

	this.Data["json"] = fail
	this.ServeJSON()
	this.StopRun()
}
