package conf

type ThirdParty struct {
	// 请填写您的AccessKeyId。
	AccessKeyId string `mapstructure:"accessKeyId"`

	// 请填写您的AccessKeySecret。
	AccessKeySecret string `mapstructure:"accessKeySecret"`

	// host的格式为 bucketname.endpoint ，请替换为您的真实信息。
	Host string `mapstructure:"host"`
}
