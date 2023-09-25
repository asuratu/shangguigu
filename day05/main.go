package main

import (
	"fmt"
	"unsafe"
)

func main() {
	// 演示 Golang 中, 整数类型的使用

	var i = 1 // int 类型, 有符号, int32
	// %T: 查看变量类型
	// %d: 十进制
	// %b: 二进制
	// %s: 字符串
	// %c: 字符
	// %p: 指针
	// %v: 值
	// %f: 浮点数
	// %t: 布尔值
	fmt.Printf("i 的类型是 %T, 值是 %v\n", i, i)

	var j int8 = 127 // -128 ~ 127
	fmt.Printf("j 的类型是 %T, 值是 %v\n", j, j)

	var k uint8 = 255 // 0 ~ 255
	fmt.Printf("k 的类型是 %T, 值是 %v\n", k, k)

	// int, uint, rune, byte 的使用
	// 1. int: 根据操作系统决定是 int32 还是 int64
	// 2. uint: 根据操作系统决定是 uint32 还是 uint64

	// 3. rune: int32 的别名, 表示一个 Unicode 码
	var name = '张' // 注意这里是单引号
	fmt.Printf("name 的类型是 %T, 值是 %v\n", name, name)
	// 4. byte: uint8 的别名, 表示一个 ASCII 码
	var char byte = 'a' // 注意这里是单引号
	fmt.Printf("char 的类型是 %T, 值是 %v\n", char, char)

	// 5. 整形的使用细节
	// 5.1. 在 Go 中, int 和 int32 是不同的类型, 但是可以相互赋值
	var n1 int32 = 10
	var n2 = int(n1)
	fmt.Printf("n2 的类型是 %T, 值是 %v\n", n2, n2)
	// 5.2. 整形的默认类型是 int
	var n3 = 10
	fmt.Printf("n3 的类型是 %T, 值是 %v\n", n3, n3)

	// 如何在程序查看某个变量的占用字节大小和数据类型
	fmt.Printf("n3 的类型是 %T, 字节大小是 %d\n", n3, unsafe.Sizeof(n3))

	// 在保证程序正确运行的前提下, 尽量使用占用空间小的数据类型

	// bit是计算机中最小的存储单位, 1 byte = 8 bit
}
