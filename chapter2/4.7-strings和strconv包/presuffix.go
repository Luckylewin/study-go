package main

import (
	"fmt"
	"strings"
)

func main() {
	var str string = "This is an example of a string"
	fmt.Printf("True/Flase? Does the string \"%s\" have prefix %s?  ",str, "Th");
	fmt.Printf("%t\n",strings.HasPrefix(str, "Th"))
}

// [前缀和后缀的判断]
// HasPrefix 判断字符串是否以 prefix 开头
// strings.HasPrefix(s, prefix string) bool

// HasSuffix(s, suffix string) bool
// strings.HasSuffix(s, suffix string) bool