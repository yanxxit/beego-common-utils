package httputil

import (
	"github.com/astaxie/beego/httplib"
	"time"
	"beego-common-utils/utils/comutil"
	"beego-common-utils/utils/logger"
	"strconv"
	"github.com/astaxie/beego"
	"encoding/json"
)

//返回的结果 为[]byte  error
//入参 reqUrl string + map 或者结构体
var httpCT time.Duration
var httpRWT time.Duration

func init() {
	httpConnectTimeout, _ := strconv.Atoi(beego.AppConfig.String("httpConnectTimeout")) //time.Duration
	httpReadWriteTimeout, _ := strconv.Atoi(beego.AppConfig.String("httpReadWriteTimeout"))
	httpCT = time.Duration(httpConnectTimeout) * time.Second
	httpRWT = time.Duration(httpReadWriteTimeout) * time.Second
}

func getResponseTime(start time.Time) time.Duration {
	end := time.Now()
	return end.Sub(start)
}

func Get(reqUrl string) (content []byte, err error) {
	start := time.Now()
	logger.Info("request-get-url: " + reqUrl)
	req := httplib.Get(reqUrl).SetTimeout(httpCT, httpRWT)
	data, err := req.Bytes()
	logger.Info("request-get-data: "+string(data), err)
	logger.Info("request-get-response-time:", getResponseTime(start))
	return data, err
}

func Post(reqUrl string, params interface{}, token string) (content []byte, err error) {
	start := time.Now()
	logger.Info("request-post-url: " + reqUrl)
	logger.Info("request-post-params: ", params)
	req, err := httplib.Post(reqUrl).Header("Content-Type", "application/json").Header("token", token).SetTimeout(httpCT, httpRWT).JSONBody(params)
	data, err := req.Bytes()
	logger.Info("request-post-data: "+string(data), err)
	logger.Info("request-post-response-time:", getResponseTime(start))
	return data, err
}

func PostForm(reqUrl string, params interface{}) (content []byte, err error) {
	start := time.Now()
	logger.Info("request-postForm-url: " + reqUrl)
	logger.Info("request-postForm-params: ", params)

	byts, err := json.Marshal(params)
	if err != nil {

	}
	mapParams := make(map[string]interface{});
	json.Unmarshal(byts, &mapParams)

	req := httplib.Post(reqUrl).SetTimeout(httpCT, httpRWT)
	for k, v := range mapParams {
		req.Param(k, comutil.TransInterfaceToString(v))
	}
	data, err := req.Bytes()
	logger.Info("request-postForm-data: "+string(data), err)
	logger.Info("request-postForm-response-time:", getResponseTime(start))
	return data, err
}

func Put(reqUrl string, params interface{}) (content []byte, err error) {
	logger.Info("request-put-url: " + reqUrl)
	logger.Info("request-put-params: ", params)
	start := time.Now()
	req, err := httplib.Put(reqUrl).SetTimeout(httpCT, httpRWT).JSONBody(params)
	data, err := req.Bytes()
	strs, _ := req.String()
	logger.Info("request-put-string: ", strs, err)
	logger.Info("request-put-data: "+string(data), err)
	logger.Info("request-put-response-time:", getResponseTime(start))
	return data, err
}
