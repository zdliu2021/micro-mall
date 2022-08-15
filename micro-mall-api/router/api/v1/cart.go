package v1

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"mall-demo/micro-mall-api/model"
	"mall-demo/micro-mall-api/model/response"
	proto_cart "mall-demo/micro-mall-api/proto/micro-mall-cart-proto"
	rpc_client "mall-demo/micro-mall-api/rpc-client"
	"mall-demo/micro-mall-api/utils/idgenerator"
	"strconv"
)

type CartLoginUser struct {
	UserId int64
	TempId string
}

func CartList(ctx *gin.Context) {
	rpcClient := rpc_client.GetCartClient()
	cartLoginUser := preHandle(ctx)
	resp, err := rpcClient.GetCart(context.TODO(), &proto_cart.GetCartRequest{
		UserId:  cartLoginUser.UserId,
		UserKey: cartLoginUser.TempId,
	})
	if err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	postHandle(ctx, cartLoginUser)
	response.OkWithData(resp, ctx)
}

func GetCartItem(ctx *gin.Context) {
	skuId, _ := strconv.ParseInt(ctx.Query("skuId"), 10, 64)
	rpcClient := rpc_client.GetCartClient()
	cartLoginUser := preHandle(ctx)
	resp, err := rpcClient.GetCartItem(context.TODO(), &proto_cart.GetCartItemRequest{
		UserKey: cartLoginUser.TempId,
		UserId:  cartLoginUser.UserId,
		SkuId:   int64(skuId),
	})
	if err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	postHandle(ctx, cartLoginUser)
	response.OkWithData(resp, ctx)
}

func AddCartItem(ctx *gin.Context) {
	skuId, _ := strconv.ParseInt(ctx.Query("skuId"), 10, 64)
	num, _ := strconv.ParseInt(ctx.Query("num"), 10, 64)
	rpcClient := rpc_client.GetCartClient()
	cartLoginUser := preHandle(ctx)
	_, err := rpcClient.AddToCart(context.TODO(), &proto_cart.AddToCartRequest{
		UserKey: cartLoginUser.TempId,
		UserId:  cartLoginUser.UserId,
		SkuId:   int32(skuId),
		Num:     int32(num),
	})
	if err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	postHandle(ctx, cartLoginUser)
	response.Ok(ctx)
}

func CheckItem(ctx *gin.Context) {
	skuId, _ := strconv.ParseInt(ctx.Query("skuId"), 10, 64)
	rpcClient := rpc_client.GetCartClient()
	cartLoginUser := preHandle(ctx)
	_, err := rpcClient.CheckItem(context.TODO(), &proto_cart.CheckItemRequest{
		UserKey: cartLoginUser.TempId,
		UserId:  cartLoginUser.UserId,
		SkuId:   int64(skuId),
	})
	if err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	postHandle(ctx, cartLoginUser)
	response.Ok(ctx)
}

func DeleteItem(ctx *gin.Context) {
	skuId, _ := strconv.ParseInt(ctx.Query("skuId"), 10, 64)
	rpcClient := rpc_client.GetCartClient()
	cartLoginUser := preHandle(ctx)
	_, err := rpcClient.DeleteIdCartInfo(context.TODO(), &proto_cart.DeleteIdCartInfoRequest{
		UserKey: cartLoginUser.TempId,
		UserId:  cartLoginUser.UserId,
		SkuId:   int64(skuId),
	})
	if err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	postHandle(ctx, cartLoginUser)
	response.Ok(ctx)
}

func CurrentUserCartItems(ctx *gin.Context) {
	rpcClient := rpc_client.GetCartClient()
	cartLoginUser := preHandle(ctx)
	resp, err := rpcClient.GetUserCartItems(context.TODO(), &proto_cart.GetUserCartItemsRequest{
		UserKey: cartLoginUser.TempId,
		UserId:  cartLoginUser.UserId,
	})
	if err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	postHandle(ctx, cartLoginUser)
	response.OkWithData(resp, ctx)
}

func preHandle(ctx *gin.Context) *CartLoginUser {
	var cartLoginUser CartLoginUser
	user := model.LoginUser{}
	session := sessions.Default(ctx)
	if session.Get("current_user") != nil {
		fmt.Println("current user:", session.Get("current_user").([]byte))
		json.Unmarshal(session.Get("current_user").([]byte), &user)
		cartLoginUser.UserId = user.Id
	}
	if session.Get("cart_tmp_user") != nil {
		fmt.Println("current cart tmp user:" + session.Get("cart_tmp_user").(string))
		cartLoginUser.TempId = session.Get("cart_tmp_user").(string)
		fmt.Println(cartLoginUser.TempId)
	} else {
		// 如果没有临时用户，分配一个临时用户
		id := idgenerator.GetSnowflakeId()
		cartLoginUser.TempId = id
	}
	return &cartLoginUser
}

func postHandle(ctx *gin.Context, user *CartLoginUser) {
	session := sessions.Default(ctx)
	if session.Get("cart_tmp_user") == nil {
		fmt.Println("new set tmp user:" + user.TempId)
		session.Set("cart_tmp_user", user.TempId)
		session.Save()
	}
	ctx.SetCookie("gin_cookie", "test", 3600, "/", "localhost", false, true)
}
