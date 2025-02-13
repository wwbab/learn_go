package main

import "fmt"

// main 函数是程序的入口点
func main() {
    // 定义一个整数切片 nums
    nums := []int{2, 3, 4}
    // 初始化变量 sum 用于存储总和
    sum := 0
    // 使用 range 循环遍历切片 nums
    for i, num := range nums {
        // 将当前元素加到 sum 上
        sum += num
        // 如果当前元素是 2，打印其索引和值
        if num == 2 {
            fmt.Println("index:", i, "num:", num) // index: 0 num: 2
        }
    }
    // 打印总和
    fmt.Println(sum) // 9

    // 定义一个映射 m，键和值都是字符串类型
    m := map[string]string{"a": "A", "b": "B"}
    // 使用 range 循环遍历映射 m 的键值对
    for k, v := range m {
        // 打印键和值
        fmt.Println(k, v) // b 8; a A
    }
    // 使用 range 循环遍历映射 m 的键
    for k := range m {
        // 打印键
        fmt.Println("key", k) // key a; key b
    }
}