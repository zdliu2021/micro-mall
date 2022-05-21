package global

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
	"mall-demo/micro-mall-product/conf"
)

var (
	GVA_LOG      *zap.Logger
	GVA_CONFIG   *conf.Server
	PmsMysqlConn *gorm.DB
	OmsMysqlConn *gorm.DB
	SmsMysqlConn *gorm.DB
	UmsMysqlConn *gorm.DB
	WmsMysqlConn *gorm.DB
)
