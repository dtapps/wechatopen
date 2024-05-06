package wechatopen

import (
	"context"
	"go.dtapp.net/gojson"
	"go.dtapp.net/gorequest"
	"net/http"
)

type WxaGetDefaultamsInfoGetShareRatioResponse struct {
	Ret    int    `json:"ret"`
	ErrMsg string `json:"err_msg"`
}

type WxaGetDefaultamsInfoGetShareRatioResult struct {
	Result WxaGetDefaultamsInfoGetShareRatioResponse // 结果
	Body   []byte                                    // 内容
	Http   gorequest.Response                        // 请求
}

func newWxaGetDefaultamsInfoGetShareRatioResult(result WxaGetDefaultamsInfoGetShareRatioResponse, body []byte, http gorequest.Response) *WxaGetDefaultamsInfoGetShareRatioResult {
	return &WxaGetDefaultamsInfoGetShareRatioResult{Result: result, Body: body, Http: http}
}

// WxaGetDefaultamsInfoGetShareRatio 查询分账比例
// https://developers.weixin.qq.com/doc/oplatform/openApi/OpenApiDoc/ams/percentage/GetShareRatio.html
func (c *Client) WxaGetDefaultamsInfoGetShareRatio(ctx context.Context, authorizerAppid, authorizerAccessToken string, notMustParams ...gorequest.Params) (*WxaGetDefaultamsInfoGetShareRatioResult, error) {
	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	params.Set("appid", authorizerAppid)
	// 请求
	request, err := c.request(ctx, apiUrl+"/wxa/getdefaultamsinfo?action=get_share_ratio&access_token="+authorizerAccessToken, params, http.MethodPost)
	if err != nil {
		return newWxaGetDefaultamsInfoGetShareRatioResult(WxaGetDefaultamsInfoGetShareRatioResponse{}, request.ResponseBody, request), err
	}
	// 定义
	var response WxaGetDefaultamsInfoGetShareRatioResponse
	err = gojson.Unmarshal(request.ResponseBody, &response)
	return newWxaGetDefaultamsInfoGetShareRatioResult(response, request.ResponseBody, request), err
}

// ErrcodeInfo 错误描述
func (resp *WxaGetDefaultamsInfoGetShareRatioResult) ErrcodeInfo() string {
	switch resp.Result.Ret {
	case -202:
		return "内部错误"
	case 1700:
		return "参数错误"
	case 1701:
		return "参数错误"
	case 1735:
		return "商户未完成协议签署流程"
	case 1737:
		return "操作过快"
	case 2056:
		return "服务商未在变现专区开通账户"
	default:
		return resp.Result.ErrMsg
	}
}
