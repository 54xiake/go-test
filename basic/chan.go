package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
)

func addNumberToChan(chanName chan int) {
	for {
		chanName <- 1
		time.Sleep(1 * time.Second)
	}
}
func main() {
	/*	test1()
		var chan1 = make(chan int, 10)
		var chan2 = make(chan int, 10)
		go addNumberToChan(chan1)
		go addNumberToChan(chan2)
		for {
			select {
			case e := <-chan1:
				fmt.Printf("Get element from chan1: %d\n", e)
				break
			case e := <-chan2:
				fmt.Printf("Get element from chan2: %d\n", e)
			default:
				fmt.Printf("No element in chan1 and chan2.\n")
				//time.Sleep(1 * time.Second)
			}
		}*/

	int_chan := make(chan int, 10)
	//string_chan := make(chan string, 1)

	go func() {
		i := 1
		for {
			int_chan <- i
			fmt.Println("test")
			fmt.Println(len(int_chan))
			//string_chan <- "hello"
			i++
		}
	}()

	_filePath := "./test.txt"
	_file, _err := os.OpenFile(_filePath, os.O_WRONLY|os.O_CREATE, 0666)
	if _err != nil {
		fmt.Printf("打开文件错误=%v\n", _err)
	}
	//提前关闭文件
	defer _file.Close()
	//写入文件
	_writer := bufio.NewWriter(_file)

	for {
		fmt.Println(len(int_chan))
		c := <-int_chan
		//time.Sleep(30000)
		_writer.WriteString(strconv.Itoa(c) + "\n")
		_writer.Flush()
	}

	select {
	case value := <-int_chan:
		fmt.Println(value)
		//case value := <-string_chan:
		//	panic(value)
	}

}

func test1() {
	fmt.Println("===")
	a := []int{1, 3, 4}
	fmt.Println(a)

	var b string
	fmt.Println(b)
}
