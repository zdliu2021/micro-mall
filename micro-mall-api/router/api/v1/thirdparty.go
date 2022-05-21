package v1

import (
	"context"
	"github.com/gin-gonic/gin"
	"mall-demo/micro-mall-api/global"
	response2 "mall-demo/micro-mall-api/model/response"
	"mall-demo/micro-mall-api/proto/micro-mall-thirdparty-proto"
	"mall-demo/micro-mall-api/rpc-client"
)

func GetOSSPolicy(ctx *gin.Context) {
	rpcClient := rpc_client.GetThirdPartyClient()
	var req proto_thirdparty.GetOSSTokenRequest
	resp, err := rpcClient.GetOSSToken(context.TODO(), &req)
	if err != nil {
		global.GVA_LOG.Error("getosstoken error:" + err.Error())
		response2.FailWithMessage(err.Error(), ctx)
		return
	}

	res := &response2.GetOSSTokenResponse{
		AccessId:  resp.AccessId,
		Host:      resp.Host,
		Expire:    resp.Expire,
		Signature: resp.Signature,
		Policy:    resp.Policy,
		Dir:       resp.Dir,
		Callback:  resp.Callback,
	}
	response2.OkWithData(res, ctx)
}
