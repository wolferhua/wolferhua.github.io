package main

import (
	"gitee.com/wolferhua/go-micro-demo/demo001/rpc1/handler"
	"gitee.com/wolferhua/go-micro-demo/demo001/rpc1/proto/rpc1"

	micro "github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/util/log"
)

func main() {
	// 创建新服务
	service := micro.NewService(
		micro.Name("com.anycps.wolferhua.api.rpc1"),
	)

	// 初始化
	service.Init()
	// 注册处理器
	rpc1.RegisterRpc1Handler(service.Server(), new(handler.Rpc1))
	// 执行服务
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
