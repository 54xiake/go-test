package main

/*
extern int* getGoPtr();

static void Main() {
	int* p = getGoPtr();
	*p = 42;
}
*/
import "C"

func main() {
	C.Main()
}

//导出C函数不能返回Go内存
//export getGoPtr
func getGoPtr() *C.int {
	return new(C.int)
}
