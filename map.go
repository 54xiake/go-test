package main

import (
	"fmt"
	"sort"
)

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

	//map实现排序
	m := make(map[int]int, 10)
	keys := make([]int, 0, 10)
	for i := 0; i < 10; i++ {
		m[i] = i
		keys = append(keys, i)
	}
	//降序
	sort.Slice(keys, func(i, j int) bool {
		if keys[i] > keys[j] {
			return true
		}
		return false
	})
	for _, v := range keys {
		fmt.Println(m[v])
	}

}
