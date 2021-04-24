package main

import (
	"fmt"
	"strings"
)

func main() {
	var Address = "四川\t成都\t高新区"
	Content := "四川是一个好地方\n成都是一座来了就离不开的城市"

	fmt.Println(Address)
	fmt.Println(Content)

	// 字符串拼接 +

	var Language = "Golang Java Python "

	var webLanguage = "PHP"

	allLanguage := Language + webLanguage + " Shell"

	fmt.Println(allLanguage)

	// 字符串+=
	var webUrl = "https://www.baidu.com"

	webUrl += "/search?name=aaaa"

	fmt.Println(webUrl)

	helloWorld := strings.Join([]string{"hello", "world"}, ",")
	fmt.Println(helloWorld)

	myName := fmt.Sprintf("%s=%s", "name", "roddy")
	fmt.Println(myName)

	// 多行字符串定义

	var multiLine = `
	你好
	golang
	我是你的粉丝
	谢谢	
	`
	fmt.Println(multiLine)

}
