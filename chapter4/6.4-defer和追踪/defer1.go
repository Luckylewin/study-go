package main

import "fmt"

func main() {
	i := a()
	println("value is ", i)
}

// 使用 defer 的语句同样可以接受参数，下面这个例子就会在执行 defer 语句时打印 10
func a() int {
	i := 10
	defer fmt.Println(i)
	i++
	return i
}