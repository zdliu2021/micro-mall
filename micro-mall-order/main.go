package main

import (
	"flag"
	"mall-demo/micro-mall-order/conf"
	"mall-demo/micro-mall-order/global"
	"mall-demo/micro-mall-order/utils"
	"mall-demo/micro-mall-order/utils/db"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func main() {
	var serviceName string
	flag.StringVar(&serviceName, "svc", "", "用户名,默认为空")
	flag.Parse()

	global.GVA_CONFIG = conf.ReadConfig()
	global.GVA_LOG = utils.Zap()
	global.OmsMysqlConn = db.InitMysqlClient()
	db.InitMysqlTables()

	//关闭信号处理
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL, syscall.SIGHUP, syscall.SIGQUIT)
	go func() {
		s := <-ch
		unRegisterEtcd()
		if i, ok := s.(syscall.Signal); ok {
			os.Exit(int(i))
		} else {
			os.Exit(0)
		}
	}()

	var wg sync.WaitGroup
	for _, addr := range addrs {
		wg.Add(1)
		go func(addr string) {
			defer wg.Done()
			RegisterEtcd(addr)
			//RegisterNacos(addr)
			StartServer(addr)
		}(addr)
	}
	wg.Wait()

	global.GVA_LOG.Error("program exit.")
}
