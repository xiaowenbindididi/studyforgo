package main

import "fmt"

func main() {
	var a int = 10
	var b *int = &a
	fmt.Printf("%v %p\n",a,&a)
	fmt.Printf("%v %p\n",b, &b)
	var c int = *b
	fmt.Println(c)
}


