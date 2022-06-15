package main

import "fmt"

type person struct {
	name,city string
	age int
}

func newPerson(name, city string, age int) *person {
	return &person{
		name: name,
		city: city,
		age:  age,
	}
}

func main() {
	//p是个地址
	p := newPerson("张三", "沙河", 11)
	fmt.Printf("%#v\n",p)

}
