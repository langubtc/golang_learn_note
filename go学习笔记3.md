go学习笔记3


### 多个变量赋值
多个变量赋值，多重赋值时，变量的左值和右值按从左到右的顺序赋值。

``a,b := 1,2``

复制多个变量还可以使用``var ()``来定义

```go
package main

import "fmt"

func main() {
	var a, b int = 1, 2
	fmt.Println(a, b)

	var name string = "roddy"

	cname := name
	fmt.Println(name, cname)


	c,d := 4,5
	fmt.Println(c,d,c+d)
}
```

可以使用:= 定义，也可以直接在var初始化的时候定义
```go
package main

import "fmt"

func main(){
	var a ,b int = 1,2
	fmt.Println(a,b)


	var name string = "roddy"
	
	cname := name
	fmt.Println(name,cname)


}
```


### 匿名变量

匿名变量的特点是一个下画线```“_”```，``“_”``本身就是一个特殊的标识符，被称为空白标识符。又称为占位符，可以赋予任何变量。
匿名变量不占用内存空间，不会分配内存。匿名变量与匿名变量之间也不会因为多次声明而无法使用。

```go
package main

import "fmt"

func Add(a int, b int)(int, int){
	return a+b,a
}

func main(){
	result,_:= Add(1,2) // _表示匿名变量
	fmt.Println(result)
}

```

### 变量的作用域

变量的作用域分为几个级别

- 包级别  又称为全局变量
- 函数级别 又称为局部变量

注意：

- ``函数级别变量在调用完函数后自动销毁``
- ``包级别变量可以被外包引用，但是变量名首字母必须大写``

例：

```go
package main

import "fmt"

// 这个是包级别变量
var level string = "debug"

func main(){
	// 这个是函数级别变量
	var funcName string = "roddy"

	fmt.Println(funcName)
	fmt.Println(level)
}
```

### block中的变量
golang 使用``{}`` 定义一个block，子block中的函数无法被外面的方法调用。子块里面可以调用函数界别的变量。


```go
package main

import "fmt"

// 这个是包级别变量
var level string = "debug"

func main(){
	// 这个是函数级别变量
	var funcName string = "roddy"

	fmt.Println(funcName)
	fmt.Println(level)
	
	{  //使用大括号定义一个block,该block中也可以定义函数，
		
		var cc string = "sSDD"
		
		fmt.Println(cc,funcName)
		
	}
	// 子块里面定义的函数，外面无法调用，只能在block中调用
	//fmt.Println(cc) 这里会报错

}
```

### 同名变量

Go语言程序中全局变量与局部变量名称可以相同，但是函数体内的局部变量会被优先考虑。

```go
package main

import "fmt"

// 这个是包级别变量
var level string = "debug"
var funcName string = "roddy1"
func main(){



	// 这个是函数级别变量
	var funcName string = "roddy"
	fmt.Println(funcName)
	fmt.Println(level)

	{  //使用大括号定义一个block,该block中也可以定义函数，

		var cc string = "sSDD"

		fmt.Println(cc,funcName)

	}
	// 子块里面定义的函数，外面无法调用，只能在block中调用
	//fmt.Println(cc) 这里会报错

}
```

此时会输出``roddy``



### 总结

- 匿名变量使用``_`` 来充当占位，可以赋予任何类型
- 变量是有作用域的，作用域可以为全局的，局部的
- go使用``{}``来声明一个块，块中定义的变量不能被外面(函数，包等)调用






