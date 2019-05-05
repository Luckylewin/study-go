package main

import "fmt"

func main() {
	work(10)
}

// 练习 6.5 使用递归函数从 10 打印到 1
func work(n int) int {
	if n == 1 {
		fmt.Println(1)
		return 1
	}

	fmt.Println(n)
	next := work(n - 1)

	return next
}