package wechatopen

import (
	"context"
	"go.dtapp.net/gorequest"
	"net/http"
)

type ModifyThirdpartyServerDomainResponse struct {
	Errcode                  int    `json:"errcode"`                     // 返回码
	Errmsg                   string `json:"errmsg"`                      // 返回码信息
	PublishedWxaServerDomain string `json:"published_wxa_server_domain"` // 目前生效的 “全网发布版”第三方平台“小程序服务器域名”。如果修改失败，该字段不会返回。如果没有已发布的第三方平台，该字段也不会返回。
	TestingWxaServerDomain   string `json:"testing_wxa_server_domain"`   // 目前生效的 “测试版”第三方平台“小程序服务器域名”。如果修改失败，该字段不会返回
	InvalidWxaServerDomain   string `json:"invalid_wxa_server_domain"`   // 未通过验证的域名。如果不存在未通过验证的域名，该字段不会返回。
}

type ModifyThirdpartyServerDomainResult struct {
	Result ModifyThirdpartyServerDomainResponse // 结果
	Body   []byte                               // 内容
	Http   gorequest.Response                   // 请求
}

func newModifyThirdpartyServerDomainResult(result ModifyThirdpartyServerDomainResponse, body []byte, http gorequest.Response) *ModifyThirdpartyServerDomainResult {
	return &ModifyThirdpartyServerDomainResult{Result: result, Body: body, Http: http}
}

// ModifyThirdpartyServerDomain 设置第三方平台服务器域名
// https://developers.weixin.qq.com/doc/oplatform/openApi/OpenApiDoc/thirdparty-management/domain-mgnt/modifyThirdpartyServerDomain.html
func (c *Client) ModifyThirdpartyServerDomain(ctx context.Context, componentAccessToken string, action string, notMustParams ...gorequest.Params) (*ModifyThirdpartyServerDomainResult, error) {

	// OpenTelemetry链路追踪
	ctx, span := TraceStartSpan(ctx, "cgi-bin/component/modify_wxa_server_domain")
	defer span.End()

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	params.Set("action", action)

	// 请求
	var response ModifyThirdpartyServerDomainResponse
	request, err := c.request(ctx, span, "cgi-bin/component/modify_wxa_server_domain?access_token="+componentAccessToken, params, http.MethodPost, &response)
	return newModifyThirdpartyServerDomainResult(response, request.ResponseBody, request), err
}
