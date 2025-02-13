package main

import "fmt"

// main 函数是程序的入口点
func main() {
    // 使用 make 函数创建一个空的 map，键的类型是 string，值的类型是 int
    m := make(map[string]int)
    // 向 map 中添加键值对
    m["one"] = 1
    m["two"] = 2
    // 打印 map 的内容
    fmt.Println(m)           // map[one:1 two:2]
    // 打印 map 的长度
    fmt.Println(len(m))      // 2
    // 根据键获取对应的值
    fmt.Println(m["one"])    // 1
    // 尝试获取不存在的键的值，返回该类型的零值
    fmt.Println(m["unknow"]) // 0

    // 使用 ok-idiom 来安全地获取键的值，并检查键是否存在
    r, ok := m["unknow"]
    // 打印获取到的值和键是否存在的布尔值
    fmt.Println(r, ok) // 0 false

    // 使用 delete 函数从 map 中删除指定的键值对
    delete(m, "one")

    // 使用字面量创建并初始化一个 map
    m2 := map[string]int{"one": 1, "two": 2}
    // 使用 var 关键字声明并初始化一个 map
    var m3 = map[string]int{"one": 1, "two": 2}
    // 打印两个 map 的内容
    fmt.Println(m2, m3)
}