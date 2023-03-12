package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	addOneHour := t.Add(time.Hour)
	addTwoHour := t.Add(2 * time.Hour)
	fmt.Println("增加1小时：", addOneHour)
	fmt.Println("增加2小时：", addTwoHour)

	subTwoHour := t.Add(-2 * time.Hour)
	fmt.Println("减去2小时：", subTwoHour)

	// int 转 duration
	date, _ := time.ParseInLocation("2006-01-02 15:04:05", "2022-01-18 08:30:00", time.Local)
	three := 72
	subThreeDay := date.Add(-time.Duration(three) * time.Hour)
	fmt.Println(subThreeDay)

	addDate := t.AddDate(1, 0, 0)
	fmt.Println("增加1年：", addDate) // 2021-10-24 22:10:53.328973 +0800 CST

	subDate := t.AddDate(-1, 0, 0)
	fmt.Println("减去1年：", subDate) // 2019-10-24 22:10:53.328973 +0800 CST

	before := t.Before(t.Add(time.Hour))
	fmt.Println("before：", before)

	after := t.After(t.Add(time.Hour))
	fmt.Println("after：", after)

	beforeDay := t.AddDate(0, 0, -1)   // 三个参数分别是年月日，此处获取的是前一天的日期
	beforeMonth := t.AddDate(0, -1, 0) // 前一个月的日期
	beforeYear := t.AddDate(-1, 0, 0)  // 去年的当天日期
	fmt.Println(beforeDay, beforeMonth, beforeYear)

	m := t.Format("2006-01-02 15:04:05") //获取当前格式的日期
	fmt.Println(m)
}
