package main

import "fmt"

func main() {
	var x interface{}
	x = "hello"
	fmt.Printf("%#v\n", x)
	x = 100
	fmt.Printf("%#v\n", x)
	x = false
	fmt.Printf("%#v\n", x)

	//var m = make(map[string]interface{}, 15)
	//m["name"] = "扎不都得勒"
	//m["age"] = 18
	//m["hobby"] = []string{"篮球", "足球"}
	//fmt.Printf("%#v\n", m)
	ret, ok := x.(string)
	if ok {
		fmt.Printf("是字符串类型\n")
	} else {
		fmt.Println("不是字符串类型\n")
	}
	fmt.Printf("%#v\n", ret)
	ret1, ok := x.(bool)
	if ok {
		fmt.Printf("是布尔类型\n")
	} else {
		fmt.Println("不是布尔类型\n")
	}
	fmt.Printf("%#v\n", ret1)

	//使用Switch进行类型断言
	switch v := x.(type) {
	case int:
		fmt.Printf("int类型,value: %v\n",v)
	case string:
		fmt.Printf("string类型,value:%v\n",v)
	case bool:
		fmt.Printf("bool类型,value:%v\n",v)
	default:
		fmt.Printf("猜不到了。艹,value:%v\n",v)
	}
}
