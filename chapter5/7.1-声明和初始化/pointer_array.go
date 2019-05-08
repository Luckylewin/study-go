package main

import "fmt"

// 所以在函数中数组作为参数传入时，如 func1(arr2)，会产生一次数组拷贝，func1 方法不会修改原始的数组 arr2。
// 如果你想修改原数组，那么 arr2 必须通过&操作符以引用方式传过来，例如 func1(&arr2），下面是一个例子
func main() {
	var ar [3]int
	f(ar)
	fp(&ar)
	fmt.Println(ar)
}

func f(a [3]int) {
	a[0] = 1
	fmt.Println(a)
}

func fp(a *[3]int) {
	a[0] = 3
	fmt.Println(a)
}