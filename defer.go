package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	/*
		先defer的后执行
		recover后输出panic中的信息
	*/
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("no")
		}
	}()
	defer func() {
		//panic("1111111111111")
		fmt.Println("defer ======")
	}()

	go func() {
		//defer func() {
		//	if err := recover(); err != nil {
		//		fmt.Println("go ", err)
		//	}
		//}()
		panic("go func panic")
	}()

	panic("22222222222")

	deferFuncParameter()

	fmt.Printf("%#v", parseStudent())
	deferCall()
	i := 0
	for {
		//go func(x int) {
		//	fmt.Println("=============" + strconv.Itoa(x))
		//}(i)
		time.Sleep(3 * time.Second)
		fmt.Println("=============" + strconv.Itoa(i))
		i++
	}

}

func deferCall() {
	defer func() { fmt.Println("打印前") }()
	defer func() { fmt.Println("打印中") }()
	defer func() { fmt.Println("打印后") }()

	//panic("触发异常")
}

type student struct {
	Name string
	Age  int
}

func parseStudent() map[string]student {
	m := make(map[string]student)
	stus := []student{
		{Name: "zhou", Age: 24},
		{Name: "li", Age: 23},
		{Name: "wang", Age: 22},
	}
	for _, stu := range stus {
		m[stu.Name] = stu
	}
	return m
}

func deferFuncParameter() {
	var aInt = 1
	defer fmt.Println(aInt)
	aInt = 2
	return
}
