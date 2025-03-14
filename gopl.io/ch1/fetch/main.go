// fetch 输出从URL获取的内容
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)


func main() {
	for _, url := range os.Args[1:] {
		// strings.HasPrefix 检查 url 的前缀是否为 “http://”
		if !strings.HasPrefix(url, "http://") {
			url = "http://" + url
		} 
		// http.Get 函数产生一个 HTTP 请求，
		// 如果没有出错， 
		// 返回结果存在响应结构体 resp 里面。
		resp, err := http.Get(url)
		if err !=nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
	/*	// resp 的 Body 域 包含服务器端响应的一个可读取数据流。
		// io.ReadAll 读取整个响应结果并存入 b，内存消耗大。
		b ,err := io.ReadAll(resp.Body)
		// 关闭 Body 数据流来避免资源泄漏。
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
		fmt.Printf("%s", b)
	*/
	// io.copy 从 resp.Body 读，并写入 os.Stdout (标准输出流) ，
	// 流式处理，内存消耗小。
	_, err = io.Copy(os.Stdout, resp.Body)
	// 输出 HTTP 状态码。
	fmt.Printf("\nHTTP status: %d", resp.StatusCode)
	resp.Body.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
		os.Exit(1)
	}
	}

		
		

}
