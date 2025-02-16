package main

import (
	// 导入 bufio 包，用于高效地读取和写入数据
	"bufio"
	// 导入 log 包，用于记录日志
	"log"
	// 导入 net 包，用于网络编程
	"net"
)

// main 函数作为程序的入口点，监听本地的 1080 端口，接受客户端连接，并为每个客户端连接启动一个新的 goroutine 进行处理。
func main() {
    // 监听本地的 1080 端口
    server, err := net.Listen("tcp", "127.0.0.1:1080")
    // 如果监听失败，打印错误并 panic
    if err!= nil {
        panic(err)
    }
    // 无限循环，接受客户端连接
    for {  
        // 接受客户端连接
        client, err := server.Accept()
        // 如果接受连接失败，打印错误并继续循环
        if err!= nil {
            log.Printf("Accept failed %v", err)
            continue
        }
        // 为每个客户端连接启动一个新的 goroutine 进行处理
        go process(client)
    }
}

// process 函数用于处理客户端的连接，读取客户端发送的数据，并将其原样写回客户端。
func process(conn net.Conn) {
    // 延迟关闭连接
    defer conn.Close() 
    // 创建一个新的 bufio.Reader，用于从连接中读取数据
    reader := bufio.NewReader(conn)
    // 无限循环，读取客户端发送的数据
    for {
        // 从连接中读取一个字节的数据
        b, err := reader.ReadByte()
        // 如果读取失败，跳出循环
        if err!= nil {
            break
        }
        // 将读取到的字节数据写回客户端
        _, err = conn.Write([]byte{b})
        // 如果写入失败，跳出循环
        if err!= nil {
            break
        }
    }
}