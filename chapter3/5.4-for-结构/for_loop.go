package main

import "fmt"

func main() {
	// for 循环 十五遍
	for i:=0; i<15; i++ {
		fmt.Printf("(use for)The counter is at %d\n", i)
	}
	// 不使用 for 使用goto进行循环
	i := 0
START:
	fmt.Printf("(use goto)The counter is at %d\n", i)
	i++
	if i < 15 {
		goto START
	}	
}