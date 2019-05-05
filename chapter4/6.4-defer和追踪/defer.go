package main

import "fmt"

func main() {
	// defer 用法类似于 Java 中的finally 语句块，一般用于释放某些已分配的资源

	// 使用defer 对一些函数执行完成后的收尾工作
	// 1.关闭文件流
	// defer file.Close()

	// 2.解锁一个加锁的资源
	// defer mu.Unlock

	// 3.打印最终报告
	// defer printFooter()

	// 4.关闭数据库链接
	// disconnectFromDB() 
	
	function1()
}

func function1() {
	fmt.Printf("In function1 at the top\n")
	defer function2()
	fmt.Printf("In function1 at the bottom\n")
}

func function2() {
	fmt.Printf("Function2:Deferred until the end of the calling function!")
}