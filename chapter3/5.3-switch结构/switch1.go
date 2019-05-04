package main

import "fmt"

func main() {
	var num1 int = 100
	// [switch 形式1]
	switch num1 {
		case 98,99:
			fmt.Println("It's equal to 98 or 99")
		case 100:
			fmt.Println("It's equal to 100")
		default:
			fmt.Println("It's not equal to 98 or 100")		
	}

	// [switch 形式2] 任何支持进行相等判断的类型都可以作为测试表达式的条件 包括 int, string, 指针等
	var num2 int = 7
	switch {
		case num2 < 0:
			fmt.Println("Number is negative")
		case num2 > 0 && num2 < 10:
			fmt.Println("Number is between 0 and 10")
		default:
			fmt.Println("Number is 10 or greater")
	}

	// [switch 形式3] 包含初始化语句

	switch result := calculate(); {
		case result < 0:
			fmt.Println("result lower than 0")
		case result > 0 && result < 100:
			fmt.Println("result bigger than 0 and lower than 100")
		default:
			fmt.Println("result bigger than 100")	
	}
}

func calculate() int {
	return 10
}