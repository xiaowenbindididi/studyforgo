package main

import "fmt"

type person struct {
	name, city string
	age int
}

func main() {
	p := new(person)
	fmt.Printf("%p,%#v\n",p,p)
	p.name = "11"
	p.city = "xxx"
	p.age = 33
	fmt.Printf("%p,%#v\n",p,p)
}
