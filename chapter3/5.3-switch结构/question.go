package main

import "fmt"

func main() {
	k := 6
	switch k {
		case 4:
			fmt.Println("was <= 4")
			fallthrough
		case 5:
			fmt.Println("was <= 5")
			fallthrough
		case 6:
			fmt.Println("was <= 6")
			fallthrough
		case 7:
			fmt.Println("was <= 7")
			fallthrough
		case 8:
			fmt.Println("was <= 8")
			fallthrough
		default:
			fmt.Println("default case")
	}

	// 在匹配的记录中遇到fallthrough后，后面的分支无条件进入，如果再遇到 fallthrough 继续无条件进入后面的分支
}