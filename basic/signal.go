package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

//在 go 的安装目录修改 Go\src\syscall\types_windows.go，增加如下代码：
//var signals = [...]string{
//	// 这里省略N行。。。。
//
//	/** 找到此位置添加如下 */
//	16: "SIGUSR1",
//	17: "SIGUSR2",
//	18: "SIGTSTP",
//
//}
//
///** 兼容windows start */
//func Kill(...interface{}) {
//	return;
//}
//const (
//	SIGUSR1 = Signal(0x10)
//	SIGUSR2 = Signal(0x11)
//	SIGTSTP = Signal(0x12)
//)

func main() {
	//创建监听退出chan
	c := make(chan os.Signal)
	//监听指定信号 ctrl+c kill
	signal.Notify(c, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM,
		syscall.SIGQUIT, syscall.SIGUSR1, syscall.SIGUSR2)
	go func() {
		for s := range c {
			switch s {
			case syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT:
				fmt.Println("Program Exit...", s)
				GracefullExit()
			case syscall.SIGUSR1:
				fmt.Println("usr1 signal", s)
			case syscall.SIGUSR2:
				fmt.Println("usr2 signal", s)
			default:
				fmt.Println("other signal", s)
			}
		}
	}()

	fmt.Println("Program Start...")
	sum := 0
	for {
		sum++
		fmt.Println("sum:", sum)
		time.Sleep(time.Second)
	}
}

func GracefullExit() {
	fmt.Println("Start Exit...")
	fmt.Println("Execute Clean...")
	fmt.Println("End Exit...")
	os.Exit(0)
}
