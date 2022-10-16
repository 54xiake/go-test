package main

import (
	"container/list"
	"fmt"
	"os"
)

func main() {
	var myList list.List
	fmt.Println(myList.Len())
	myList.PushBack("back")
	myList.PushFront("front")
	element := myList.PushBack("first")
	myList.InsertAfter("high", element)
	myList.InsertBefore("zero", element)
	fmt.Println(element.Value)
	myList.Remove(element)
	fmt.Println(myList)
	fmt.Println(element.Value)
	for i := myList.Front(); i != nil; i = i.Next() {
		fmt.Println(i.Value)
		if i.Value == "zero" {
			myList.Remove(i)
		}
	}
	fmt.Println(myList)

	os.Exit(500)
}
