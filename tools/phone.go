package main

import (
	"fmt"

	"github.com/xluohome/phonedata"
)

func main2() {
	pr, err := phonedata.Find("15512261234")
	if err != nil {
		panic(err)
	}
	fmt.Print(pr)
}
