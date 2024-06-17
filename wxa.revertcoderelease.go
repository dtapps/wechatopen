package wechatopen

import (
	"context"
	"go.dtapp.net/gorequest"
	"net/http"
)

type WxaRevertCodeReleaseResponse struct {
	Errcode     int    `json:"errcode"` // 错误码
	Errmsg      string `json:"errmsg"`  // 错误信息
	VersionList []struct {
		CommitTime  int    `json:"commit_time"`  // 更新时间，时间戳
		UserVersion string `json:"user_version"` // 模板版本号，开发者自定义字段
		UserDesc    string `json:"user_desc"`    // 模板描述，开发者自定义字段
		AppVersion  int    `json:"app_version"`  // 小程序版本
	} `json:"version_list"` // 模板信息列表
}

type WxaRevertCodeReleaseResult struct {
	Result WxaRevertCodeReleaseResponse // 结果
	Body   []byte                       // 内容
	Http   gorequest.Response           // 请求
}

func newWxaRevertCodeReleaseResult(result WxaRevertCodeReleaseResponse, body []byte, http gorequest.Response) *WxaRevertCodeReleaseResult {
	return &WxaRevertCodeReleaseResult{Result: result, Body: body, Http: http}
}

// WxaRevertCodeRelease 小程序版本回退
// https://developers.weixin.qq.com/doc/oplatform/openApi/OpenApiDoc/miniprogram-management/code-management/revertCodeRelease.html
func (c *Client) WxaRevertCodeRelease(ctx context.Context, authorizerAccessToken string, notMustParams ...gorequest.Params) (*WxaRevertCodeReleaseResult, error) {

	// OpenTelemetry链路追踪
	ctx, span := TraceStartSpan(ctx, "wxa/revertcoderelease")
	defer span.End()

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)

	// 请求
	var response WxaRevertCodeReleaseResponse
	request, err := c.request(ctx, span, "wxa/revertcoderelease?access_token="+authorizerAccessToken, params, http.MethodGet, &response)
	return newWxaRevertCodeReleaseResult(response, request.ResponseBody, request), err
}

// ErrcodeInfo 错误描述
func (resp *WxaRevertCodeReleaseResult) ErrcodeInfo() string {
	switch resp.Result.Errcode {
	case 40001:
		return "获取 access_token 时 AppSecret 错误，或者 access_token 无效。请开发者认真比对 AppSecret 的正确性，或查看是否正在为恰当的公众号调用接口"
	default:
		return resp.Result.Errmsg
	}
}
