package main

import "fmt"

type Person1 struct {
	Name string
	Age  int
}

func main() {
	//panic: runtime error: index out of range [0] with length 0
	//var s []string
	//fmt.Println(s)
	//fmt.Println(s[0])

	//panic: runtime error: invalid memory address or nil pointer dereference
	//var p *Person
	//fmt.Println(p)
	//fmt.Println(p.Name)

	//panic: interface conversion: interface {} is string, not int
	//add(20, 18)
	//add(1, "hello")

	//panic: close of nil channel
	//var ch chan int
	//close(ch)

	//panic: send on closed channel
	//var ch chan int
	//ch = make(chan int,0)
	//close(ch)
	//ch <- 108

}

func add(a, b interface{}) {
	i := a.(int)
	j := b.(int)
	fmt.Println(i + j)
}
