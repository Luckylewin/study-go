package main

import "strings"

func main() {
	suffix := MakeAddSuffix(".png")
	println(suffix("hello"))	
	println(suffix("world"))	
}

// 一个返回值为另一个函数的函数被称之为工厂函数

// 下面的函数演示了如何动态追加后缀的函数
func MakeAddSuffix(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}

	// 可以返回其它函数的函数和接受其它函数作为参数的函数均被称之为高阶函数
	// 是函数式语言的特点

}