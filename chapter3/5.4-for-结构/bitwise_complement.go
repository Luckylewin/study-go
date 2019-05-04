package main

import "fmt"

func main() {
	//练习: 按位补码从0到10 使用位表达式 %b 格式化输出
	for i:=0; i<=10; i++ {
		fmt.Printf("the complement of %b is: %b\n", i, ^i)
	}
}