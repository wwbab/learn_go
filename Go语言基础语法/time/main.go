package main

import (
	"fmt"
	"time"
)

// main 函数是程序的入口点
func main() {
    // 获取当前时间
    now := time.Now()
    // 打印当前时间
    fmt.Println(now) // 2022-03-27 18:04:59.433297 +0800 CST m=+0.000087933
    // 创建一个指定日期和时间的 Time 对象
    t := time.Date(2022, 3, 27, 1, 25, 36, 0, time.UTC)
    // 创建另一个指定日期和时间的 Time 对象
    t2 := time.Date(2022, 3, 27, 2, 30, 36, 0, time.UTC)
    // 打印 t 的日期和时间
    fmt.Println(t)                                                  // 2022-03-27 01:25:36 +0000 UTC
    // 打印 t 的年、月、日、时、分
    fmt.Println(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute()) // 2022 March 27 1 25
    // 使用指定的格式字符串格式化 t 的日期和时间，并打印结果
    fmt.Println(t.Format("2006-01-02 15:04:05"))                    // 2022-03-27 01:25:36
    // 计算 t2 和 t 之间的时间差
    diff := t2.Sub(t)
    // 打印时间差
    fmt.Println(diff)                           // 1h5m0s
    // 打印时间差的分钟数和秒数
    fmt.Println(diff.Minutes(), diff.Seconds()) // 65 3900
    // 解析一个字符串表示的日期和时间，并将其转换为 Time 对象
    t3, err := time.Parse("2006-01-02 15:04:05", "2022-03-27 01:25:36")
    // 如果发生错误，抛出异常
    if err != nil {
        panic(err)
    }
    // 打印 t3 是否等于 t
    fmt.Println(t3 == t)    // true
    // 打印当前时间的 Unix 时间戳（秒数）
    fmt.Println(now.Unix()) // 1648738080
}