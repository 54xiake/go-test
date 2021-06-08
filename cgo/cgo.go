package main

/*
#include <stdio.h>
#include <errno.h>


void printint(int v) {
    printf("printint: %d\n", v);
}

static int add(int a, int b) {
	return a+b;
}

int sum(int a, int b) { return a+b; }

static int div(int a, int b) {
	if(b == 0) {
		errno = EINVAL;
		return 0;
	}
	return a/b;
}

static void noreturn() {}


#cgo windows CFLAGS: -DX86=1
#cgo !windows LDFLAGS: -lm

//#cgo CFLAGS: -DPNG_DEBUG=1
//#define CGO_OS_DARWIN 1

#cgo windows CFLAGS: -DCGO_OS_WINDOWS=1
#cgo darwin CFLAGS: -DCGO_OS_DARWIN=1
#cgo linux CFLAGS: -DCGO_OS_LINUX=1

#if defined(CGO_OS_WINDOWS)
	const char* os = "windows";
#elif defined(CGO_OS_DARWIN)
	const char* os = "darwin";
#elif defined(CGO_OS_LINUX)
	const char* os = "linux";
#else
	const char* os = "unknown";
#endif



*/
import "C"
import (
	"fmt"
	"strconv"
	"unsafe"
)

//保证环境变量CGO_ENABLED被设置为1，这表示CGO是被启用的状态
func main() {
	v := 42
	C.printint(C.int(v))

	fmt.Println(C.add(1, 1))
	fmt.Println(C.sum(2, 1))

	fmt.Println(C.GoString(C.os))

	v0, err0 := C.div(2, 1)
	fmt.Println(v0, err0)

	v1, err1 := C.div(1, 0)
	fmt.Println(v1, err1)

	_, err := C.noreturn()
	fmt.Println(err)

	vv, _ := C.noreturn()
	fmt.Printf("%#v", vv)

	var x int = 100
	fmt.Println(unsafe.Sizeof(x)) // 8

	var y int64 = 1
	fmt.Println(unsafe.Sizeof(y)) // 8
	var y1 int32 = 1
	fmt.Println(unsafe.Sizeof(y1)) // 4
	var z uint64 = 1
	fmt.Println(unsafe.Sizeof(z)) // 8
	var z1 uint32 = 1
	fmt.Println(unsafe.Sizeof(z1)) // 4

	fmt.Println(convert2Bin(x, 8))

	fmt.Println(
		convert2Bin(5, 8),  //101
		convert2Bin(13, 8), //1101
		convert2Bin(11111, 8),
		convert2Bin(0, 8),
		convert2Bin(1, 8),
		convert2Bin(-5, 8),
		convert2Bin(-11111, 8),
	)

}

//bin表示转化后的位数
func convert2Bin(n int, bin int) string {
	var b string
	switch {
	case n == 0:
		for i := 0; i < bin; i++ {
			b += "0"
		}
	case n > 0:
		//strcov.Itoa 将 1 转为 "1" , string(1)直接转为assic码
		for ; n > 0; n /= 2 {
			b = strconv.Itoa(n%2) + b
		}
		//加0
		j := bin - len(b)
		for i := 0; i < j; i++ {
			b = "0" + b
		}
	case n < 0:
		n = n * -1
		// fmt.Println("变为整数：",n)
		s := convert2Bin(n, bin)
		// fmt.Println("bin:",s)
		//取反
		for i := 0; i < len(s); i++ {
			if s[i:i+1] == "1" {
				b += "0"
			} else {
				b += "1"
			}
		}
		// fmt.Println("~bin :",b)
		//转化为整形，之后加1 这里必须要64，否则在转化过程中可能会超出范围
		n, err := strconv.ParseInt(b, 2, 64)
		if err != nil {
			fmt.Println(err)
		}
		//转为bin
		//+1
		b = convert2Bin(int(n+1), bin)
	}
	return b
}
