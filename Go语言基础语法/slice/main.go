package main

import "fmt"

func main() {
    // 创建一个长度为3的字符串切片
    s := make([]string, 3)
    // 向切片中添加元素
    s[0] = "a"
    s[1] = "b"
    s[2] = "c"
    // 打印切片的第三个元素
    fmt.Println("get:", s[2]) // c
    // 打印切片的长度
    fmt.Println("len:", len(s)) // 3

    // 向切片中添加一个元素
    s = append(s, "d")
    // 向切片中添加两个元素
    s = append(s, "e", "f")
    // 打印切片
    fmt.Println(s) // [a b c d e f]

    // 创建一个新的切片，长度与s相同
    c := make([]string, len(s))
    // 将s中的元素复制到c中
    copy(c, s)
    // 打印切片c
    fmt.Println(c) // [a b c d e f]

    // 打印切片s的第3个到第5个元素
    fmt.Println(s[2:5]) // [c d e]
    // 打印切片s的前5个元素
    fmt.Println(s[:5]) // [a b c d e]
    // 打印切片s的第3个元素到最后一个元素
    fmt.Println(s[2:]) // [c d e f]

    // 创建一个字符串切片，包含"g", "o", "o", "d"
    good := []string{"g", "o", "o", "d"}
    // 打印切片good
    fmt.Println(good) // [g o o d]
}