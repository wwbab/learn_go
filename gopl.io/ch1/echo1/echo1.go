// echo1 输出其命令行参数
package main

import (
	"fmt"
	"os"
)


func main() {
	var s, seq string
	for i := 1; i < len(os.Args); i++ {
		s += seq + os.Args[i]
		seq = ""
	}
	fmt.Println(s)
}