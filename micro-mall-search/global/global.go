package global

import (
	"github.com/olivere/elastic/v7"
	"go.uber.org/zap"
	"mall-demo/micro-mall-search/conf"
)

var (
	GVA_LOG    *zap.Logger
	GVA_CONFIG *conf.Server
	GVA_ES     *elastic.Client
)

const ProductIndex = "mall_product"
