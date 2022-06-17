package main

import (
	"fmt"
	"time"
)

func main() {

	now := time.Now()
	fmt.Println(now)
	year := now.Year()
	month := now.Month()
	day := now.Day()
	minute := now.Minute()
	second := now.Second()
	fmt.Println(year, month, day, minute, second)
	timestamp1 := now.Unix()
	timestamp2 := now.UnixNano()
	fmt.Println(timestamp1, timestamp2)
	//将时间戳转化为具体的时间格式
	t := time.Unix(timestamp1, 0)
	fmt.Println(t)
	//时间间隔
	
}
