package main

/*
#include <stdio.h>

void printint(int v) {
    printf("printint: %d\n", v);
}


#cgo windows CFLAGS: -DCGO_OS_WINDOWS=1
#cgo darwin CFLAGS: -DCGO_OS_DARWIN=1
#cgo linux CFLAGS: -DCGO_OS_LINUX=1

#if defined(CGO_OS_WINDOWS)
	static char* os = "windows";
#elif defined(CGO_OS_DARWIN)
	static char* os = "darwin";
#elif defined(CGO_OS_LINUX)
	static char* os = "linux";
#else
#	error(unknown os)
#endif


*/
import "C"

//保证环境变量CGO_ENABLED被设置为1，这表示CGO是被启用的状态
func main() {
	v := 42
	C.printint(C.int(v))

	print(C.GoString(C.os))
}
