package main

import (
	"fmt"
	"strconv"
)

// main 函数是程序的入口点
func main() {
    // 将字符串 "1.234" 解析为 64 位浮点数
    f, _ := strconv.ParseFloat("1.234", 64)
    fmt.Println(f) // 输出解析后的浮点数 1.234

    // 将字符串 "111" 解析为 10 进制的 64 位整数
    n, _ := strconv.ParseInt("111", 10, 64)
    fmt.Println(n) // 输出解析后的整数 111

    // 将字符串 "0x1000" 解析为 0 进制（自动识别为 16 进制）的 64 位整数
    n, _ = strconv.ParseInt("0x1000", 0, 64)
    fmt.Println(n) // 输出解析后的整数 4096

    // 将字符串 "123" 转换为整数
    n2, _ := strconv.Atoi("123")
    fmt.Println(n2) // 输出转换后的整数 123

    // 将字符串 "AAA" 转换为整数，由于无法转换，会返回错误
    n2, err := strconv.Atoi("AAA")
    fmt.Println(n2, err) // 输出转换结果 0 和错误信息 "strconv.Atoi: parsing "AAA": invalid syntax"
}