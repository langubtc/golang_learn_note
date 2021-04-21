package main

import "fmt"

func Add(a int, b int) (int, int) {
	return a + b, a
}

func main() {
	result, _ := Add(1, 2)
	fmt.Println(result)
}
