package main

import (
	"github.com/gin-contrib/pprof"
	"mall-demo/micro-mall-api/conf"
	"mall-demo/micro-mall-api/global"
	"mall-demo/micro-mall-api/router"
	"mall-demo/micro-mall-api/utils"
)

func main() {
	//f, _ := os.OpenFile("api_cpu.pprof", os.O_CREATE|os.O_RDWR, 0644)
	//defer f.Close()
	//if err := pprof.StartCPUProfile(f); err != nil {
	//	fmt.Println(err)
	//	os.Exit(1)
	//}
	//defer pprof.StopCPUProfile()

	global.GVA_CONFIG = conf.ReadConfig()
	global.GVA_LOG = utils.Zap()
	// 开启Api 服务
	router.InitRpcClients()
	route := router.InitRouter()
	pprof.Register(route, "debug/pprof")
	route.Run(":8080")
}
