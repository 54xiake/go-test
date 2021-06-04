package main

import (
	"encoding/json"
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

}
