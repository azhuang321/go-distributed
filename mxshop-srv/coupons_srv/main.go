package main

import (
	"context"
	"coupons_srv/global"
	"coupons_srv/handler"
	"coupons_srv/initialize"
	"coupons_srv/library"
	"coupons_srv/proto"
	"coupons_srv/register"
	"encoding/json"
	"fmt"
	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/consumer"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/apache/rocketmq-client-go/v2/rlog"
	"github.com/satori/go.uuid"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health/grpc_health_v1"
	"net"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	library.InitConfig()
	library.InitLogger("debug", "logs/test.log", 1, 5, 1, false)
	initialize.InitConfig()
	args := library.GetArgs()

	g := grpc.NewServer()
	//注册用户服务
	userSrv := &handler.CouponsService{}
	proto.RegisterCouponsServer(g, userSrv)

	//注册健康检查服务
	healthCheckSrv := &handler.HealthCheckSrv{Status: grpc_health_v1.HealthCheckResponse_SERVING, Reason: "running"}
	grpc_health_v1.RegisterHealthServer(g, healthCheckSrv)

	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", args["host"], args["port"]))
	//lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", "0.0.0.0", 8000))
	if err != nil {
		zap.S().Errorf("启动服务失败:%s\n", err.Error())
		return
	}

	go func() {
		zap.S().Infof("启动服务成功:%s:%d", args["host"], args["port"])
		if err := g.Serve(lis); err != nil {
			zap.S().Errorf("启动服务失败:%s\n", err.Error())
		}
	}()

	u2 := uuid.NewV4()
	zap.S().Info("开始注册服务中心....")
	consulRegister, err := register.NewConsulRegister()
	if err != nil {
		zap.S().Errorf("注册服务中心失败:%s", err.Error())
	} else {
		if consulRegister.Register(global.ServerConfig.Name, fmt.Sprintf("%s", u2), args["host"].(string), args["port"].(int), global.ServerConfig.Tags, nil) {
			zap.S().Info("注册服务中心成功")
		}
	}

	rlog.SetLogLevel("error")
	c, _ := rocketmq.NewPushConsumer(
		consumer.WithGroupName("testGroup"),
		consumer.WithNsResolver(primitive.NewPassthroughResolver([]string{"172.18.0.1:9876"})),
	)

	err = c.Subscribe("test", consumer.MessageSelector{}, handler.DealMsg())
	if err != nil {
		fmt.Println(err.Error())
	}
	// Note: start after subscribe
	err = c.Start()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(-1)
	}

	nc, _ := rocketmq.NewPushConsumer(
		consumer.WithGroupName("testGroup1"),
		consumer.WithNsResolver(primitive.NewPassthroughResolver([]string{"172.18.0.1:9876"})),
	)

	err = nc.Subscribe("test1", consumer.MessageSelector{}, func(ctx context.Context, msgs ...*primitive.MessageExt) (consumer.ConsumeResult, error) {
		for i, _ := range msgs {
			msgBody := msgs[i].Body
			msg := make(map[string]interface{})
			_ = json.Unmarshal(msgBody, &msg)
			fmt.Println(msg)
			//todo  消费消息  定时取消未领取优惠券

		}

		return consumer.ConsumeSuccess, nil
	})
	if err != nil {
		fmt.Println(err.Error())
	}
	// Note: start after subscribe
	err = nc.Start()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(-1)
	}

	//主进程信号退出
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	zap.S().Info("服务关闭中 ...")
	zap.S().Info("注销服务中心...")
	if consulRegister.Deregister(fmt.Sprintf("%s", u2)) {
		zap.S().Info("注销服务中心成功")
	}

	err = c.Shutdown()
	if err != nil {
		fmt.Printf("shutdown Consumer error: %s", err.Error())
	}
	err = nc.Shutdown()
	if err != nil {
		fmt.Printf("Shutdown Consumer error: %s", err.Error())
	}
}
