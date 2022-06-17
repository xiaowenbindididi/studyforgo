package main

import "fmt"

type mover interface {
	move()
}

type people struct {
	name string
	age  int
}

func (p *people) move() {
	fmt.Printf("%s在跑\n", p.name)
}

func (p *people) say() {
	fmt.Printf("%s在叫\n", p.name)
}
func main() {
	var m mover
	fmt.Printf("%T\n", m)
	p1 := &people{
		name: "啊啊啊",
		age:  19,
	}
	p2 := &people{name: "鹅鹅鹅",
		age: 11}

	m = p1
	fmt.Printf("%#v\n",m)
	m.move()
	m = p2
	fmt.Printf("%#v\n",m)
	m.move()
}
