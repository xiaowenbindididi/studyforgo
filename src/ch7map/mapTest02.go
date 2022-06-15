package main

import "fmt"

func main() {
	scoreMap := make(map[string]int)
	scoreMap["张三"] = 90
	scoreMap["小明"] = 100
	// 如果key存在ok为true,v为对应的值；不存在ok为false,v为值类型的零值
	v, ok := scoreMap["张si"]
	m, yes := scoreMap["小明"]
	if ok {
		fmt.Println(v)
	}
	if yes {
		fmt.Println(m)
	}
}
