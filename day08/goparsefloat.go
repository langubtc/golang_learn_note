package main

import (
	"fmt"
	"strconv"
)

func main() {
	// 将字符串的浮点类型转为一个float类型

	var strDiff string = "7343438.232"
	//bitSize 指定了返回值的类型，32 表示 float32，64 表示 float64；
	floatDiff, err := strconv.ParseFloat(strDiff, 64)
	if err != nil {
		fmt.Print("error message:", err)
	}
	fmt.Printf("float:%T,value:%v", floatDiff, floatDiff)

}
