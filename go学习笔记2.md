golang自学笔记2

本节主要讲述golang的变量声明方式已经变量的调用

### 变量
变量就是可变化的数，没有固定的值。但是在计算机中是一段或多段存储数据的内存

#### 声明方式

变量声明有以下三种形式
- 使用``var``关键字声明单个变量
- 使用``var ()`` 批量声明变量
- 简单格式``:=`` 声明

#### Go语言的基本类型
- bool
- string
- int、int8、int16、int32、int64
- uint、uint8、uint16、uint32、uint64、uintptr
- byte // uint8 的别名
- rune // int32 的别名 代表一个 Unicode 码
- float32、float64
- complex64、complex128

### 说明

变量被声明之后，系统自动赋予它该类型的零值：int 为 0，float 为 0.0，bool 为 false，string 为空字符串，指针为 nil 等。所有的内存在 Go 中都是经过初始化的。
变量的命名规则遵循骆驼命名法，即首个单词小写，每个新单词的首字母大写，例如：numShips 和 startDate 。


### 定义变量var

使用var 关键字定义变量，语法 ``var [变量名] [类型]``

```go
package main

import "fmt"

func main(){

	var name string   //声明一个变量name 并且是string类型

	name = "golang"  // 给name赋值 为 golang

	fmt.Println(name)  //打印输出

}
```

使用var 声明后直接赋予值 ``var name = roddy``
```go
package main

import "fmt"

func main() {
	var myName = "roddy" //声明一个变量myName 值是"roddy"，golang会根据值自动声明变量的类型
	fmt.Println(myName)
}

```



### 定义多个变量

var 后面加()定义，在()中声明多个变量

```go
package main

import "fmt"

func main(){

	var name string   //声明一个变量name 并且是string类型

	var (
		age int
		phone int64
		amount float32
	)
	// 给变量赋值
	age = 28
	phone = 13232323232
	amount  = 0.23232
	
	name = "golang"  // 给name赋值 为 golang
	fmt.Println(name)  //打印输出
	fmt.Println("age:",age,"phone:",phone,"amount:",amount)

}
```


### 简短声明

简短声明使用``:=``声明一个变量，并且有如下要求.

- 定义变量，同时显式初始化。 前面是变量名:=变量值
- 不能提供数据类型。无需什么是什么类型，根据变量的值确定
- 只能用在函数内部。 无法在函数体外包，只能在函数体内部,main也是函数体

也可以用来声明和初始化一组变量

```go
address := "Sichun Chengdu"

Job,Title := "IT","golang开发"
```

因为简短声明的方式比较方便，所以在开发中会大量使用。

### 总结
- 变量的声明有三种方式,var var() :=
- :=只能在函数体内部且要同时初始化(赋值)
- 变量声明建议使用驼峰法
- 变量声明系统会模式赋予零值



