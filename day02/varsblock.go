package main

import "fmt"

// 这个是包级别变量
var level string = "debug"
var funcName string = "roddy1"

func main() {

	// 这个是函数级别变量
	var funcName string = "roddy"
	fmt.Println(funcName)
	fmt.Println(level)

	{ //使用大括号定义一个block,该block中也可以定义函数，

		var cc string = "sSDD"

		fmt.Println(cc, funcName)

	}
	// 子块里面定义的函数，外面无法调用，只能在block中调用
	//fmt.Println(cc) 这里会报错

}
