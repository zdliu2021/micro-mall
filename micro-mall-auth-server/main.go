package main

import (
	"mall-demo/micro-mall-auth-server/conf"
	"mall-demo/micro-mall-auth-server/global"
	rpc_client "mall-demo/micro-mall-auth-server/rpc-client"
	"mall-demo/micro-mall-auth-server/utils"
	"mall-demo/micro-mall-auth-server/utils/db"
	"sync"
)

func main() {
	global.GVA_CONFIG = conf.ReadConfig()
	global.GVA_LOG = utils.Zap()
	global.GVA_REDIS = db.GetRedisInstance(db.DefaultRedisOption)

	rpc_client.InitThirdPartyClient()
	rpc_client.InitMemberClient()
	
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
