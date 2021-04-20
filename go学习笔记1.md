golang自学笔记1

本节主要讲述golang的程序构成及如何编写第一个go程序，并且讲解如何给代码添加注释

## 如何学习
- 知识点分解
- 刻意练习（练习不会的，多看多想，不要复制代码）
- 反馈 （主动反馈/被动反馈/copy/借鉴别人代码）

## 概述
golang 安装掠过...

### 第一节

#### 打印 hello world

```go
package main

import "fmt"

func main(){

	fmt.Println("Hello World")

}
```

运行程序

```shell script
$ go build helloworld.go 
$ ls
helloworld    helloworld.go #生成了可执行文件 helloworld
$ ./helloworld   # 执行完后直接可执行，如果不可执行需要chmod +x 加可执行权限
Hello World
```


#### 代码解释

- golang以包作为管理单位，每个go文件必须声明他所属的包。以关键字``package``来声明。
- 导入包使用``import``关键字进行导入，需要注意的是import导入的包必须要使用，否者编译器会报错。如果导入多个包可以使用()包括起来，参考以下
- ``main``函数是go程序的入口，也就是说是程序启动运行的第一个函数，``main`` 函数只能声明在``main``包中且一个main包中只能有一个``main``函数
- ``fmt.Println`` Println是fmt包中的一个函数``print``意为打印/输出 ``ln``表示 ``line``的缩写。表示打印后自动换行。
- 其中``.``在golang中表示调用的意思，表示调用fmt包中的Println方法。

```go
package  main
// 导入包参考
import (
	"fmt"
	"os"
)
```

```go
package  main   // 声明main包

import "fmt"    // 导入fmt包

func main(){    //定义main函数

	fmt.Println("Hello World")  //调用fmt包中的Println方法进行输出

}
```

#### 编译运行
go程序运行需要先进行编译，然后才可以进行使用。

在build时候可以不加文件，直接在当前目录下执行后会编译当前包中的文件. 
- ``-x`` 查看golang 编译的过程
- ``-o`` 指定编译后生成的文件名

```shell script
go build
```

```shell script
go build helloworld.go
```

```shell script
go help #查看go命令帮助
```

格式化golang的代码

一般编辑器会自动帮我们格式化

```shell script
go fmt helloworld.go #自动格式代码
```




先编译后运行，会在编译之后立即执行Go语言程序，但是不会生成可执行文件。
```shell script
go run hello.go
```

指定生成的文件名``hello``

```shell script
$ go build -o hello helloworld.go 
$ ll
total 8488
drwxr-xr-x  5 luodi  staff      160 Apr 20 22:58 ./
drwxr-xr-x  8 luodi  staff      256 Apr 20 22:57 ../
-rwxr-xr-x  1 luodi  staff  2169864 Apr 20 22:58 hello*
-rwxr-xr-x  1 luodi  staff  2169864 Apr 20 22:56 helloworld*
-rw-r--r--  1 luodi  staff       79 Apr 20 22:43 helloworld.go
```


### 注释

golang注释使用以下两种注释方式
- 块注释
- 行注释

行注释使用``//``在前，注释在后
```go
package main

import (
	"fmt"
)
//这里是main入口函数
func main() {

	fmt.Println("hello world")

}

``` 

块注释使用``/* */``，注释内容在中间

```go
package main

import (
	"fmt"
)
//这里是main入口函数
func main() {
	/*
		这里执行主要方法
		方法1.
		方法2.
	 */
	fmt.Println("hello world")

}
```


### 总结

- go的每个文件必须声明一个包名，一个main包中有且只有一个main入口函数
- 使用import 导入包
- 使用.调用包中的方法
- go语言是静态编译型语言，需要先编译后使用的，使用的命令为build




