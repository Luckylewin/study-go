package main

import "fmt"

/*
	实现一个输出前 30 个整数的阶乘的程序。
	n! 的阶乘定义为：n! = n * (n-1)!, 0! = 1，因此它非常适合使用递归函数来实现。
	然后，使用命名返回值来实现这个程序的第二个版本。
	特别注意的是，使用 int 类型最多只能计算到 12 的阶乘，因为一般情况下 int 类型的大小为 32 位，继续计算会导致溢出错误。那么，如何才能解决这个问题呢？
	最好的解决方案就是使用 big 包（详见第 9.4 节）。
*/

func main() {
	for i:=0;i<=12;i++ {
		fmt.Printf("%d! = %d \n", i, factorial(i));
	}
}

func factorial(n int) (factorialVal int) {
	if n == 0 {
		return 1
	}

	factorialVal = n * factorial(n - 1)
	return
}