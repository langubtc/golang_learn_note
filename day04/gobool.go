package main

import "fmt"

func main() {

	var isBig bool // 默认值为false

	fmt.Printf("%T", isBig) //打印isBig的类型
	fmt.Println()

	fmt.Println(isBig)

	a, b, c, d := true, false, true, false

	fmt.Println("a&&b:", a && b) // false
	fmt.Println("a&&c:", a && c) // true
	fmt.Println("c&&b:", c && b) // false
	fmt.Println("a&&d:", a && d) // false

	fmt.Println("a||b:", a || b) // true
	fmt.Println("a||c:", a || c) // true
	fmt.Println("c||b:", c || b) // true
	fmt.Println("a||d:", a || d) // true

	fmt.Println("!a", !a) // a = true 取反 true等于false flase 等于true

	fmt.Println("!b", !b)

	fmt.Println("a == b", a == b) //false
	fmt.Println("a!=c", a != c)   // false
	fmt.Println("a!=b", a != b)   // true

}
