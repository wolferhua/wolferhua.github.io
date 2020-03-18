package main

import (
	"gitee.com/wolferhua/go-micro-demo/demo001/meta1/handler"
	"gitee.com/wolferhua/go-micro-demo/demo001/meta1/proto/api1"
	"github.com/micro/go-micro/v2/api"
	"github.com/micro/go-micro/v2/api/handler/rpc"

	micro "github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/util/log"
)

func main() {
	// 创建新服务
	service := micro.NewService(
		micro.Name("com.anycps.wolferhua.api.api1"),
	)

	// 初始化
	service.Init()
	// 只需改造注册处理器的代码。
	api1.RegisterApi1Handler(service.Server(), new(handler.Api1),
		api.WithEndpoint(
			&api.Endpoint{
				// 接口方法，一定要在proto接口中存在，不能是类的自有方法
				Name:    "Api1.Get",
				// 该接口使用的API转发模式
				Handler: rpc.Handler,
				// 支持的请求方法
				Method:  []string{"GET"},
				// http请求路由
				Path:    []string{"/api1"},
			}),
		api.WithEndpoint(
			&api.Endpoint{
				Name:    "Api1.Post",
				Handler: rpc.Handler,
				Host:    nil,
				Method:  []string{"POST"},
				Path:    []string{"/api1"},
			}),
	)
	// 执行服务
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}

//micro api --namespace com.anycps.wolferhua.api --handler event