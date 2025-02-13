package main

import "fmt"

// main 函数是程序的入口点
func main() {

    // 声明一个长度为 5 的整型数组 a，并将其所有元素初始化为默认值 0
    var a [5]int
    // 将数组 a 的第五个元素赋值为 100
    a[4] = 100
    // 打印数组 a 的第三个元素的值
    fmt.Println("get:", a[2])
    // 打印数组 a 的长度
    fmt.Println("len:", len(a))

    // 声明一个长度为 5 的整型数组 b，并使用初始化列表进行初始化
    b := [5]int{1, 2, 3, 4, 5}
    // 打印数组 b 的所有元素
    fmt.Println(b)

    // 声明一个二维整型数组 twoD，大小为 2x3，并将其所有元素初始化为默认值 0
    var twoD [2][3]int
    // 使用嵌套循环遍历二维数组 twoD，并为每个元素赋值为其行和列的索引之和
    for i := 0; i < 2; i++ {
        for j := 0; j < 3; j++ {
            twoD[i][j] = i + j
        }
    }
    // 打印二维数组 twoD 的所有元素
    fmt.Println("2d: ", twoD)
}