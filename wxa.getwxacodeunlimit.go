package wechatopen

import (
	"context"
	"encoding/json"
	"fmt"
	"go.dtapp.net/gorequest"
	"net/http"
)

type WxaGetWxaCodeUnLimitResponse struct {
	Errcode     int         `json:"errcode"`
	Errmsg      string      `json:"errmsg"`
	ContentType string      `json:"contentType"`
	Buffer      interface{} `json:"buffer"`
}

type WxaGetWxaCodeUnLimitResult struct {
	Result WxaGetWxaCodeUnLimitResponse // 结果
	Body   []byte                       // 内容
	Http   gorequest.Response           // 请求
}

func newWxaGetWxaCodeUnLimitResult(result WxaGetWxaCodeUnLimitResponse, body []byte, http gorequest.Response) *WxaGetWxaCodeUnLimitResult {
	return &WxaGetWxaCodeUnLimitResult{Result: result, Body: body, Http: http}
}

// WxaGetWxaCodeUnLimit 获取小程序码，适用于需要的码数量极多的业务场景。通过该接口生成的小程序码，永久有效，数量暂无限制
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/qr-code/wxacode.getUnlimited.html
func (c *Client) WxaGetWxaCodeUnLimit(ctx context.Context, notMustParams ...gorequest.Params) (*WxaGetWxaCodeUnLimitResult, error) {
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
	params := gorequest.NewParamsWith(notMustParams...)
	// 请求
	request, err := c.request(ctx, fmt.Sprintf(apiUrl+"/wxa/getwxacodeunlimit?access_token=%s", c.GetAuthorizerAccessToken(ctx)), params, http.MethodPost)
	if err != nil {
		return nil, err
	}
	// 定义
	var response WxaGetWxaCodeUnLimitResponse
	// 判断内容是否为图片
	if request.HeaderIsImg() {
	} else {
		err = json.Unmarshal(request.ResponseBody, &response)
		if err != nil {
			return nil, err
		}
	}
	return newWxaGetWxaCodeUnLimitResult(response, request.ResponseBody, request), nil
}

// ErrcodeInfo 错误描述
func (resp *WxaGetWxaCodeUnLimitResult) ErrcodeInfo() string {
	switch resp.Result.Errcode {
	case 45009:
		return "调用分钟频率受限(目前5000次/分钟，会调整)，如需大量小程序码，建议预生成"
	case 41030:
		return "page 不合法（页面不存在或者小程序没有发布、根路径前加 /或者携带参数）"
	case 40097:
		return "env_version 不合法"
	}
	return "系统繁忙"
}
