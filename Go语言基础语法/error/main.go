package main

import (
	"errors"
	"fmt"
)

// user 定义了一个用户结构体，包含用户名和密码字段
type user struct {
    name     string
    password string
}

// findUser 在给定的用户切片中查找指定用户名的用户，如果找到则返回该用户的指针和 nil 错误，否则返回 nil 和一个自定义的错误
func findUser(users []user, name string) (v *user, err error) {
    for _, u := range users {
        if u.name == name {
            return &u, nil
        }
    }
    return nil, errors.New("not found")
}

func main() {
    // 在用户切片中查找名为 "wang" 的用户
    u, err := findUser([]user{{"wang", "1024"}}, "wang")
    if err != nil {
        // 如果发生错误，打印错误信息并返回
        fmt.Println(err)
        return
    }
    // 打印找到的用户的名字
    fmt.Println(u.name) // wang

    // 在用户切片中查找名为 "li" 的用户
    if u, err := findUser([]user{{"wang", "1024"}}, "li"); err != nil {
        // 如果发生错误，打印错误信息并返回
        fmt.Println(err) // not found
        return
    } else {
        // 如果没有错误，打印找到的用户的名字
        fmt.Println(u.name)
    }
}