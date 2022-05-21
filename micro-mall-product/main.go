package main

import (
	"mall-demo/micro-mall-product/conf"
	"mall-demo/micro-mall-product/global"
	"mall-demo/micro-mall-product/utils"
	"mall-demo/micro-mall-product/utils/db"
	"sync"
)

func main() {
	global.GVA_CONFIG = conf.ReadConfig()
	global.GVA_LOG = utils.Zap()
	global.PmsMysqlConn = db.InitMysqlClient()
	db.InitMysqlTables()

	InitRpcClients()
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
