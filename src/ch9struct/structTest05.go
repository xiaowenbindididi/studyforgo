package main

import "fmt"
// 接收者
type Person struct {
	name string
	age int
}

func NewPerson(name string, age int) *Person {
	return &Person{
		name: name,
		age: age,
	}
}

func (p Person) Dream() {
	fmt.Println("测试")
}

func main() {
	newPerson := NewPerson("王五", 18)
	newPerson.Dream()
}
