package main

import "fmt"

func main() {
	var ret int = calc(11,12,add)
	fmt.Println(ret)
}
func add(a, b int) int{
	return a + b;
}
func sub(a,b int) int{
	return a - b;
}
//函数作为参数的情况
func calc(x, y int, op func(int, int) int) int{
	return op(x, y)
}
