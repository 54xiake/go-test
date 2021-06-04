package main

import "fmt"

func main() {
	/*
		先defer的后执行
		recover后输出panic中的信息
	*/
	//defer func() {
	//	if err := recover(); err != nil {
	//		fmt.Print(err)
	//	} else {
	//		fmt.Print("no")
	//	}
	//}()
	//defer func() {
	//	panic("1111111111111")
	//}()
	//panic("22222222222")

	fmt.Printf("%#v", parseStudent())
	deferCall()
}

func deferCall() {
	defer func() { fmt.Println("打印前") }()
	defer func() { fmt.Println("打印中") }()
	defer func() { fmt.Println("打印后") }()

	panic("触发异常")
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
