package main

import "fmt"

func main() {
	// 演示 Golang 中, 浮点类型使用

	// 浮点类型包含两种: float32, float64
	// 1. float32 的精度只有 7 位小数, float64 的精度只有 15 位小数
	// 默认是 float64
	// float32 单精度, 占据 4 个字节, 范围 -3.403E38 ~ 3.403E38
	// float64 双精度, 占据 8 个字节, 范围 -1.798E308 ~ 1.798E308

	// 1. 浮点数的表示
	// 1.1. 小数部分和指数部分
	var f1 float32 = 123.456
	fmt.Printf("f1 的类型是 %T, f1 = %v\n", f1, f1)

	var f2 = 123.456
	fmt.Printf("f2 的类型是 %T, f2 = %v\n", f2, f2)
}
