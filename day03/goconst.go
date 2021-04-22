package main

import "fmt"

// 声明常量的两种方式
const (
	Name string = "Roddy" // 类型 = 值
	Age         = 10      // = 值
)

const (
	A = 100
	B //使用前一个值进行初始化为100
	C //使用前一个值进行初始化
	D = 1000
	F //使用前一个值进行初始化为1000
	E //使用前一个值进行初始化

)

const (
	BTC = iota // 必须在常量组内使用，初始化是0，每次调用+1
	LTC        // 1
	ETH        // 2
	BCH = iota
	EOS // =4 中间iota没有作用，上一个值为3

)

func main() {

	// 使用const定义

	const CPU = 2
	const Mem = 8
	// 常量不可以修改值
	//CPU = 5
	fmt.Println("CPU:", CPU, "内存:", Mem)
	fmt.Println(Name, Age)
	fmt.Println(B, C)
	fmt.Println(F, E)
	fmt.Println(BTC, ETH, LTC, BCH, EOS)

	fmt.Println("打印并换行")
	fmt.Print("只打印不换行")
	fmt.Println("看看是不是不换行")
	fmt.Printf("Name is :%s", Name)
	bitcoinInfo := fmt.Sprintf("The bitcoin is very good ,now price:%.2f,diff:%d", 6500.2322, 644232323222)
	fmt.Print(bitcoinInfo)
}
