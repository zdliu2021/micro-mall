package service

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/jinzhu/copier"
	"mall-demo/micro-mall-cart/global"
	"mall-demo/micro-mall-cart/model"
	proto_cart "mall-demo/micro-mall-cart/proto/micro-mall-cart-proto"
	"mall-demo/micro-mall-cart/proto/micro-mall-product-proto"
	rpc_client "mall-demo/micro-mall-cart/rpc-client"
	"strconv"
	"sync"
)

const (
	CART_PREFIX              = "gulimall-cart-"
	TEMP_USER_COOKIE_TIMEOUT = 60 * 60 * 24 * 30
	TEMP_USER_COOKIE_NAME    = "user-key"
	CART                     = "cart"
)

type BaseService struct {
}

func (bs *BaseService) AddToCart(ctx context.Context, in *proto_cart.AddToCartRequest) (*proto_cart.AddToCartResponse, error) {
	var key string
	if in.UserId != 0 {
		fmt.Println("用户登陆了")
		// 登陆了
		key = CART_PREFIX + strconv.Itoa(int(in.UserId))
	} else {
		//未登陆
		key = CART_PREFIX + in.UserKey
	}
	productRedisValue, _ := global.GVA_REDIS.HGet(key, strconv.Itoa(int(in.SkuId))).Result()
	if productRedisValue == "" {
		var item model.CartInfo
		//购物车没有此商品
		var wg sync.WaitGroup
		// 远程查询商品的信息
		f1 := func(wg *sync.WaitGroup) {
			defer wg.Done()
			rpcClient := rpc_client.GetProductClient()
			resp, _ := rpcClient.GetSkuItem(context.TODO(), &proto_product.GetSkuItemRequest{
				SkuId: int64(in.SkuId),
			})
			if resp.SkuInfoEntity != nil {
				item.SkuID = resp.SkuInfoEntity.SkuId
				item.Price = resp.SkuInfoEntity.Price
				item.Check = 0
				item.DefaultImage = resp.SkuInfoEntity.SkuDefaultImg
				item.Title = resp.SkuInfoEntity.SkuTitle
				item.UserID = in.UserKey
			}

		}

		// 远程查询skuAttrValues组合信息
		f2 := func(wg *sync.WaitGroup) {
			defer wg.Done()
			//TODO
		}

		wg.Add(1)
		go f1(&wg)
		wg.Add(1)
		go f2(&wg)
		wg.Wait()
		s, _ := json.Marshal(&item)
		global.GVA_REDIS.HSet(key, strconv.Itoa(int(in.SkuId)), s)
	} else {
		// 购物车有此商品，改变商品数量即可
		var cartInfo model.CartInfo
		json.Unmarshal([]byte(productRedisValue), &cartInfo)
		cartInfo.Count += int(in.Num)
		s, _ := json.Marshal(&cartInfo)
		global.GVA_REDIS.HSet(key, strconv.Itoa(int(in.SkuId)), s)
	}
	return &proto_cart.AddToCartResponse{}, nil
}

func (bs *BaseService) GetCartItem(ctx context.Context, in *proto_cart.GetCartItemRequest) (*proto_cart.GetCartItemResponse, error) {
	var key string
	if in.UserId != 0 {
		// 登陆了
		key = CART_PREFIX + strconv.Itoa(int(in.UserId))
	} else {
		//未登陆
		key = CART_PREFIX + in.UserKey
	}
	productRedisValue, _ := global.GVA_REDIS.HGet(key, strconv.Itoa(int(in.SkuId))).Result()

	print("raw data:" + productRedisValue)
	var cartItem model.CartInfo
	if err := json.Unmarshal([]byte(productRedisValue), &cartItem); err != nil {
		fmt.Println(err)
	}

	fmt.Println("data:", cartItem)

	tmp := proto_cart.CartInfo{
		SkuId:        cartItem.SkuID,
		Check:        int32(cartItem.Check),
		Title:        cartItem.Title,
		DefaultImage: cartItem.DefaultImage,
		UserId:       cartItem.UserID,
		Price:        float32(cartItem.Price),
	}
	return &proto_cart.GetCartItemResponse{Item: &tmp}, nil
}

