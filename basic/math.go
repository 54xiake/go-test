package main

import (
	"fmt"
	"math"
)

func main() {
	//var i int
	//for {
	//	i = i + 2
	//	fmt.Println(i)
	//	time.Sleep(1*time.Second)
	//}
	var n int
	for i := 1; i <= 200; i++ {
		n = n + i
	}
	fmt.Println(n)
	return
	sin := math.Sin(45)
	fmt.Println(sin)

	// 面积
	r := 5.0
	s := math.Pi * r * r
	fmt.Println(s)

	fmt.Println(math.Abs(-1))

	// 平方
	fmt.Println(math.Sqrt(3))

	// 立方
	fmt.Println(math.Pow(2, 3))
	fmt.Println(math.Pow(3, 2))

	/*
	   取绝对值,函数签名如下:
	       func Abs(x float64) float64
	*/
	fmt.Printf("[-3.14]的绝对值为:[%.2f]\n", math.Abs(-3.14))

	/*
	   取x的y次方，函数签名如下:
	       func Pow(x, y float64) float64
	*/
	fmt.Printf("[2]的16次方为:[%.f]\n", math.Pow(2, 16))

	/*
	   取余数，函数签名如下:
	       func Pow10(n int) float64
	*/
	fmt.Printf("10的[3]次方为:[%.f]\n", math.Pow10(3))

	/*
	   取x的开平方，函数签名如下:
	       func Sqrt(x float64) float64
	*/
	fmt.Printf("[64]的开平方为:[%.f]\n", math.Sqrt(64))

	/*
	   取x的开立方，函数签名如下:
	       func Cbrt(x float64) float64
	*/
	fmt.Printf("[27]的开立方为:[%.f]\n", math.Cbrt(27))

	/*
	   向上取整，函数签名如下:
	       func Ceil(x float64) float64
	*/
	fmt.Printf("[3.14]向上取整为:[%.f]\n", math.Ceil(3.14))

	/*
	   向下取整，函数签名如下:
	       func Floor(x float64) float64
	*/
	fmt.Printf("[8.75]向下取整为:[%.f]\n", math.Floor(8.75))

	/*
	   取余数，函数签名如下:
	       func Floor(x float64) float64
	*/
	fmt.Printf("[10/3]的余数为:[%.f]\n", math.Mod(10, 3))

	/*
	   分别取整数和小数部分,函数签名如下:
	       func Modf(f float64) (int float64, frac float64)
	*/
	Integer, Decimal := math.Modf(3.14159265358979)
	fmt.Printf("[3.14159265358979]的整数部分为:[%.f],小数部分为:[%.14f]\n", Integer, Decimal)

}
