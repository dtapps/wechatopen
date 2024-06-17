package wechatopen

import (
	"context"
	"go.dtapp.net/gorequest"
	"net/http"
)

type WxaSetDefaultamsInfoAgencySetCustomShareRatioResponse struct {
	Ret    int    `json:"ret"`
	ErrMsg string `json:"err_msg,omitempty"`
}

type WxaSetDefaultamsInfoAgencySetCustomShareRatioResult struct {
	Result WxaSetDefaultamsInfoAgencySetCustomShareRatioResponse // 结果
	Body   []byte                                                // 内容
	Http   gorequest.Response                                    // 请求
}

func newWxaSetDefaultamsInfoAgencySetCustomShareRatioResult(result WxaSetDefaultamsInfoAgencySetCustomShareRatioResponse, body []byte, http gorequest.Response) *WxaSetDefaultamsInfoAgencySetCustomShareRatioResult {
	return &WxaSetDefaultamsInfoAgencySetCustomShareRatioResult{Result: result, Body: body, Http: http}
}

// WxaSetDefaultamsInfoAgencySetCustomShareRatio
// 设置自定义分账比例
// https://developers.weixin.qq.com/doc/oplatform/openApi/OpenApiDoc/ams/percentage/SetCustomShareRatio.html
func (c *Client) WxaSetDefaultamsInfoAgencySetCustomShareRatio(ctx context.Context, authorizerAccessToken string, appid string, shareRatio int64, notMustParams ...gorequest.Params) (*WxaSetDefaultamsInfoAgencySetCustomShareRatioResult, error) {

	// OpenTelemetry链路追踪
	ctx, span := TraceStartSpan(ctx, "wxa/setdefaultamsinfo")
	defer span.End()

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	params.Set("appid", appid)
	params.Set("share_ratio", shareRatio)

	// 请求
	var response WxaSetDefaultamsInfoAgencySetCustomShareRatioResponse
	request, err := c.request(ctx, span, "wxa/setdefaultamsinfo?action=agency_set_custom_share_ratio&access_token="+authorizerAccessToken, params, http.MethodPost, &response)
	return newWxaSetDefaultamsInfoAgencySetCustomShareRatioResult(response, request.ResponseBody, request), err
}

// ErrcodeInfo 错误描述
func (resp *WxaSetDefaultamsInfoAgencySetCustomShareRatioResult) ErrcodeInfo() string {
	switch resp.Result.Ret {
	case -202:
		return "内部错误"
	case 1700:
		return "参数错误"
	case 1701:
		return "参数错误"
	case 1737:
		return "操作过快"
	case 2056:
		return "服务商未在变现专区开通账户"
	default:
		return resp.Result.ErrMsg
	}
}
