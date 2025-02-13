package main

import "fmt"

// add 函数接受两个整数参数 a 和 b，并返回它们的和
func add(a int, b int) int {
    return a + b
}

// add2 函数与 add 函数功能相同，但参数类型声明简化
func add2(a, b int) int {
    return a + b
}

// exists 函数检查给定的键 k 是否存在于 map m 中
// 如果键存在，它将返回对应的值 v 和布尔值 true
// 如果键不存在，它将返回零值和布尔值 false
func exists(m map[string]string, k string) (v string, ok bool) {
    v, ok = m[k]
    return v, ok
}

func main() {
    // 调用 add 函数并打印结果
    res := add(1, 2)
    fmt.Println(res) // 3

    // 调用 exists 函数并打印结果
    v, ok := exists(map[string]string{"a": "A"}, "a")
    fmt.Println(v, ok) // A True
}