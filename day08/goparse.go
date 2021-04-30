package main

import (
	"fmt"
	"strconv"
)

func main() {

	strBool, err := strconv.ParseBool("1") //只能是1、0、t、f、T、F、true、false、True、False、TRUE、FALSE
	if err != nil {
		fmt.Printf("error message:", err)
	}

	fmt.Printf("type:%T, value:%v", strBool, strBool)

	str := "-11"
	num, err := strconv.ParseInt(str, 10, 0)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(num)
		fmt.Printf("%T", num)
	}

}