func (bs *BaseService) GetCart(ctx context.Context, in *proto_cart.GetCartRequest) (*proto_cart.GetCartResponse, error) {
	if in.UserId != 0 {
		// 登陆了
		key := CART_PREFIX + strconv.Itoa(int(in.UserId))
		ans := GetCartItems(key)
		var items []*proto_cart.CartInfo
		copier.CopyWithOption(&items, &ans, copier.Option{IgnoreEmpty: true, DeepCopy: true})
		return &proto_cart.GetCartResponse{
			Items: items,
		}, nil
	} else {
		//未登陆
		key := CART_PREFIX + in.UserKey
		ans := GetCartItems(key)
		var items []*proto_cart.CartInfo = make([]*proto_cart.CartInfo, 0, 10)
		for i := 0; i < len(ans); i++ {
			cartItem := ans[i]
			tmp := proto_cart.CartInfo{
				SkuId:        cartItem.SkuID,
				Check:        int32(cartItem.Check),
				Title:        cartItem.Title,
				DefaultImage: cartItem.DefaultImage,
				UserId:       cartItem.UserID,
				Price:        float32(cartItem.Price),
			}
			items = append(items, &tmp)
		}
		return &proto_cart.GetCartResponse{
			Items: items,
		}, nil
	}
}
func (bs *BaseService) ClearCartInfo(ctx context.Context, in *proto_cart.ClearCartInfoRequest) (*proto_cart.ClearCartInfoResponse, error) {

	return &proto_cart.ClearCartInfoResponse{}, nil
}
func (bs *BaseService) CheckItem(ctx context.Context, in *proto_cart.CheckItemRequest) (*proto_cart.CheckItemResponse, error) {
	var key string
	if in.UserId != 0 {
		// 登陆了
		key = CART_PREFIX + strconv.Itoa(int(in.UserId))
	} else {
		//未登陆
		key = CART_PREFIX + in.UserKey
	}
	productRedisValue, _ := global.GVA_REDIS.HGet(key, strconv.Itoa(int(in.SkuId))).Result()

	var cartItem model.CartInfo
	json.Unmarshal([]byte(productRedisValue), &cartItem)
	cartItem.Check = 1
	s, _ := json.Marshal(&cartItem)
	global.GVA_REDIS.HSet(key, strconv.Itoa(int(in.SkuId)), s)
	return &proto_cart.CheckItemResponse{}, nil
}
func (bs *BaseService) ChangeItemCount(ctx context.Context, in *proto_cart.ChangeItemCountRequest) (*proto_cart.ChangeItemCountResponse, error) {
	return &proto_cart.ChangeItemCountResponse{}, nil
}

func (bs *BaseService) DeleteIdCartInfo(ctx context.Context, in *proto_cart.DeleteIdCartInfoRequest) (*proto_cart.DeleteIdCartInfoResponse, error) {
	if in.UserId != 0 {
		// 登陆了
		global.GVA_REDIS.HDel(CART_PREFIX+strconv.Itoa(int(in.UserId)), strconv.Itoa(int(in.SkuId)))
	} else {
		//未登陆
		global.GVA_REDIS.HDel(CART_PREFIX+in.UserKey, strconv.Itoa(int(in.SkuId)))
	}
	return &proto_cart.DeleteIdCartInfoResponse{}, nil
}

func (bs *BaseService) GetUserCartItems(ctx context.Context, in *proto_cart.GetUserCartItemsRequest) (*proto_cart.GetUserCartItemsResponse, error) {
	if in.GetUserId() == 0 {
		return &proto_cart.GetUserCartItemsResponse{}, nil
	}
	ans := GetCartItems(CART_PREFIX + strconv.Itoa(int(in.UserId)))
	resp := make([]*proto_cart.CartInfo, 0, 10)
	for i := 0; i < len(ans); i++ {
		if ans[i].Check == 1 || ans[i].Check == 0 {
			tmp := proto_cart.CartInfo{}
			copier.CopyWithOption(&tmp, &ans[i], copier.Option{IgnoreEmpty: true, DeepCopy: true})
			resp = append(resp, &tmp)
		}
	}
	return &proto_cart.GetUserCartItemsResponse{
		Items: resp,
	}, nil
}

func GetCartItems(cartKey string) []model.CartInfo {
	ans := make([]model.CartInfo, 0, 10)
	values := global.GVA_REDIS.HGetAll(cartKey).Val()
	for _, v := range values {
		var cartInfo model.CartInfo
		json.Unmarshal([]byte(v), &cartInfo)
		ans = append(ans, cartInfo)
	}
	return ans
}
