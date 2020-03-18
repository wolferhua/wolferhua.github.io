package handler

import (
	"context"

	"gitee.com/wolferhua/go-micro-demo/demo001/rpc1/proto/srv1"

	"gitee.com/wolferhua/go-micro-demo/demo001/rpc1/srv"

	"github.com/micro/go-micro/v2/util/log"
)

type Rpc1 struct {
}

// Rpc1.Get 通过API向外暴露为/rpc1/get，接收http请求
// 即：/rpc1/get请求会调用 com.anycps.wolferhua.api.rpc1 服务的 Rpc1.Get方法
// curl 'http://localhost:8080/rpc1/say' -X POST -H 'Content-Type: application/json' -d '{"name":"wolferhua"}'
func (a *Rpc1) Say(ctx context.Context, in *srv1.HelloRequest, out *srv1.HelloResponse) error {

	log.Log("Rpc1.Say接口收到请求")

	// 调用服务,其目的是达到服务相互调用的效果演示。
	rep, err := srv.Srv1(in.Name)
	if err != nil {
		log.Error("服务端响应失败:" + err.Error())
		return nil //errors.InternalServerError("com.anycps.wolferhua.api.rpc1", "服务端响应失败:"+err.Error())
	}
	// 返回数据
	out.Greeting = "srv1 response : " + rep

	return nil
}

// 其他方法，无法被api调用
func (a *Rpc1) Test() {
	log.Log("Rpc1.Test接口收到请求")
}
