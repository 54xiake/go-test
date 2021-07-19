package main

import "fmt"

const (
	mutexLocked = 1 << iota // mutex is locked
	mutexWoken
	mutexStarving
	mutexWaiterShift      = iota
	starvationThresholdNs = 1e6
)

func main() {
	//a := new make(map[string][string]{}, 10)

	fmt.Println(mutexLocked)
	fmt.Println(mutexWoken)
	fmt.Println(mutexStarving)
	fmt.Println(mutexWaiterShift)
	fmt.Println(int(starvationThresholdNs))

	fmt.Println(1 << 3)
	fmt.Println(1 << 16)

	a := map[string]string{"key": "value"}
	fmt.Println(a["key"])

}
