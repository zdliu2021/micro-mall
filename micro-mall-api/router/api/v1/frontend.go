package v1

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"mall-demo/micro-mall-api/model"
	"mall-demo/micro-mall-api/model/response"
	proto_auth_server "mall-demo/micro-mall-api/proto/micro-mall-auth-server"
	rpc_client "mall-demo/micro-mall-api/rpc-client"
)

func SendCode(ctx *gin.Context) {
	phoneNumber := ctx.Query("phone")
	rpcClient := rpc_client.GetAuthServerRpcClient()
	_, err := rpcClient.SendCode(context.TODO(), &proto_auth_server.SendCodeRequest{PhoneNumber: phoneNumber})
	if err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	response.Ok(ctx)
}

func Register(ctx *gin.Context) {
	rpcClient := rpc_client.GetAuthServerRpcClient()
	username := ctx.PostForm("userName")
	password := ctx.PostForm("password")
	phone := ctx.PostForm("phone")
	code := ctx.PostForm("code")
	_, err := rpcClient.Register(context.TODO(), &proto_auth_server.RegisterRequest{
		Username:    username,
		Password:    password,
		PhoneNumber: phone,
		Code:        code,
	})

	if err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	response.Ok(ctx)
}

func Login(ctx *gin.Context) {
	rpcClient := rpc_client.GetAuthServerRpcClient()
	username := ctx.PostForm("loginacct")
	password := ctx.PostForm("password")
	_, err := rpcClient.Login(context.TODO(), &proto_auth_server.LoginRequest{
		Username: username,
		Password: password,
	})

	if err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	response.Ok(ctx)
}

func OAuthGitteSuccess(ctx *gin.Context) {
	user := model.LoginUser{}
	session := sessions.Default(ctx)
	if session.Get("current_user") != nil {
		json.Unmarshal(session.Get("current_user").([]byte), &user)
		fmt.Println(user)
		response.Ok(ctx)
		return
	}

	rpcClient := rpc_client.GetAuthServerRpcClient()
	code := ctx.Query("code")
	resp, err := rpcClient.OAuthGitteSuccess(context.TODO(), &proto_auth_server.OAuthGitteSuccessRequest{
		Code: code,
	})
	user.UserName = resp.GetUserName()
	user.Id = resp.GetId()

	if err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	_user, _ := json.Marshal(&user)
	session.Set("current_user", _user)
	session.Save()

	response.Ok(ctx)
}
