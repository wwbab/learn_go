package main

import (
	"fmt"
	"os"
	"os/exec"
)

// main 函数是程序的入口点
func main() {
    // 打印命令行参数
    fmt.Println(os.Args)           // [/var/folders/8p/n34xxfnx38dg8bv_x8l62t_m0000gn/T/go-build3406981276/b001/exe/main a b c d]
    // 打印环境变量 PATH 的值
    fmt.Println(os.Getenv("PATH")) // /usr/local/go/bin...
    // 设置环境变量 AA 的值为 BB
    fmt.Println(os.Setenv("AA", "BB"))

    // 执行 grep 命令，在 /etc/hosts 文件中查找 "127.0.0.1"
    buf, err := exec.Command("grep", "127.0.0.1", "/etc/hosts").CombinedOutput()
    // 如果执行命令时发生错误，抛出异常
    if err != nil {
        panic(err)
    }
    // 打印 grep 命令的输出结果
    fmt.Println(string(buf)) // 127.0.0.1       localhost
}