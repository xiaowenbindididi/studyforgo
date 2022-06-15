package main

import "fmt"

func main() {
	//var a *int = new(int)
	//fmt.Println(*a)
	//*a = 100
	//fmt.Println(*a)

	var b map[string]int
	b = make(map[string]int,10)
	b["沙河娜扎"] = 100
	fmt.Println(b)

}
