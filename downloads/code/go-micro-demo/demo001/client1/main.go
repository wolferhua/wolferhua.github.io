package main

import (
	"context"
	"fmt"

	"github.com/micro/go-micro/v2/util/log"

	"gitee.com/wolferhua/go-micro-demo/demo001/client1/proto/srv1"
	micro "github.com/micro/go-micro/v2"
)

func main() {
	// 创建服务

	service := micro.NewService()

	// 获取客户端对象
	client := service.Client()

	// 创建srv1 客户端
	srv1Client := srv1.NewSrv1Service("com.anycps.wolferhua.srv1", client)

	// 在Srv1 handler上请求调用Hello方法
	rsp, err := srv1Client.Hello(context.TODO(), &srv1.HelloRequest{
		Name: "Srv1",
	})
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println("服务端返回的内容为：", rsp.Greeting)
}
