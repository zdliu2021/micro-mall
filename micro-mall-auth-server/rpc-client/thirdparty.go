package rpc_client

import (
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"mall-demo/micro-mall-auth-server/global"
	proto_thirdparty "mall-demo/micro-mall-auth-server/proto/micro-mall-thirdparty-proto"
	"math/rand"
	"strconv"
)

var ThirdPartyClients []proto_thirdparty.ThirdPartyRpcClient

func GetThirdPartyClient() proto_thirdparty.ThirdPartyRpcClient {
	return ThirdPartyClients[rand.Int()%len(ThirdPartyClients)]
}

func InitThirdPartyClient() {
	// 创建clientConfig
	clientConfig := constant.ClientConfig{
		//NamespaceId:         "e525eafa-f7d7-4029-83d9-008937f9d468", // 如果需要支持多namespace，我们可以场景多个client,它们有不同的NamespaceId。当namespace是public时，此处填空字符串。
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              "/Users/zongdianliu/go/src/mall-demo/micro-mall-api/log/api_thirdparty/",
		CacheDir:            "/Users/zongdianliu/go/src/mall-demo/micro-mall-api/log/api_thirdparty/",
		LogLevel:            "debug",
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

	// 创建服务发现客户端的另一种方式 (推荐)
	namingClient, _ := clients.NewNamingClient(
		vo.NacosClientParam{
			ClientConfig:  &clientConfig,
			ServerConfigs: serverConfigs,
		},
	)
	// SelectAllInstance可以返回全部实例列表,包括healthy=false,enable=false,weight<=0
	instances, _ := namingClient.SelectAllInstances(vo.SelectAllInstancesParam{
		ServiceName: "thirdparty_server",
		GroupName:   "group-a",             // 默认值DEFAULT_GROUP
		Clusters:    []string{"cluster-a"}, // 默认值DEFAULT
	})

	ThirdPartyClients = make([]proto_thirdparty.ThirdPartyRpcClient, 0, len(instances))

	for _, instance := range instances {
		addr := instance.Ip + ":" + strconv.Itoa(int(instance.Port))
		global.GVA_LOG.Info("addr", zap.Any("addr", addr))
		grpcConn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

		if err != nil {
			global.GVA_LOG.Error(err.Error())
		}
		c := proto_thirdparty.NewThirdPartyRpcClient(grpcConn)
		ThirdPartyClients = append(ThirdPartyClients, c)

	}
	global.GVA_LOG.Info(strconv.Itoa(len(ThirdPartyClients)))
}
