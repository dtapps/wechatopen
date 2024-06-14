package wechatopen

import (
	"context"
	"errors"
	"go.opentelemetry.io/otel/codes"
	"net/http"
)

// ServeHttpAuthorizerAppid 授权跳转
func (c *Client) ServeHttpAuthorizerAppid(ctx context.Context, w http.ResponseWriter, r *http.Request, componentAccessToken string) (resp CgiBinComponentApiQueryAuthResponse, agentUserId string, err error) {

	// OpenTelemetry链路追踪
	ctx = c.TraceStartSpan(ctx, "ServeHttpAuthorizerAppid")
	defer c.TraceEndSpan()

	var (
		query = r.URL.Query()

		authCode  = query.Get("auth_code")
		expiresIn = query.Get("expires_in")
	)

	agentUserId = query.Get("agent_user_id")

	if authCode == "" {
		err = errors.New("找不到授权码参数")
		c.TraceRecordError(err)
		c.TraceSetStatus(codes.Error, err.Error())
		return resp, agentUserId, err
	}

	if expiresIn == "" {
		err = errors.New("找不到过期时间参数")
		c.TraceRecordError(err)
		c.TraceSetStatus(codes.Error, err.Error())
		return resp, agentUserId, err
	}

	info, err := c.CgiBinComponentApiQueryAuth(ctx, componentAccessToken, authCode)
	if err != nil {
		c.TraceRecordError(err)
		c.TraceSetStatus(codes.Error, err.Error())
		return resp, agentUserId, err
	}
	if info.Result.AuthorizationInfo.AuthorizerAppid == "" {
		err = errors.New("获取失败")
		c.TraceRecordError(err)
		c.TraceSetStatus(codes.Error, err.Error())
		return resp, agentUserId, err
	}

	return info.Result, agentUserId, nil
}
