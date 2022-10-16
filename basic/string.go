package main

import (
	"bytes"
	"fmt"
	"reflect"
	"strings"
)

var abc = 123

func main() {
	//for循环不支持以逗号为间隔的多个赋值语句，必须使用平行赋值的方式来初始化多个变量
	//for i:=0,j:=10; i<j; i++ {
	for i, j := 0, 10; i < j; i++ {
		a, b := "abc", "bcd"
		fmt.Println(a)
		continue
		fmt.Println(b)
		break

	}

	//break 语句在多层嵌套语句块中时，可以通过标签指明终止哪一层语句块；
lable1: // 设置跳出for 循环的标签
	for i := 0; i < 4; i++ {
		for j := 0; j < 10; j++ {
			if j == 2 {
				break lable1
			}
			fmt.Println("j=", j)
		}
	}

	t1 := []string{"abc", "bcd"}
	fmt.Println(strings.Join(t1, ","))

	//buffer := new(bytes.Buffer)
	//t2, _ := buffer.WriteString("abc")
	//fmt.Println(t2)

	var buffer bytes.Buffer

	buffer.WriteString("test")
	buffer.WriteString("\n")
	buffer.WriteString("welcome!")
	buffer.WriteString("ITers happy amusement park!")

	fmt.Println("buffer = \n", buffer)
	fmt.Println("buffer type = \n", reflect.TypeOf(buffer))

	fmt.Println("the string of buffer = \n", buffer.String())
	fmt.Println("the type of the string of buffer = \n", reflect.TypeOf(buffer.String()))

}
