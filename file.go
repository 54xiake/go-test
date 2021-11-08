package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	filePath := "1G.txt"
	filePath2 := "1G.txt"

	//只读的形式打开文件。
	//ps: os.OpenFile可以以读写方式打开
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err)
		return
	}
	file2, err := os.Open(filePath2)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	//分片读取

	//1024bit*8=1024byte=1kb
	//1kb*1024=1m
	//读取100M的内容
	buffer := make([]byte, 1024*1024*100)
	i := 0
	for {
		n, err := file.Read(buffer)
		if err != nil && err != io.EOF {
			fmt.Println("错误", err)
			return
		}
		if n == 0 {
			fmt.Println("结束")
			break
		}
		i += n / 1024
		fmt.Println("大小", i/1024, "KB")
	}
	fmt.Println("文件大小：", i/1024, "KB")

	//流处理：
	scanner := bufio.NewScanner(file2)
	for scanner.Scan() {
		fmt.Println(scanner.Bytes())
	}
}
