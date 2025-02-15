package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

// 定义一个结构体，用于存储 JSON 数据
type DictRequest struct {
	TransType string `json:"trans_type"`
	Source    string `json:"source"`
	User_id   string `json:"user_id"`
}
// 定义一个结构体，用于存储 JSON 响应数据
type DictResponse struct {
	Rc int `json:"rc"`
	Wiki struct {
	} `json:"wiki"`
	Dictionary struct {
		Prons struct {
			EnUs string `json:"en-us"`
			En string `json:"en"`
		} `json:"prons"`
		Explanations []string `json:"explanations"`
		Synonym []string `json:"synonym"`
		Antonym []string `json:"antonym"`
		WqxExample [][]string `json:"wqx_example"`
		Entry string `json:"entry"`
		Type string `json:"type"`
		Related []interface{} `json:"related"`
		Source string `json:"source"`
	} `json:"dictionary"`
}

// 定义一个函数，用于查询单词并打印结果
func query(word string) {
    client := &http.Client{}
	request := DictRequest{TransType: "en2zh", Source: word}
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
	//检查响应状态码是否为 200
	if resp.StatusCode != 200 {
		log.Fatal("bad Statuscode:",resp.StatusCode,"body",string(bodyText))
	}
	// 定义一个名为 dictResponse 的变量，类型为 DictResponse
	var dictResponse DictResponse
	// 将 JSON 格式的 bodyText 解析到 dictResponse 变量中
	err = json.Unmarshal(bodyText, &dictResponse)
	// 如果解析过程中发生错误，记录错误并终止程序
	if err!= nil {
    	log.Fatal(err)
	}
	// 打印单词的英式和美式发音
	fmt.Println(word, "UK:", dictResponse.Dictionary.Prons.En, "US:", dictResponse.Dictionary.Prons.EnUs)
	// 遍历解释列表并打印每个解释项
	for _, item := range dictResponse.Dictionary.Explanations {
    	fmt.Println(item)
		}
}
func main() {
    // 检查命令行参数数量是否正确
    if len(os.Args) != 2 {
        // 如果参数数量不正确，打印使用帮助信息到标准错误输出
        fmt.Fprintf(os.Stderr, `usage: simpleDict WORD
example: simpleDict hello
        `)
        // 程序以非零状态码 1 退出，表示有错误发生
        os.Exit(1)
    }
    // 获取命令行参数中的单词
    word := os.Args[1]
    // 调用 query 函数查询单词
    query(word)
}