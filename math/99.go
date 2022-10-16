package main

import "fmt"

/*
99乘法表
1*1=2
2*1=2 2*2=3
.。。。
*/
func main() {
	//写两层循环，第一层为前面的数 前*后   第二层为后面的数
	for i := 0; i < 10; i++ {
		//i最大是9 j最大也是9 所以j<=i
		for j := 1; j <= i; j++ {
			//为保证输出格式统一，所以用%2d占位符
			fmt.Printf("%2d *%2d =%2d", j, i, i*j)
		}
		//每次运行结果后必须要有回车进行换行处理
		fmt.Println("")
	}
}
