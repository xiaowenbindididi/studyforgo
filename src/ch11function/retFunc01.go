package main

import (
	"fmt"
	"strings"
)

func a() func() {
	return func() {
		fmt.Println("函数作为返回值")
	}
}

func makeSuffixFunc(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}

func calc(base int) (func(int) int, func(int) int) {
	add := func(i int) int {
		base += i
		return base
	}

	sub := func(i int) int {
		base -= i
		return base
	}
	return add, sub
}

func main() {
	add, sub := calc(1)
	ret1 := add(2)
	fmt.Println(ret1)
	j := sub(1)
	fmt.Println(j)
}
