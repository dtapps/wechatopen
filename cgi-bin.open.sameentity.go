package wechatopen

import (
	"context"
	"go.dtapp.net/gojson"
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
	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	// 请求
	request, err := c.request(ctx, apiUrl+"/cgi-bin/open/sameentity?access_token="+componentAccessToken, params, http.MethodGet)
	if err != nil {
		return newCgiBinOpenSameEnTityResult(CgiBinOpenSameEnTityResponse{}, request.ResponseBody, request), err
	}
	// 定义
	var response CgiBinOpenSameEnTityResponse
	err = gojson.Unmarshal(request.ResponseBody, &response)
	return newCgiBinOpenSameEnTityResult(response, request.ResponseBody, request), err
}
