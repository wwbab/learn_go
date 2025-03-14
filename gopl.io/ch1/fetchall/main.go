// fetchall 并发获取 URL 并报告他们的时间和大小
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)


func main() {
	start := time.Now()
	// 创建字符串通道
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		// goroutine 轻量级（2KB栈空间）；
		// 高效调度（调度器自动调度）；
		// 通信安全（利用 channel 传递数据）。
		// main 函数在一个 goroutine 中执行，
		// go 语句创建额外的 goroutine。
		go fetch(url, ch)
	}
	for range os.Args[1:] {
		// 从通道 ch 接收
		fmt.Println(<-ch)
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		// 发送到通道 ch
		ch <- fmt.Sprint(err)
		return
	}
	// io.Discard 忽略所有写入的数据，
	// 也就是把写入的数据丢弃掉，
	// 不会进行实际的存储或者处理。
	// io.Copy 返回字节数以及可能出现的错误。
	nbytes, err := io.Copy(io.Discard, resp.Body)
	// 避免资源泄漏
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	// 发送汇总信息到通道 ch。
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)

}