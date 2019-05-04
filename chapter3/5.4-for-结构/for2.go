package main

import "fmt"

func main() {
	// for 结构的第二种形式时没有头部的条件判断迭代
	// 类似其它语言中的while循环，基本形式时 for 条件语句{}
	// 因为没有初始化以及修饰语句 所以省略;;
	var i int = 5

	for i >= 0 {
		i = i -1
		fmt.Printf("The variable i is now: %d\n", i)
	}

	// 无限循环的方式是
	// for {}
}