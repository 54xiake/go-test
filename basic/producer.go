package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func Produce(header string, pipe chan string) {
	for {
		pipe <- fmt.Sprintf("%s: %v", header, rand.Int31())
		time.Sleep(time.Second)
	}

}

func Consume(pipe chan string) {
	for {
		message := <-pipe
		fmt.Println(message)
	}

}

func main() {
	channel := make(chan string)
	go Produce("dog", channel)
	go Produce("cat", channel)
	go Consume(channel)

	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)
	//设置要接收的信号
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		sig := <-sigs
		fmt.Println(sig)
		done <- true
	}()

	fmt.Println("等待信号")
	<-done
	fmt.Println("进程被终止")

}
