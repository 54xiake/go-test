package main

import (
	"fmt"
	"net"
)

func main() {
	//服务器端ip和端口
	addr, _ := net.ResolveTCPAddr("tcp4", "localhost:8899")
	//申请连接客户端
	//第二个参数:本地地址  第三个参数:远程地址
	conn, _ := net.DialTCP("tcp4", nil, addr)
	//向服务端发送数据
	count, _ := conn.Write([]byte("客户端传递的数据"))
	fmt.Println("客户端向服务端发送的数据量为:", count)
	//通过休眠测试客户端对象不关闭,服务器是否能接收到对象
	//time.Sleep(10 * time.Second)
	// 关闭连接
	conn.Close()
	//fmt.Println("客户端结束")

}
