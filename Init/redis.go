//   File Name:  redis.go
//    Description:
//    Author:      Chenghu
//    Date:       2024/6/13 15:51
//    Change Activity:

package Init

import (
	"context"
	"github.com/redis/go-redis/v9"
)

var RedisClient *redis.Client

type RedisConfig struct {
	Host        string `json:"host"`
	Port        string `json:"port"`
	Password    string `json:"password"`
	Db          int    `json:"db"`
	MaxIdle     int    `json:"maxIdle"`
	IdleTimeout int    `json:"idleTimeout"`
	PoolSize    int    `json:"PoolSize"`
}

func initRedis(cf RedisConfig) *redis.Client {
	addr := cf.Host + ":" + cf.Port
	redisOpt := &redis.Options{
		Addr:         addr,
		Password:     cf.Password,
		DB:           cf.Db,
		PoolSize:     cf.PoolSize,
		MaxIdleConns: cf.MaxIdle,
		MinIdleConns: cf.MaxIdle,
	}
	RedisClient = redis.NewClient(redisOpt)
	cmd := RedisClient.Ping(context.Background())
	if cmd.Err() != nil {
		panic("redis connect fail")
	}

	return RedisClient
}
