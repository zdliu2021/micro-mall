package main

import (
	"mall-demo/micro-mall-auth-server/conf"
	"mall-demo/micro-mall-auth-server/global"
	rpc_client "mall-demo/micro-mall-auth-server/rpc-client"
	"mall-demo/micro-mall-auth-server/utils"
	"mall-demo/micro-mall-auth-server/utils/db"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func main() {
	global.GVA_CONFIG = conf.ReadConfig()
	global.GVA_LOG = utils.Zap()
	global.GVA_REDIS = db.GetRedisInstance(db.DefaultRedisOption)

	rpc_client.InitThirdPartyRpcClient()
	rpc_client.InitMemberRpcClient()

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
