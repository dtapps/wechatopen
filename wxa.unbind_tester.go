package wechatopen

import (
	"context"
	"go.dtapp.net/gojson"
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
}

func newWxaUnbindTesterResult(result WxaUnbindTesterResponse, body []byte, http gorequest.Response) *WxaUnbindTesterResult {
	return &WxaUnbindTesterResult{Result: result, Body: body, Http: http}
}

// WxaUnbindTester 解除绑定体验者
// https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/Mini_Program_AdminManagement/unbind_tester.html
func (c *Client) WxaUnbindTester(ctx context.Context, authorizerAccessToken, wechatid, userstr string, notMustParams ...gorequest.Params) (*WxaUnbindTesterResult, error) {
	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	if wechatid != "" {
		params.Set("wechatid", wechatid)
	}
	params.Set("userstr", userstr)
	// 请求
	request, err := c.request(ctx, apiUrl+"/wxa/unbind_tester?access_token="+authorizerAccessToken, params, http.MethodPost)
	if err != nil {
		return newWxaUnbindTesterResult(WxaUnbindTesterResponse{}, request.ResponseBody, request), err
	}
	// 定义
	var response WxaUnbindTesterResponse
	err = gojson.Unmarshal(request.ResponseBody, &response)
	return newWxaUnbindTesterResult(response, request.ResponseBody, request), err
}
