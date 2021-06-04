package main

import (
	"fmt"
	"net"
	"os"
	"reflect"
	"runtime"
	"strconv"
	"unsafe"
)

const (
	i = 1 << iota
	j = 3 << iota
	k
	l
)

func init() {
	fmt.Println("init==============")
}
func main() {
	fmt.Fprintf(os.Stderr, "dup2: %v\n", "aaa")
	fmt.Println(runtime.GOOS)
	fmt.Println(runtime.GOARCH)
	fmt.Println(runtime.Version())

	name, err := os.Hostname()
	if err == nil {
		fmt.Println(name)
	}

	fmt.Println(reflect.TypeOf(name))

	fmt.Println(runtime.GOMAXPROCS(0))
	fmt.Println(net.Interfaces())

	typeSize()

	fmt.Println("Google" + " " + "Runoob")

	var age int = 123
	fmt.Println(age + 456)

	var t bool = true
	fmt.Println(t)

	//var b, c int = 1, 2
	//fmt.Println(b, c)

	var d int
	fmt.Println(d)

	var e bool
	fmt.Println(e)

	//var i int
	//var f float64
	//var b bool
	//var s string
	//fmt.Printf("%v %v %v %q\n", i, f, b, s)
	//var intVal int = 1
	intVal := 1
	fmt.Println(intVal)
	fmt.Println(&intVal)

	r1 := 1
	var r2 = &r1
	fmt.Println(r1)
	fmt.Println(*r2)

	r1 = 4
	fmt.Println(r1)
	fmt.Println(*r2)

	const LENGTH int = 10
	const WIDTH int = 5
	fmt.Println(LENGTH * WIDTH)

	fmt.Println("i=", i)
	fmt.Println("j=", j)
	fmt.Println("k=", k)
	fmt.Println("l=", l)

	fmt.Println(convertToBin1(int(3)))

	var c rune = 'Z'
	fmt.Println(convertToBin1(int(c)))
	fmt.Println(int(c))
	//121103
	//898882
	//898890

	c1 := rune(100)
	fmt.Println(string(c1))

	sum := 0
	for i := 1; i < 100; i++ {
		sum += i
	}
	fmt.Println(sum)

	sum = 1
	for sum <= 10 {
		sum += sum
	}
	fmt.Println(sum)

	// 这样写也可以，更像 While 语句形式
	for sum <= 10 {
		sum += sum
	}
	fmt.Println(sum)

	//sum = 0
	//for {
	//	sum++ // 无限循环下去
	//	fmt.Println(sum)
	//}
	//fmt.Println(sum) // 无法输出

	strings := []string{"google", "runoob"}
	for i, s := range strings {
		fmt.Println(i, s)
	}

	var balance = []float32{1000.0, 2.0, 3.4, 7.0, 50.0}
	for i, s := range balance {
		fmt.Println(i, s)
	}

	var a = [5][2]int{{0, 0}, {1, 2}, {2, 4}, {3, 6}, {4, 8}}
	var i, j int

	/* 输出数组元素 */
	for i = 0; i < 5; i++ {
		for j = 0; j < 2; j++ {
			fmt.Printf("a[%d][%d] = %d\n", i, j, a[i][j])
		}
	}

}

func typeSize() {
	var p *struct{} = nil
	fmt.Println(unsafe.Sizeof(p)) // 8

	var s []int = nil
	fmt.Println(unsafe.Sizeof(s)) // 24

	var m map[int]bool = nil
	fmt.Println(unsafe.Sizeof(m)) // 8

	var c chan string = nil
	fmt.Println(unsafe.Sizeof(c)) // 8

	var f func() = nil
	fmt.Println(unsafe.Sizeof(f)) // 8

	var i interface{} = nil
	fmt.Println(unsafe.Sizeof(i)) // 16

	var test = map[string]string{"姓名": "李四", "性别": "男"}
	name, ok := test["姓名"] // 假如key存在,则name = 李四 ，ok = true,否则，ok = false
	if ok {
		fmt.Println(name)
	}
	delete(test, "姓名") //删除为姓名为key的值，不存在没关系
	fmt.Println(test)
	a := make(map[string]string)
	a["b"] = "c" //这样才不会错
	fmt.Println(a)
}

// 将十进制数字转化为二进制字符串
func convertToBin1(num int) string {
	s := ""

	if num == 0 {
		return "0"
	}

	// num /= 2 每次循环的时候 都将num除以2  再把结果赋值给 num
	for ; num > 0; num /= 2 {
		lsb := num % 2
		// strconv.Itoa() 将数字强制性转化为字符串
		s = strconv.Itoa(lsb) + s
	}
	return s
}
