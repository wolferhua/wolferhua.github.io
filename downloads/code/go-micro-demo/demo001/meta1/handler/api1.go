package handler

import (
	"context"
	"encoding/json"
	"strings"

	"gitee.com/wolferhua/go-micro-demo/demo001/meta1/srv"

	proto1 "github.com/micro/go-micro/api/proto"
	"github.com/micro/go-micro/v2/errors"
	"github.com/micro/go-micro/v2/util/log"
)

type Api1 struct {
}

// Api1.Get 通过API向外暴露为/api1/get，接收http请求
// 即：/api1/get请求会调用com.anycps.wolferhua.api.api1服务的Api1.Get方法
//curl 'http://localhost:8080/api1/get?name=wolferhua'
func (a *Api1) Get(ctx context.Context, in *proto1.Request, out *proto1.Response) error {

	log.Log("Api1.Get接口收到请求")

	name, ok := in.Get["name"]

	if !ok || len(name.Values) == 0 {
		return errors.BadRequest("com.anycps.wolferhua.api.api1", "参数不正确")
	}

	// 打印请求头
	for k, v := range in.Header {
		log.Log("请求头信息，", k, " : ", v)
	}

	out.StatusCode = 200

	// 调用服务,其目的是达到服务相互调用的效果演示。
	sname := strings.Join(name.Values, "")
	srvRsp, err := srv.Srv1(sname)
	if err != nil {
		return errors.InternalServerError("com.anycps.wolferhua.api.api1", "服务端响应失败:"+err.Error())
	}

	//业务代码
	b, _ := json.Marshal(map[string]string{
		"message": "我们已经收到你的请求，" + srvRsp,
	})

	// 设置返回值
	out.Body = string(b)

	return nil
}

// Api1.Post 通过API向外暴露为/api1/get，接收http请求
// 即：/api1/post请求会调用com.anycps.wolferhua.api.api1服务的Api1.Post方法
// curl http://localhost:8080/api1/post -X POST -H 'Content-Type: application/json' -d '{"name":"wolferhua"}'

func (a *Api1) Post(ctx context.Context, in *proto1.Request, out *proto1.Response) error {
	log.Log("Api1.Post接口收到请求")
	// 判断请求类型
	if in.Method != "POST" {
		return errors.BadRequest("com.anycps.wolferhua.api.api1", "require post")
	}

	// 获取请求头信息
	ct, ok := in.Header["Content-Type"]
	if !ok || len(ct.Values) == 0 {
		return errors.BadRequest("com.anycps.wolferhua.api.api1", "need content-type")
	}

	// 判断请求头类型
	if ct.Values[0] != "application/json" {
		return errors.BadRequest("com.anycps.wolferhua.api.api1", "expect application/json")
	}

	// 获取请求数据
	var body map[string]interface{}
	json.Unmarshal([]byte(in.Body), &body)

	// 设置返回值
	out.Body = "收到消息：" + string([]byte(in.Body))

	return nil
}

// 其他方法，无法被api调用
func (a *Api1) Test() {
	log.Log("Api1.Test接口收到请求")
}
