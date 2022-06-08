package main

import "fmt"

func main() {
	var a, b string = swap("a","b")
	fmt.Println(a,b)
}

func swap(x, y string) (string, string) {
	return y,x
}

