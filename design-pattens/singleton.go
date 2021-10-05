package main

import (
	"fmt"
	"sync"
)

// 定义：单例对象的类必须保证只有一个实例存在，全局有唯一接口访问。
// 分类：
// 懒汉方式：指全局的单例实例在第一次被使用时构建。
// 饿汉方式：指全局的单例实例在类装载时构建。
type singleton struct{}

var ins *singleton
var mu sync.Mutex

// GetIns 避免了每次加锁，提高代码效率
func GetIns() *singleton {
	if ins == nil {
		mu.Lock()
		defer mu.Unlock()
		if ins == nil {
			fmt.Println("创建")
			ins = &singleton{}
		}
	}
	return ins
}

var once sync.Once

// GetIns2 sync.Once实现
func GetIns2() *singleton {
	once.Do(func() {
		ins = &singleton{}
	})
	return ins
}

func main() {
	for i := 0; i < 10; i++ {
		ins := GetIns()
		fmt.Println(ins)
	}

	for i := 0; i < 10; i++ {
		ins := GetIns2()
		fmt.Println(ins)
	}
}
