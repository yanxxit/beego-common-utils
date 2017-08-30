package common

import (
	"github.com/astaxie/beego"
	"github.com/shawflying/beego-common-utils/utils/logger"
	"github.com/shawflying/beego-common-utils/utils/comutil"
	"github.com/shawflying/beego-common-utils/open"
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

//通用json数据响应
func (this *CommonController) ResJson(data interface{}) {
	if open.IsOpenResponseLog {
		logger.Info(" ResData:", comutil.InterfaceToJSON(data))
	}
	this.Data["json"] = comutil.InterfaceTo2(data)
	this.ServeJSON()
	this.StopRun()
}
