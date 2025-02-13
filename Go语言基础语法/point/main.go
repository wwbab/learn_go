package main

import "fmt"

// add2 函数接受一个整数参数，将其值加 2，但不改变原始值
func add2(n int) {
    n += 2
}

// add2ptr 函数接受一个整数指针参数，将其指向的值加 2，改变原始值
func add2ptr(n *int) {
    *n += 2
}

func main() {
    // 初始化变量 n 为 5
    n := 5
    // 调用 add2 函数，传入 n 的值，但不会改变 n 的值
    add2(n)
    // 输出 n 的值，仍然是 5
    fmt.Println(n) // 5
    // 调用 add2ptr 函数，传入 n 的地址，会改变 n 的值
    add2ptr(&n)
    // 输出 n 的值，现在是 7
    fmt.Println(n) // 7
}