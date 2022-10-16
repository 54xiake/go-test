package main

import "fmt"

func main() {
	//ch := make(chan int)
	//
	////fmt.Println(<-ch)
	//
	//go func() {
	//	ch <- 1
	//}()
	//
	//fmt.Println(<-ch)

	//fatal error: all goroutines are asleep - deadlock!
	//var ch chan int
	//ch = make(chan int,1)
	//ch <- 108
	//ch <- 109

	//ch := make(chan int)
	//go func() {
	//	for {
	//		ch <- 1
	//	}
	//}()

	//多次读取
	//fmt.Println(<-ch, <-ch)

	//for v := range ch {
	//	fmt.Println(v)
	//}

	//ch1 := make(chan int)
	//
	//ch2 := make(chan int)
	//
	//go func() {
	//	ch1 <- 1
	//	ch2 <- 2
	//}()
	//
	//fmt.Println(<-ch2) //这里读取的顺序颠倒
	//fmt.Println(<-ch1)

	//ch1 := make(chan int)
	//ch2 := make(chan int)
	//
	//go func() {
	//	select {
	//	case <-ch1:
	//		ch2 <- 20
	//	}
	//}()
	//
	//select {
	//case <-ch2:
	//	ch1 <- 10
	//}

	ch := make(chan int) //声明
	//向ch中写入数据，ch此时等待读出数据，造成堵塞，使得下方读出数据部分的代码不执行，造成死锁
	ch <- 89
	num := <-ch
	fmt.Println(num)

}
