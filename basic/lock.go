package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	//mutex()
	rwmutex()
}

func mutex() {
	var mutex sync.Mutex
	count := 0

	for r := 0; r < 50; r++ {
		go func() {
			mutex.Lock()
			count += 1
			mutex.Unlock()
			fmt.Println(count)
		}()
	}

	time.Sleep(time.Second)
	fmt.Println("the count is : ", count)
}

func rwmutex() {
	var mutex sync.RWMutex
	arr := []int{1, 2, 3}
	go func() {
		fmt.Println("Try to lock writing operation.")
		mutex.Lock()
		fmt.Println("Writing operation is locked.")

		arr = append(arr, 4)

		fmt.Println("Try to unlock writing operation.")
		mutex.Unlock()
		fmt.Println("Writing operation is unlocked.")
	}()

	go func() {
		fmt.Println("Try to lock reading operation.")
		mutex.RLock()
		fmt.Println("The reading operation is locked.")

		fmt.Println("The len of arr is : ", len(arr))

		fmt.Println("Try to unlock reading operation.")
		mutex.RUnlock()
		fmt.Println("The reading operation is unlocked.")
	}()

	time.Sleep(time.Second * 2)
	return
}
