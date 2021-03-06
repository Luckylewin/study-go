package main

import "fmt"

func main() {
	var first int = 10
	var cond int

	if first <= 0 {
		fmt.Printf("first is less than or equal to 0\n")
	} else if first > 0 && first < 5 {
		fmt.Printf("first is between 0 and 5\n")
	} else {
		fmt.Printf("first is 5 or greater\n")
	}

	if cond = 5; cond > 10 {
		fmt.Printf("first is greater than 10\n")
	} else {
		fmt.Printf("first is not greater than 10\n")
	}

	// 下面的代码片段展示了如何通过在初始化语句中获取函数 process() 的返回值
	// 并在条件语句中作为判定条件来决定是否执行 if 结构中的代码
	if value := process("dev"); value > 50 {
		fmt.Printf("value is greate than 50\n")
	}
}

func process(data string) int {
	if data == "dev" {
		return 100
	}
	return 0
}