package main

import "fmt"

//定义结构体
type person struct {
	name string
	age  int
	city string
}

func main() {
	var p person
	p.name = "滴滴"
	p.city = "北京"
	p.age = 18
	fmt.Printf("p = %v\n",p)
	fmt.Printf("p = %#v\n",p)

	var student struct{
		name string
		grade float32
	}

	student.name = "李四"
	student.grade = 11.11
	fmt.Printf("%#v\n",student)
}
