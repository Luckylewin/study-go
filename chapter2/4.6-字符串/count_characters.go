package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	// 创建一个统计字节和字符(rune)的程序 
	// 并对字符串asSASA ddd dsjkdsjs dk 进行分析
	// 然后再分析 asSASA ddd dsjkdsjsこん dk，

	str1 := "asSASA ddd dsjkdsjs dk"
	fmt.Printf("The number of bytes in string str1 is %d\n",len(str1))
	fmt.Printf("The number of characters in string str1 is %d\n",utf8.RuneCountInString(str1))

	// 不一致的结果是因为非ASCII存储的字符使用了两字节以上进行存储
	str2 := "asSASA ddd dsjkdsjsこん dk"
	fmt.Printf("The number of bytes in string str2 is %d\n",len(str2))
	fmt.Printf("The number of characters in string str2 is %d\n",utf8.RuneCountInString(str2))
}