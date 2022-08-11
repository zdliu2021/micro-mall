package main

import (
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"mall-demo/micro-mall-coupon/global"
	proto_coupon "mall-demo/micro-mall-coupon/proto/micro-mall-coupon-proto"
	"mall-demo/micro-mall-coupon/server"
	"net"
	"strconv"
	"strings"
)

var (
	addrs = []string{"127.0.0.1:50081", "127.0.0.1:50082"}
)

func StartServer(addr string) {
	grpcServer := grpc.NewServer()

	// 2.注册服务
	proto_coupon.RegisterCouponRpcServer(grpcServer, new(server.Server))

	// 3.设置监听, 指定 IP/port
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		global.GVA_LOG.Error("Listen err:" + err.Error())
		return
	}
	defer listener.Close()

	global.GVA_LOG.Info("服务启动", zap.Any("addr", addr))
	// 4. 启动服务
	grpcServer.Serve(listener)
}

func RegisterNacos(addr string) {
	// 创建clientConfig
	clientConfig := constant.ClientConfig{
		//NamespaceId:         "e525eafa-f7d7-4029-83d9-008937f9d468", // 如果需要支持多namespace，我们可以场景多个client,它们有不同的NamespaceId。当namespace是public时，此处填空字符串。
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              "/Users/zongdianliu/go/src/mall-demo/micro-mall-coupon/log",
		CacheDir:            "/Users/zongdianliu/go/src/mall-demo/micro-mall-coupon/log",
		LogLevel:            "warn",
	}

	// 至少一个ServerConfig
	serverConfigs := []constant.ServerConfig{
		{
			IpAddr:      "162.105.88.120",
			ContextPath: "/nacos",
			Port:        8848,
			Scheme:      "http",
		},
	}

	// 创建服务发现客户端 (推荐)
	namingClient, err := clients.NewNamingClient(
		vo.NacosClientParam{
			ClientConfig:  &clientConfig,
			ServerConfigs: serverConfigs,
		},
	)
	if err != nil {
		global.GVA_LOG.Error("clients.NewNamingClient err," + err.Error())
	}
	ip := strings.Split(addr, ":")[0]
	port, _ := strconv.ParseUint(strings.Split(addr, ":")[1], 10, 64)

	global.GVA_LOG.Info("reg coupon instance", zap.Any("ip", ip), zap.Any("port", port))

	success, err := namingClient.RegisterInstance(vo.RegisterInstanceParam{
		Ip:          ip,
		Port:        port,
		ServiceName: "coupon_server",
		Weight:      10,
		Enable:      true,
		Healthy:     true,
		Ephemeral:   true,
		Metadata:    map[string]string{"idc": "shanghai"},
		ClusterName: "cluster-a", // 默认值DEFAULT
		GroupName:   "group-a",   // 默认值DEFAULT_GROUP
	})
	if !success {
		return
	} else {
		global.GVA_LOG.Info("namingClient.RegisterInstance Success")
	}
}
