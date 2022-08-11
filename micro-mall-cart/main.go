package main

import (
	"mall-demo/micro-mall-cart/conf"
	"mall-demo/micro-mall-cart/global"
	"mall-demo/micro-mall-cart/utils"
	"sync"
)

func main() {
	global.GVA_CONFIG = conf.ReadConfig()
	global.GVA_LOG = utils.Zap()

	// 开启 Rpc 服务
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
