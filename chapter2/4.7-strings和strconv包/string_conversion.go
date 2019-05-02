package main

import (
	"fmt"
	"strconv"
)

func main() {
	var orig string = "666"
	var an int
	var newStr string

	fmt.Printf("The size of ints is: %d\n", strconv.IntSize)

	an, _ = strconv.Atoi(orig)
	fmt.Printf("The integer is: %d\n",an)
	an += 5
	newStr = strconv.Itoa(an)
	fmt.Printf("The new string is:%s\n",newStr)
}