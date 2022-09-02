package wechatopen

import (
	"context"
	"encoding/base64"
	"encoding/xml"
	"errors"
	"fmt"
	"strings"
)

type SignDecryptParams struct {
	Signature    string `json:"signature"`
	Timestamp    string `json:"timestamp"`
	Nonce        string `json:"nonce"`
	EncryptType  string `json:"encrypt_type"`
	MsgSignature string `json:"msg_signature"`
	AppId        string `json:"app_id"`
	Encrypt      string `json:"encrypt"`
}

// SignDecrypt 解密
func (c *Client) SignDecrypt(ctx context.Context, params SignDecryptParams, strXml interface{}) ([]byte, error) {

	if params.Signature == "" {
		return nil, errors.New("找不到签名参数")
	}

	if params.Timestamp == "" {
		return nil, errors.New("找不到时间戳参数")
	}

	if params.Nonce == "" {
		return nil, errors.New("未找到随机数参数")
	}

	wantSignature := Sign(c.GetMessageToken(), params.Timestamp, params.Nonce)
	if params.Signature != wantSignature {
		return nil, errors.New("签名错误")
	}

	// 进入事件执行
	if params.EncryptType != "aes" {
		return nil, errors.New("未知的加密类型: " + params.EncryptType)
	}
	if params.Encrypt == "" {
		return nil, errors.New("找不到签名参数")
	}

	cipherData, err := base64.StdEncoding.DecodeString(params.Encrypt)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Encrypt 解码字符串错误：%v", err))
	}

	AesKey, err := base64.StdEncoding.DecodeString(c.GetMessageKey() + "=")
	if err != nil {
		return nil, errors.New(fmt.Sprintf("messageKey 解码字符串错误：%v", err))
	}

	msg, err := AesDecrypt(cipherData, AesKey)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("AES解密错误：%v", err))
	}

	str := string(msg)

	left := strings.Index(str, "<xml>")
	if left <= 0 {
		return nil, errors.New(fmt.Sprintf("匹配不到<xml>：%v", left))
	}
	right := strings.Index(str, "</xml>")
	if right <= 0 {
		return nil, errors.New(fmt.Sprintf("匹配不到</xml>：%v", right))
	}
	msgStr := str[left:right]
	if len(msgStr) == 0 {
		return nil, errors.New(fmt.Sprintf("提取错误：%v", msgStr))
	}

	strByte := []byte(msgStr + "</xml>")
	err = xml.Unmarshal(strByte, strXml)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("解析错误：%v", err))
	}

	return strByte, nil
}
