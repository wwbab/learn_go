package main

import (
	"fmt"
	"time"
)

func main() {

    // 定义一个变量 a 并赋值为 2
    a := 2
    switch a {
    case 1:
        // 如果 a 的值为 1，则打印 "one"
        fmt.Println("one")
    case 2:
        // 如果 a 的值为 2，则打印 "two"
        fmt.Println("two")
    case 3:
        // 如果 a 的值为 3，则打印 "three"
        fmt.Println("three")
    case 4, 5:
        // 如果 a 的值为 4 或 5，则打印 "four or five"
        fmt.Println("four or five")
    default:
        // 如果 a 的值不是上述值，则打印 "other"
        fmt.Println("other")
    }

    // 获取当前时间
    t := time.Now()
    switch {
    case t.Hour() < 12:
        // 如果当前时间小于中午 12 点，则打印 "It's before noon"
        fmt.Println("It's before noon")
    default:
        // 如果当前时间大于等于中午 12 点，则打印 "It's after noon"
        fmt.Println("It's after noon")
    }
}