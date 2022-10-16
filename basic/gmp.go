package main

import (
	"fmt"
	"runtime"
)

func main() {
	go func() {
		for {
			fmt.Println("456")
		}
	}()

	for {
		fmt.Println(getNumGoroutine())
		go func() {
			fmt.Println("================")
		}()
	}

}

func getNumGoroutine() int {
	return runtime.NumGoroutine()
}
