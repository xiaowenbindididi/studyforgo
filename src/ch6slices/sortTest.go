package main

import (
	"fmt"
	"sort"
)

func main() {
	var a []int
	a = append(a, 1, 9, 1, 4, 5, 6)
	sort.Ints(a)
	fmt.Println(a)
}
