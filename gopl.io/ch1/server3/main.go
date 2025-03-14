// server3
package main

import (
	"fmt"
	"log"
	"net/http"
)


func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000",nil))
}

// 处理程序回显 HTTP 请求
func handler(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf 的作用是将格式化的字符串输出到一个 io.Writer 接口类型的对象里,
	// 返回值成功写入的字节数，err
	fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)
	for k, v := range r.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}
	fmt.Fprintf(w, "Host = %q\n", r.Host)
	fmt.Fprintf(w, "RemoteAddr = %q\n", r.RemoteAddr)
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	for k, v := range r.Form {
		fmt.Fprintf(w, "Form[%q] = %q\n", k, v)
	}
}
