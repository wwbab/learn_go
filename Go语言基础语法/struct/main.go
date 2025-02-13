package main

import "fmt"

// user 定义了一个用户结构体，包含用户名和密码字段
type user struct {
    name     string
    password string
}

func main() {
    // 创建一个名为 a 的 user 结构体实例，初始化其 name 和 password 字段
    a := user{name: "wang", password: "1024"}
    // 创建一个名为 b 的 user 结构体实例，使用顺序初始化其 name 和 password 字段
    b := user{"wang", "1024"}
    // 创建一个名为 c 的 user 结构体实例，只初始化其 name 字段
    c := user{name: "wang"}
    // 为结构体实例 c 的 password 字段赋值
    c.password = "1024"
    // 声明一个名为 d 的 user 结构体变量
    var d user
    // 为结构体变量 d 的 name 字段赋值
    d.name = "wang"
    // 为结构体变量 d 的 password 字段赋值
    d.password = "1024"

    // 打印结构体实例 a、b、c、d 的值
    fmt.Println(a, b, c, d)                 // {wang 1024} {wang 1024} {wang 1024} {wang 1024}
    // 调用 checkPassword 函数，检查用户 a 的密码是否为 "haha"，并打印结果
    fmt.Println(checkPassword(a, "haha"))   // false
    // 调用 checkPassword2 函数，检查用户 a 的密码是否为 "haha"，并打印结果
    fmt.Println(checkPassword2(&a, "haha")) // false
}

// checkPassword 函数接受一个 user 结构体实例和一个字符串作为参数，检查密码是否匹配，并返回布尔值
func checkPassword(u user, password string) bool {
    return u.password == password
}

// checkPassword2 函数接受一个指向 user 结构体的指针和一个字符串作为参数，检查密码是否匹配，并返回布尔值
func checkPassword2(u *user, password string) bool {
    return u.password == password
}