package main

import (
	"encoding/json"
	"fmt"
	jsoniter "github.com/json-iterator/go"
)

// 成员变量名首字母必须大写
type IT struct {
	Company  string `json:"company"`
	Subjects []string
	IsOk     bool
	Price    float64
}

func main() {
	//定义一个结构体变量，同时初始化
	s := &IT{"itcast", []string{"Go", "C++", "Python", "Test"}, true, 666.666}

	//编码，根据内容生成json文本
	buf, err := json.Marshal(s)
	if err != nil {
		fmt.Println("err = ", err)
		return
	}
	fmt.Println("buf = ", string(buf))

	a := &IT{}
	json.Unmarshal(buf, a)
	fmt.Printf("%#v", a)

	fmt.Println()
	// go get github.com/json-iterator/go
	buf2, _ := jsoniter.Marshal(s)
	fmt.Println("buf2 = ", string(buf2))

	jsoniter.Unmarshal(buf2, a)
	fmt.Printf("%#v", a)

}
