package main

import (
	"mall-demo/micro-mall-thirdparty/conf"
	"mall-demo/micro-mall-thirdparty/global"
	"mall-demo/micro-mall-thirdparty/service"
	"mall-demo/micro-mall-thirdparty/utils"
	"sync"
)

func main() {
	global.GVA_CONFIG = conf.ReadConfig()
	global.GVA_LOG = utils.Zap()

	service.InitArgs()

	var wg sync.WaitGroup
	for _, addr := range addrs {
		wg.Add(1)
		go func(addr string) {
			defer wg.Done()
			RegisterNacos(addr)
			StartServer(addr)
		}(addr)
	}
	wg.Wait()
	global.GVA_LOG.Error("program exit.")

}
