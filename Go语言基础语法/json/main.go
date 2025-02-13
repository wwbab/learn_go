package main

import (
	"encoding/json"
	"fmt"
)

// userInfo 定义了一个用户信息的结构体
type userInfo struct {
    Name  string
    Age   int    `json:"age"` // 结构体标签，指定在JSON编码时使用"age"作为字段名
    Hobby []string
}

func main() {
    // 创建一个 userInfo 类型的实例 a
    a := userInfo{Name: "wang", Age: 18, Hobby: []string{"Golang", "TypeScript"}}

    // 将实例 a 编码为JSON格式
    buf, err := json.Marshal(a)
    if err != nil {
        panic(err)
    }
    // 打印JSON编码后的字节切片
    fmt.Println(buf)         // [123 34 78 97...]
    // 将字节切片转换为字符串并打印
    fmt.Println(string(buf)) // {"Name":"wang","age":18,"Hobby":["Golang","TypeScript"]}

    // 将实例 a 编码为JSON格式，并使用缩进格式化输出
    buf, err = json.MarshalIndent(a, "", "\t")
    if err != nil {
        panic(err)
    }
    // 打印格式化后的JSON字符串
    fmt.Println(string(buf))

    // 创建一个 userInfo 类型的实例 b
    var b userInfo
    // 将JSON字符串解码到实例 b 中
    err = json.Unmarshal(buf, &b)
    if err != nil {
        panic(err)
    }
    // 使用格式化输出打印实例 b 的详细信息
    fmt.Printf("%#v\n", b) // main.userInfo{Name:"wang", Age:18, Hobby:[]string{"Golang", "TypeScript"}}
}