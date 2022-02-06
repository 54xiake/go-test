package main

import (
	"fmt"
	"sort"
)

type Person struct {
	Age  int
	Name string
}


func main() {
	values := []int32{42, 9, 101, 95, 27, 25}

	sort.Slice(values, func(i, j int) bool {
		return values[i] < values[j]
	})
	fmt.Println(values)
	var personArr []*Person

	p1 := &Person{Age: 1, Name: "小明"}
	p2 := &Person{Age: 5, Name: "小红"}
	p3 := &Person{Age: 4, Name: "小李"}
	p4 := &Person{Age: 2, Name: "小杰"}

	personArr = append(personArr, p1)
	personArr = append(personArr, p2)
	personArr = append(personArr, p3)
	personArr = append(personArr, p4)

	//按照最大值排序
	sort.Slice(personArr,func (i int,j int) bool {
		return personArr[i].Age>personArr[j].Age
	})
	for _,v := range personArr {
		fmt.Println(v)
	}

}
