package main

import "fmt"

func main() {
	// q: Go 转义字符
	// a: 1. \n: 换行符
	fmt.Println("Hello \nWorld!")

	//    2. \r: 回车符: 从当前行的最前面开始输出, 覆盖掉以前的内容
	fmt.Println("天龙八部雪山飞狐\r张飞")

	//    3. \t: 制表符: 一个制表符相当于 8 个空格
	fmt.Println("Hello \tWorld!")
	fmt.Println("a \tWorld!")
	fmt.Println("123 \tWorld!")

	//    4. \': 单引号
	//    5. \": 双引号
	fmt.Println("Hello \"World!\"")

	//    6. \\: 反斜杠
	fmt.Println("Hello \\World!")

}
