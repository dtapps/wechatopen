package wechatopen

import (
	"context"
	"github.com/gin-gonic/gin"
)

// ResponseServeHttpGin 推送信息
type ResponseServeHttpGin struct {
	MsgSignature string `form:"msg_signature" json:"msg_signature" xml:"msg_signature" uri:"msg_signature" binding:"omitempty"` // 签名串，对应 URL 参数的msg_signature
	Timestamp    string `form:"timestamp" json:"timestamp" xml:"timestamp" uri:"timestamp" binding:"omitempty"`                 // 时间戳，对应 URL 参数的timestamp
	Nonce        string `form:"nonce" json:"nonce" xml:"nonce" uri:"nonce" binding:"omitempty"`                                 // 随机串，对应 URL 参数的nonce
	Signature    string `form:"signature" json:"signature" xml:"signature" uri:"signature" binding:"omitempty"`
	EncryptType  string `form:"encrypt_type" json:"encrypt_type" xml:"encrypt_type" uri:"encrypt_type" binding:"omitempty"` // 加密类型
	AppId        string `form:"AppId" json:"AppId" xml:"AppId" uri:"AppId" binding:"omitempty"`                             // 第三方平台 appid
	Encrypt      string `form:"Encrypt" json:"Encrypt" xml:"Encrypt" uri:"Encrypt" binding:"omitempty"`                     // 加密内容
}

// ServeHttpGin 验证票据推送
func (c *Client) ServeHttpGin(ctx context.Context, ginCtx *gin.Context) (ResponseServeHttpGin, error) {

	query := ginCtx.Request.URL.Query()

	// 声明接收的变量
	var validateJson struct {
		AppId   string `form:"AppId" json:"AppId" xml:"AppId" uri:"AppId" binding:"omitempty"`         // 第三方平台 appid
		Encrypt string `form:"Encrypt" json:"Encrypt" xml:"Encrypt" uri:"Encrypt" binding:"omitempty"` // 加密内容
	}

	err := ginCtx.ShouldBind(&validateJson)

	return ResponseServeHttpGin{
		MsgSignature: query.Get("msg_signature"),
		Timestamp:    query.Get("timestamp"),
		Nonce:        query.Get("nonce"),
		Signature:    query.Get("signature"),
		EncryptType:  query.Get("encrypt_type"),
		AppId:        validateJson.AppId,
		Encrypt:      validateJson.Encrypt,
	}, err
}
