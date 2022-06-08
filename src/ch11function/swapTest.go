package main

import "fmt"

func main() {
	var a int = 100
	var b int = 200
	fmt.Printf("before exchange a = %d\n", a)
	fmt.Printf("before exchange b = %d\n", b)
	swap(&a, &b)
	fmt.Printf("before exchange a = %d\n", a)
	fmt.Printf("before exchange b = %d\n", b)
}
func swap(x *int, y *int)  {
	var temp int
	temp = *x
	*x = *y
	*y = temp
}
