package conf

type CommonEtcd struct {
	Host              string `mapstructure:"host"`
	BasePath          string `mapstructure:"basePath"`
	ServerPathCoupon  string `mapstructure:"ServerPathCoupon"`
	ServerPathMember  string `mapstructure:"ServerPathMember"`
	ServerPathProduct string `mapstructure:"ServerPathProduct"`
	ServerPathOrder   string `mapstructure:"ServerPathOrder"`
}
