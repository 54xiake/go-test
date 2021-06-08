package main

import (
	"fmt"
	"sync/atomic"
)

func main() {
	var value int32
	//atomic.AddUint32(&value, 1)

	atomic.StoreInt32(&value, 1)

	for i := 1; i <= 10; i++ {
		atomic.AddInt32(&value, int32(i))
	}

	v := atomic.LoadInt32(&value)
	fmt.Println(v)

}
