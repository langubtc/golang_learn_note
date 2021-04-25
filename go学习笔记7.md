go学习笔记7

### 浮点数类型
Go语言提供了两种精度的浮点数 float32 和 float64

浮点数类型具有以下特点

- Golang 的浮点型默认声明为 float64 类型
- 常用的表示有十进制表示，一种为科学计数法
- 开发中建议使用float64，因为比32精度高
- Golang 浮点类型有固定的范围和字段长度，不受具体 OS(操作系统)的影响。


定义浮点数

默认类型是float64
```go
package main

import "fmt"

func main(){
	var floatNumber = 1.23232 // float默认类型为float64
	fmt.Printf("%T",floatNumber)  // %T 输出变量类型，

}
```

浮点数在声明的时候可以只写整数部分或者小数部分，像下面这样：

```go
const SecondFloat = 1.222222222222222212222  // float64 可以提供约 15 个十进制数的精度 1.2222222222222223
const float32Value = float32(1.23222222222323)// float32 只提供6个十进制数的精度 1.2322222
```

科学计数发定义
```text
var newFloat = 1e4
fmt.Println(newFloat)
```

float32和float64的取值范围

```text
float32	3.4e38	math.MaxFloat32
float64	1.8e308	math.MaxFloat64
```

输出是保留小数使用```%f``` 来控制

```go
fmt.Printf("%.2f\n",SecondFloat)
```
 

### 总结

- 默认float为float64，推荐使用float64,因为精度较高
- float可以是10进制数也可以是科学计数法定义




