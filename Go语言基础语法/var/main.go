package main

import (
	"fmt"
	"math"
)

// main 函数是程序的入口点
func main() {

    // 声明一个字符串变量 a，并初始化为 "initial"
    var a = "initial"

    // 声明两个整型变量 b 和 c，并分别初始化为 1 和 2
    var b, c int = 1, 2

    // 声明一个布尔型变量 d，并初始化为 true
    var d = true

    // 声明一个浮点型变量 e，未初始化，默认值为 0
    var e float64

    // 将变量 e 转换为 float32 类型，并赋值给变量 f
    f := float32(e)

    // 将变量 a 和字符串 "foo" 拼接，并赋值给变量 g
    g := a + "foo"

    // 打印变量 a、b、c、d、e、f 的值
    fmt.Println(a, b, c, d, e, f) // initial 1 2 true 0 0

    // 打印变量 g 的值
    fmt.Println(g) // initialapple

    // 声明一个字符串常量 s，并初始化为 "constant"
    const s string = "constant"

    // 声明一个整型常量 h，并初始化为 500000000
    const h = 500000000

    // 声明一个浮点型常量 i，其值为 3e20 除以 h
    const i = 3e20 / h

    // 打印常量 s、h、i 的值，以及 h 和 i 的正弦值
    fmt.Println(s, h, i, math.Sin(h), math.Sin(i))
}