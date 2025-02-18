package main

import (
	"bufio"
	"context"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"log"
	"net"
)

// 定义SOCKS5协议的版本号
const socks5Ver = 0x05
// 定义CONNECT请求的命令码
const cmdBind = 0x01
// 定义IPv4地址类型的代码
const atypeIPV4 = 0x01
// 定义域名地址类型的代码
const atypeHOST = 0x03
// 定义IPv6地址类型的代码
const atypeIPV6 = 0x04

func main() {
    // 在本地127.0.0.1的1080端口监听TCP连接
    server, err := net.Listen("tcp", "127.0.0.1:1080")
    if err != nil {
        // 如果监听失败，程序崩溃并输出错误信息
        panic(err)
    }
    // 进入无限循环，不断接受新的客户端连接
    for {
        // 接受一个新的客户端连接
        client, err := server.Accept()
        if err != nil {
            // 如果接受连接失败，记录错误信息并继续接受下一个连接
            log.Printf("Accept failed %v", err)
            continue
        }
        // 为每个客户端连接启动一个新的goroutine来处理
        go process(client)
    }
}

// process函数用于处理客户端连接，包括认证和连接目标服务器
func process(conn net.Conn) {
    // 确保在函数结束时关闭客户端连接
    defer conn.Close()
    // 创建一个带缓冲的读取器，用于从连接中读取数据
    reader := bufio.NewReader(conn)
    // 进行认证操作
    err := auth(reader, conn)
    if err != nil {
        // 如果认证失败，记录错误信息并返回
        log.Printf("client %v auth failed:%v", conn.RemoteAddr(), err)
        return
    }
    // 进行连接目标服务器的操作
    err = connect(reader, conn)
    if err != nil {
        // 如果连接目标服务器失败，记录错误信息并返回
        log.Printf("client %v auth failed:%v", conn.RemoteAddr(), err)
        return
    }
}

// auth函数用于处理SOCKS5协议的认证过程
func auth(reader *bufio.Reader, conn net.Conn) (err error) {
    // SOCKS5协议认证请求的格式：
    // +----+----------+----------+
    // |VER | NMETHODS | METHODS  |
    // +----+----------+----------+
    // | 1  |    1     | 1 to 255 |
    // +----+----------+----------+
    // VER: 协议版本，socks5为0x05
    // NMETHODS: 支持认证的方法数量
    // METHODS: 对应NMETHODS，NMETHODS的值为多少，METHODS就有多少个字节。RFC预定义了一些值的含义，内容如下:
    // X’00’ NO AUTHENTICATION REQUIRED
    // X’02’ USERNAME/PASSWORD

    // 读取协议版本号
    ver, err := reader.ReadByte()
    if err != nil {
        // 如果读取失败，返回错误信息
        return fmt.Errorf("read ver failed:%w", err)
    }
    if ver != socks5Ver {
        // 如果版本号不是SOCKS5，返回不支持的错误信息
        return fmt.Errorf("not supported ver:%v", ver)
    }
    // 读取支持的认证方法数量
    methodSize, err := reader.ReadByte()
    if err != nil {
        // 如果读取失败，返回错误信息
        return fmt.Errorf("read methodSize failed:%w", err)
    }
    // 创建一个字节切片，用于存储支持的认证方法
    method := make([]byte, methodSize)
    // 读取所有支持的认证方法
    _, err = io.ReadFull(reader, method)
    if err != nil {
        // 如果读取失败，返回错误信息
        return fmt.Errorf("read method failed:%w", err)
    }

    // SOCKS5协议认证响应的格式：
    // +----+--------+
    // |VER | METHOD |
    // +----+--------+
    // | 1  |   1    |
    // +----+--------+
    // 向客户端发送认证响应，表示不需要认证
    _, err = conn.Write([]byte{socks5Ver, 0x00})
    if err != nil {
        // 如果写入失败，返回错误信息
        return fmt.Errorf("write failed:%w", err)
    }
    return nil
}

