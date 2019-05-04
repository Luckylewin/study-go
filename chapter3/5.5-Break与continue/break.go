package main

import "fmt"

func main() {
	// 使用 break 重写上节中的for2.go
	var i int = 5

	for {
		fmt.Printf("The variable i is now: %d\n", i)
		i = i - 1
		if i < 0 {
			break
		}
	}

	// break 跳出最里层的for循环
	// continue 仅跳过本次循环
}