package wechatopen

import (
	"go.dtapp.net/dorm"
	"go.dtapp.net/golog"
	"go.dtapp.net/gorequest"
)

type ConfigClient struct {
	ComponentAccessToken   string // 第三方平台 access_token
	ComponentVerifyTicket  string // 微信后台推送的 ticket
	PreAuthCode            string // 预授权码
	AuthorizerAccessToken  string // 接口调用令牌
	AuthorizerRefreshToken string // 刷新令牌
	AuthorizerAppid        string // 授权方 appid
	ComponentAppId         string // 第三方平台 appid
	ComponentAppSecret     string // 第三方平台 app_secret
	MessageToken           string
	MessageKey             string
	RedisClient            *dorm.RedisClient // 缓存数据库
	GormClient             *dorm.GormClient  // 日志数据库
	LogClient              *golog.ZapLog     // 日志驱动
	LogDebug               bool              // 日志开关
}

// Client 微信公众号服务
type Client struct {
	requestClient *gorequest.App    // 请求服务
	redisClient   *dorm.RedisClient // 缓存服务
	logClient     *golog.ApiClient  // 日志服务
	config        *ConfigClient     // 配置
}

func NewClient(config *ConfigClient) (*Client, error) {

	var err error
	c := &Client{config: config}

	c.requestClient = gorequest.NewHttp()

	c.redisClient = config.RedisClient

	if c.config.GormClient.Db != nil {
		c.logClient, err = golog.NewApiClient(&golog.ApiClientConfig{
			GormClient: c.config.GormClient,
			TableName:  logTable,
			LogClient:  c.config.LogClient,
			LogDebug:   c.config.LogDebug,
		})
		if err != nil {
			return nil, err
		}
	}

	return c, nil
}

// ConfigComponent 配置
func (c *Client) ConfigComponent(componentAppId, componentAppSecret string) *Client {
	c.config.ComponentAppId = componentAppId
	c.config.ComponentAppSecret = componentAppSecret
	return c
}

// ConfigAuthorizer 配置第三方
func (c *Client) ConfigAuthorizer(authorizerAppid string) *Client {
	c.config.AuthorizerAppid = authorizerAppid
	return c
}
