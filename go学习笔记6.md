go学习笔记6

### 概述
字符串在golang中是一种类型，且值不可变，一旦创建某个文本后就无法修改这个文本内容

### 定义字符串
字符使用``""`` 来定义，字符串中间可以使用以下的转义关键字。

- 换行符: \n
- 回车符: \r
- tab键: \t
- Unicode字符: \u 或者\U
- 反斜杠: \\\   相当转义

```go
package main

import "fmt"

func main(){
	var Address = "四川\t成都\t高新区"
	Content := "四川是一个好地方\n成都是一座来了就离不开的城市"

	fmt.Println(Address)
	fmt.Println(Content)

}
```

Out输出
```text
四川    成都    高新区
四川是一个好地方
成都是一座来了就离不开的城市
```

### 字符串拼接
- 字符串拼接可以使用``+``号来连接左右两边的字符串
- 还可以使用``+=``来拼接字符串
- 可以使用strings.Join()方法
- 使用fmt.Sprintf()方法

#### 使用+号

```go
// 字符串拼接

var Language = "Golang Java Python "

var webLanguage = "PHP"

allLanguage := Language + webLanguage + " Shell"

fmt.Println(allLanguage)

```

Out输出:
```text
Golang Java Python PHP Shell
```

#### 使用+=

```go
// 字符串+=
var webUrl = "https://www.baidu.com"

webUrl += "/search?name=aaaa"

fmt.Println(webUrl)
```

#### 使用strings.Join()

Out输出:
```text
https://www.baidu.com/search?name=aaaa
```


使用加号拼接字符串，每次都会必须重新分配内存。如果是频繁的使用+进行拼接会存在严重的性能问题。 
解决方法是使用``strings.Join()``预分配足够大的内存空间,它会统计所有参数的长度，并一次性完成内存的分配操作。

```go

package main

import (
	"fmt"
	"strings"
)

func main(){

	helloWorld := strings.Join([]string{"hello", "world"}, ",") // ","为连接字符 可以是空格等。
	fmt.Println(helloWorld)
}

```

Out输出
```text
hello,world

```

### 使用fmt方法
使用的是fmt.Sprintf方法，里面可以将字符串格式化后赋予一个变量，该变量就是连接后的值。
```go
package main

import (
	"fmt"

)
func main(){
	myName:=fmt.Sprintf("%s=%s","name","roddy")
	fmt.Println(myName)
}
```

Out输出:
```text
name=roddy
```

### 多行字符串

一般定义字符串使用的是双引号定义，但是不能跨行定义字符串。如果要定义跨行字符需要使用`反引号定义

```go
package main

import (
	"fmt"
)

func main(){
	var multiLine = `
	你好
	golang
	我是你的粉丝
	谢谢	
	`
	fmt.Println(multiLine)

}

```

Out输出:

```text
        你好
        golang
        我是你的粉丝
        谢谢    

```

### 总结
- golang字符串可以使用双引号和反引号定义
- 字符串拼接可以使用+号等多种方法，方法根据不同条件下自己选择，性能最优选strings
- 可以在字符串中使用常用的转义关键字