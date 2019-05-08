package main

import "fmt"

// 练习7.1：array_value.go: 证明当数组赋值时，发生了数组内存拷贝。
func main() {
	var arr1 [4]int
	arr2 := arr1
	fmt.Println("arr1的内存地址",&arr1[0])
	fmt.Println("arr2的内存地址", &arr2[0])

}