package main

import "fmt"

type person struct {
	name string
	city string
	age  int8
}

func main() {
	//var user struct{Name string; Age int}
	//user.Name = "小王子"
	//user.Age = 18
	//fmt.Printf("%#v\n", user)
	//var p2 = new(person)
	//fmt.Printf("%T\n", p2)     //*main.person
	//fmt.Printf("p2=%#v\n", p2) //p2=&main.person{name:"", city:"", age:0}
	p3 := &person{}
	fmt.Printf("%T\n", p3)     //*main.person
	fmt.Printf("p3=%#v\n", p3) //p3=&main.person{name:"", city:"", age:0}
	p3.name = "七米"
	p3.age = 30
	p3.city = "成都"
	fmt.Printf("p3=%#v\n", p3) //p3=&main.person{name:"七米", city:"成都", age:30}
}
