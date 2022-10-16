package main

import (
	"fmt"
	"math"
)

func main() {
	sin := math.Sin(45)
	fmt.Println(sin)

	// 面积
	r := 5.0
	s := math.Pi * r * r
	fmt.Println(s)
}
