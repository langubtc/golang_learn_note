package main

import "fmt"

func main() {

	var name string //声明一个变量name 并且是string类型
	var myName = "roddy"

	var (
		age    int
		phone  int64
		amount float32
	)
	// 给变量赋值
	age = 28
	phone = 13232323232
	amount = 0.23232

	fmt.Println(myName)

	name = "golang"   // 给name赋值 为 golang
	fmt.Println(name) //打印输出
	fmt.Println("age:", age, "phone:", phone, "amount:", amount)

	address := "Sichun Chengdu"
	Job, Title := "IT", "golang开发"
	fmt.Println(address)
	fmt.Println(Job, Title)

}
