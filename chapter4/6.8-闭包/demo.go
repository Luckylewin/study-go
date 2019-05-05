package main

import "fmt"

func main() {
	func() {
		sum := 0 
		for i:=0; i<=100; i++ {
			sum += i
		}
		fmt.Printf("1+...+100 = %d\n", sum)
	}()
}