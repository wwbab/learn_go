package main

import "fmt"

// main 函数是程序的入口点
func main() {
    // 使用 if 语句检查 7 是否为偶数
    if 7%2 == 0 {
        fmt.Println("7 is even")
    } else {
        // 如果 7 不是偶数，则打印它是奇数
        fmt.Println("7 is odd")
    }

    // 使用 if 语句检查 8 是否能被 4 整除
    if 8%4 == 0 {
        fmt.Println("8 is divisible by 4")
    }

    // 使用 if 语句检查 num 是否小于 0
    if num := 9; num < 0 {
        fmt.Println(num, "is negative")
    } else if num < 10 {
        // 如果 num 在 0 到 10 之间，则打印它是一位数
        fmt.Println(num, "has 1 digit")
    } else {
        // 如果 num 大于等于 10，则打印它是多位数
        fmt.Println(num, "has multiple digits")
    }
}