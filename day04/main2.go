package main

import (
	"fmt"
)

func main() {
	// q: Go 语言的基本数据类型?
	// a: 1. 整型: int8, int16, int32, int64, uint8, uint16, uint32, uint64, int, uint
	//    2. 浮点型: float32, float64
	//    3. 复数: complex64, complex128
	//    4. 布尔型: bool (true, false)
	//    5. 字符串: string
	//    6. 字符: byte, rune (rune是int32的别名, 用于表示Unicode字符)
	//    7. 指针: pointer
	//    8. 数组: array
	//    9. 切片: slice
	//    10. 映射: map
	//    11. 通道: channel
	//    12. 结构体: struct
	//    13. 接口: interface (interface{}是空接口, 任何类型都实现了空接口, any 可以替代 interface{})
	//    14. 函数: function (func)
	//	  15. 类型别名: type
	type MyInt int32
	var a MyInt = 10
	fmt.Println(a)
	fmt.Printf("%T\n", a)

	// + 号的使用
	// 1. 如果 + 两边都是数值型, 则直接相加
	var n1 = 10
	var n2 = 20
	var c = n1 + n2
	fmt.Println(c)
	// 2. 如果 + 两边都是字符串, 则表示拼接
	var s1 = "hello"
	var s2 = "world"
	var s3 = s1 + s2
	fmt.Println(s3)

}
