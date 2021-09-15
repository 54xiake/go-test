package main

import (
	"encoding/json"
	"errors"
	"fmt"
)

type Order struct {
	OrderId    string  `json:"order_id"`
	OrderPrice float64 `json:"order_price"`
	Goods      []Good  `json:"good"`
}

type Good struct {
	GoodsName  string  `json:"goods_name"`
	GoodsPrice float64 `json:"goods_price"`
	GoodsColor []Color `json:"Goods_color"`
}

type Color struct {
	GoodColor string `json:"good_color"`
}

func main() {
	var OrderInfo Order
	OrderInfo.OrderId = "20190707212318"
	OrderInfo.OrderPrice = 26.87

	var csli []Color
	csli = append(csli, Color{GoodColor: "红色"})
	csli = append(csli, Color{GoodColor: "蓝色"})

	OrderInfo.Goods = append(OrderInfo.Goods, Good{GoodsName: "手机", GoodsPrice: 23.9, GoodsColor: csli})

	OrderInfo.Goods = append(OrderInfo.Goods, Good{GoodsName: "电脑", GoodsPrice: 123.9, GoodsColor: csli})

	data, _ := json.Marshal(OrderInfo)

	fmt.Println(string(data))

	/*
		{
			"order_id": "20190707212318",
			"order_price": 26.87,
			"good": [{
				"goods_name": "手机",
				"goods_price": 23.9,
				"Goods_color": [{
					"good_color": "红色"
				}, {
					"good_color": "蓝色"
				}]
			}, {
				"goods_name": "电脑",
				"goods_price": 123.9,
				"Goods_color": [{
					"good_color": "红色"
				}, {
					"good_color": "蓝色"
				}]
			}]
		}
	*/

	var temp Order
	err := json.Unmarshal(data, &temp)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(temp)
	fmt.Println(temp.Goods[0].GoodsName)

	//{20190707212318 26.87 [{手机 23.9 [{红色} {蓝色}]} {电脑 123.9 [{红色} {蓝色}]}]}

	//结构体比较

	TestStructJunior()

	testReturn()
	testReturn2()
}

func testReturn() (i int) {
	i = 1
	defer func() {
		i = i + 1
		fmt.Println(i)
	}()
	return 0
}

func testReturn2() int {
	i := 1
	defer func(x int) {
		x = x + 1
		fmt.Println(x)
	}(i)
	return 0
}

type Dog struct {
	name string
	//friends map[int]string  //若加上这个属性，则会报错 invalid operation: dog1 == dog2 (struct containing map[int]string cannot be compared)
}

func (dog *Dog) SetName(name string) {
	dog.name = name
}

func (dog Dog) Name() string {
	return dog.name
}

func TestStructJunior() {
	dog1 := Dog{name: "one"}
	fmt.Println(dog1)
	var dog2 Dog = dog1
	//dog1.SetName("two")  //若加上这个语句， 下面的值就不会相等
	fmt.Println(dog2)

	if dog1 == dog2 {
		fmt.Println("dog1==dog2")
	} else {
		fmt.Println("dog1!=dog2")
	}

	a := errors.New("abc")
	b := errors.New("abc")

	if a == b {
		fmt.Println("===")
	} else {
		fmt.Println("!==")
	}
}
