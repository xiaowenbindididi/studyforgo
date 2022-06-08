package main

import "fmt"

func main() {
	var a int = 4;
	var ptr *int;
	ptr = &a
	*ptr++;
	fmt.Printf("%d\n", *ptr)
	fmt.Printf("%d\n",a)
}
