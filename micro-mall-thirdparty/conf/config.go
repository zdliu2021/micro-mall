package conf

import (
	"github.com/spf13/viper"
	"runtime"
	"strings"
	"sync"
)

var once sync.Once

var c *Server

type Server struct {
	Zap         Zap         `mapstructure:"common-zap"`
	CommonEtcd  CommonEtcd  `mapstructure:"common-etcd"`
	CommonMysql MysqlOption `mapstructure:"common-mysql"`
	ThirdParty  `mapstructure:"thirdparty-base"`
}

func getCurrentDir() string {
	_, fileName, _, _ := runtime.Caller(1)
	aPath := strings.Split(fileName, "/")
	dir := strings.Join(aPath[0:len(aPath)-1], "/")
	return dir
}

func init() {
	once.Do(func() {
		realPath := getCurrentDir()
		configFilePath := realPath + "/conf"
		viper.SetConfigType("toml")
		viper.SetConfigName("/common")
		viper.AddConfigPath(configFilePath)
		err := viper.ReadInConfig()
		if err != nil {
			panic(err)
		}
		viper.SetConfigName("/api")
		err = viper.MergeInConfig()
		if err != nil {
			panic(err)
		}
		viper.SetConfigName("/coupon")
		err = viper.MergeInConfig()
		if err != nil {
			panic(err)
		}
		viper.SetConfigName("/member")
		err = viper.MergeInConfig()
		if err != nil {
			panic(err)
		}
		viper.SetConfigName("/order")
		err = viper.MergeInConfig()
		if err != nil {
			panic(err)
		}
		viper.SetConfigName("/product")
		err = viper.MergeInConfig()
		if err != nil {
			panic(err)
		}
		viper.SetConfigName("/security")
		err = viper.MergeInConfig()
		if err != nil {
			panic(err)
		}

		c = new(Server)
		viper.Unmarshal(&c)
	})
}

func ReadConfig() *Server {
	return c
}
