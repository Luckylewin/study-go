package main

import (
	"fmt"
	"strings"
)

func main() {
	var str = "Good morning, and in case I don't see you, good afternoon, good evening, and good night!"
	var lowerStr,upperStr string
	
	lowerStr = strings.ToLower(str)
	upperStr  = strings.ToUpper(str)

	fmt.Printf("origin string is \"%s\"\n",str)	
	fmt.Printf("lowerStr string is \"%s\"\n",lowerStr)	
	fmt.Printf("uperStr string is \"%s\"\n",upperStr)	
}