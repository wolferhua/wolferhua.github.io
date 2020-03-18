package main

import (
	"gitee.com/wolferhua/go-micro-demo/demo001/api1/handler"
	"gitee.com/wolferhua/go-micro-demo/demo001/api1/proto/api1"

	micro "github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/util/log"
)

func main() {
	// 创建新服务
	service := micro.NewService(
		micro.Name("com.anycps.wolferhua.api.api1"),
	)
	//细心的朋友可以发现：我们之前定义服务时使用的是：com.anycps.wolferhua.***，现在却使用：com.anycps.wolferhua.api.***。因为我们会定义非常多的接口，需要一个统一的命名空间（Namespace）方便管理。至于命名空间，我们之后会讲解，go-micro 中使用的非常多。

	// 初始化
	service.Init()
	// 注册处理器
	api1.RegisterApi1Handler(service.Server(), new(handler.Api1))
	// 执行服务
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
