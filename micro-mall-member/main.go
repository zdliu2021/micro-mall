package main

import (
	"mall-demo/micro-mall-member/conf"
	"mall-demo/micro-mall-member/global"
	"mall-demo/micro-mall-member/utils"
	"mall-demo/micro-mall-member/utils/db"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func main() {
	global.GVA_CONFIG = conf.ReadConfig()
	global.GVA_LOG = utils.Zap()
	global.UmsMysqlConn = db.InitMysqlClient()
	db.InitMysqlTables()

	// 开启 member 服务

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
