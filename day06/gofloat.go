package main

import "fmt"

func main() {
	var floatNumber = 1.23232     // float默认类型为float64
	fmt.Printf("%T", floatNumber) // %T 输出变量类型，
	fmt.Println()
	const firstFloat = .23232
	const SecondFloat = 1.222222222222222212222    // float64 可以提供约 15 个十进制数的精度 1.2222222222222223
	const float32Value = float32(1.23222222222323) // float32 只提供6个十进制数的精度 1.2322222

	var newFloat = 1e4
	fmt.Println(newFloat)
	fmt.Println(firstFloat)
	fmt.Println(SecondFloat)
	fmt.Println(float32Value)

	fmt.Printf("%.2f\n", SecondFloat)

}
