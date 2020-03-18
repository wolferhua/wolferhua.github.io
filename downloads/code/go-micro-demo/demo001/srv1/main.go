package main

import (
	"os"

	"gitee.com/wolferhua/go-micro-demo/demo001/srv1/handler"
	"gitee.com/wolferhua/go-micro-demo/demo001/srv1/proto/srv1"
	micro "github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/util/log"
	// 引入proto定义
)

func main() {
	log.Info("pid is:", os.Getpid())
	// 创建服务
	service := micro.NewService(
		//com.anycps.wolferhua
		micro.Name("com.anycps.wolferhua.srv1"),
		micro.Version("latest"),
	)
	// 注册处理器
	srv1.RegisterSrv1Handler(service.Server(), new(handler.Srv1))

	if err := service.Run(); err != nil {
		// exit

		log.Fatal(err)
	}

}
