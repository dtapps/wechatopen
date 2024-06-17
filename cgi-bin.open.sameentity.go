package wechatopen

import (
	"context"
	"go.dtapp.net/gorequest"
	"net/http"
)

type CgiBinOpenSameEnTityResponse struct {
	Errcode    int    `json:"errcode"`
	Errmsg     string `json:"errmsg"`
	SameEntity bool   `json:"same_entity"` // 是否同主体；true表示同主体；false表示不同主体
}

type CgiBinOpenSameEnTityResult struct {
	Result CgiBinOpenSameEnTityResponse // 结果
	Body   []byte                       // 内容
	Http   gorequest.Response           // 请求
}

func newCgiBinOpenSameEnTityResult(result CgiBinOpenSameEnTityResponse, body []byte, http gorequest.Response) *CgiBinOpenSameEnTityResult {
	return &CgiBinOpenSameEnTityResult{Result: result, Body: body, Http: http}
}

// CgiBinOpenSameEnTity 获取授权绑定的商户号列表
// https://developers.weixin.qq.com/doc/oplatform/openApi/OpenApiDoc/cloudbase-common/wechatpay/getWechatPayList.html
func (c *Client) CgiBinOpenSameEnTity(ctx context.Context, componentAccessToken string, notMustParams ...gorequest.Params) (*CgiBinOpenSameEnTityResult, error) {

	// OpenTelemetry链路追踪
	ctx, span := TraceStartSpan(ctx, "cgi-bin/open/sameentity")
	defer span.End()

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)

	// 请求
	var response CgiBinOpenSameEnTityResponse
	request, err := c.request(ctx, span, "cgi-bin/open/sameentity?access_token="+componentAccessToken, params, http.MethodGet, &response)
	return newCgiBinOpenSameEnTityResult(response, request.ResponseBody, request), err
}
