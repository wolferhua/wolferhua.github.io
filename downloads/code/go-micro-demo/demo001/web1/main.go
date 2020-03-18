package main

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strings"

	"github.com/gorilla/websocket"

	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/transport/grpc"
	"github.com/micro/go-micro/v2/util/log"
	"github.com/micro/go-micro/v2/web"
)

// websocket 服务
var upgrader = websocket.Upgrader{CheckOrigin: func(r *http.Request) bool {
	return true
}}

var sysName = `系统通知：`

// 消息列表
type Message struct {
	Code    int    `json:"code"`
	Name    string `json:"name"`
	Type    int    `json:"type"`
	Message string `json:"message"`
}

var messages = make(chan *Message, 100)

// 链接列表

var connections = map[string]*websocket.Conn{}

func main() {
	// 这里使用的是go-micro web 包，因为它已经实现了服务注册，和发现的能力。不需要我们再去处理了。
	service := web.NewService(
		web.Name("com.anycps.wolferhua.api.web1"),
		web.MicroService(micro.NewService(micro.Transport(grpc.NewTransport()))),
	)
	service.Options().Service.Client()
	// 初始化服务
	if err := service.Init(); err != nil {
		log.Fatal("Init", err)
	}

	// 路由注册

	// 注册静态资源目录（页面）
	service.Handle("/web1/", http.StripPrefix("/web1/", http.FileServer(http.Dir("html"))))

	// 注册websocket监听（socket）
	service.HandleFunc("/web1/ws", ws)

	// 获取用户列表接口（接口）
	service.HandleFunc("/web1/ws/users", users)

	go wsSend()
	// 启动服务
	if err := service.Run(); err != nil {
		log.Fatal("Run: ", err)
	}
}

// 获取用户列表接口
func users(w http.ResponseWriter, r *http.Request) {
	// 读取列表
	list := make([]string, 0, len(connections))
	for name, _ := range connections {
		list = append(list, name)
	}
	// 返回数据

	w.Header().Add("Content-Type", "application/json; charset=utf-8")

	w.WriteHeader(200)

	data, _ := json.Marshal(list)
	w.Write(data)
}

// websocket 服务
func ws(w http.ResponseWriter, r *http.Request) {
	vals, err := url.ParseQuery(r.URL.RawQuery)
	name := strings.TrimSpace(vals.Get("name"))

	//fmt.Println(vals.Get("name"))

	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Errorf("upgrade: %s", err)
		return
	}
	// 退出
	defer c.Close()

	// 判断内容
	if len(name) < 1 {
		// 返回错误信息
		// http.Error(w, `{"code":500,"message":"请输入名称！"}`, 500)
		c.WriteMessage(websocket.TextMessage, []byte(`{"code":500,"message":"请输入名称！"}`))

		return
	}
	// 验证重复状态
	if _, ok := connections[name]; ok {
		// 返回错误信息
		// http.Error(w, `{"code":500,"message":"请输入名称！"}`, 500)
		c.WriteMessage(websocket.TextMessage, []byte(`{"code":501,"message":"名称已经存在了！"}`))
		return
	}

	defer func() {

		// 退出客户列表
		delete(connections, name)

		messages <- &Message{Type: websocket.TextMessage, Code: 119, Name: sysName, Message: ` ` + name + ` 退出了 `}
		// 关闭链接
		c.Close()
	}()
	connections[name] = c
	messages <- &Message{Type: websocket.TextMessage, Code: 120, Name: sysName, Message: ` ` + name + ` 上线了 `}

	for {
		// 读取消息
		mt, message, err := c.ReadMessage()
		if err != nil {
			log.Error("read:", err)
			break
		}
		log.Errorf("recv: %s", message)
		// 写入消息列表
		messages <- &Message{Name: name, Code: 200, Type: mt, Message: string(message)}
	}
}

// 消息消费
func wsSend() {
	// 读取消息
	for message := range messages {
		// 遍历发给所有人。
		for name, conn := range connections {
			if conn != nil {
				err := conn.WriteJSON(message)
				if err != nil {
					log.Error("send to ", name, "err :", err)
				}
			}
		}
	}
}
