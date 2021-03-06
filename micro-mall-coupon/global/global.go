package global

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
	"mall-demo/micro-mall-coupon/conf"
)

var (
	GVA_LOG      *zap.Logger
	GVA_CONFIG   *conf.Server
	SmsMysqlConn *gorm.DB
)
