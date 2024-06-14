package wechatopen

import (
	"context"
	"go.dtapp.net/gorequest"
	"net/http"
)

type WxaSecurityApplyPrivacyInterfaceResponse struct {
	Errcode int    `json:"errcode"`  // 返回码
	Errmsg  string `json:"errmsg"`   // 返回码信息
	AuditId int64  `json:"audit_id"` // 审核单id
}

type WxaSecurityApplyPrivacyInterfaceResult struct {
	Result WxaSecurityApplyPrivacyInterfaceResponse // 结果
	Body   []byte                                   // 内容
	Http   gorequest.Response                       // 请求
}

func newWxaSecurityApplyPrivacyInterfaceResult(result WxaSecurityApplyPrivacyInterfaceResponse, body []byte, http gorequest.Response) *WxaSecurityApplyPrivacyInterfaceResult {
	return &WxaSecurityApplyPrivacyInterfaceResult{Result: result, Body: body, Http: http}
}

// WxaSecurityApplyPrivacyInterface 申请接口
// https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/apply_api/apply_privacy_interface.html
func (c *Client) WxaSecurityApplyPrivacyInterface(ctx context.Context, authorizerAccessToken string, notMustParams ...gorequest.Params) (*WxaSecurityApplyPrivacyInterfaceResult, error) {

	// OpenTelemetry链路追踪
	ctx = c.TraceStartSpan(ctx, "wxa/security/apply_privacy_interface")
	defer c.TraceEndSpan()

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)

	// 请求
	var response WxaSecurityApplyPrivacyInterfaceResponse
	request, err := c.request(ctx, "wxa/security/apply_privacy_interface?access_token="+authorizerAccessToken, params, http.MethodPost, &response)
	return newWxaSecurityApplyPrivacyInterfaceResult(response, request.ResponseBody, request), err
}

// ErrcodeInfo 错误描述
func (resp *WxaSecurityApplyPrivacyInterfaceResult) ErrcodeInfo() string {
	switch resp.Result.Errcode {
	case 61031:
		return "审核中，请不要重复申请"
	case 61032:
		return "视频格式不对，要传mp4格式的"
	case 61033:
		return "视频下载失败"
	case 61034:
		return "必填的参数没填，检查后重新提交"
	case 61035:
		return "输入的api（api_name严格区分大小写）无需申请，可以直接使用"
	case 61036:
		return "该帐号不可申请，请检查类目是否符合"
	case 61037:
		return "需要以ntf-8的编码格式提交"
	default:
		return resp.Result.Errmsg
	}
}
