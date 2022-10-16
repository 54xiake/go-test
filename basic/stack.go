package main

import "fmt"

type Student1 struct {
	Name string
	Age  int
}

func StudentRegister(name string, age int) *Student1 {
	s := new(Student1) //局部变量s逃逸到堆
	//s := &Student1{}
	s.Name = name
	s.Age = age
	return s
}

func Slice() {
	s := make([]int, 10000, 10000)
	for index, _ := range s {
		s[index] = index
	}
}

//go build -gcflags=-m stack.go
func main() {
	StudentRegister("Jim", 18)

	Slice()

	s := "Escape"
	fmt.Println(s)
}
