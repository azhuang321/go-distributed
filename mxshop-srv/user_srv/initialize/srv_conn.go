package initialize

import (
	"fmt"

	_ "github.com/mbobakov/grpc-consul-resolver"
	"go.uber.org/zap"
	"google.golang.org/grpc"

	"user_srv/global"
	"user_srv/proto"
)

func InitSrvConn() {
	consulInfo := global.ServerConfig.Consul
	// 通过负载均衡器 去注册中心拿用户服务
	couponsSrvConn, err := grpc.Dial(
		fmt.Sprintf("consul://%s:%d/%s", consulInfo.Host, consulInfo.Port, global.ServerConfig.CouponsSrv.Name),
		grpc.WithInsecure(),
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy":"round_robin"}`),
	)
	if err != nil {
		zap.S().Errorw("[InitSrvConn] 连接 [user_srv] 失败", "msg", err.Error())
		return
	}
	pointsSrvClient := proto.NewCouponsClient(couponsSrvConn)
	global.CouponsClient = pointsSrvClient
}
