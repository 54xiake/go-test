package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	//var wg sync.WaitGroup
	//wg.Add(1)
	//go func() {
	//	defer func() {
	//		if err := recover(); err != nil {
	//			println(err.(string))
	//		}
	//	}()
	//	//p()
	//	//wg.Done()
	//	//在延迟函数中使用defer对WaitGroup对象的调用可以防止死锁
	//	defer wg.Done()
	//	p()
	//
	//}()
	//wg.Wait()

	runtime.GOMAXPROCS(runtime.NumCPU())
	wg := sync.WaitGroup{}
	wg.Add(20)
	for i := 0; i < 10; i++ {
		go func() {
			fmt.Println("i: ", i)
			wg.Done()
		}()
	}
	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println("i: ", i)
			wg.Done()
		}(i)
	}
	wg.Wait()
}
func p() {
	panic("foo")
}
