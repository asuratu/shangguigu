package main

func main() {
	// 演示 Golang 中, 字符类型的使用

	// Golang 中没有专门的字符类型, 如果要存储单个字符(字母), 一般使用 byte 来保存

	// 1. 字符使用单引号
	// 2. 字符使用 UTF-8 编码, UTF-8 是 Unicode 的实现
	// 3. 字符的本质是整数, 直接输出时, 是该字符对应的 UTF-8 编码的码值
	// 4. 可以直接给某个变量赋一个数字, 然后按格式化输出时 %c, 会输出该数字对应的 unicode 字符
	// 5. 字符串是由多个字符组成的, 因此, 字符串也是可以进行遍历的

}
