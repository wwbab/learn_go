package main

import (
	"fmt"
	"strings"
)

func main() {
    a := "hello"
    // 判断字符串a是否包含子串"ll"
    fmt.Println(strings.Contains(a, "ll"))
    // 统计字符串a中子串"l"出现的次数
    fmt.Println(strings.Count(a, "l"))
    // 判断字符串a是否以"he"为前缀
    fmt.Println(strings.HasPrefix(a, "he"))
    // 判断字符串a是否以"llo"为后缀
    fmt.Println(strings.HasSuffix(a, "llo"))
    // 获取字符串a中子串"ll"第一次出现的位置
    fmt.Println(strings.Index(a, "ll"))
    // 使用"-"连接字符串数组
    fmt.Println(strings.Join([]string{"he", "llo"}, "-"))
    // 重复字符串a两次
    fmt.Println(strings.Repeat(a, 2))
    // 将字符串a中的"e"替换为"E"，替换所有匹配的字符
    fmt.Println(strings.Replace(a, "e", "E", -1))
    // 按照"-"分割字符串
    fmt.Println(strings.Split("a-b-c", "-"))
    // 将字符串a转换为小写
    fmt.Println(strings.ToLower(a))
    // 将字符串a转换为大写
    fmt.Println(strings.ToUpper(a))
    // 获取字符串a的长度
    fmt.Println(len(a))
    // 定义一个包含中文字符的字符串b
    b := "你好"
    // 获取字符串b的长度，由于中文字符占用多个字节，所以长度为6
    fmt.Println(len(b))
}