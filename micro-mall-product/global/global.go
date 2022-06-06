package global

import (
	"github.com/go-redis/redis"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"mall-demo/micro-mall-product/conf"
)

var (
	GVA_LOG      *zap.Logger
	GVA_CONFIG   *conf.Server
	GVA_REDIS    *redis.Client
	PmsMysqlConn *gorm.DB
	OmsMysqlConn *gorm.DB
	SmsMysqlConn *gorm.DB
	UmsMysqlConn *gorm.DB
	WmsMysqlConn *gorm.DB
)
