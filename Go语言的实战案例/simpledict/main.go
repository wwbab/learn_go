package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

// 定义一个结构体，用于存储 JSON 数据
type DictRequest struct {
	TransType string `json:"trans_type"`
	Source    string `json:"source"`
	User_id   string `json:"user_id"`
}

// main 函数是程序的入口点
func main() {
    client := &http.Client{}
	request := DictRequest{TransType: "en2zh", Source: "good"}
	buf, err := json.Marshal(request)
	// 检查错误是否发生
if err!= nil {
    // 如果发生错误，使用 log.Fatal 记录错误并终止程序
    log.Fatal(err)
}
    // 创建一个新的 Reader，用于读取 JSON 数据
    var data = bytes.NewReader(buf)
    // 创建一个 HTTP POST 请求，目标 URL 是彩云小译的翻译 API
    req, err := http.NewRequest("POST", "https://api.interpreter.caiyunai.com/v1/dict", data)
    // 如果创建请求时发生错误，记录错误并终止程序
    if err!= nil {
        log.Fatal(err)
    }
    // 设置请求头，接受 JSON、纯文本和所有类型的响应
    req.Header.Set("accept", "application/json, text/plain, */*")
    // 设置请求头，接受语言为中文
    req.Header.Set("accept-language", "zh")
    // 设置请求头，应用名称为 xiaoyi
    req.Header.Set("app-name", "xiaoyi")
    // 设置请求头，授权类型为 Bearer
    req.Header.Set("authorization", "Bearer")
    // 设置请求头，内容类型为 JSON，字符编码为 UTF-8
    req.Header.Set("content-type", "application/json;charset=UTF-8")
    // 设置请求头，设备 ID 为空
    req.Header.Set("device-id", "")
    // 设置请求头，源地址为彩云小译的翻译网页
    req.Header.Set("origin", "https://fanyi.caiyunapp.com")
    // 设置请求头，操作系统类型为 web
    req.Header.Set("os-type", "web")
    // 设置请求头，操作系统版本为空
    req.Header.Set("os-version", "")
    // 设置请求头，优先级为 u=1, i
    req.Header.Set("priority", "u=1, i")
    // 设置请求头，引用地址为彩云小译的翻译网页
    req.Header.Set("referer", "https://fanyi.caiyunapp.com/")
    // 设置请求头，浏览器标识为 Microsoft Edge 131, Chromium 131, Not_A Brand 24
    req.Header.Set("sec-ch-ua", `"Microsoft Edge";v="131", "Chromium";v="131", "Not_A Brand";v="24"`)
    // 设置请求头，是否为移动设备为否
    req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("sec-ch-ua-platform", `"Windows"`)
	req.Header.Set("sec-fetch-dest", "empty")
	req.Header.Set("sec-fetch-mode", "cors")
	req.Header.Set("sec-fetch-site", "cross-site")
	req.Header.Set("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/131.0.0.0 Safari/537.36 Edg/131.0.0.0")
	req.Header.Set("x-authorization", "token:qgemv4jr1y38jyq6vhvi")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", bodyText)
}