package wechatopen

import (
	"context"
	"go.dtapp.net/gojson"
	"go.dtapp.net/gorequest"
	"net/http"
)

type WxaApiWxAembeddedGetOwnListResponse struct {
	Errcode         int    `json:"errcode"`       // 错误码
	Errmsg          string `json:"errmsg"`        // 错误信息
	EmbeddedFlag    int    `json:"embedded_flag"` // 授权方式。0表示需要管理员确认，1表示自动通过，2表示自动拒绝
	WxaEmbeddedList []struct {
		Appid       string `json:"appid"`       // 半屏小程序appid
		Create_time int64  `json:"create_time"` // 添加时间
		Headimg     string `json:"headimg"`     // 头像url
		Nickname    string `json:"nickname"`    // 半屏小程序昵称
		Reason      string `json:"reason"`      // 申请理由
		Status      string `json:"status"`      // 申请状态
	} `json:"wxa_embedded_list"` // 半屏小程序列表
}

type WxaApiWxAembeddedGetOwnListResult struct {
	Result WxaApiWxAembeddedGetOwnListResponse // 结果
	Body   []byte                              // 内容
	Http   gorequest.Response                  // 请求
}

func newWxaApiWxAembeddedGetOwnListResult(result WxaApiWxAembeddedGetOwnListResponse, body []byte, http gorequest.Response) *WxaApiWxAembeddedGetOwnListResult {
	return &WxaApiWxAembeddedGetOwnListResult{Result: result, Body: body, Http: http}
}

// WxaApiWxAembeddedGetOwnList 获取半屏小程序授权列表
// https://developers.weixin.qq.com/doc/oplatform/openApi/OpenApiDoc/miniprogram-management/embedded-management/getOwnList.html
func (c *Client) WxaApiWxAembeddedGetOwnList(ctx context.Context, authorizerAccessToken string, notMustParams ...gorequest.Params) (*WxaApiWxAembeddedGetOwnListResult, error) {
	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	// 请求
	request, err := c.request(ctx, apiUrl+"/wxaapi/wxaembedded/get_own_list?access_token="+authorizerAccessToken, params, http.MethodGet)
	if err != nil {
		return newWxaApiWxAembeddedGetOwnListResult(WxaApiWxAembeddedGetOwnListResponse{}, request.ResponseBody, request), err
	}
	// 定义
	var response WxaApiWxAembeddedGetOwnListResponse
	err = gojson.Unmarshal(request.ResponseBody, &response)
	return newWxaApiWxAembeddedGetOwnListResult(response, request.ResponseBody, request), err
}

// ErrcodeInfo 错误描述
func (resp *WxaApiWxAembeddedGetOwnListResult) ErrcodeInfo() string {
	switch resp.Result.Errcode {
	case 89408:
		return "半屏小程序系统错误"
	case 89409:
		return "获取半屏小程序列表参数错误"
	}
	return "系统繁忙"
}
