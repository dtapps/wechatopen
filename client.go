package wechatopen

import (
	"go.dtapp.net/dorm"
	"go.dtapp.net/golog"
	"go.dtapp.net/gorequest"
)

// client *dorm.GormClient
type gormClientFun func() *dorm.GormClient

// client *dorm.MongoClient
// databaseName string
type mongoClientFun func() (*dorm.MongoClient, string)

// ClientConfig 实例配置
type ClientConfig struct {
	AuthorizerAppid    string // 授权方 appid
	ComponentAppId     string // 第三方平台 appid
	ComponentAppSecret string // 第三方平台 app_secret
	MessageToken       string
	MessageKey         string
	RedisClient        *dorm.RedisClient // 缓存数据库
	GormClientFun      gormClientFun     // 日志配置
	MongoClientFun     mongoClientFun    // 日志配置
	Debug              bool              // 日志开关
}

// Client 实例
type Client struct {
	requestClient *gorequest.App // 请求服务
	config        struct {
		componentAccessToken   string // 第三方平台 access_token
		componentVerifyTicket  string // 微信后台推送的 ticket
		preAuthCode            string // 预授权码
		authorizerAccessToken  string // 接口调用令牌
		authorizerRefreshToken string // 刷新令牌
		authorizerAppid        string // 授权方 appid
		componentAppId         string // 第三方平台appid
		componentAppSecret     string // 第三方平台app_secret
		messageToken           string
		messageKey             string
	}
	cache struct {
		redisClient *dorm.RedisClient // 缓存数据库
	}
	log struct {
		gorm           bool              // 日志开关
		gormClient     *dorm.GormClient  // 日志数据库
		logGormClient  *golog.ApiClient  // 日志服务
		mongo          bool              // 日志开关
		mongoClient    *dorm.MongoClient // 日志数据库
		logMongoClient *golog.ApiClient  // 日志服务
	}
}

// NewClient 创建实例化
func NewClient(config *ClientConfig) (*Client, error) {

	var err error
	c := &Client{}

	c.config.componentAppId = config.ComponentAppId
	c.config.componentAppSecret = config.ComponentAppSecret
	c.config.messageToken = config.MessageToken
	c.config.messageKey = config.MessageKey

	c.requestClient = gorequest.NewHttp()

	gormClient := config.GormClientFun()
	if gormClient != nil && gormClient.Db != nil {
		c.log.logGormClient, err = golog.NewApiGormClient(func() (*dorm.GormClient, string) {
			return gormClient, logTable
		}, config.Debug)
		if err != nil {
			return nil, err
		}
		c.log.gorm = true
		c.log.gormClient = gormClient
	}

	mongoClient, databaseName := config.MongoClientFun()
	if mongoClient != nil && mongoClient.Db != nil {
		c.log.logMongoClient, err = golog.NewApiMongoClient(func() (*dorm.MongoClient, string, string) {
			return mongoClient, databaseName, logTable
		}, config.Debug)
		if err != nil {
			return nil, err
		}
		c.log.mongo = true
		c.log.mongoClient = mongoClient
	}

	c.cache.redisClient = config.RedisClient

	return c, nil
}
