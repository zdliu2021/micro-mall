package main

import (
	"mall-demo/micro-mall-member/conf"
	"mall-demo/micro-mall-member/global"
	"mall-demo/micro-mall-member/utils"
	"mall-demo/micro-mall-member/utils/db"
	"sync"
)

func main() {
	global.GVA_CONFIG = conf.ReadConfig()
	global.GVA_LOG = utils.Zap()
	global.UmsMysqlConn = db.InitMysqlClient()
	db.InitMysqlTables()

	// 开启 member 服务

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

	//switch serviceName {
	//case "micro-mall-api":
	//	router.InitRpcClients()
	//	route := router.InitRouter()
	//	route.Run(":8080")
	//case "micro-mall-micro-mall-member-proto":
	//	micro_mall_member.StartMemberServices()
	//case "micro-mall-micro-mall-product-proto":
	//	micro_mall_product.StartProductServices()
	//case "micro-mall-micro-mall-order-proto":
	//	micro_mall_order.StartOrderServices()
	//case "micro-mall-micro-mall-coupon-proto":
	//	micro_mall_coupon.StartCouponServices()
	//case "micro-mall-micro-mall-thirdparty-proto":
	//	thirdparty_service.InitArgs()
	//	micro_mall_thirdparty.StartThirdPartyServices()
	//case "micro-mall-micro-mall-ware-proto":
	//	micro_mall_ware.StartWareServices()
	//default:
	//	global.GVA_LOG.Error("错误的参数")
	//}
}
