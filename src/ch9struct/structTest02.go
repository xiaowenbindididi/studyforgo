package main

import "fmt"

type person struct {
	name string
	city string
	age  int8
}

func main() {
	//var p1 person
	//p1.name = "沙河娜扎"
	//p1.city = "北京"
	//p1.age = 18
	//fmt.Printf("p1=%v\n", p1)  //p1={沙河娜扎 北京 18}
	//fmt.Printf("p1=%#v\n", p1) //p1=main.person{name:"沙河娜扎", city:"北京", age:18}

	var p2 = new(person)
	fmt.Printf("%T\n", p2)     //*main.person
	fmt.Printf("p2=%#v\n", p2) //p2=&main.person{name:"", city:"", age:0}
}
