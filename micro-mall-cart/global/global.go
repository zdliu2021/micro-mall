package global

import (
	"github.com/go-redis/redis"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"mall-demo/micro-mall-cart/conf"
)

var (
	GVA_LOG      *zap.Logger
	GVA_CONFIG   *conf.Server
	GVA_REDIS    *redis.Client
	SmsMysqlConn *gorm.DB
)
