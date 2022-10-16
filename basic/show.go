package main

import "fmt"

//type People struct{}
//
//func (p *People) ShowA() {
//	fmt.Println("showA")
//	p.ShowB()
//}
//func (p *People) ShowB() {
//	fmt.Println("showB")
//}
//
//type Teacher struct {
//	People
//}
//
//func (t *Teacher) ShowB() {
//	fmt.Println("teacher showB")
//}
//
//func main() {
//	t := Teacher{}
//	t.ShowA()
//
//	runtime.GOMAXPROCS(1)
//	intChan := make(chan int, 1)
//	stringChan := make(chan string, 1)
//	intChan <- 1
//	stringChan <- "hello"
//	select {
//	case value := <-intChan:
//		fmt.Println(value)
//	case value := <-stringChan:
//		panic(value)
//	}
//}

//func calc(index string, a, b int) int {
//	ret := a + b
//	fmt.Println(index, a, b, ret)
//	return ret
//}
//
//func main() {
//	a := 1
//	b := 2
//	defer calc("1", a, calc("10", a, b))
//	a = 0
//	defer calc("2", a, calc("20", a, b))
//	b = 1
//
//	s := make([]int, 5)
//	s = append(s, 1, 2, 3)
//	fmt.Println(s)
//
//	var ua UserAges
//	ua.Add("aa", 11)
//
//	fmt.Println(ua.Get("bb"))
//
//
//	var set ThreadSafeSet
//	set.s = []string{"a","b"}
//	intChan := set.Iter()
//	value := <- intChan
//	fmt.Println("value : ", value)
//}
//
//
//type UserAges struct {
//	ages map[string]int
//	sync.Mutex
//}
//
//func (ua *UserAges) Add(name string, age int) {
//	ua.Lock()
//	defer ua.Unlock()
//	ua.ages = make(map[string]int)
//	ua.ages[name] = age
//}
//
//func (ua *UserAges) Get(name string) int {
//	if age, ok := ua.ages[name]; ok {
//		return age
//	}
//	return -1
//}
//
//type ThreadSafeSet struct {
//	s []string
//	sync.RWMutex
//}
//func (set *ThreadSafeSet) Iter() <-chan interface{} {
//	ch := make(chan interface{})
//	go func() {
//		set.RLock()
//
//		for _, elem := range set.s {
//			ch <- elem
//		}
//
//		close(ch)
//		set.RUnlock()
//
//	}()
//	return ch
//}

//import (
//	"fmt"
//)
//
//type People interface {
//	Speak(string) string
//}
//
//type Stduent struct{
//	People
//}
//
//func (stu *Stduent) Speak(think string) (talk string) {
//	if think == "bitch" {
//		talk = "You are a good boy"
//	} else {
//		talk = "hi"
//	}
//	return
//}
//
//func main() {
//	var peo Stduent
//	think := "bitch"
//	fmt.Println(peo.Speak(think))
//}

type People interface {
	Show()
}

type Student struct{}

func (stu *Student) Show() {
	fmt.Println("===")
}

func live() People {
	var stu *Student
	return stu
}

func main() {
	aa := live()
	fmt.Printf("%#v", aa)
	fmt.Printf("%#v", nil)
	if aa == nil {
		fmt.Println("AAAAAAA")
	} else {
		fmt.Println("BBBBBBB")
	}
}
