package main

import ( 
	"os"
	"log"
    "net/http"
)

type helloHandler struct{}
var host = ""
func (h *helloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=UTF-8" )
	w.Header().Set("Content-Transfer-Encoding", "quoted-printable")
    w.Write([]byte("Host："+host+"<br> Hello, world! from golang.")) //输出内容
}

func main() {
	host, _ = os.Hostname() //获取当前机器名称
	http.Handle("/", &helloHandler{})
	log.Println("Start web service: http://127.0.0.1:8080")
    log.Fatal(http.ListenAndServe(":8080", nil)) //指定端口
} 