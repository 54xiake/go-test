package main

/*
#include <stdlib.h>
#include <stdio.h>

void printString(const char* s) {
	printf("%s", s);
}
*/
import "C"
import (
	"fmt"
	"unsafe"
)

func printString(s string) {
	cs := C.CString(s)
	defer C.free(unsafe.Pointer(cs))

	C.printString(cs)
}

func main() {
	s := "hello"
	printString(s)

	u := uint32(32)
	i := int32(1)
	fmt.Println(&u, &i) // 打印出地址
	p := &i             // p 的类型是 *int32
	p = (*int32)(unsafe.Pointer(&u))
	fmt.Println(p)
	fmt.Println(*p)

	data := []byte("abcd")
	for i := 0; i < len(data); i++ {
		fmt.Println(uintptr(i))
		fmt.Println(unsafe.Sizeof(data[0]))
		fmt.Println(uintptr(i) * unsafe.Sizeof(data[0]))
		ptr := unsafe.Pointer(uintptr(unsafe.Pointer(&data[0])) + uintptr(i)*unsafe.Sizeof(data[0]))
		fmt.Printf("%c,", *(*byte)(unsafe.Pointer(ptr)))
	}
	fmt.Printf("\n")
}
