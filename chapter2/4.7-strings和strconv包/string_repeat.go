package main

import (
	"strings"
	"fmt"
)

func main() {
	var str = "Hello World "
	var newStr string

	newStr = strings.Repeat(str, 3)
	fmt.Printf("orgin string         : \"%s\"\n", str)
	fmt.Printf("after Repeat 3 times : \"%s\"\n", newStr)
}