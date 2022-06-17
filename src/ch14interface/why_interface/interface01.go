package main

import "fmt"

type cat struct {}
type dog struct {}
type person struct {
	name string
}
type sayer interface {
	say()
}

func (c cat) say()  {
	fmt.Println("喵喵喵")
}
func (d dog) say()  {
	fmt.Println("汪汪汪")
}
func (p person)say()  {
	fmt.Println(p.name + "啊啊啊")
}

func bit(say sayer)  {
	say.say()
}
func main() {
	c := cat{}
	bit(c)
	d := dog{}
	bit(d)
	p := person{name: "王五"}
	bit(p)

}
