package wechatopen

import (
	"encoding/json"
	"fmt"
	"go.dtapp.net/gorequest"
	"net/http"
)

type WxaUnbindTesterResponse struct {
	Errcode int    `json:"errcode"` // 错误码
	Errmsg  string `json:"errmsg"`  // 错误信息
}

type WxaUnbindTesterResult struct {
	Result WxaUnbindTesterResponse // 结果
	Body   []byte                  // 内容
	Http   gorequest.Response      // 请求
	Err    error                   // 错误
}

func NewWxaUnbindTesterResult(result WxaUnbindTesterResponse, body []byte, http gorequest.Response, err error) *WxaUnbindTesterResult {
	return &WxaUnbindTesterResult{Result: result, Body: body, Http: http, Err: err}
}

// WxaUnbindTester 解除绑定体验者
// https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/Mini_Program_AdminManagement/unbind_tester.html
func (app *App) WxaUnbindTester(wechatid, userstr string) *WxaUnbindTesterResult {
	// 参数
	params := NewParams()
	if wechatid != "" {
		params["wechatid"] = wechatid
	}
	params["userstr"] = userstr
	// 请求
	request, err := app.request(fmt.Sprintf("https://api.weixin.qq.com/wxa/unbind_tester?access_token=%s", app.GetAuthorizerAccessToken()), params, http.MethodPost)
	// 定义
	var response WxaUnbindTesterResponse
	err = json.Unmarshal(request.ResponseBody, &response)
	return NewWxaUnbindTesterResult(response, request.ResponseBody, request, err)
}
