package srv

import (
	"context"
	"log"

	"gitee.com/wolferhua/go-micro-demo/demo001/api1/proto/srv1"
	"github.com/micro/go-micro/v2"
)

func Srv1(name string) (string, error) {
	service := micro.NewService()

	// 获取客户端对象
	client := service.Client()

	// 创建srv1 客户端
	srv1Client := srv1.NewSrv1Service("com.anycps.wolferhua.srv1", client)

	// 在Srv1 handler上请求调用Hello方法
	rsp, err := srv1Client.Hello(context.TODO(), &srv1.HelloRequest{
		Name: name,
	})
	if err != nil {
		log.Fatal(err)
		return "", err
	}
	// 返回服务端
	return rsp.Greeting, nil
}
