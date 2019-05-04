package main

import "fmt"

func main() {
	// 使用嵌套循环
	for i:=0; i<25; i++ {
		for j:=0; j<i; j++ {
			fmt.Printf("G")
		} 
		fmt.Println()
	}

	str := "G"
	// 使用一次循环
	for i:=1; i<=25; i++ {
		fmt.Println(str)
		str += "G"
	}
}