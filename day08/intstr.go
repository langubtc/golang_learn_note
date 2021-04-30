package main

import (
	"fmt"
	"strconv"
)

func main() {

	intStr := strconv.Itoa(100)
	fmt.Printf("type:%T,value:%s", intStr, intStr)

}
