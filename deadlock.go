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

	ch := make(chan int)
	go func() {
		for {
			ch <- 1
		}
	}()

	//多次读取
	//fmt.Println(<-ch, <-ch)

	for v := range ch {
		fmt.Println(v)
	}

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

}
