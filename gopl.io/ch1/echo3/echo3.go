package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// strings.Join 是 Go 语言里 strings 包提供的一个实用函数，
	// 其功能是把字符串切片中的所有元素用指定的分隔符连接起来，最后返回一个连接好的字符串。
	// func Join(a []string, sep string) string
	fmt.Println(strings.Join(os.Args[1:],""))
	fmt.Println(os.Args[1:])
	fmt.Println(os.Args[0]) // os.Args 的第一个元素是 os.Args[0],它是命令本身的名字，
							// 另外的元素是程序开始执行时的参数。
}