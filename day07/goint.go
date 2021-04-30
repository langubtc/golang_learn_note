package main

import "fmt"

func main() {
	a := 10
	fmt.Println(a)
	fmt.Printf("%T", a)
	fmt.Println()

	//输出
	//10
	//int

	v := 123_456 // 下划线也会自动格式化成

	fmt.Println(v)

	var d uint = 12

	fmt.Printf("value:%v,%T", d, d)

}
