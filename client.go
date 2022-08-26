package wechatopen

import (
	"go.dtapp.net/dorm"
	"go.dtapp.net/golog"
	"go.dtapp.net/gorequest"
)

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
		gormClient     *dorm.GormClient  // 日志数据库
		gorm           bool              // 日志开关
		logGormClient  *golog.ApiClient  // 日志服务
		mongoClient    *dorm.MongoClient // 日志数据库
		mongo          bool              // 日志开关
		logMongoClient *golog.ApiClient  // 日志服务
	}
}

// client *dorm.GormClient
type gormClientFun func() *dorm.GormClient

// client *dorm.MongoClient
// databaseName string
type mongoClientFun func() (*dorm.MongoClient, string)

// NewClient 创建实例化
// componentAppId 第三方平台appid
// componentAppSecret 第三方平台app_secret
// messageToken
// messageKey
// redisClient 缓存数据库
func NewClient(componentAppId, componentAppSecret, messageToken, messageKey string, redisClient *dorm.RedisClient, gormClientFun gormClientFun, mongoClientFun mongoClientFun, debug bool) (*Client, error) {

	var err error
	c := &Client{}

	c.config.componentAppId = componentAppId
	c.config.componentAppSecret = componentAppSecret
	c.config.messageToken = messageToken
	c.config.messageKey = messageKey

	c.requestClient = gorequest.NewHttp()

	gormClient := gormClientFun()
	if gormClient.Db != nil {
		c.log.logGormClient, err = golog.NewApiGormClient(func() (client *dorm.GormClient, tableName string) {
			return gormClient, logTable
		}, debug)
		if err != nil {
			return nil, err
		}
		c.log.gorm = true
	}
	c.log.gormClient = gormClient

	mongoClient, databaseName := mongoClientFun()
	if mongoClient.Db != nil {
		c.log.logMongoClient, err = golog.NewApiMongoClient(func() (*dorm.MongoClient, string, string) {
			return mongoClient, databaseName, logTable
		}, debug)
		if err != nil {
			return nil, err
		}
		c.log.mongo = true
	}
	c.log.mongoClient = mongoClient

	c.cache.redisClient = redisClient

	return c, nil
}
