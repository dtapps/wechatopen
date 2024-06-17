package wechatopen

import (
	"context"
	"go.dtapp.net/gorequest"
	"net/http"
)

type WxaGetUserRiskRankResponse struct {
	Errcode  int    `json:"errcode"`   // 错误码
	Errmsg   string `json:"errmsg"`    // 错误信息
	RiskRank int    `json:"risk_rank"` // 用户风险等级，合法值为0,1,2,3,4，数字越大风险越高。
	UnoinId  int64  `json:"unoin_id"`  // 唯一请求标识，标记单次请求
}

type WxaGetUserRiskRankResult struct {
	Result WxaGetUserRiskRankResponse // 结果
	Body   []byte                     // 内容
	Http   gorequest.Response         // 请求
}

func newWxaGetUserRiskRankResult(result WxaGetUserRiskRankResponse, body []byte, http gorequest.Response) *WxaGetUserRiskRankResult {
	return &WxaGetUserRiskRankResult{Result: result, Body: body, Http: http}
}

// WxaGetUserRiskRank 获取用户安全等级
// https://developers.weixin.qq.com/miniprogram/dev/OpenApiDoc/sec-center/safety-control-capability/getUserRiskRank.html
func (c *Client) WxaGetUserRiskRank(ctx context.Context, authorizerAppid, authorizerAccessToken string, notMustParams ...gorequest.Params) (*WxaGetUserRiskRankResult, error) {

	// OpenTelemetry链路追踪
	ctx, span := TraceStartSpan(ctx, "wxa/getuserriskrank")
	defer span.End()

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	params.Set("appid", authorizerAppid)

	// 请求
	var response WxaGetUserRiskRankResponse
	request, err := c.request(ctx, span, "wxa/getuserriskrank?access_token="+authorizerAccessToken, params, http.MethodPost, &response)
	return newWxaGetUserRiskRankResult(response, request.ResponseBody, request), err
}
