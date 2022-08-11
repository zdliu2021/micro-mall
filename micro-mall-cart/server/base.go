package server

import proto_cart "mall-demo/micro-mall-cart/proto/micro-mall-cart-proto"

type Server struct {
	proto_cart.UnimplementedCartRpcServer
}
