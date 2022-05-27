package db

import (
	"fmt"
	"github.com/go-redis/redis"
	"mall-demo/micro-mall-search/conf"
	"time"
)

var DefaultRedisOption conf.RedisOption

func init() {
	DefaultRedisOption.Address = "127.0.0.1:6399"
	DefaultRedisOption.DB = 0
	DefaultRedisOption.Password = ""
}

func GetRedisInstance(redisOpt conf.RedisOption) *redis.Client {
	address := redisOpt.Address
	db := redisOpt.DB
	password := redisOpt.Password
	addr := fmt.Sprintf("%s", address)
	client := redis.NewClient(&redis.Options{
		Addr:       addr,
		Password:   password,
		DB:         db,
		MaxConnAge: 20 * time.Second,
	})
	return client
}
