package main

import (
	"fmt"
	"time"
)

const LIM = 41

var fibs [LIM]uint64

func main() {
	// 当在进行大量计算时 提升性能最直接有效的一种方式时避免重复计算
	// 通过在内存中缓存和重复利用相同计算结果 称之为内存缓存
	// 最明显的例子时生成斐波那契数列
	// 要计算数列前n个数字，需要先得到之前两个数的值
	// 但是很显然绝大多数情况下,前两个数的值都是已经计算过的
	// 而我们要做的就是将第n个数的值存到数组中 索引为n的位置
	// 然后在数组中查找是否已经计算过，如果没有找到则再进行计算

	var result uint64 = 0
	start := time.Now()
	for i:=0; i < LIM; i++ {
		result = fibonacci(i)
		fmt.Printf("fibonacci(%d) = %d\n", i, result)
	}
	end := time.Now()
	delta := end.Sub(start)
	fmt.Printf("longCalculation took this amount of time: %s\n", delta)
	// 使用了内存缓存
	// longCalculation took this amount of time: 90.997µs
	// 不使用内存缓存
	// longCalculation took this amount of time : 6.508103372s
}

func fibonacci(n int) (res uint64) {
	// 判断 fibonacci(n) 是否已经存在于数组中
	if fibs[n] != 0 {
		res = fibs[n]
		return
	}

	if n <= 1 {
		res = 1
	} else {
		res = fibonacci(n-1) + fibonacci(n-2)
	}

	fibs[n] = res
	return
}