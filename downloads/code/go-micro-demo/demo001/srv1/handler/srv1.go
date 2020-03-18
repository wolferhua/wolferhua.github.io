package handler

import (
	context "context"
	"os"
	"strconv"

	// 引入proto定义
	srv1 "gitee.com/wolferhua/go-micro-demo/demo001/srv1/proto/srv1"
)

type Srv1 struct{}

func (g *Srv1) Hello(ctx context.Context, req *srv1.HelloRequest, rsp *srv1.HelloResponse) error {
	// 业务逻辑代码
	rsp.Greeting = "Hello " + req.Name + " ,from pid " + strconv.Itoa(os.Getpid())
	return nil
}
