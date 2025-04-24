package main

import (
	"flag"
	"fmt"
	"goproject/handler"
	"goproject/initialize"

	"goproject/proto"

	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func main() {
	IP := flag.String("ip", "0.0.0.0", "ip地址")
	Port := flag.Int("port", 0, "端口号")
	//初始化日志
	initialize.InitLogger()
	//初始化路由
	Router := initialize.Routers()
	zap.S().Info("router init success,端口：", Port)

	if err := Router.Run(fmt.Sprintf(":%d",Port)); err != nil {
		zap.S().Panic("listen and serve error:%s", err.Error())
	}
	flag.Parse()
	zap.S().Info("ip: ", *IP)
	zap.S().Info("port: ", *Port)
	if *Port == 0{
		*Port, _ = utils.GetFreePort()
	}
	zap.S().Info("port: ", *Port)
	server := grpc.NewServer()
	proto.RegisterUserServer(server, &handler.UserServer{})
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", *IP, *Port))
	if err != nil {
		panic("failed to listen:" + err.Error())
	}
	
}
