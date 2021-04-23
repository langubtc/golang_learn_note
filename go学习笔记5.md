go学习笔记5

### 基本数据类型

### bool类型(布尔类型)

- 布尔类型无非有两种: true 或者false 。if 和 for 语句的条件部分都是布尔类型的值，并且==和<等比较操作也会产生布尔型的值。
- 布尔型无法参与数值运算，也无法与其他类型进行转换
- 布尔类型的常用操作是做逻辑运算：与、或、非

运算

- 与：左与右都为true，结果为true
- 或：左或右只要一个为true,结果为true
- 非：取反 true ==> false  false ==> true


例：

```go
package main

import "fmt"

func main(){

	var isBig bool // 默认值为false

	fmt.Printf("%T",isBig) //打印isBig的类型
	fmt.Println()
	fmt.Println(isBig)
}

```


### &&与运算

使用两个``&``符号表示
a && b  表示，a = true 并且b= true 结果为true

```go
package main

import "fmt"

func main(){

	a,b,c,d := true,false,true,false

	fmt.Println("a&&b:",a && b) // false
	fmt.Println("a&&c:",a && c) // true
	fmt.Println("c&&b:",c && b) // false
	fmt.Println("a&&d:",a && d) // false
	
}
```

### || 或
使用``||``符号表示
a || b  表示，a = true 或者 b= true 结果为true。两者为false=false


```go
package main

import "fmt"

func main(){
	a,b,c,d := true,false,true,false
	fmt.Println("a||b:",a || b) // true
	fmt.Println("a||c:",a || c) // true
	fmt.Println("c||b:",c || b) // true
	fmt.Println("a||d:",a || d) // true

}

```

### !非运算

非运算使用``!`` 放在变量前作取反操作

```go
package main

import "fmt"

func main(){
	a,b,c,d := true,false,true,false

	fmt.Println("!a",!a) // a = true 取反 true等于false flase 等于true

	fmt.Println("!b",!b)
}
```

### 关系运算

== 来判断是否true或者false

```go
fmt.Println("a == b",a==b) //false
fmt.Println("a!=c",a!=c) // false
fmt.Println( "a!=b",a!=b) // true
```

### 总结

- golang的bool的初始值（零值）为false
- golang中逻辑运算&& || !/与/或/非
- 关系判断使用==来表示






