package main

import "fmt"

func main() {
	// 打印宽为20,高位10的矩形
	w, h := 20, 10
	for i:=1; i<=h; i++ {
		for j:=1; j<=w; j++ {
			fmt.Printf("*")
		}
		fmt.Println()
	}
}