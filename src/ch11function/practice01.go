package main

import (
	"fmt"
	"strings"
)

/*
你有50枚金币，需要分配给以下几个人：Matthew,Sarah,Augustus,Heidi,Emilie,Peter,Giana,Adriano,Aaron,Elizabeth。
分配规则如下：
a. 名字中每包含1个'e'或'E'分1枚金币
b. 名字中每包含1个'i'或'I'分2枚金币
c. 名字中每包含1个'o'或'O'分3枚金币
d: 名字中每包含1个'u'或'U'分4枚金币
写一个程序，计算每个用户分到多少金币，以及最后剩余多少金币？
程序结构如下，请实现 ‘dispatchCoin’ 函数
*/
var (
	coins = 50
	users = []string{
		"Matthew", "Sarah", "Augustus", "Heidi", "Emilie", "Peter", "Giana", "Adriano", "Aaron", "Elizabeth",
	}
	distribution = make(map[string]int, len(users))
)

func dispatchCoin() int {
	for _, user := range users {
		distribution[user] = rule(user,0)
	}
	return coins
}
/*
name : name转小写
 */
func rule(name string,input int) int {
	var output int = input
	if strings.Contains(name,"e") {
		coins -= 1
		output += 1
	}
	if strings.Contains(name, "i") {
		coins -= 2
		output += 2
	}
	if strings.Contains(name, "o") {
		coins -= 3
		output += 3
	}
	if strings.Contains(name, "u") {
		coins -= 4
		output += 4
	}
	return output
}

func main() {
	left := dispatchCoin()
	fmt.Println("剩下：", left)
}
