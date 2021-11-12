package main

import (
	"fmt"
	"sync"
)
var wg sync.WaitGroup

func main() {
	var a string
	var b string
	scoreMap := make(map[string]string, 8)
	scoreMap["1"] = "qwerqwer"
	scoreMap["2"] = "wwwwwwww"
	fmt.Println(scoreMap)
	fmt.Println(scoreMap["小明"])
	fmt.Printf("type of a:%T\n", scoreMap)

	fmt.Println("请输入账号")
	fmt.Scanf("%s", &a)
	fmt.Println("请输入密码")
	fmt.Scanf("%s", &b)
	if(scoreMap[a]==b)  {
		fmt.Println("密码正确")
	} else{
		fmt.Println("密码错误")
	}
}
