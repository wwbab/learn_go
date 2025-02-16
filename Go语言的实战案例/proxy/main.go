package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
)

// SOCKS5协议版本常量定义
const (
	socks5Ver   = 0x05    // SOCKS5协议版本号
	cmdBind     = 0x01    // 建立连接请求命令（实际未在本代码中使用）
	atypeIPV4   = 0x01    // IPv4地址类型
	atypeHOST   = 0x03    // 域名地址类型
	atypeIPV6   = 0x04    // IPv6地址类型
)

func main() {
	// 在本地1080端口启动TCP监听（标准SOCKS5端口）
	server, err := net.Listen("tcp", "127.0.0.1:1080")
	if err != nil {
		panic(err) // 启动失败直接终止程序
	}
	
	// 主循环处理客户端连接
	for {
		client, err := server.Accept()
		if err != nil {
			log.Printf("Accept failed %v", err) // 记录错误但保持服务运行
			continue
		}
		// 为每个客户端连接创建独立goroutine处理
		go process(client)
	}
}

// 处理客户端连接主流程
func process(conn net.Conn) {
	defer conn.Close() // 确保连接最终关闭
	reader := bufio.NewReader(conn) // 创建带缓冲的读取器
	
	// 认证阶段处理
	if err := auth(reader, conn); err != nil {
		log.Printf("client %v auth failed:%v", conn.RemoteAddr(), err)
		return
	}
	log.Println("auth success")
	// 此处可添加后续请求处理逻辑（当前代码仅完成认证阶段）
}

// 处理SOCKS5认证流程
func auth(reader *bufio.Reader, conn net.Conn) (err error) {
	/* 认证请求报文格式（客户端 -> 服务端）
	+----+----------+----------+
	|VER | NMETHODS | METHODS  |
	+----+----------+----------+
	| 1  |    1     | 1 to 255 |
	+----+----------+----------+
	VER: 协议版本，SOCKS5应为0x05
	NMETHODS: 支持的认证方法数量
	METHODS: 认证方法列表，每个方法占1字节
	*/
	
	// 读取协议版本
	ver, err := reader.ReadByte()
	if err != nil {
		return fmt.Errorf("read ver failed:%w", err)
	}
	if ver != socks5Ver {
		return fmt.Errorf("not supported ver:%v", ver) // 版本不兼容错误
	}
	
	// 读取方法数量
	methodSize, err := reader.ReadByte()
	if err != nil {
		return fmt.Errorf("read methodSize failed:%w", err)
	}
	
	// 读取所有支持的认证方法
	method := make([]byte, methodSize)
	_, err = io.ReadFull(reader, method)
	if err != nil {
		return fmt.Errorf("read method failed:%w", err)
	}
	log.Println("client auth methods:", "ver", ver, "method", method)
	
	/* 认证响应报文格式（服务端 -> 客户端）
	+----+--------+
	|VER | METHOD |
	+----+--------+
	| 1  |   1    |
	+----+--------+
	METHOD: 选择的认证方法，0x00表示无需认证
	*/
	
	// 发送响应：选择无需认证（0x00）
	_, err = conn.Write([]byte{socks5Ver, 0x00})
	if err != nil {
		return fmt.Errorf("write response failed:%w", err)
	}
	return nil
}

/*
当前代码实现功能：
1. 实现SOCKS5协议认证阶段的握手处理
2. 支持无认证方式（始终返回0x00方法）
3. 基础错误处理和日志记录

待完善功能：
1. 命令解析（CONNECT/BIND/UDP ASSOCIATE）
2. 目标地址解析（IPv4/IPv6/域名处理）
3. 数据转发功能
4. 错误状态码返回（如认证失败等）
5. UDP支持
*/