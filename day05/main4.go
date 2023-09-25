package main

import "fmt"

func main() {
	// 演示 Golang 中, 布尔类型的使用

	// 布尔类型的值只可以是常量 true 或者 false
	// 布尔类型占 1 个字节
	// 布尔类型适用于逻辑运算, 一般用于程序流程控制

	var b = false
	fmt.Printf("b 的类型是 %T, b = %v\n", b, b)

}
