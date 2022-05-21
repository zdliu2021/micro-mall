package main

import (
	"mall-demo/micro-mall-ware/conf"
	"mall-demo/micro-mall-ware/global"
	"mall-demo/micro-mall-ware/utils"
	"mall-demo/micro-mall-ware/utils/db"
	"sync"
)

func main() {
	global.GVA_CONFIG = conf.ReadConfig()
	global.GVA_LOG = utils.Zap()
	global.WmsMysqlConn = db.InitMysqlClient()
	db.InitMysqlTables()

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
