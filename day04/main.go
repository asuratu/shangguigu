package main

import (
	"fmt"
)

func main() {
	// 变量介绍
	// 1. 变量的概念
	// 1.1. 变量是计算机中存储数据的一块内存空间
	// 1.2. 变量是程序运行期间可以改变的量
	// 1.3. Go 语言是静态语言, 变量的类型一旦确定, 在程序运行期间不可改变

	var name string
	name = "john"
	fmt.Println(name)

	// 2. 变量的三要素
	// 2.1. 变量名 (标识符)
	// 2.2. 变量的类型 (数据类型)
	// 2.3. 变量的值 (字面量)

	// 3. 变量的使用步骤
	// 3.1. 声明变量
	// 3.2. 给变量赋值
	// 3.3. 使用变量

	// 4. 变量的声明
	// 4.1. var 变量名 数据类型
	// 4.2. var 变量名 数据类型 = 变量值
	// 4.3  变量名 := 变量值 (只能在函数内部使用)

	// 5. 变量使用的注意事项
	// 5.1. 变量表示内存一个存储区域, 该区域有自己的名称(变量名)和类型(数据类型)
	// 5.2. 一次性声明多个变量
	// 5.2.1. var a, b, c int

}
