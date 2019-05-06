package main

import "fmt"

func main() {
	var f = Adder()
	fmt.Println(f(1))   // 1
	fmt.Println(f(20))  // 21
	fmt.Println(f(300)) // 321
	// 三次调用函数 f 的过程中函数 Adder() 中变量的 delta 的值分别为 1,20,300
	// 我们可以看到 在多次调用中，变量x的值是被保留的，即 0 + 1 = 1 ,1 + 20 = 21 ,21 + 300 = 321
	// 闭包函数保存并积累其中变量的值，不管外部函数退出与否，他都能继续操作外部函数中的局部变量
}

func Adder() func(int) int {
	var x int
	return func(delta int) int {
		x += delta
		return x
	}
}