go学习笔记9

### 字符串和数据类型的相互转换
Go语言中的 strconv 包为我们提供了字符串和基本数据类型之间的转换功能。

### 整型转字符串
字符串和整型之间的转换是我们平时编程中使用的最多的
strconv.Itoa 用于整数类型转换为字符串类型

参考样例:

```go
package main

import (
	"fmt"
	"strconv"
)

func main(){

	intStr := strconv.Itoa(100)
	fmt.Printf("type:%T,value:%s",intStr,intStr)

}

```

Out输出:
此时看到输出类型为string 值为100
```text
type:string,value:100
```

### 字符串转整型
strconv.Atoi 用于字符串类型转换为整数类型,Atoi() 函数有两个返回值，i 为转换成功的整型，err 在转换成功是为空转换失败时为相应的错误信息

参考样例：
```go
package main

import (
	"fmt"
	"strconv"
)

func main(){

	var age  string = "23"

	intAge,err := strconv.Atoi(age) // 第一个为转换的值，第二个参数为错误消息

	if err != nil{
		fmt.Println("error message:",err)
	}
	fmt.Printf("age:%T,value:%v",intAge,intAge)

}

```

如果转换失败提示错误消息

```text
error message: strconv.Atoi: parsing "23w": invalid syntax
age:int,value:0

```

默认成功

```text
age:int,value:23

```

### Parse 系列函数

Parse 系列函数用于将字符串转换为指定类型的值，其中包括 ParseBool()、ParseFloat()、ParseInt()、ParseUint()。


#### ParseBool

ParseBool() 函数用于将字符串转换为 bool 类型的值，它只能接受 1、0、t、f、T、F、true、false、True、False、TRUE、FALSE，其它的值均返回错误，函数签名如下。

同样ParseBool也是有两个返回值

```go
package main

import (
	"fmt"
	"strconv"
)

func main(){

	strBool,err := strconv.ParseBool("t")  //只能是1、0、t、f、T、F、true、false、True、False、TRUE、FALSE
	if err != nil{
		fmt.Printf("error message:",err)
	}

	fmt.Printf("type:%T, value:%v",strBool,strBool)

}
```

#### ParseFloat

将一个表示浮点数的字符串转换为 float 类型
 ``strconv.ParseFloat(strvalue,bitsize) ``
bitSize 指定了返回值的类型，32 表示 float32，64 表示 float64；

```go
package main
import (
	"fmt"
	"strconv"
)

func main(){
	// 将字符串的浮点类型转为一个float类型

	var strDiff string = "7343438.232"
	//bitSize 指定了返回值的类型，32 表示 float32，64 表示 float64；
	floatDiff,err := strconv.ParseFloat(strDiff,64)
	if err!=nil{
		fmt.Print("error message:",err)
	}
	fmt.Printf("float:%T,value:%v",floatDiff,floatDiff)

}
```

#### ParseInt

ParseInt() 函数用于返回字符串表示的整数值（可以包含正负号）

- base 指定进制，取值范围是 2 到 36。如果 base 为 0，则会从字符串前置判断，“0x”是 16 进制，“0”是 8 进制，否则是 10 进制。
- bitSize 指定结果必须能无溢出赋值的整数类型，0、8、16、32、64 分别代表 int、int8、int16、int32、int64。
- 返回的 err 是 *NumErr 类型的，如果语法有误，err.Error = ErrSyntax，如果结果超出类型范围 err.Error = ErrRange。

```go
package  main
import (
	"fmt"
	"strconv"
)

func main() {
    str := "-11"
    num, err := strconv.ParseInt(str, 10, 0) // 0表示默认的int 
    if err != nil {
        fmt.Println(err)
    } else {
        fmt.Println(num)
    }
}
```



#### ParseUint

ParseUint() 函数的功能类似于 ParseInt() 函数，但 ParseUint() 函数不接受正负号，用于无符号整型

```go
package  main
import (
	"fmt"
	"strconv"
)

func main() {
    str := "11"
    num, err := strconv.ParseUint(str, 10, 0)
    if err != nil {
        fmt.Println(err)
    } else {
        fmt.Println(num)
    }
}
```


### 总结

- 字符串转int使用 strconv.Atoi
- int转字符串使用 strconv.Itoa
- Parse系列可以将字符串转为指定的类型格式






