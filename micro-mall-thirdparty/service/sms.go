package service

import (
	"context"
	"fmt"
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	dysmsapi20170525 "github.com/alibabacloud-go/dysmsapi-20170525/v2/client"
	console "github.com/alibabacloud-go/tea-console/client"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
	proto_thirdparty "mall-demo/micro-mall-thirdparty/proto/micro-mall-thirdparty-proto"
)

const (
	sms_accessKeyId     = "LTAI5tSAc4FtJcMyJESqVHAV"
	sms_accessKeySecret = "peEgSzkKrNQa7VhPWm329kK8xDiNiO"
)

type Sms struct {
}

/**
 * 使用AK&SK初始化账号Client
 * @param accessKeyId
 * @param accessKeySecret
 * @return Client
 * @throws Exception
 */
func CreateClient(accessKeyId *string, accessKeySecret *string) (_result *dysmsapi20170525.Client, _err error) {
	config := &openapi.Config{
		// 您的 AccessKey ID
		AccessKeyId: accessKeyId,
		// 您的 AccessKey Secret
		AccessKeySecret: accessKeySecret,
	}
	// 访问的域名
	config.Endpoint = tea.String("dysmsapi.aliyuncs.com")
	_result = &dysmsapi20170525.Client{}
	_result, _err = dysmsapi20170525.NewClient(config)
	return _result, _err
}

func (sms *Sms) SendSms(ctx context.Context, req *proto_thirdparty.SendSmsRequest) (*proto_thirdparty.SendSmsResponse, error) {
	client, _ := CreateClient(tea.String(sms_accessKeyId), tea.String(sms_accessKeySecret))

	sendSmsRequest := &dysmsapi20170525.SendSmsRequest{
		SignName:      tea.String("阿里云短信测试"),
		TemplateCode:  tea.String("SMS_154950909"),
		PhoneNumbers:  tea.String("15098793212"),
		TemplateParam: tea.String(fmt.Sprintf("{'code':'%s'}", req.Code)),
	}

	runtime := &util.RuntimeOptions{}
	resp, _ := client.SendSmsWithOptions(sendSmsRequest, runtime)

	console.Log(util.ToJSONString(tea.ToMap(resp)))

	return &proto_thirdparty.SendSmsResponse{}, nil
}

//
//func main() {
//	err := _main()
//	if err != nil {
//		panic(err)
//	}
//}
