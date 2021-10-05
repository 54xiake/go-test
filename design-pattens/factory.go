package main

import (
	"fmt"
)

//玩具工厂
type ToyFactory interface {
	Create() Toy //制造功能，返回一个玩具
}

//所有玩具都有一个功能
type Toy interface {
	Speak() //都会说话，但是每种玩具说的话不一样
}

//车间1 可以制作玩具猫，就叫猫工厂吧
type CatFactory struct {
}

//玩具猫模型
type Cat struct {
}

//玩具猫会说话
func (this *Cat) Speak() {
	fmt.Println("hello,I am a Cat...")
}

//车间1 制作玩具猫的方法
func (this *CatFactory) Create() Toy {
	return &Cat{}
}

//车间2 可以制作玩具狗，就叫狗工厂吧
type DogFactory struct {
}

//狗模具
type Dog struct {
}

//狗也会说话
func (this *Dog) Speak() {
	fmt.Println("hello, i am a Dog...")
}

//车间2 制作玩具狗的方法
func (this *DogFactory) Create() Toy {
	return &Dog{}
}

func main() {
	cat := &CatFactory{} //进入猫工厂
	cat.Create().Speak() //制作猫猫，说话试音

	dog := &DogFactory{} //进入狗工厂
	dog.Create().Speak() //制作小狗，说话试音
}
