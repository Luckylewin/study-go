package main

import "fmt"

const NUM = 50

func main() {
	fmt.Println(Fibonacci(NUM))
}

// 练习7.3：fibonacci_array.go
// [数组方法] 实现斐波那契数列
func Fibonacci(n int) (arr []uint64) {
	arr = make([]uint64, n)
	arr[0] , arr[1] = 1, 1

	for i:=2; i<NUM; i++ {
		arr[i] = arr[i-1] + arr[i-2]
	}

	return
}