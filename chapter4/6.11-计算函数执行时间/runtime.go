package main

import (
	"fmt"
	"time"
)

func main() {
	// 在对比基准测试中
	// 最简单的一个办法就是在计算机开始之前设置一个时间 计算结束时的结束时间，最后取出他们的差值
	// 就是这个计算所消耗的时间 想要实现这样的做法，可以使用 time 包中的Now() 和 Sub 函数
	
	start := time.Now()
	longCalculation()
	end := time.Now()
	delta := end.Sub(start)
	fmt.Printf("longCalculation took this amount of time : %s\n", delta)

}

func longCalculation() {
	for i:=0;i<=41;i++ {
		fmt.Printf("%d:%d \n", i, fibonacci(i))
	}
}

func fibonacci(n int) (res int) {
	if n <= 1 {
		res = 1
	} else {
		res = fibonacci(n-1) + fibonacci(n-2)
	}
	return
}
