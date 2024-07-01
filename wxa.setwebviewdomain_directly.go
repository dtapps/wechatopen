package wechatopen

import (
	"context"
	"go.dtapp.net/gorequest"
	"net/http"
)

type modifyJumpDomainDirectlyResponse struct {
	Errcode int    `json:"errcode"` // 错误码
	Errmsg  string `json:"errmsg"`  // 错误信息
}

type modifyJumpDomainDirectlyResult struct {
	Result modifyJumpDomainDirectlyResponse // 结果
	Body   []byte                           // 内容
	Http   gorequest.Response               // 请求
}

func newmodifyJumpDomainDirectlyResult(result modifyJumpDomainDirectlyResponse, body []byte, http gorequest.Response) *modifyJumpDomainDirectlyResult {
	return &modifyJumpDomainDirectlyResult{Result: result, Body: body, Http: http}
}

// modifyJumpDomainDirectly 快速配置小程序业务域名
// https://developers.weixin.qq.com/doc/oplatform/openApi/OpenApiDoc/miniprogram-management/domain-management/modifyJumpDomainDirectly.html
func (c *Client) modifyJumpDomainDirectly(ctx context.Context, authorizerAccessToken string, action string, webviewdomain []string, notMustParams ...gorequest.Params) (*modifyJumpDomainDirectlyResult, error) {

	// OpenTelemetry链路追踪
	ctx, span := TraceStartSpan(ctx, "wxa/setwebviewdomain_directly")
	defer span.End()

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	params.Set("action", action)
	params.Set("webviewdomain", webviewdomain)

	// 请求
	var response modifyJumpDomainDirectlyResponse
	request, err := c.request(ctx, span, "wxa/setwebviewdomain_directly?access_token="+authorizerAccessToken, params, http.MethodPost, &response)
	return newmodifyJumpDomainDirectlyResult(response, request.ResponseBody, request), err
}

// ErrcodeInfo 错误描述
func (resp *modifyJumpDomainDirectlyResult) ErrcodeInfo() string {
	switch resp.Result.Errcode {
	case 86103:
		return "check confirmfile fail! 检查检验文件失败"
	case 506015:
		return "域名绑定的小程序超出上限"
	default:
		return resp.Result.Errmsg
	}
}
