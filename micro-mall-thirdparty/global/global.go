package global

import (
	"go.uber.org/zap"
	"mall-demo/micro-mall-thirdparty/conf"
)

var (
	GVA_LOG    *zap.Logger
	GVA_CONFIG *conf.Server
)
