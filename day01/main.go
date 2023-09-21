package main

func main() {
	// q: 什么是程序?
	// a: 一堆指令的集合

	// q: 什么是指令?
	// a: 一条条的命令

	// q: Google 为什么要开发 Go 语言?
	// a: 合理利用多核计算机的计算能力

	// q: Go 语言的特点?
	// a: 1. 语法简单
	//    2. 并发编程 (协程: goroutine, 通信: channel)
	//    2.1. goroutine是一种轻量级的线程, 由Go运行时环境管理
	//    3. 内存管理
	//    4. 丰富的标准库
	//    5. 开源
	//    6. 静态类型
	//    7. 编译型
	//    8. 垃圾回收
	//    9. 丰富的内置数据类型
	//    10. 函数多返回值
	//    11. 错误处理
	//    12. 丰富的内置数据类型
	// ...

	// q: Go 的GMP模型?
	// a: G: goroutine, M: machine, P: processor
	//    1. G: goroutine, 一个goroutine就是一个轻量级的线程
	//    2. M: machine, 一个M就是一个内核线程, 由操作系统调度
	//    3. P: processor, 一个P就是一个goroutine的执行上下文环境
	//    4. GMP模型的特点: M与P是一对一的关系, M与G是多对多的关系, P与G是一对多的关系
	//    5. GMP模型的好处: 1. 降低创建线程的成本 2. 降低调度线程的成本 3. 提高程序的并发性能

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

}
