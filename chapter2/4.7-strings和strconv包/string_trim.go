package main

import (
	"fmt"
	"strings"
)

func main() {
	var str = " We use Golang as our API language "
	var str2 = "golang I studying Golang now"
	var newStr string

	newStr = strings.TrimSpace(str)
	fmt.Printf("orgin string :\"%s\"\n", str)
	fmt.Printf("after trim space :\"%s\"\n", newStr)

	fmt.Printf("orgin string :\"%s\"\n", str2)
	newStr = strings.TrimLeft(str2,"golang")
	fmt.Printf("after trim Left 'golang':\"%s\"\n", newStr)
}