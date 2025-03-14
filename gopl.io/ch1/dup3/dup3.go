package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	counts := make(map[string]int)
	for _, filename := range os.Args[1:] {
		// os.ReadFile函数用于一次性读取整个文件的内容并返回字节切片。
		data, err := os.ReadFile(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup3:%v\n", err)
			continue
		}
		// string.Split函数用于将字符串按照指定的分隔符分割成多个子字符串，
		// 并返回一个字符串切片。
		for _, line := range strings.Split(string(data), "\n") {
			counts[line]++
		}
	}
	for line, n := range counts {
		if n >1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}