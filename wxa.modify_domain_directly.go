package wechatopen

import (
	"context"
	"encoding/json"
	"fmt"
	"go.dtapp.net/gorequest"
	"net/http"
)

type WxaModifyDomainDirectlyResponse struct {
	Errcode int    `json:"errcode"` // 错误码
	Errmsg  string `json:"errmsg"`  // 错误信息
}

type WxaModifyDomainDirectlyResult struct {
	Result WxaModifyDomainDirectlyResponse // 结果
	Body   []byte                          // 内容
	Http   gorequest.Response              // 请求
}

func newWxaModifyDomainDirectlyResult(result WxaModifyDomainDirectlyResponse, body []byte, http gorequest.Response) *WxaModifyDomainDirectlyResult {
	return &WxaModifyDomainDirectlyResult{Result: result, Body: body, Http: http}
}

// WxaModifyDomainDirectly 快速设置小程序服务器域名
// https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/Mini_Program_Basic_Info/modify_domain_directly.html
func (c *Client) WxaModifyDomainDirectly(ctx context.Context, notMustParams ...gorequest.Params) (*WxaModifyDomainDirectlyResult, error) {
	// 检查
	err := c.checkComponentIsConfig()
	if err != nil {
		return nil, err
	}
	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	// 请求
	request, err := c.request(ctx, fmt.Sprintf(apiUrl+"/wxa/modify_domain_directly?access_token=%s", c.GetAuthorizerAccessToken(ctx)), params, http.MethodPost)
	if err != nil {
		return nil, err
	}
	// 定义
	var response WxaModifyDomainDirectlyResponse
	err = json.Unmarshal(request.ResponseBody, &response)
	if err != nil {
		return nil, err
	}
	return newWxaModifyDomainDirectlyResult(response, request.ResponseBody, request), nil
}

// ErrcodeInfo 错误描述
func (resp *WxaModifyDomainDirectlyResult) ErrcodeInfo() string {
	switch resp.Result.Errcode {
	case 85015:
		return "该账号不是小程序账号"
	case 86100:
		return "该 URL 的协议头有误"
	case 45082:
		return "域名需要 icp 备案，否则无法添加"
	case 86101:
		return "不支持配置api.weixin.qq.com"
	case 85016:
		return "域名数量超限制"
	case 86102:
		return "每个月只能修改50次，超过域名修改次数限制"
	}
	return "系统繁忙"
}
