// dup2 打印输入中多次出现的行的个数和文本
// 它从 stdin 或指定的文件列表读取
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			// os.Open()主要功能是打开一个文件并返回一个文件对象，
			// 之后就可以对这个文件进行读取操作,
			// 返回两个值。第一个是打开的文件（*os.File），该文件随后被Scanner读取
			// 第二个返回值是error类型。
			f, err := os.Open(arg)
			// 如果出错在标准错误流（os.Stderr）上输出一条消息
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts)
			// 文件读到结尾时，Close函数关闭文件，
			// 然后释放相应的资源（如内存）。
			f.Close()
		}
	}

	for line, n := range counts {
		if n > 1 {
			
		
			fmt.Printf("%d\t%s\n", n, line)
		}
	}

}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
	// 注意忽略input.Err() 中可能的错误 
}