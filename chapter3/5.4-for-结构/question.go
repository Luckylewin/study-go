package main

import "fmt"

func main() {
	// 练习5.9 以下程序输出的结果是什么
	for i:=0; i<5; i++ {
		var v int
		fmt.Printf("%d", v)
		v = 5
	}
	// 00000 这是因为每次循环 v是重新声明的

}