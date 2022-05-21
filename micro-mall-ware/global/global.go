package global

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
	"mall-demo/micro-mall-ware/conf"
)

var (
	GVA_LOG      *zap.Logger
	GVA_CONFIG   *conf.Server
	WmsMysqlConn *gorm.DB
)
