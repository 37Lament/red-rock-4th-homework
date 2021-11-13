package main

import (
	"fmt"
	"io/ioutil"

)



type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type user map[string]string //用map建立映射关系
func MapWrite(m map[string]string, key, value string) {
	m[key] = value
}


type register struct {
	uh            user   // 用户信息
	registerUsers []User // 注册了但未保存的用户
}

func Read() {
	content, err := ioutil.ReadFile("./users.data")
	if err != nil {
		fmt.Println("read file failed, err:", err)
		return
	}
	fmt.Println(string(content))
}
func Write()  {
	var username, password string
	fmt.Scan(&username,&password)
	m := make(map[string]string)
	MapWrite(m, username, password)
	err := ioutil.WriteFile("./user.data", []string(m), 0666)
	if err != nil {
		fmt.Println("write file failed, err:", err)
		return
	}
}

func say() {
	fmt.Println("请选择您要进行的操作：\n1.登录\n2.注册\n3.退出")
	fmt.Scanf("%d", &num)
	switch num {
	case 1:
		{
			fmt.Println("请输入账号")
			fmt.Scanf("%s", &a)
			fmt.Println("请输入密码")
			fmt.Scanf("%s", &b)
			fmt.Println("密码正确")
			else{
			fmt.Println("密码错误") }
		}
	case 2:
		{

		}
	case 3:
		{

		}
	default:
		{
			fmt.Println("输入错误，请重新输入")
		}

	}

}
var num int
func main() {

	for {
		say()
	}

}
