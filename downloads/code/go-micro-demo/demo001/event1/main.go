package main

import (
	"context"

	"github.com/micro/go-micro/v2"
	proto "github.com/micro/go-micro/v2/api/proto"
	"github.com/micro/go-micro/v2/util/log"
)

// 切记，事件订阅结构的所有公有方法都会被执行，方法名没有限制，但是方法一定要接收ctx，event
type Event1 struct{}

func (e *Event1) Handler(ctx context.Context, event *proto.Event) error {
	log.Info( "公有方法Handler Id，", event.Id)
	log.Log("公有方法Handler 收到事件，", event.Name)
	log.Log("公有方法Handler 数据", event.Data)
	return nil
}

func (e *Event1) Handler1(ctx context.Context, event *proto.Event) error {
	log.Info( "公有方法Handler1 Id，", event.Id)
	log.Log("公有方法Handler1 收到事件，", event.Name)
	log.Log("公有方法Handler1 数据，", event.Data)
	return nil
}

/**
打开本注释后，会导致侦听器无法工作。方法一定要接收ctx，event才能正常运行。
func (e *Event1) Handler2() error {
    log.Log("公有方法Handler2 收到事件，不解析参数")
    return nil
}**/

func (e *Event1) handler(ctx context.Context, event *proto.Event) error {
	log.Log("私有方法 handler，收到事件，", event.Name)
	log.Log("私有方法 handler，数据", event.Data)
	return nil
}

func main() {
	service := micro.NewService(
		// 服务名可以随意
		micro.Name("event1"), // 在event 处理器里边关注的是 topic 的名称，服务名称不用在意。
	)
	service.Init()

	// 订阅事件
	micro.RegisterSubscriber("com.anycps.wolferhua.api.event1", service.Server(), new(Event1))

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}