package main

import (
	"fmt"
	"unicode/utf8"
	"unsafe"
)

func main() {
	testReturn1()
	fmt.Println("test main")

	marginPart := 23 % 30
	fmt.Println(marginPart)
	marginPart1 := 33 % 30
	fmt.Println(marginPart1)
	fmt.Println(int(-5) / 40)
	fmt.Println(0 / 40)

	seq := []string{"a", "b", "c", "d", "e"}

	// 指定删除位置
	index := 2

	// 查看删除位置之前的元素和之后的元素
	fmt.Println(seq[:index], seq[index+1:])

	// 将删除点前后的元素连接起来
	seq = append(seq[:index], seq[index+1:]...)

	fmt.Println(seq)

	for n := 1; n < 1000; n++ {
		if (8+1*n)%40 == 0 {
			fmt.Println(n)
			fmt.Println(60 + 12*n)
			break
		}
	}

	// 字节数
	a := "    "
	b := "\t"
	c := "aa中国"
	fmt.Println(unsafe.Sizeof(a))
	fmt.Println(unsafe.Sizeof(b))
	fmt.Println(unsafe.Sizeof(c))

	fmt.Println(len(a))
	fmt.Println(len(b))
	fmt.Println(len(c))

	//在UTF-8编码中，一个英文为一个字节，一个中文为三个字节。
	//字节长度
	str := "Golang梦工厂"
	fmt.Println(len(str))
	fmt.Println(len([]rune(str)))
	// 字符个数使用utf8.RuneCountInString
	fmt.Println(utf8.RuneCountInString(str))
	//使用下标遍历获取的是ASCII字符，而使用Range遍历获取的是Unicode字符。
	for k, v := range str {
		fmt.Printf("v type: %T index,val: %v,%v \n", v, k, v)
	}
	for i := 0; i < len(str); i++ {
		fmt.Printf("v type: %T index,val:%v,%v \n", str[i], i, str[i])
	}

	return
}

func testReturn1() {
	fmt.Println("test return 1")
	return
}
