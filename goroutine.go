package main

import (
	"errors"
	"fmt"
	"github.com/54xiake/gomodone"
	"time"
)

//func main() {
//	st := time.Now()
//	ch := make(chan bool)
//	go func ()  {
//		time.Sleep(time.Second * 2)
//		<-ch
//	}()
//	ch <- true  // 无缓冲，发送方阻塞直到接收方接收到数据。
//	fmt.Printf("cost %.1f s\n", time.Now().Sub(st).Seconds())
//	time.Sleep(time.Second * 5)
//}

//func main() {
//	st := time.Now()
//	ch := make(chan bool, 2)
//	go func ()  {
//		time.Sleep(time.Second * 2)
//		<-ch
//	}()
//	ch <- true
//	ch <- true // 缓冲区为 2，发送方不阻塞，继续往下执行
//	fmt.Printf("cost %.1f s\n", time.Now().Sub(st).Seconds()) // cost 0.0 s
//	ch <- true // 缓冲区使用完，发送方阻塞，2s 后接收方接收到数据，释放一个插槽，继续往下执行
//	fmt.Printf("cost %.1f s\n", time.Now().Sub(st).Seconds()) // cost 2.0
//	time.Sleep(time.Second * 5)
//}

func query() error {
	err := make(chan error)
	go func() {
		time.Sleep(2 * time.Second)
		err <- errors.New("test error")
	}()
	return <-err
}

func main() {
	//for i := 0; i < 10; i++ {
	//	err := query()
	//	if err != nil {
	//		fmt.Println(err.Error())
	//	}
	//	fmt.Printf("goroutines: %d\n", runtime.NumGoroutine())
	//}

	a := []int{2: 1}
	fmt.Println(a)

	var arr = [10]int{1, 2, 3, 4, 7, 5, 6}
	sl := arr[2:5] //创建有3个元素的slice
	fmt.Println(sl)

	s := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	s = s[1:9:10]
	fmt.Println(s)
	fmt.Println(len(s))
	fmt.Println(cap(s))

	avengersArray := [...]string{"Captain America", "Hulk"}
	fmt.Printf("The Avengers are: %#v\n", avengersArray)

	var slice1 = []int{1, 2, 3, 4, 5, 6}
	var slice2 = []int{8, 9, 10, 11, 12, 13, 14, 15}
	copy(slice2, slice1)
	fmt.Printf("len:%d cap:%d %#v\n", len(slice2), cap(slice2), slice2)

	str := gomodone.SayHi("YuGang")
	fmt.Println(str)

	for i := 0; i < 4; i++ {
		defer fmt.Printf("%d", i)
	}
}
