package main

import "fmt"

func main() {
	p2 := Add2()
	fmt.Printf("Call Add2 for 3 gives:%v\n", p2(3))

	adder := Adder(2)
	fmt.Printf("Call Add2 for 3 :%v\n",adder(3))
}

func Add2() func(b int) int {
	return func(b int) int {
		return b + 2
	}
}

func Adder(a int) func(b int) int {
	return func(b int) int {
		return a + b
	}
}