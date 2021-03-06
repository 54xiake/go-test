package main

import (
	"fmt"
	"time"
)

func addNumberToChan(chanName chan int) {
	for {
		chanName <- 1
		time.Sleep(1 * time.Second)
	}
}
func main() {
	/*	test1()
		var chan1 = make(chan int, 10)
		var chan2 = make(chan int, 10)
		go addNumberToChan(chan1)
		go addNumberToChan(chan2)
		for {
			select {
			case e := <-chan1:
				fmt.Printf("Get element from chan1: %d\n", e)
				break
			case e := <-chan2:
				fmt.Printf("Get element from chan2: %d\n", e)
			default:
				fmt.Printf("No element in chan1 and chan2.\n")
				//time.Sleep(1 * time.Second)
			}
		}*/

	int_chan := make(chan int, 1)
	string_chan := make(chan string, 1)
	int_chan <- 1
	string_chan <- "hello"
	select {
	case value := <-int_chan:
		fmt.Println(value)
	case value := <-string_chan:
		panic(value)
	}

}

func test1() {
	fmt.Println("===")
	a := []int{1, 3, 4}
	fmt.Println(a)

	var b string
	fmt.Println(b)
}
