package main

//extern int go_qsort_compare(void* a, void* b);
import "C"

import (
	"fmt"
	"github.com/54xiake/go-test/qsort"
	"unsafe"
)

//export go_qsort_compare
func go_qsort_compare(a, b unsafe.Pointer) C.int {
	pa, pb := (*C.int)(a), (*C.int)(b)
	return C.int(*pa - *pb)
}

func main() {
	values := []int32{42, 9, 101, 95, 27, 25}

	qsort.Slice(values, func(i, j int) bool {
		return values[i] < values[j]
	})
	fmt.Println(values)
}
