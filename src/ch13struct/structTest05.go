package main

import "fmt"

type person struct {
	name, city string
	age        int
}

func main() {
	p1 := person{
		name: "名字",
		city: "城市",
		age:  11,
	}
	fmt.Printf("%#v\n",p1)

	p2 := &person{
		"李四",
		"南京",
		19,
	}
	fmt.Printf("%#v\n",p2)
}
