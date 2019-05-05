package main

import "fmt"

// 重写本节中生成斐波那契数列的程序并返回两个命名返回值（详见第 6.2 节），即数列中的位置和对应的值，例如 5 与 4，89 与 10。
func main() {
	index,val := 0, 0
	for i:=0; i<=10; i++ {
		index,val = fibonacci(i)
		fmt.Printf("%d-%d\n",index,val)
	}
}

func fibonacci(n int) (n int, val int) {
	if n < 1 {
		return 1,1
	}

	_,temp1 := fibonacci(n - 1)
	_,temp2 := fibonacci(n - 2)
	val = temp1 + temp2
	ix = n

	return 
}