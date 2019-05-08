package main

import (
	"runtime"
	"log"
)

// 使用闭包调试
// 当您分析和调试复杂程序时，无数个函数在不同的代码文件中相互调用
// 如果这时候能够准确的知道哪个文件中具体哪个函数正在执行
// 那么对于调试是十分有帮助的,您可以使用 runtime 或 log 包中特殊的函数来实现这样的功能
// 包 runtime 中的函数 Caller() 提供了相应的信息，
// 因此可以在需要的时候实现一个 where() 闭包函数来打印函数执行的位置

func main() {
	println("Hello world")

	where := func() {
		_, file, line, _ := runtime.Caller(1)
		log.Printf("%s:%d", file, line)
	}

	where()

	// 也可以设置Log包中flag参数来实现
	// log.Llongfile
	// log.Print("")

}