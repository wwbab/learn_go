package main

import "fmt"

// user 定义了一个用户结构体，包含用户名和密码字段
type user struct {
    name     string
    password string
}

// checkPassword 是 user 结构体的一个方法，用于检查给定的密码是否与用户的密码匹配
func (u user) checkPassword(password string) bool {
    return u.password == password
}

// resetPassword 是 user 结构体的一个方法，用于重置用户的密码
func (u *user) resetPassword(password string) {
    u.password = password
}

func main() {
    // 创建一个名为 "wang"，密码为 "1024" 的用户实例
    a := user{name: "wang", password: "1024"}
    // 重置用户实例 a 的密码为 "2048"
    a.resetPassword("2048")
    // 打印检查密码是否为 "2048" 的结果，预期输出为 true
    fmt.Println(a.checkPassword("2048")) // true
}