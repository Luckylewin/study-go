package main

import "fmt"

func main() {
	// 程序在内存中存储值
	// 每个内存块有一个地址，通常用十六进制表示，如0x6b0820,0xf8001d7f0
	// 取地址符是 &,放到一个变量前就会返回相应变量的内存地址
	var year = 1993
	fmt.Printf("an integer: %d, it's location in memory:%p\n",year, &year)

	// 这个地址可以存储到一个叫做指针的特殊数据类型中，在本例中这是一个指向int的指针，即int1:
	// 此处使用 *int 表示，如果我们想调用指针 intP 我们可以这样声明
	var intP *int
	intP = &year
	fmt.Printf("use var intP pointer to the location of var year: %p\n",intP)

	// 一个指针变量可以指向任何一个值的内存地址
	// 在32位机器上占用4字节，在64位机器上占用8个字节
	// 在指针类型前面加上 * 号(前缀)来获取指针所指向的内容,这里的*号是一个类型更改器
	// 一个指针的简写是 ptr
	fmt.Printf("use * to get the value of the pointer value that it point to：%d \n", *intP)

	// 指针针对string的例子
	// 它展示了分配给一个新的值 并且更改了这个变量自己的值(这里是一个字符串)

	s := "hello Golang"
	fmt.Printf("Here is the string s p: %s\n", s)

	var p *string = &s
	*p = "Hello World"

	fmt.Printf("Here is the pointer p: %p\n", p)
	fmt.Printf("Here is the string *p: %s\n", *p)
	fmt.Printf("Here is the string s p: %s\n", s)
	
	// [注意事项]
	// 无法得到一个文字或常量的地址
	// 指针的一个高级应用是你可以传递一个变量的引用（如函数的参数），这样不会传递变量的拷贝
	// 当一个指针被定义后没有分配到任何变量时，它的值为 nil。
	// 对一个空指针的反向引用是不合法的，并且会使程序崩溃 如
	// var p *int = nil
	// *p = 0 

}