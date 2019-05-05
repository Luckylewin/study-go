package main

import (
	"fmt"
	"strings"
)

func main() {
	originStr := "🙂Hello world!你好世界"
	fmt.Printf("Origin string is: %s\n", originStr)
	finalStr := MyChange1(originStr,isASCIIChar)
	fmt.Printf("(diy callback) after replace is: %s\n", finalStr)
	finalStr = MyChange2(originStr)
	fmt.Printf("(strings map)  after replace is: %s\n", finalStr)
}

// 包 strings 中的 Map 函数和 strings.IndexFunc() 一样都是非常好的使用例子。
// 请学习它的源代码并基于该函数书写一个程序，要求将指定文本内的所有非 ASCII 字符替换成 ? 或空格。您需要怎么做才能删除这些字符呢？

// 使用回调函数
func MyChange1(s string, f func(s rune) bool) string {
	var newS string
	for _,char := range s {
		if res := f(char); res {
			newS += string(char)
		} else {
			newS += "?"
		}
	}
	return newS
}

func isASCIIChar(s rune) bool {
	 if int(s) <= 255 {
	 	return true
	 }
	 return false
}

// 使用strings.Map
func MyChange2(s string) string {
	return strings.Map(replace,s)
}

func replace(s rune) rune {
	 if int(s) <= 255 {
	 	return s
	 }
	 return '?'
}