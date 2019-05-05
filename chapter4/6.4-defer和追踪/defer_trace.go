package main

import "fmt"

func main() {
	b()
}

func a() {
	trace("a")
	defer untrace("a")
	fmt.Println("in a")
}

func b() {
	trace("b")
	defer untrace("b")
	fmt.Println("in b")
	a()
}

// 使用 defer 语句实现代码追踪
// 一个基础但十分实用的实现代码执行追踪的方案就是在进入和离开某个函数打印相关的消息
// 即可以提炼为下面两个函数

func trace(s string) { fmt.Println("Entering:", s) }
func untrace(s string) { fmt.Println("Leaving:", s) }