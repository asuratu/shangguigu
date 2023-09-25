package main

import "fmt"

func main() {
	// 演示 Golang 中, 字符类型的使用

	// Golang 中没有专门的字符类型, 如果要存储单个字符(字母), 一般使用 byte 来保存

	// 1. 字符使用单引号
	// 2. 字符使用 UTF-8 编码, UTF-8 是 Unicode 的实现
	// 3. 字符的本质是整数, 直接输出时, 是该字符对应的 UTF-8 编码的码值
	// 4. 可以直接给某个变量赋一个数字, 然后按格式化输出时 %c, 会输出该数字对应的 unicode 字符
	// 5. 字符串是由多个字符组成的, 因此, 字符串也是可以进行遍历的

	var c1 byte = 'a'
	var c2 byte = '0'
	var c3 = '北'
	// 当我们直接输出 byte 值, 就是输出了对应的字符的码值
	fmt.Println("c1 =", c1, "c2 =", c2, "c3 =", c3)
	// 如果我们希望输出对应字符, 需要使用格式化输出
	fmt.Printf("c1 = %c, c2 = %c, c3 = %c\n", c1, c2, c3)
	fmt.Printf("c3的类型是 %T\n", c3) // int32

	// 英文字母是一个字节
	// 汉字是三个字节

	var c4 = 98
	fmt.Printf("c4 = %c\n", c4)

	// 字符类型是可以运算的, 相当于一个整数, 因为它都对应有 Unicode 码
	var n1 = 10 + 'a' // 10 + 97 = 107
	fmt.Println("n1 =", n1)

}
