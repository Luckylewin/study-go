package main

import "fmt"

func main() {
	// for 初始化语句; 条件语句; 修饰语句 {}
	for i := 0; i < 5; i++ {
		fmt.Printf("This is the %d iteration\n", i)
	}

	// 可以在for循环使用多个计数器
	// for i, j := 0, N; i < j; i, j = i+1,j-1 {}

	// 可以将两个 for 循环嵌套起来
	for i:=0; i<5; i++ {
		for j:=0;j<10;j++ {
			fmt.Printf("%d * %d = %d\n", i, j, i*j)
		}
	}
}