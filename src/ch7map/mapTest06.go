package main

import "fmt"

func main() {
	//Map : string - []int
	type Map map[string][]int
	m := make(Map)
	s := []int{1, 2}
	s = append(s, 3)
	//s - 1,2,3
	fmt.Printf("%+v\n", s)
	m["q1mi"] = s
	//m["q1mi"] - 1,2,3
	s = append(s[:1], s[2:]...)
	//s - 1,3
	fmt.Printf("%+v\n", s)
	fmt.Printf("%+v\n", m["q1mi"])
}
