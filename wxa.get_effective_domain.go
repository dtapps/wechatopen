package wechatopen

import (
	"encoding/json"
	"fmt"
	"go.dtapp.net/gorequest"
	"net/http"
)

type WxaGetEffectiveDomainResponse struct {
	Errcode  int    `json:"errcode"`
	Errmsg   string `json:"errmsg"`
	MpDomain struct {
		Requestdomain   []interface{} `json:"requestdomain"`
		Wsrequestdomain []interface{} `json:"wsrequestdomain"`
		Uploaddomain    []interface{} `json:"uploaddomain"`
		Downloaddomain  []interface{} `json:"downloaddomain"`
		Udpdomain       []interface{} `json:"udpdomain"`
		Tcpdomain       []interface{} `json:"tcpdomain"`
	} `json:"mp_domain"`
	ThirdDomain struct {
		Requestdomain   []interface{} `json:"requestdomain"`
		Wsrequestdomain []interface{} `json:"wsrequestdomain"`
		Uploaddomain    []interface{} `json:"uploaddomain"`
		Downloaddomain  []interface{} `json:"downloaddomain"`
		Udpdomain       []interface{} `json:"udpdomain"`
		Tcpdomain       []interface{} `json:"tcpdomain"`
	} `json:"third_domain"`
	DirectDomain struct {
		Requestdomain   []interface{} `json:"requestdomain"`
		Wsrequestdomain []interface{} `json:"wsrequestdomain"`
		Uploaddomain    []interface{} `json:"uploaddomain"`
		Downloaddomain  []interface{} `json:"downloaddomain"`
		Udpdomain       []interface{} `json:"udpdomain"`
		Tcpdomain       []interface{} `json:"tcpdomain"`
	} `json:"direct_domain"`
}

type WxaGetEffectiveDomainResult struct {
	Result WxaGetEffectiveDomainResponse // 结果
	Body   []byte                        // 内容
	Http   gorequest.Response            // 请求
	Err    error                         // 错误
}

func newWxaGetEffectiveDomainResult(result WxaGetEffectiveDomainResponse, body []byte, http gorequest.Response, err error) *WxaGetEffectiveDomainResult {
	return &WxaGetEffectiveDomainResult{Result: result, Body: body, Http: http, Err: err}
}

// WxaGetEffectiveDomain 获取发布后生效服务器域名列表
// https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/Mini_Program_Basic_Info/get_effective_domain.html
func (c *Client) WxaGetEffectiveDomain(notMustParams ...gorequest.Params) *WxaGetEffectiveDomainResult {
	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	// 请求
	request, err := c.request(fmt.Sprintf(apiUrl+"/wxa/get_effective_domain?access_token=%s", c.GetAuthorizerAccessToken()), params, http.MethodPost)
	// 定义
	var response WxaGetEffectiveDomainResponse
	err = json.Unmarshal(request.ResponseBody, &response)
	return newWxaGetEffectiveDomainResult(response, request.ResponseBody, request, err)
}
