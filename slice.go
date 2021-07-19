package main

import "fmt"

func AddElement(slice []int, e int) []int {
	return append(slice, e)
}

func main() {
	var numbers4 = [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	myslice := numbers4[3:6:8]
	fmt.Printf("myslice为 %d, 其容量为: %d\n", myslice, cap(myslice))

	myslice2 := myslice[:cap(myslice)]
	fmt.Printf("myslice: %d", myslice2)
	fmt.Printf("myslice的第四个元素为: %d\n", myslice2[3])

	myslice3 := numbers4[3:6]
	fmt.Printf("myslice为 %d, 其容量为: %d\n", myslice3, cap(myslice3))

	var slice []int
	fmt.Println(cap(slice))
	slice = append(slice, 1, 2, 3)
	fmt.Println(cap(slice))
	newSlice := AddElement(slice, 4)
	fmt.Println(cap(newSlice))
	fmt.Println(&slice[0] == &newSlice[0])

	a := make([]string, 5)
	a = append(a, "a")
	fmt.Println(cap(a))
	fmt.Println(len(a))
	fmt.Printf("%#v\n", a)

	b := make([]int, 5)
	b = append(b, 1)
	fmt.Println(cap(b))
	fmt.Println(len(b))
	fmt.Printf("%#v\n", b)
	b = append(b, 2, 3, 4, 5, 6)
	fmt.Println(cap(b))
	fmt.Println(len(b))
	fmt.Printf("%#v\n", b)

	arr := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	var s1 []int = arr[1:4]
	var s2 []int = arr[7:]
	fmt.Println(cap(arr))
	fmt.Println(cap(s1))
	fmt.Println(cap(s2))
	fmt.Println(s1)
	fmt.Println(s2)

	orderLen := 5
	order := make([]uint16, 2*orderLen)
	pollorder := order[:orderLen:orderLen]
	lockorder := order[orderLen:][:orderLen:orderLen]
	fmt.Println("len(pollorder) = ", len(pollorder))
	fmt.Println("cap(pollorder) = ", cap(pollorder))
	fmt.Println("len(lockorder) = ", len(lockorder))
	fmt.Println("cap(lockorder) = ", cap(lockorder))

	v := []int{1, 2, 3}
	for i := range v {
		v = append(v, i)
	}

	fmt.Printf("%#v\n", v)

	var aa = []int{1, 2, 3, 4, 5}
	bb := aa                                  //此时a，b都指向了内存中的[1 2 3 4 5]的地址
	bb[1] = 10                                //相当于修改同一个内存地址，所以a的值也会改变
	cc := make([]int, 5, 5)                   //切片的初始化
	copy(cc, aa)                              //将切片acopy到c
	cc[1] = 20                                //copy是值类型，所以a不会改变
	fmt.Printf("a的值是%v，a的内存地址是%p\n", aa, &aa) //a的值是[1 10 3 4 5]，a的内存地址是0xc42000a180
	fmt.Printf("b的值是%v，b的内存地址是%p\n", bb, &bb) //b的值是[1 10 3 4 5]，b的内存地址是0xc42000a1a0
	fmt.Printf("c的值是%v，c的内存地址是%p\n", cc, &cc) //c的值是[1 20 3 4 5]，c的内存地址是0xc42000a1c0
	dd := &aa                                 //将a的内存地址赋值给d，取值用*d
	aa[1] = 11
	fmt.Printf("d的值是%v，d的内存地址是%p\n", *dd, dd) //d的值是[1 11 3 4 5]，d的内存地址是0xc42000a180
	fmt.Printf("a的值是%v，a的内存地址是%p\n", aa, &aa) //a的值是[1 11 3 4 5]，a的内存地址是0xc42000a180
}
