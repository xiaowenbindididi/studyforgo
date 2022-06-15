package main

import "fmt"

type MyInt int
func main() {
	var i MyInt;
	fmt.Println(i)
	fmt.Printf("type:%T",i)
}
