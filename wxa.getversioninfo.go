package wechatopen

import (
	"context"
	"encoding/json"
	"fmt"
	"go.dtapp.net/gorequest"
	"net/http"
)

type WxaGetVersionInfoResponse struct {
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
	ExpInfo struct {
		ExpTime    int64  `json:"exp_time"`    // 提交体验版的时间
		ExpVersion string `json:"exp_version"` // 体验版版本信息
		ExpDesc    string `json:"exp_desc"`    // 体验版版本描述
	} `json:"exp_info"` // 体验版信息
	ReleaseInfo struct {
		ReleaseTime    int64  `json:"release_time"`    // 发布线上版的时间
		ReleaseVersion string `json:"release_version"` // 线上版版本信息
		ReleaseDesc    string `json:"release_desc"`    // 线上版本描述
	} `json:"release_info"` // 线上版信息
}

type WxaGetVersionInfoResult struct {
	Result WxaGetVersionInfoResponse // 结果
	Body   []byte                    // 内容
	Http   gorequest.Response        // 请求
}

func newWxaGetVersionInfoResult(result WxaGetVersionInfoResponse, body []byte, http gorequest.Response) *WxaGetVersionInfoResult {
	return &WxaGetVersionInfoResult{Result: result, Body: body, Http: http}
}

// WxaGetVersionInfo 查询小程序版本信息
// https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/code/get_versioninfo.html
func (c *Client) WxaGetVersionInfo(ctx context.Context) (*WxaGetVersionInfoResult, error) {
	// 检查
	err := c.checkComponentIsConfig()
	if err != nil {
		return nil, err
	}
	err = c.checkAuthorizerIsConfig()
	if err != nil {
		return nil, err
	}
	// 参数
	params := gorequest.NewParams()
	// 请求
	request, err := c.request(ctx, fmt.Sprintf(apiUrl+"/wxa/getversioninfo?access_token=%s", c.GetAuthorizerAccessToken(ctx)), params, http.MethodPost)
	if err != nil {
		return nil, err
	}
	// 定义
	var response WxaGetVersionInfoResponse
	err = json.Unmarshal(request.ResponseBody, &response)
	if err != nil {
		return nil, err
	}
	return newWxaGetVersionInfoResult(response, request.ResponseBody, request), nil
}
