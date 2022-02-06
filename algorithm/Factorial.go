package main

import "fmt"

func Factorial(n uint64)(result uint64) {
	if n > 0 {
		result = n * Factorial(n-1)
		return result
	}
	return 1
}

// Go实现算法：N的阶乘
func main() {
	var i int = 0
	fmt.Printf("%d 的阶乘是 %d\n", i, Factorial(uint64(i)))
}
