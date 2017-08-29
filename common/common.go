package common

import (
	"github.com/astaxie/beego"
	"beego-common-utils/utils/restful"
	"beego-common-utils/utils/logger"
	"encoding/json"
	"beego-common-utils/utils/comutil"
)

type CommonController struct {
	beego.Controller
}

//获取接口中的cookieid
func (this *CommonController) GetCookieId() string {
	cookieValue, err := this.Ctx.Request.Cookie(beego.AppConfig.String("sessionname"))
	cookiev := ""
	if err == nil {
		cookiev = cookieValue.Value
	}
	return cookiev
}

//将接口转换为json
func (this *CommonController) InterfaceToJSON(data interface{}) string {
	content, _ := json.Marshal(data)
	return comutil.StringsToJSON(string(content))
}

//通用json数据响应
func (this *CommonController) ResJson(data interface{}) {
	logger.Info(" ResData:", comutil.InterfaceToJSON(data))
	this.Data["json"] = comutil.InterfaceTo2(data)
	this.ServeJSON()
	this.StopRun()
}

func (this *CommonController) ResResult(status int, info string, data interface{}) {
	logger.Info(" ResData:", comutil.InterfaceToJSON(data))
	var result map[string]interface{} = make(map[string]interface{})
	result["status"] = status
	result["info"] = info
	result["data"] = data
	this.Data["json"] = result
	this.ServeJSON()
	this.StopRun()
}

//返回通用成功数据
func (this *CommonController) ResSuccess(data interface{}) {
	logger.Info(" ResData:", comutil.InterfaceToJSON(data))
	var success map[string]interface{} = make(map[string]interface{})
	success["status"] = 200
	success["info"] = ""
	success["data"] = data
	this.Data["json"] = success
	this.ServeJSON()
	this.StopRun()
}

//失败
func (this *CommonController) ResFail(desc string) {
	logger.Info(" ResData:", desc)
	this.Data["json"] = restful.Fail(desc)
	this.ServeJSON()
	this.StopRun()
}
