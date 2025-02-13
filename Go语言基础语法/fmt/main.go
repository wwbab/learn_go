package main

import "fmt"

// 定义一个结构体 point，包含两个整型字段 x 和 y
type point struct {
    x, y int
}

func main() {
    // 定义一个字符串变量 s，初始值为 "hello"
    s := "hello"
    // 定义一个整型变量 n，初始值为 123
    n := 123
    // 定义一个 point 结构体变量 p，初始值为 {1, 2}
    p := point{1, 2}
    // 使用 Println 函数打印 s 和 n 的值
    fmt.Println(s, n) // hello 123
    // 使用 Println 函数打印 p 的值
    fmt.Println(p)    // {1 2}

    // 使用 Printf 函数格式化打印 s 的值
    fmt.Printf("s=%v\n", s)  // s=hello
    // 使用 Printf 函数格式化打印 n 的值
    fmt.Printf("n=%v\n", n)  // n=123
    // 使用 Printf 函数格式化打印 p 的值
    fmt.Printf("p=%v\n", p)  // p={1 2}
    // 使用 Printf 函数格式化打印 p 的值，并显示字段名
    fmt.Printf("p=%+v\n", p) // p={x:1 y:2}
    // 使用 Printf 函数格式化打印 p 的值，并显示结构体类型
    fmt.Printf("p=%#v\n", p) // p=main.point{x:1, y:2}

    // 定义一个浮点型变量 f，初始值为 3.141592653
    f := 3.141592653
    // 使用 Println 函数打印 f 的值
    fmt.Println(f)          // 3.141592653
    // 使用 Printf 函数格式化打印 f 的值，保留两位小数
    fmt.Printf("%.2f\n", f) // 3.14
}