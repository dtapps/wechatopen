package wechatopen

import (
	"encoding/json"
	"fmt"
	"go.dtapp.net/gorequest"
	"net/http"
)

type CgiBinWxOpenQrCodeJumpGetResponse struct {
	Errcode  int    `json:"errcode"`
	Errmsg   string `json:"errmsg"`
	RuleList []struct {
		Prefix        string   `json:"prefix"`          // 二维码规则
		PermitSubRule int      `json:"permit_sub_rule"` // 是否独占符合二维码前缀匹配规则的所有子规 1 为不占用，2 为占用
		Path          string   `json:"path"`            // 小程序功能页面
		OpenVersion   int      `json:"open_version"`    // 测试范围
		DebugUrl      []string `json:"debug_url"`       // 测试链接（选填）可填写不多于 5 个用于测试的二维码完整链接，此链接必须符合已填写的二维码规则。
		State         int      `json:"state"`           // 发布标志位，1 表示未发布，2 表示已发布
	} `json:"rule_list"` // 二维码规则详情列表
	QrcodejumpOpen     int `json:"qrcodejump_open"`      // 是否已经打开二维码跳转链接设置
	ListSize           int `json:"list_size"`            // 二维码规则数量
	QrcodejumpPubQuota int `json:"qrcodejump_pub_quota"` // 本月还可发布的次数
}

type CgiBinWxOpenQrCodeJumpGetResult struct {
	Result CgiBinWxOpenQrCodeJumpGetResponse // 结果
	Body   []byte                            // 内容
	Http   gorequest.Response                // 请求
	Err    error                             // 错误
}

func NewCgiBinWxOpenQrCodeJumpGetResult(result CgiBinWxOpenQrCodeJumpGetResponse, body []byte, http gorequest.Response, err error) *CgiBinWxOpenQrCodeJumpGetResult {
	return &CgiBinWxOpenQrCodeJumpGetResult{Result: result, Body: body, Http: http, Err: err}
}

// CgiBinWxOpenQrCodeJumpGet 获取已设置的二维码规则
// https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/qrcode/qrcodejumpadd.html
func (app *App) CgiBinWxOpenQrCodeJumpGet() *CgiBinWxOpenQrCodeJumpGetResult {
	// 参数
	params := NewParams()
	// 请求
	request, err := app.request(fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/wxopen/qrcodejumpget?access_token=%s", app.GetAuthorizerAccessToken()), params, http.MethodPost)
	// 定义
	var response CgiBinWxOpenQrCodeJumpGetResponse
	err = json.Unmarshal(request.ResponseBody, &response)
	return NewCgiBinWxOpenQrCodeJumpGetResult(response, request.ResponseBody, request, err)
}
