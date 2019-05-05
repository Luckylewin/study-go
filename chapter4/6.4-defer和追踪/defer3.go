package main

import "fmt"

func main() {
	f()
}

func f() {
	// 当多个defer 行为被注册，它们会以逆序执行(类似栈,即后进先出)
	for i:=0; i<=5; i++ {
		defer fmt.Printf("defer %d\n", i)
	}
}