package main

import "fmt"

func main() {
	var a, b int = 1, 2
	fmt.Println(a, b)

	var name string = "roddy"

	cname := name
	fmt.Println(name, cname)

	c, d := 4, 5
	fmt.Println(c, d, c+d)
}
