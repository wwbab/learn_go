// dup1 输出标准输入中出现次数大于 1 的行，前面是次数
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	// bufio.NewScanner 是 Go 语言中 bufio 包提供的一个函数，用于创建一个 Scanner 对象，
	// 该对象可以方便地对输入进行扫描，逐行、逐个单词或者按自定义分隔符读取数据。
	// 这在处理大文件或者流式输入时非常有用，因为它可以避免一次性将所有数据加载到内存中。
	// 标准输入（os.Stdin）
	input := bufio.NewScanner(os.Stdin)
	// 调用input.Scan()读取下一行，并且将结尾的换行符去掉，
	// Scan函数在读到新行的时候返回 true，
	// 在没有更多内容时返回 false。
	for input.Scan() {
		// 调用input.Text() 来获取读到的内容。
		counts[input.Text()]++
	}
	// 注意：忽略input.Err()中可能的错误
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}