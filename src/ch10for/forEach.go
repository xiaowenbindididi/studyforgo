package main

import "fmt"

func main() {
	strings := []string{"google","hello"}
	for i, s:= range strings{
		fmt.Println(i,s)
	}
}
