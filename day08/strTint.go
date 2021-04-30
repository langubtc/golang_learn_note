package main

import (
	"fmt"
	"strconv"
)

func main() {

	var age string = "23"

	intAge, err := strconv.Atoi(age)

	if err != nil {
		fmt.Println("error message:", err)
	}
	fmt.Printf("age:%T,value:%v", intAge, intAge)

}
