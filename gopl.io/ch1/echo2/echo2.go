// echo2 输出其命令行参数
package main

import (
	"fmt"
	"os"
)

func main() {
	s, seq := "", ""
	// 每次迭代，range产生一对值：索引和索引处的元素值
	// Go语言不允许出现无用的临时变量
	for _, arg := range os.Args[1:] {
		s += seq + arg
		seq = ""
	}
	fmt.Println(s)

	for i, arg := range os.Args[:] {
		fmt.Println(i, arg)
	}
}