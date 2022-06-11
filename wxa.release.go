package wechatopen

import (
	"encoding/json"
	"fmt"
	"go.dtapp.net/gorequest"
	"net/http"
)

type WxaReleaseResponse struct {
	Errcode int    `json:"errcode"` // 错误码
	Errmsg  string `json:"errmsg"`  // 错误信息
}

type WxaReleaseResult struct {
	Result WxaReleaseResponse // 结果
	Body   []byte             // 内容
	Http   gorequest.Response // 请求
	Err    error              // 错误
}

func NewWxaReleaseResult(result WxaReleaseResponse, body []byte, http gorequest.Response, err error) *WxaReleaseResult {
	return &WxaReleaseResult{Result: result, Body: body, Http: http, Err: err}
}

// WxaRelease 发布已通过审核的小程序
// https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/code/release.html
func (app *App) WxaRelease() *WxaReleaseResult {
	// 参数
	params := NewParams()
	// 请求
	request, err := app.request(fmt.Sprintf("https://api.weixin.qq.com/wxa/release?access_token=%s", app.GetAuthorizerAccessToken()), params, http.MethodPost)
	// 定义
	var response WxaReleaseResponse
	err = json.Unmarshal(request.ResponseBody, &response)
	return NewWxaReleaseResult(response, request.ResponseBody, request, err)
}

// ErrcodeInfo 错误描述
func (resp *WxaReleaseResult) ErrcodeInfo() string {
	switch resp.Result.Errcode {
	case 85019:
		return "没有审核版本"
	case 85020:
		return "审核状态未满足发布"
	}
	return "系统繁忙"
}
