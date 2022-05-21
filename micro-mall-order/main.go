package main

import (
	"flag"
	"mall-demo/micro-mall-order/conf"
	"mall-demo/micro-mall-order/global"
	"mall-demo/micro-mall-order/utils"
	"mall-demo/micro-mall-order/utils/db"
	"sync"
)

func main() {
	var serviceName string
	flag.StringVar(&serviceName, "svc", "", "用户名,默认为空")
	flag.Parse()

	global.GVA_CONFIG = conf.ReadConfig()
	global.GVA_LOG = utils.Zap()
	global.OmsMysqlConn= db.InitMysqlClient()
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
