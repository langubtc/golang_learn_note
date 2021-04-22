golang学习笔记4

### 概述

本节需要了解的是golang的常量和golang的打印输出


### 常量
常量就是不能变化的量，声明方式使用``const`` 关键字。常量一旦声明，那么就不可以修改

- 常量也是和变量一样，可以在包级别、函数级别、块级定义常量。
- 常量声明后可以不使用，不会报错

```go
package main

import "fmt"

func main(){
	// 使用const定义
	const CPU = 2
	const Mem = 8
	// 常量不可以修改值
	//CPU = 5
	fmt.Println("CPU:",CPU,"内存:",Mem)
}
```


声明常量的两种方式
- const关键字 + 变量名 + 类型 + 值
- const关键字 + 变量名 + 值

```go

// 声明常量的两种方式
const (
	Name string = "Roddy" // 类型 = 值
	Age = 10   // = 值
)

```

常量声明还可以省略声明的方式,如果不赋值，那么会使用前一个变量的值进行初始化

```go
const (
	A = 100
	B		//使用前一个值进行初始化为100
	C		//使用前一个值进行初始化
	D = 1000
	F		//使用前一个值进行初始化为1000
	E		//使用前一个值进行初始化

)
```

### iota常量生成器

常量声明可以使用 iota 常量生成器初始化，它用于生成一组以相似规则初始化的常量，但是不用每行都写一遍初始化表达式。在一个 const 声明语句中，在第一个声明的常量所在的行，iota 将会被置为 0，然后在每一个有常量声明的行加一。
通常是用在枚举类型

```go
const (
    Sunday  = iota // 周日为0
    Monday         // 周一等于1
    Tuesday
    Wednesday
    Thursday
    Friday
    Saturday
)
```

如果中间又定义了iota,那么会自动继承上一个常量的值
```go

const (
	BTC  = iota // 必须在常量组内使用，初始化是0，每次调用+1
	LTC			// 1
	ETH			// 2
	BCH	= iota
	EOS			// =4 中间iota没有作用，上一个值为3 
```



### Print打印输出

Println和Print区别
```go
fmt.Println("打印并换行")
fmt.Print("只打印不换行")
fmt.Print("看看是不是不换行")
```


格式化打印``Printf``
```go
fmt.Printf("Name is :%s",Name)
```

fmt.Sprintf（格式化输出）
变量名 := fmt.Sprintf('xxxx')的内容

```go
bitcoinInfo :=fmt.Sprintf("The bitcoin is very good ,now price:%.2f,diff:%d",6500.2322,644232323222)
fmt.Print(bitcoinInfo)
```
### golang打印输出的动词
参考以下格式化方式

| 动词 | 功能 |
|--- |--- |
|%v |按值的本来值输出|
|%+v	|在 %v 基础上，对结构体字段名和值进行展开|
|%#v	|输出 Go 语言语法格式的值|
|%T	|输出 Go 语言语法格式的类型和值|
|%%	|输出 % 本体|
|%b	|整型以二进制方式显示|
|%o	|整型以八进制方式显示|
|%d	|整型以十进制方式显示|
|%x	|整型以十六进制方式显示|
|%X	|整型以十六进制、字母大写方式显示|
|%U	|Unicode 字符|
|%f	|浮点数|
|%p|	指针，十六进制方式显示|

### 总结
- 常量定义为const关键字
- iota默认初始化值为0
- 格式化使用fmt包
- 输出格式参考以上表格






