package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

// main 函数作为程序的入口点
func main() {
    // 设置随机数的最大值为 100
    maxNum := 100
    // 以当前时间的纳秒值作为随机数种子
    rand.Seed(time.Now().UnixNano())
    // 生成一个 0 到 maxNum 之间的随机数，作为秘密数字
    secretNumber := rand.Intn(maxNum)
    // 打印秘密数字，用于调试
    fmt.Println("The secret number is ", secretNumber)

    // 提示用户输入猜测的数字
    fmt.Println("Please input your guess")
    // 从标准输入创建一个读取器
    reader := bufio.NewReader(os.Stdin)
    // 读取用户输入的字符串，直到遇到换行符为止
    input, err := reader.ReadString('\n')
    // 如果读取过程中发生错误，打印错误信息并返回
    if err != nil {
        fmt.Println("An error occured while reading input. Please try again", err)
        return
    }
    // 去掉输入字符串末尾的回车和换行符
    input = strings.Trim(input, "\r\n")

    // 将输入的字符串转换为整数
    guess, err := strconv.Atoi(input)
    // 如果转换失败，打印错误信息并返回
    if err != nil {
        fmt.Println("Invalid input. Please enter an integer value")
        return
    }
    // 打印用户猜测的数字
    fmt.Println("You guess is", guess)
}