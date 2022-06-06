package main

import (
	"github.com/olivere/elastic/v7"
	"log"
	"mall-demo/micro-mall-search/conf"
	"mall-demo/micro-mall-search/global"
	"mall-demo/micro-mall-search/utils"
	"os"
	"sync"
)

func main() {
	var err error
	global.GVA_CONFIG = conf.ReadConfig()
	global.GVA_LOG = utils.Zap()
	global.GVA_ES, err = elastic.NewClient(
		//elastic 服务地址
		elastic.SetURL("http://127.0.0.1:9200"),
		// 设置错误日志输出
		elastic.SetErrorLog(log.New(os.Stderr, "ELASTIC ", log.LstdFlags)),
		// 设置info日志输出
		elastic.SetInfoLog(log.New(os.Stdout, "", log.LstdFlags)),
		elastic.SetSniff(false),
		elastic.SetTraceLog(log.New(os.Stdout, "[go-elastic] ", 0)), // <- log requests/response to stdout
	)
	if err != nil {
		global.GVA_LOG.Error(err.Error())
	}

	if global.GVA_ES == nil {
		global.GVA_LOG.Error("init error")
		os.Exit(1)
	}
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
