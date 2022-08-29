package wechatopen

import (
	"context"
	"encoding/json"
	"fmt"
	"go.dtapp.net/gorequest"
	"net/http"
)

type CgiBinComponentApiQueryAuthResponse struct {
	AuthorizationInfo struct {
		AuthorizerAppid        string `json:"authorizer_appid"`         // 授权方 appid
		AuthorizerAccessToken  string `json:"authorizer_access_token"`  // 接口调用令牌（在授权的公众号/小程序具备 API 权限时，才有此返回值）
		ExpiresIn              int64  `json:"expires_in"`               // authorizer_access_token 的有效期（在授权的公众号/小程序具备API权限时，才有此返回值），单位：秒
		AuthorizerRefreshToken string `json:"authorizer_refresh_token"` // 刷新令牌（在授权的公众号具备API权限时，才有此返回值），刷新令牌主要用于第三方平台获取和刷新已授权用户的 authorizer_access_token。一旦丢失，只能让用户重新授权，才能再次拿到新的刷新令牌。用户重新授权后，之前的刷新令牌会失效
		FuncInfo               []struct {
			FuncscopeCategory struct {
				Id int `json:"id"`
			} `json:"funcscope_category"`
			ConfirmInfo struct {
				NeedConfirm    int `json:"need_confirm"`
				AlreadyConfirm int `json:"already_confirm"`
				CanConfirm     int `json:"can_confirm"`
			} `json:"confirm_info,omitempty"`
		} `json:"func_info"`
	} `json:"authorization_info"`
}

type CgiBinComponentApiQueryAuthResult struct {
	Result CgiBinComponentApiQueryAuthResponse // 结果
	Body   []byte                              // 内容
	Http   gorequest.Response                  // 请求
}

func newCgiBinComponentApiQueryAuthResult(result CgiBinComponentApiQueryAuthResponse, body []byte, http gorequest.Response) *CgiBinComponentApiQueryAuthResult {
	return &CgiBinComponentApiQueryAuthResult{Result: result, Body: body, Http: http}
}

// CgiBinComponentApiQueryAuth 使用授权码获取授权信息
// https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/ThirdParty/token/authorization_info.html
func (c *Client) CgiBinComponentApiQueryAuth(ctx context.Context, authorizationCode string) (*CgiBinComponentApiQueryAuthResult, error) {
	// 检查
	err := c.checkComponentIsConfig()
	if err != nil {
		return nil, err
	}
	// 参数
	param := gorequest.NewParams()
	param["component_appid"] = c.GetComponentAppId() // 第三方平台 appid
	param["authorization_code"] = authorizationCode  // 授权码, 会在授权成功时返回给第三方平台
	params := gorequest.NewParamsWith(param)
	// 请求
	request, err := c.request(ctx, fmt.Sprintf(apiUrl+"/cgi-bin/component/api_query_auth?component_access_token=%v", c.GetComponentAccessToken(ctx)), params, http.MethodPost)
	if err != nil {
		return nil, err
	}
	// 定义
	var response CgiBinComponentApiQueryAuthResponse
	err = json.Unmarshal(request.ResponseBody, &response)
	if err != nil {
		return nil, err
	}
	return newCgiBinComponentApiQueryAuthResult(response, request.ResponseBody, request), nil
}
