package main

import "fmt"

// 编写一个函数，接收两个整数，然后返回它们的和、积与差。编写两个版本，一个是非命名返回值，一个是命名返回值。
func main() {
	sum, prod, diff := caculate1(10, 1)
	fmt.Printf("sum = %d , prod = %d ,diff = %d\n", sum, prod, diff)
	sum, prod, diff = caculate2(10, 1)
	fmt.Printf("sum = %d , prod = %d ,diff = %d\n", sum, prod, diff)
}

func caculate1(a int, b int) (int,int,int) {
	return a + b, a * b, a - b 
}

func caculate2(a int, b int) (sum int, prod int, diff int) {
	sum, prod, diff = a + b, a * b, a - b
	return
}