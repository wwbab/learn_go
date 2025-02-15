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
    // 设置最大猜测数为100
    maxNum := 100
    // 使用当前时间的纳秒作为随机数种子
    rand.Seed(time.Now().UnixNano())
    // 生成一个0到maxNum之间的随机数作为秘密数字
    secretNumber := rand.Intn(maxNum)
    // fmt.Println("The secret number is ", secretNumber)

    // 提示用户输入猜测
    fmt.Println("Please input your guess")
    // 创建一个从标准输入读取的读取器
    reader := bufio.NewReader(os.Stdin)
    // 无限循环，直到用户猜对为止
    for {
        // 读取用户输入，直到遇到换行符为止
        input, err := reader.ReadString('\n')
        // 如果发生错误，打印错误并继续循环
        if err != nil {
            fmt.Println("An error occured while reading input. Please try again", err)
            continue
        }
        // 去掉输入字符串末尾的回车和换行符
        input = strings.Trim(input, "\r\n")

        // 将输入的字符串转换为整数
        guess, err := strconv.Atoi(input)
        // 如果转换失败，打印错误并继续循环
        if err != nil {
            fmt.Println("Invalid input. Please enter an integer value")
            continue
        }
        // 打印用户的猜测
        fmt.Println("You guess is", guess)
        // 如果猜测大于秘密数字，提示用户猜测过大
        if guess > secretNumber {
            fmt.Println("Your guess is bigger than the secret number. Please try again")
            // 如果猜测小于秘密数字，提示用户猜测过小
        } else if guess < secretNumber {
            fmt.Println("Your guess is smaller than the secret number. Please try again")
            // 如果猜测等于秘密数字，提示用户猜对并结束循环
        } else {
            fmt.Println("Correct, you Legend!")
            break
        }
    }
}

