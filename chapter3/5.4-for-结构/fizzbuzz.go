package main

import "fmt"

const (
	FIZZ = 3
	BUZZ = 5
	FIZZBUZZ = 15
)

func main() {
	// 写一个从1打印到100的程序 
	// 但是每当遇到3的倍数，不打印相应的数字，但打印一次Fizz
	// 遇到5的倍数 打印Buzz
	// 对于同时为3和5的倍数的数时，打印FizzBuzz(使用switch语句)
	for i:=1; i<=100; i++ {
		switch {
			case i%FIZZ == 0:
				fmt.Println("Fizz")
			case i%BUZZ == 0:
				fmt.Println("Buzz")
			case i%FIZZBUZZ == 0:
				fmt.Println("FizzBuzz")	
			default:
				fmt.Println(i)		
		}
	}
}