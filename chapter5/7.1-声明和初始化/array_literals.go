package main

import "fmt"

func main() {
	var arrAge = [5]int{1,2,3,4,5}
	var arrLazy = [...]int{5,6,7,8,22}
	var arrLazy2 = []int{5,6,7,8,22} // 得到的是切片
	var arrKeyValue = [5]string{3: "Chris", 4: "Ron"}
	var arrKeyValue2 = []string{3: "Chris", 4: "Ron"} // 初始化得到的是切片

	// 第一种变化
	var arrAge = [5]int{18, 20, 15, 22, 16}
	// 第二种变化
	var arrLazy = [...]int{5,6,7,8,22}
	// 第三种变化 key:value
	var arrKeyValue = [5]string{3: "Chris", 4: "Ron"}


}