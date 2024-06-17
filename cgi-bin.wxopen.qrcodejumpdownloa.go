package wechatopen

import (
	"context"
	"go.dtapp.net/gorequest"
	"net/http"
)

type CgiBinWxOpenQrCodeJumpDownloadResponse struct {
	Errcode     int    `json:"errcode"`
	Errmsg      string `json:"errmsg"`
	FileName    string `json:"file_name"`
	FileContent string `json:"file_content"`
}

type CgiBinWxOpenQrCodeJumpDownloadResult struct {
	Result CgiBinWxOpenQrCodeJumpDownloadResponse // 结果
	Body   []byte                                 // 内容
	Http   gorequest.Response                     // 请求
}

func newCgiBinWxOpenQrCodeJumpDownloadResult(result CgiBinWxOpenQrCodeJumpDownloadResponse, body []byte, http gorequest.Response) *CgiBinWxOpenQrCodeJumpDownloadResult {
	return &CgiBinWxOpenQrCodeJumpDownloadResult{Result: result, Body: body, Http: http}
}

// CgiBinWxOpenQrCodeJumpDownload 获取校验文件名称及内容
// https://developers.weixin.qq.com/doc/oplatform/openApi/OpenApiDoc/miniprogram-management/jumpqrcode-config/downloadQRCodeText.html
func (c *Client) CgiBinWxOpenQrCodeJumpDownload(ctx context.Context, authorizerAccessToken string, notMustParams ...gorequest.Params) (*CgiBinWxOpenQrCodeJumpDownloadResult, error) {

	// OpenTelemetry链路追踪
	ctx, span := TraceStartSpan(ctx, "cgi-bin/wxopen/qrcodejumpdownload")
	defer span.End()

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)

	// 请求
	var response CgiBinWxOpenQrCodeJumpDownloadResponse
	request, err := c.request(ctx, span, "cgi-bin/wxopen/qrcodejumpdownload?access_token="+authorizerAccessToken, params, http.MethodPost, &response)
	return newCgiBinWxOpenQrCodeJumpDownloadResult(response, request.ResponseBody, request), err
}
