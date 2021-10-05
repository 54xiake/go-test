package main

import "fmt"

// 首先定义接口，所有的策略都是基于一套标准，这样策略(类)才有可替换性。声明一个计算策略接口
type ICompute interface {
	Compute(a, b int) int
}

// 接着两个接口实现类。复习golang语言实现接口是非侵入式设计。
type Add struct {
}

func (p *Add) Compute(a, b int) int {
	return a + b
}

type Sub struct {
}

func (p *Sub) Compute(a, b int) int {
	return a - b
}

//声明一个策略类。复习golang中规定首字母大写是public，小写是private。
//如果A,B改为小写a,b，在客户端调用时会报unknown field 'a' in struct literal of type strategy.Context

var compute ICompute

type Context struct {
	A, B int
}

func (p *Context) SetContext(o ICompute) {
	compute = o
}

func (p *Context) Result() int {
	return compute.Compute(p.A, p.B)
}

//客户端调用
func main() {
	c := Context{A: 15, B: 7}
	// 用户自己决定使用什么策略
	c.SetContext(new(Add))
	fmt.Println(c.Result())
}
