package main

import "fmt"

func main() {
	// 从切片中删除元素
	a := []int{30, 31, 32, 33, 34, 35, 36, 37}
	// 要删除索引为2的元素
	//a = append(a[:2], a[3:]...)
	a = append(a[:3], a[4:]...)
	fmt.Println(a) //[30 31 33 34 35 36 37]
}
