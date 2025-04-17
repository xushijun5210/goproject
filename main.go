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
	// //监听端口
	// port := 8021
	// //初始化日志
	initialize.InitLogger()
	// //初始化路由
	// Router := initialize.Routers()
	// zap.S().Info("router init success,端口：", port)

	// if err := Router.Run(fmt.Sprintf(":%d", port)); err != nil {
	// 	zap.S().Panic("listen and serve error:%s", err.Error())
	// }
	IP := flag.String("ip", "localhost", "ip地址")
	Port := flag.Int("port", 8021, "端口号")
	flag.Parse()
	fmt.Println("启动成功，监听地址：", *IP, "，端口号：", *Port)
	//初始化grpc服务
	server := grpc.NewServer()
	//启动grpc服务
	proto.RegisterGreeterServer(server, &handler.UserServiceServer{})
	listen, err := grpc.Listen("tcp", fmt.Sprintf("%s:%d", *IP, *Port))
	if err != nil {
		zap.S().Panic("listen error:%s", err.Error())
	}
	if err := server.Serve(listen); err != nil {
		zap.S().Panic("grpc serve error:%s", err.Error())
	}
}
