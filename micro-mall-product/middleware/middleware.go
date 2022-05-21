package middleware

import (
	"context"
	"encoding/json"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/peer"
	"log"
	"mall-demo/micro-mall-product/global"
	"time"
)

func UnaryServerInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	remote, _ := peer.FromContext(ctx)
	remoteAddr := remote.Addr.String()

	in, _ := json.Marshal(req)
	inStr := string(in)
	log.Println("ip", remoteAddr, "access_start", info.FullMethod, "in", inStr)
	global.GVA_LOG.Info("", zap.String("access_start", info.FullMethod), zap.Any("in", inStr))
	start := time.Now()
	defer func() {
		duration := int64(time.Since(start) / time.Millisecond)
		if duration >= 500 {
			global.GVA_LOG.Info("[Long time request]. ", zap.String("access_end", info.FullMethod), zap.Any("err", err), zap.Int64("duration/ms", duration))
		} else {
			global.GVA_LOG.Info("[Short time request]. ", zap.String("access_end", info.FullMethod), zap.Any("err", err), zap.Int64("duration/ms", duration))
		}
	}()

	resp, err = handler(ctx, req)

	return
}
