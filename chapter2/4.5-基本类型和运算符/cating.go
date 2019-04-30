package main

import "fmt"

func main() {
	var a int64 = 100
	var b int32

	// 因为不存在隐式转换所以要 类型转换
	b = int32(a)

	fmt.Printf("64 bit int is: %d\n", a)
	fmt.Printf("32 bit int is: %d\n", b)
}