// connect函数用于处理SOCKS5协议的连接请求
func connect(reader *bufio.Reader, conn net.Conn) (err error) {
    // SOCKS5协议连接请求的格式：
    // +----+-----+-------+------+----------+----------+
    // |VER | CMD |  RSV  | ATYP | DST.ADDR | DST.PORT |
    // +----+-----+-------+------+----------+----------+
    // | 1  |  1  | X'00' |  1   | Variable |    2     |
    // +----+-----+-------+------+----------+----------+
    // VER 版本号，socks5的值为0x05
    // CMD 0x01表示CONNECT请求
    // RSV 保留字段，值为0x00
    // ATYP 目标地址类型，DST.ADDR的数据对应这个字段的类型。
    //   0x01表示IPv4地址，DST.ADDR为4个字节
    //   0x03表示域名，DST.ADDR是一个可变长度的域名
    // DST.ADDR 一个可变长度的值
    // DST.PORT 目标端口，固定2个字节

    // 创建一个长度为4的字节切片，用于读取连接请求的头部信息
    buf := make([]byte, 4)
    // 读取连接请求的头部信息
    _, err = io.ReadFull(reader, buf)
    if err != nil {
        // 如果读取失败，返回错误信息
        return fmt.Errorf("read header failed:%w", err)
    }
    // 从头部信息中提取版本号、命令码和地址类型
    ver, cmd, atyp := buf[0], buf[1], buf[3]
    if ver != socks5Ver {
        // 如果版本号不是SOCKS5，返回不支持的错误信息
        return fmt.Errorf("not supported ver:%v", ver)
    }
    if cmd != cmdBind {
        // 如果命令码不是CONNECT请求，返回不支持的错误信息
        return fmt.Errorf("not supported cmd:%v", cmd)
    }
    // 用于存储目标地址的字符串
    addr := ""
    // 根据地址类型处理目标地址
    switch atyp {
    case atypeIPV4:
        // 读取4字节的IPv4地址
        _, err = io.ReadFull(reader, buf)
        if err != nil {
            // 如果读取失败，返回错误信息
            return fmt.Errorf("read atyp failed:%w", err)
        }
        // 将IPv4地址转换为字符串格式
        addr = fmt.Sprintf("%d.%d.%d.%d", buf[0], buf[1], buf[2], buf[3])
    case atypeHOST:
        // 读取域名的长度
        hostSize, err := reader.ReadByte()
        if err != nil {
            // 如果读取失败，返回错误信息
            return fmt.Errorf("read hostSize failed:%w", err)
        }
        // 创建一个字节切片，用于存储域名
        host := make([]byte, hostSize)
        // 读取域名
        _, err = io.ReadFull(reader, host)
        if err != nil {
            // 如果读取失败，返回错误信息
            return fmt.Errorf("read host failed:%w", err)
        }
        // 将域名转换为字符串格式
        addr = string(host)
    case atypeIPV6:
        // 不支持IPv6地址，返回错误信息
        return errors.New("IPv6: no supported yet")
    default:
        // 无效的地址类型，返回错误信息
        return errors.New("invalid atyp")
    }
    // 读取目标端口
    _, err = io.ReadFull(reader, buf[:2])
    if err != nil {
        // 如果读取失败，返回错误信息
        return fmt.Errorf("read port failed:%w", err)
    }
    // 将端口号从字节切片转换为uint16类型
    port := binary.BigEndian.Uint16(buf[:2])

    // 连接目标服务器
    dest, err := net.Dial("tcp", fmt.Sprintf("%v:%v", addr, port))
    if err != nil {
        // 如果连接失败，返回错误信息
        return fmt.Errorf("dial dst failed:%w", err)
    }
    // 确保在函数结束时关闭与目标服务器的连接
    defer dest.Close()
    // 记录连接的目标地址和端口
    log.Println("dial", addr, port)

    // SOCKS5协议连接响应的格式：
    // +----+-----+-------+------+----------+----------+
    // |VER | REP |  RSV  | ATYP | BND.ADDR | BND.PORT |
    // +----+-----+-------+------+----------+----------+
    // | 1  |  1  | X'00' |  1   | Variable |    2     |
    // +----+-----+-------+------+----------+----------+
    // VER socks版本，这里为0x05
    // REP Relay field,内容取值如下 X’00’ succeeded
    // RSV 保留字段
    // ATYPE 地址类型
    // BND.ADDR 服务绑定的地址
    // BND.PORT 服务绑定的端口DST.PORT
    // 向客户端发送连接响应，表示连接成功
    _, err = conn.Write([]byte{0x05, 0x00, 0x00, 0x01, 0, 0, 0, 0, 0, 0})
    if err != nil {
        // 如果写入失败，返回错误信息
        return fmt.Errorf("write failed: %w", err)
    }
    // 创建一个可取消的上下文
    ctx, cancel := context.WithCancel(context.Background())
    // 确保在函数结束时取消上下文
    defer cancel()

    // 启动一个goroutine，将客户端的数据转发到目标服务器
    go func() {
        _, _ = io.Copy(dest, reader)
        // 数据转发完成后，取消上下文
        cancel()
    }()
    // 启动一个goroutine，将目标服务器的数据转发到客户端
    go func() {
        _, _ = io.Copy(conn, dest)
        // 数据转发完成后，取消上下文
        cancel()
    }()

    // 等待上下文被取消
    <-ctx.Done()
    return nil
}
