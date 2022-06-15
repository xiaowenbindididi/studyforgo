package main

import "fmt"

type student struct {
	name string
	age  int
}

func main() {
	m := make(map[string]*student)
	stus := []student{
		{name: "小王子", age: 18},
		{name: "娜扎", age: 23},
		{name: "大王八", age: 9000},
	}

	fmt.Printf("0%p,%#v\n",&stus[0],stus[0])
	fmt.Printf("1%p,%#v\n",&stus[1],stus[1])
	fmt.Printf("2%p,%#v\n",&stus[2],stus[2])

	for _, stu := range stus {
		fmt.Printf("%#v\n",stu)
		fmt.Printf("%p\n",&stu)
		m[stu.name] = &stu
	}
	for k, v := range m {
		fmt.Println(k, "=>", v.name)
	}
}
/*
因为for range循环只是值拷贝进行循环，其实每个值指向的地址是一样的，相当于用一个temp存值
 */
//小王子 => 大王八
//娜扎 => 大王八
//大王八 => 大王八