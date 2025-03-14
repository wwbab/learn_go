// server2 是一个迷你的回声和计数器服务器
package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)


var mu sync.Mutex
var count int

func main() {
	// 回声请求调用处理程序
	// http.HandleFunc 将特定的 URL 路径与对应的处理函数绑定。
	http.HandleFunc("/", handler)
	http.HandleFunc("/count", counter)
	// log.Fatal 用于输出日志信息并终止程序的执行。
	// http.ListenAndServe 的作用是启动一个 HTTP 服务器，
	// 监听指定的地址和端口，
	// 并处理客户端的 HTTP 请求。
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// 处理程序回显请求 URL r 的路径部分
// http.ResponseWriter 用于向客户端发送响应
// *http.Request 则包含了客户端的请求信息。
func handler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	count++
	mu.Unlock()
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

// counter 回显目前为止的调用次数
func counter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "Count %d\n", count)
	mu.Unlock()
}

