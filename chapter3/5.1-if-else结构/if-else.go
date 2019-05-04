package main

import (
	"fmt"
	"runtime"
)

func main() {
	bool1 := true
	// 注意: 这里不需要使用 if bool1 == true 来判断，因为 bool1 本身已经是一个布尔类型的值。
	if bool1	{
		fmt.Printf("The value is true\n")
	} else {
		fmt.Printf("The value is false\n")
	}

	var str string
	// 判断一个字符串是否为空
	if str == ""	{
		fmt.Printf("The value of str is empty\n")
	}

	if len(str) == 0	{
		fmt.Printf("The value of str is empty\n")
	}

	// 判断Go程序的操作系统，可以通过常量 runtime.GOOS 来判断
	if runtime.GOOS == "windows"	{
		fmt.Printf("Go run in windows\n")
	} else {
		fmt.Printf("Go run in Unix-like\n")
	}	

	// if 可以包含一个初始化语句(如给一个变量赋值) 这种写法具有固定的格式(在初始化语句后必须加上分号)
	
	/**
		if initialization; condition {
			// do something
		}

		例如:
		val := 10
		if val > max {
			// do something
		}

		你可以这么写
		if val := 10; val > max {
			// do something
		}

		使用简短方式 := 声明的变量的作用域只存在于if结构中 
		(在if结构的大括号之间，如果使用if-else结构 则在else代码块变量中也会存在)
		如果变量在if结构之前就已经存在，那么在if结构中，该变量原来的值就会被隐藏
		最简单的解决方案：不要在初始化语句中声明变量
	**/

}

// 函数 Abs 用于返回一个整型数字的绝对值
func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// isGreater 用于比较两个整型数字的大小
func isGreater(x,y int) bool {
	if x > y {
		return true
	}
	return false
}