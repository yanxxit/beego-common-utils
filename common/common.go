package common

import (
	"github.com/astaxie/beego"
	"github.com/shawflying/beego-common-utils/utils/logger"
	"github.com/shawflying/beego-common-utils/utils/comutil"
	"github.com/shawflying/beego-common-utils/open"
	"encoding/json"
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
		logger.Info(" ResData:", comutil.TransInterfaceToString(data))
	}
	this.Data["json"] = comutil.InterfaceTo2(data)
	this.ServeJSON()
	this.StopRun()
}

//将post 请求的body 转换为map
func (c *CommonController) RequestBodyToStringMap() (body map[string]string, err error) {
	err = json.Unmarshal(c.Ctx.Input.RequestBody, &body)
	return body, err
}

func (c *CommonController) RequestBodyToInterfaceMap() (body map[string]interface{}, err error) {
	err = json.Unmarshal(c.Ctx.Input.RequestBody, &body)
	return body, err
}
