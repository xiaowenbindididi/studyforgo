package main

import "fmt"

type person struct {
	name string
	city string
	age  int8
}
//检验没有初始化的结构体，成员变量值是多少
func main() {
	var p4 person
	fmt.Printf("addr = %p\n p4=%#v\n",&p4 ,p4) //p4=main.person{name:"", city:"", age:0}
	p4.name = "111"
	fmt.Printf("addr = %p\n p4=%#v\n",&p4 ,p4) //p4=main.person{name:"", city:"", age:0}
}
