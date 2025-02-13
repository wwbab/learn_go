package main

import "fmt"

func main() {
    // 初始化变量i为1
    i := 1
    // 无限循环，打印一次"loop"后跳出循环
    for {
        fmt.Println("loop")
        break
    }
    // 遍历7到8的整数，打印每个数
    for j := 7; j < 9; j++ {
        fmt.Println(j)
    }

    // 遍历0到4的整数，打印奇数，遇到偶数则跳过
    for n := 0; n < 5; n++ {
        if n%2 == 0 {
            continue
        }
        fmt.Println(n)
    }
    // 当i小于等于3时，打印i并自增1
    for i <= 3 {
        fmt.Println(i)
        i = i + 1
    }
}