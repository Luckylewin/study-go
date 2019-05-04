package main

import (
	"fmt"
	"strconv"
)

func main() {
	var originStr string = "ABC"
	// var an int
	var newS string
	// var err error

	fmt.Printf("The size of ints is: %d\n", strconv.IntSize)

	// 旧版的方式
	// anInt := strconv.Atoi(originStr)
	anInt, err := strconv.Atoi(originStr)
	if err != nil {
		fmt.Printf("originStr %s is not an integer - exiting with error \n", originStr)
		return
	}

	fmt.Printf("The integer is %d\n",anInt)
	anInt += 5
	newS = strconv.Itoa(anInt)
	fmt.Printf("The new string is: %s\n", newS)
}