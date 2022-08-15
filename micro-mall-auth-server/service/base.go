package service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-redis/redis"
	"io/ioutil"
	"mall-demo/micro-mall-auth-server/global"
	proto_auth_server "mall-demo/micro-mall-auth-server/proto/micro-mall-auth-server"
	proto_member "mall-demo/micro-mall-auth-server/proto/micro-mall-member-proto"
	proto_thirdparty "mall-demo/micro-mall-auth-server/proto/micro-mall-thirdparty-proto"
	rpc_client "mall-demo/micro-mall-auth-server/rpc-client"
	"mall-demo/micro-mall-auth-server/utils/idgenerator"
	"net/http"
	"strings"
	"time"
)

type BaseService struct {
}

func (bs *BaseService) SendCode(ctx context.Context, req *proto_auth_server.SendCodeRequest) (*proto_auth_server.SendCodeResponse, error) {
	rpcClient := rpc_client.GetThirdPartyClient()
	code := idgenerator.GetSnowflakeId()[:5]
	global.GVA_LOG.Info("sendsms:" + code)

	var err error
	timeLayoutStr := "2006-01-02 15:04:05"
	// 接口防刷
	if global.GVA_REDIS.Get(global.SMS_CODE_CACHE_PREFIX+req.PhoneNumber).Err() == redis.Nil {
		_, _ = rpcClient.SendSms(context.TODO(), &proto_thirdparty.SendSmsRequest{
			PhoneNumber: req.PhoneNumber,
			Code:        code,
		})
		global.GVA_REDIS.Set(global.SMS_CODE_CACHE_PREFIX+req.PhoneNumber, code+"_"+time.Now().Format(timeLayoutStr), 3600*time.Second)

	} else {
		res, _ := global.GVA_REDIS.Get(global.SMS_CODE_CACHE_PREFIX + req.PhoneNumber).Result()
		_time, _ := time.Parse(timeLayoutStr, strings.Split(res, "_")[1])
		if _time.Add(60 * time.Second).After(time.Now()) {
			_, _ = rpcClient.SendSms(context.TODO(), &proto_thirdparty.SendSmsRequest{
				PhoneNumber: req.PhoneNumber,
				Code:        code,
			})
			global.GVA_REDIS.Set(global.SMS_CODE_CACHE_PREFIX+req.PhoneNumber, code+"_"+time.Now().Format(timeLayoutStr), 3600*time.Second)
		} else {
			err = errors.New("获取验证码太频繁")
		}
	}

	return &proto_auth_server.SendCodeResponse{}, err
}

func (bs *BaseService) Register(ctx context.Context, req *proto_auth_server.RegisterRequest) (*proto_auth_server.RegisterResponse, error) {
	if req.Code == "" || req.Password == "" || req.Username == "" || req.PhoneNumber == "" {
		return &proto_auth_server.RegisterResponse{}, errors.New("输入不能为空")
	}

	res, err := global.GVA_REDIS.Get(global.SMS_CODE_CACHE_PREFIX + req.PhoneNumber).Result()
	if err != nil {
		if err == redis.Nil {
			return &proto_auth_server.RegisterResponse{}, errors.New("验证码错误")
		} else {
			return &proto_auth_server.RegisterResponse{}, errors.New("未知错误" + err.Error())
		}
	}
	//用完立即删除
	global.GVA_REDIS.Del(global.SMS_CODE_CACHE_PREFIX + req.PhoneNumber)
	if strings.Split(res, "_")[0] == req.Code {
		// 调用远程服务，注册
		rpcClient := rpc_client.GetMemberClient()
		_, err := rpcClient.Register(context.TODO(), &proto_member.RegisterRequest{
			Username:    req.Username,
			Password:    req.Password,
			PhoneNumber: req.PhoneNumber,
		})
		if err != nil {
			return &proto_auth_server.RegisterResponse{}, err
		}
		return &proto_auth_server.RegisterResponse{}, nil
	} else {
		return &proto_auth_server.RegisterResponse{}, errors.New("验证码错误")
	}
}

func (bs *BaseService) Login(ctx context.Context, req *proto_auth_server.LoginRequest) (*proto_auth_server.LoginResponse, error) {
	if req.Password == "" || req.Username == "" {
		return &proto_auth_server.LoginResponse{}, errors.New("输入不能为空")
	}
	rpcClient := rpc_client.GetMemberClient()
	resp, err := rpcClient.Login(context.TODO(), &proto_member.LoginRequest{
		Username: req.Username,
		Password: req.Password,
	})
	if err != nil {
		return &proto_auth_server.LoginResponse{}, err
	}

	// 登陆成功

	return &proto_auth_server.LoginResponse{
		Id:       resp.Id,
		UserName: resp.UserName,
	}, nil
}

func (bs *BaseService) OAuthGitteSuccess(ctx context.Context, req *proto_auth_server.OAuthGitteSuccessRequest) (*proto_auth_server.OAuthGitteSuccessResponse, error) {
	//
	client_id := "608147cc373228e1173523194beaf10a8b6ad64cf6b3b39438d5f2d7ed51c7bf"
	client_secret := "64b49b802925578fdc62131c75b297bf5411a149492c0cd9341f5c52e6de839d"
	redirect_uri := "http://127.0.0.1:8088/oauth2/gitte/success"
	params := fmt.Sprintf("grant_type=authorization_code&code=%s&client_id=%s&redirect_uri=%s&client_secret=%s", req.Code, client_id, redirect_uri, client_secret)

	httpReq, _ := http.NewRequest("POST", "https://gitee.com/oauth/token", strings.NewReader(params))
	resp, err := http.DefaultClient.Do(httpReq)
	defer resp.Body.Close()
	if err != nil {
		global.GVA_LOG.Error(err.Error())
	}

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))

	res := make(map[string]string)
	_ = json.Unmarshal(body, &res)
	accessToken := res["access_token"]
	fmt.Println(accessToken)

	rpcClient := rpc_client.GetMemberClient()
	_resp, err := rpcClient.SocialLogin(context.TODO(), &proto_member.SocialRequest{
		AccessToken: accessToken,
	})

	return &proto_auth_server.OAuthGitteSuccessResponse{
		Id:       _resp.Id,
		UserName: _resp.UserName,
	}, err
}
