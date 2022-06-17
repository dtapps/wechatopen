package goredis

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-redis/redis/v8"
	"time"
)

// ConfigClient 配置
type ConfigClient struct {
	Addr     string // 地址
	Password string // 密码
	DB       int    // 数据库
	PoolSize int    // 连接池大小
}

type Client struct {
	Db     *redis.Client
	config *ConfigClient
}

func NewClient(config *ConfigClient) *Client {

	client := &Client{config: config}

	if config.PoolSize == 0 {
		config.PoolSize = 100
	}

	client.Db = redis.NewClient(&redis.Options{
		Addr:     config.Addr,     // 地址
		Password: config.Password, // 密码
		DB:       config.DB,       // 数据库
		PoolSize: config.PoolSize, // 连接池大小
	})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := client.Db.Ping(ctx).Result()
	if err != nil {
		panic(errors.New(fmt.Sprintf("redis连接失败：%v", err)))
	}

	return client
}
