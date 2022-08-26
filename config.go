package wechatopen

// ConfigComponent 配置
func (c *Client) ConfigComponent(componentAppId, componentAppSecret string) *Client {
	c.config.componentAppId = componentAppId
	c.config.componentAppSecret = componentAppSecret
	return c
}

// ConfigAuthorizer 配置第三方
func (c *Client) ConfigAuthorizer(authorizerAppid string) *Client {
	c.config.authorizerAppid = authorizerAppid
	return c
}
