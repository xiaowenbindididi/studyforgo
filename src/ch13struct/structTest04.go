package main

import "fmt"

type person struct {
	name, city string
	age        int
}

func main() {
	var p1 = new(person)
	//(*p1).city = "上海"
	//(*p1).age = 11
	//(*p1).name = "李四"
	p1.city = "北京"
	p1.age = 12
	p1.name = "王五"
	fmt.Printf("%#v\n",p1)
	fmt.Printf("%T\n",p1)
	fmt.Printf("%v",p1)

	//取结构体的地址进行实例化
	fmt.Println("\np3:")
	p3 := &person{}
	fmt.Printf("%T\n",p3)
	fmt.Printf("%#v\n",p3)
	fmt.Printf("%#v\n",*p3)
}
