package global

import (
	"github.com/go-redis/redis"
	"go.uber.org/zap"
	"mall-demo/micro-mall-auth-server/conf"
)

var (
	GVA_LOG    *zap.Logger
	GVA_CONFIG *conf.Server
	GVA_REDIS  *redis.Client
)

const (
	SMS_CODE_CACHE_PREFIX = "sms:code:"
)
