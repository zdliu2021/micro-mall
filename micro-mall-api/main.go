package main

import (
	"mall-demo/micro-mall-api/conf"
	"mall-demo/micro-mall-api/global"
	"mall-demo/micro-mall-api/router"
	"mall-demo/micro-mall-api/utils"
)

func main() {
	global.GVA_CONFIG = conf.ReadConfig()
	global.GVA_LOG = utils.Zap()
	// 开启Api 服务
	router.InitRpcClients()
	route := router.InitRouter()
	route.Run(":8080")
}
