package main

import (
	"fmt"
)

func Multiply(a, b int, reply *int) {
	*reply = a * b
}

// 仅仅是个指导性的例子
// 当需要在函数内改变一个占用内存比较大的变量时 性能优势就更加明显了
// 但是使用不小心的话 传递一个指针也容易引发不确定的事情
// 所以要十分小心改变外部变量的函数
// 在必要时，需要添加注释以便其他人能够知道函数里面到底发生了什么

func main() {
	n := 0
	reply := &n
	Multiply(10, 5, reply)
	fmt.Println("Multiply:",*reply)
}